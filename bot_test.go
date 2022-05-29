package feishu_bot

import (
	"testing"
)

func TestSendText(t *testing.T) {
	webhook := "http://xxxx"
	bot := NewBot(webhook)
	err := bot.sendText("123")
	if err != nil {
		t.Fatalf("test send text error: %s", err.Error())
	}
}
