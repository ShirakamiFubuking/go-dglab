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
	url := dglab.WaitConnect("192.168.137.1", 8888, "", func(c *dglab.Coyote) {
		// 註冊事件監聽
		c.Handle(func(pa, pb, la, lb int) {
			// APP反饋強度設定
			log.Printf("當前強度/上限A=%d/%d, B=%d/%d\n", pa, la, pb, lb)
			// 設置強度為上限
			if pa != la {
				c.SetPower(dglab.CHANNEL_A, la)
			}
			if pb != lb {
				c.SetPower(dglab.CHANNEL_B, lb)
			}
		}, func(button int) {
			// app按鍵反饋
		})
		// A通道
		go func() {
			wave := dglab.DefaultWaves["呼吸"]
			for {
				c.SendWave(dglab.CHANNEL_A, wave)
				<-time.After(wave.Duration())
			}
		}()
		// B通道
		go func() {
			wave := dglab.DefaultWaves["快速按捏"]
			for {
				c.SendWave(dglab.CHANNEL_B, wave)
				<-time.After(wave.Duration())
			}
		}()
	})
	// 產生QRCode
	dglab.GenerateQRCode("qrcode.png", url)
	select {}
}
```