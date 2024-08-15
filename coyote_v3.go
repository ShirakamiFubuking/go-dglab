package godglab

import (
	"fmt"
	"log"
	"net/http"
	"regexp"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/google/uuid"
	"github.com/gorilla/websocket"
	"github.com/skip2/go-qrcode"
)

const (
	CHANNEL_A = iota
	CHANNEL_B
)

type Coyote struct {
	// lock for write websocket
	lock sync.Mutex
	// Websocket
	ws *websocket.Conn
	// ID
	client, target string
	// handle事件
	handleStrength func(int, int, int, int)
	handleFeedback func(int)
	// 儲存當前強度狀態
	PowerA, PowerB, LimitA, LimitB int
}

type Message struct {
	Type     string `json:"type"`
	ClientID string `json:"clientId"`
	TargetID string `json:"targetId"`
	Msg      string `json:"message"`
}

// addr ip:port
func WaitConnect(ip string, port int, id string, onConnect func(c *Coyote)) string {
	// t, _ := uuid.Parse("02e5fed7-d088-4a49-97c9-c1c238476907")
	// clientID := t.String()
	clientID := id
	if clientID == "" {
		clientID = uuid.New().String()
	}
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		var upgrader = websocket.Upgrader{}
		ws, err := upgrader.Upgrade(w, r, nil)
		if err != nil {
			log.Printf("err:%v", err)
			return
		}
		// 發送ID
		targetID := uuid.New().String()
		ws.WriteJSON(&Message{
			Type:     "bind",
			ClientID: targetID,
			TargetID: "",
			Msg:      "targetId",
		})
		var msg Message
		// 等待綁定
		err = ws.ReadJSON(&msg)
		if err != nil {
			log.Printf("讀取訊息錯誤: %v", err)
			return
		}
		log.Printf("msg:%v", msg)
		if msg.Type != "bind" {
			return
		}
		ws.WriteJSON(&Message{
			Type:     "bind",
			ClientID: clientID,
			TargetID: targetID,
			Msg:      "200",
		})
		// 綁定成功
		c := &Coyote{
			ws:     ws,
			client: clientID,
			target: targetID,
		}
		onConnect(c)
		// 心跳
		go func() {
			ticker := time.NewTicker(1 * time.Minute)
			defer ticker.Stop()
			for {
				<-ticker.C
				if err := c.sendMessage(&Message{
					Type:     "msg",
					ClientID: c.client,
					TargetID: c.target,
					Msg:      "heartbeat"}); err != nil {
					break
				}
			}
		}()
		for {
			var msg Message
			err := ws.ReadJSON(&msg)
			if err != nil {
				log.Printf("read message error: %v", err)
				break
			}
			if msg.Type == "msg" && msg.ClientID == clientID && msg.TargetID == targetID {
				// 只有strength-和feedback-消息
				if c.handleStrength != nil && strings.HasPrefix(msg.Msg, "strength-") {
					c.PowerA, c.PowerB, c.LimitA, c.LimitB = parseStrength(msg.Msg)
					c.handleStrength(c.PowerA, c.PowerB, c.LimitA, c.LimitB)
				} else if c.handleFeedback != nil && strings.HasPrefix(msg.Msg, "feedback-") {
					re := regexp.MustCompile(`\d+`)
					a, _ := strconv.Atoi(re.FindString(msg.Msg))
					c.handleFeedback(a)
				}
			}
		}
	})
	go http.ListenAndServe(fmt.Sprintf("0.0.0.0:%d", port), nil)
	return connectURL(fmt.Sprintf("%s:%d", ip, port), clientID)
}

func GenerateQRCode(filepath, url string) {
	err := qrcode.WriteFile(url, qrcode.Medium, 256, filepath)
	if err != nil {
		log.Printf("failed generating qrcode:%v", err)
	}
}

func (c *Coyote) Handle(handleStrength func(powerA int, powerB int, limitA int, limitB int), handleFeedback func(button int)) {
	c.handleStrength = handleStrength
	c.handleFeedback = handleFeedback
}

// 波形操作
// channel: 大寫A或B
func (c *Coyote) SendWave(channel int, w Wave) {
	c.sendWave(channel, w)
}

// 強度操作
func (c *Coyote) SetPower(channel, power int) {
	msg := c.newMessagePower(channel, power)
	c.sendMessage(msg)
}

// 清空佇列
func (c *Coyote) Clear(channel int) {
	var t string
	if channel == CHANNEL_A {
		t = "clear-1"
	} else {
		t = "clear-2"
	}
	c.sendMessage(&Message{
		Type:     "msg",
		ClientID: c.client,
		TargetID: c.target,
		Msg:      t,
	})
}

func parseStrength(input string) (int, int, int, int) {
	re := regexp.MustCompile(`\d+`)
	// 提取字符串中的所有数字
	matches := re.FindAllString(input, -1)
	if len(matches) != 4 {
		fmt.Println("Error: Expected exactly 4 numbers in the input string")
		return 0, 0, 0, 0
	}

	// 转换字符串为整数并赋值给变量a, b, c, d
	a, _ := strconv.Atoi(matches[0])
	b, _ := strconv.Atoi(matches[1])
	c, _ := strconv.Atoi(matches[2])
	d, _ := strconv.Atoi(matches[3])
	return a, b, c, d
}

// addr: <ip>:<port>
func connectURL(addr, clientID string) string {
	return fmt.Sprintf(`https://www.dungeon-lab.com/app-download.php#DGLAB-SOCKET#ws://%s/%s`, addr, clientID)
}

func (c *Coyote) sendMessage(msg *Message) error {
	c.lock.Lock()
	defer c.lock.Unlock()
	return c.ws.WriteJSON(*msg)
}

func (c *Coyote) sendWave(channel int, w Wave) {
	pulses := w.GetPulses()
	m := c.newMessageWave(channel, pulses)
	c.sendMessage(m)
}

func (c *Coyote) newMessageWave(channel int, ps []Pulse) *Message {
	var buf strings.Builder
	buf.WriteString("pulse-")
	if channel == CHANNEL_A {
		buf.WriteString("A")
	} else {
		buf.WriteString("B")
	}
	buf.WriteString(`:["`)
	buf.WriteString(ps[0].ToString())
	buf.WriteString(`"`)
	for _, p := range ps[1:] {
		buf.WriteString(`,"`)
		buf.WriteString(p.ToString())
		buf.WriteString(`"`)
	}
	buf.WriteString(`]`)
	return &Message{
		Type:     "msg",
		ClientID: c.client,
		TargetID: c.target,
		Msg:      buf.String()}
}

func (c *Coyote) newMessagePower(channel int, power int) *Message {
	var buf strings.Builder
	buf.WriteString("strength-")
	if channel == CHANNEL_A {
		buf.WriteString("1+")
	} else {
		buf.WriteString("2+")
	}
	buf.WriteString("2+")
	buf.WriteString(strconv.Itoa(int(power)))
	return &Message{
		Type:     "msg",
		ClientID: c.client,
		TargetID: c.target,
		Msg:      buf.String()} //`strength-x+x+x`}
}
