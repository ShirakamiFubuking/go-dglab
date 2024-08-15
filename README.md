## [DGLAB官方開源](https://github.com/DG-LAB-OPENSOURCE/DG-LAB-OPENSOURCE) APP [websocket](https://github.com/DG-LAB-OPENSOURCE/DG-LAB-OPENSOURCE/blob/main/socket/README.md)的golang實現

### !!!使用前請務必詳讀安全須知!!!
## !!!使用前請務必詳讀安全須知!!!
# !!!使用前請務必詳讀安全須知!!!
---
- 僅支持3.0
- 包含APP預設基礎波形
- 波形相關基礎請閱讀[這個](https://github.com/DG-LAB-OPENSOURCE/DG-LAB-OPENSOURCE/blob/main/coyote/extra/README.md)和[這個](https://github.com/DG-LAB-OPENSOURCE/DG-LAB-OPENSOURCE/blob/main/coyote/v3/README_V3.md)
## Example
```golang
package main

import (
	"log"
	"time"

	dglab "github.com/ShirakamiFubuking/go-dglab"
)

func main() {
	// IP是APP可直接訪問之IP
	url := dglab.WaitConnect("192.168.137.1", 8888, "", func(c *dglab.Coyote) { // id留空會自動生成uuid
		// 註冊事件監聽
		c.Handle(func(powerA, powerB, limitA, limitB int) {
			// APP反饋強度設定
			log.Printf("當前強度/上限A=%d/%d, B=%d/%d\n", powerA, powerB, limitA, limitB)
			// 設置強度為上限
			if powerA != limitA {
				c.SetPower(dglab.CHANNEL_A, limitA)
			}
			if powerB != limitB {
				c.SetPower(dglab.CHANNEL_B, limitB)
			}
		}, func(button int) {
			// app按鍵反饋
		})
		// 計時10秒
		stop := false
		go func() {
			<-time.After(10 * time.Second)
			stop = true
		}()
		// A通道
		go func() {
			wave := dglab.DefaultWaves["呼吸"]
			for !stop {
				c.SendWave(dglab.CHANNEL_A, wave)
				<-time.After(wave.Duration()) // 可以略少於波形時間
			}
			c.Clear(dglab.CHANNEL_A) // 觸發停止事件時清空queue
		}()
		// B通道
		go func() {
			wave := dglab.DefaultWaves["快速按捏"]
			for !stop {
				c.SendWave(dglab.CHANNEL_B, wave)
				<-time.After(wave.Duration() - (5 * time.Millisecond)) // 可以略少於波形時間
			}
			c.Clear(dglab.CHANNEL_B) // 觸發停止事件時清空queue
		}()
	})
	// 產生QRCode
	dglab.GenerateQRCode("qrcode.png", url)
	select {}
}
```