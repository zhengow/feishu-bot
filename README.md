# feishu-bot

```go
package main

import zFeishu "github.com/zhengow/feishu-bot"

func main() {
	webhook := "http://xxxx"
	bot := zFeishu.NewBot(webhook)
	err := bot.sendText("123")
}

```