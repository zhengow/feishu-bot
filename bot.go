package feishu_bot

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type ZBot struct {
	webhook string
}

type zParams struct {
	MsgType string   `json:"msg_type"`
	Content zContent `json:"content"`
}

type zContent struct {
	Text string `json:"text"`
}

type zResp struct {
	Code          uint64 `json:"code"`
	Msg           string `json:"msg"`
	StatusCode    uint64 `json:"StatusCode"`
	StatusMessage string `json:"StatusMessage"`
}

func NewBot(webhook string) *ZBot {
	return &ZBot{webhook: webhook}
}

func (b *ZBot) SendText(text string) error {
	p := &zParams{
		MsgType: "text",
		Content: zContent{
			Text: text,
		},
	}
	tmp, err := json.Marshal(p)
	if err != nil {
		return fmt.Errorf("json marshal params error: %s", err.Error())
	}
	req, err := http.NewRequest("POST", b.webhook, bytes.NewReader(tmp))
	if err != nil {
		return fmt.Errorf("new request error: %s", err.Error())
	}
	req.Header.Set("Content-Type", "application/json")
	client := http.Client{}
	resp, err := client.Do(req)

	if err != nil {
		return fmt.Errorf("send feishu text error: %s", err.Error())
	}

	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("read resp body error: %s", err.Error())
	}
	feishuResp := &zResp{}
	err = json.Unmarshal(data, feishuResp)
	if err != nil {
		return fmt.Errorf("json unMarshal response error: %s", err.Error())
	}

	if feishuResp.Code != 0 {
		return fmt.Errorf("feishu response error: %d %s", feishuResp.Code, feishuResp.Msg)
	}

	return nil
}
