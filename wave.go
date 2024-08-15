package godglab

import (
	"encoding/hex"
	"strings"
	"time"
)

// 發送最小單位, 0.1s
// 前4bytes為頻率 後4bytes為強度
type Pulse [8]byte // !!!!!長度等於8

// scale在[0~1)之間
func (p *Pulse) Scale(scale float32) *Pulse {
	if scale >= 1 || scale < 0 {
		return p
	}
	pulse := &Pulse{p[0], p[1], p[2], p[3]}
	for i := 4; i < 8; i++ {
		pulse[i] = byte(float32(p[i]) * scale)
	}
	return pulse
}

func (p *Pulse) ToString() string {
	return strings.ToUpper(hex.EncodeToString(p[:]))
}

// 波形,Pulse的集合
type Wave interface {
	GetPulses() []Pulse
}

type WavePatten struct {
	Name string
	Data []Pulse
}

func (w WavePatten) GetPulses() []Pulse {
	return w.Data
}

// 回傳波形總時間
func (w *WavePatten) Duration() time.Duration {
	return time.Duration(len(w.Data)) * 100 * time.Millisecond
}

func (w *WavePatten) Scale(s float32) *WavePatten {
	if s >= 1 || s < 0 {
		return w
	}
	newWave := &WavePatten{
		Data: make([]Pulse, len(w.Data)),
	}
	for _, p := range w.Data {
		newWave.Data = append(newWave.Data, *p.Scale(s))
	}
	return newWave
}
