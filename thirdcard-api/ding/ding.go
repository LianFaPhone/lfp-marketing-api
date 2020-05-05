package ding

import (
	"LianFaPhone/lfp-marketing-api/common"
	"LianFaPhone/lfp-marketing-api/config"
	"bytes"
	"encoding/json"
	"fmt"
)

const (
	robotUrl = "https://oapi.dingtalk.com/robot/send?access_token="
)

type Content struct {
	Content *string `json:"content,omitempty"`
}

type (
	ReDing struct{
		Msgtype *string `json:"msgtype,omitempty"`
		Text    Content `json:"text,omitempty"`
	}

	ResDing struct{
		Errcode int    `json:"errcode,omitempty"`
		Errmsg  string `json:"errmsg,omitempty"`
	}

)

func (this *ReDing) Send(dingBody string) ( error) {
	msg := new(ReDing)
	text := "text"
	msg.Msgtype = &text
	msg.Text.Content = &dingBody

	reqBody, err := json.Marshal(msg)
	if err != nil {
		return err
	}

	res, err := common.HttpSend(robotUrl+config.GConfig.DingDing.RobToken, bytes.NewBuffer(reqBody), "POST", nil)
	if err != nil {
		return err
	}

	robotRes := new(ResDing)
	err = json.Unmarshal(res, robotRes)
	if err != nil {
		return err
	}
	if robotRes.Errcode != 0 {
		return fmt.Errorf("%d_%s", robotRes.Errcode, robotRes.Errmsg)
	}
	return nil
}

