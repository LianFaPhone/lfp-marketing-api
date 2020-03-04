package sdk

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"

	"errors"
)

var GNotifySdk NotifySdk

type SmsSend struct {
	TempName *string  `valid:"optional" json:"temp_name,omitempty"` //groupname+lang 联合使用
	TempId   *int64   `valid:"optional" json:"temp_id,omitempty"`   //groupid+lang联合使用
	Params   []string `valid:"-" json:"params,omitempty"`           //optional
	Phone    []string `valid:"optional" json:"phone,omitempty"`     //require
	Author   *string  `valid:"optional" json:"author,omitempty"`
	PlayTp   int      `valid:"optional" json:"play_tp,omitempty"`   // 0短信，1语音
	IsRecord int      `valid:"optional" json:"is_record,omitempty"` // 0不记录，1记录
	ReTry    int      `valid:"optional" json:"retry,omitempty"`     // 0不，1是
	PlatformTp *int   `valid:"optional" json:"platform_tp,omitempty"`
}

type Response struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}

type NotifySdk struct {
	addr string
}

func (this *NotifySdk) Init(addr string) error {
	this.addr = addr
	return nil
}

func (this *NotifySdk) SendSms(Params []string, phone, tempName string, PlayTp int, PlatformTp *int) error {
	param := &SmsSend{
		TempName: &tempName,
		Params:   Params,
		Phone:    []string{phone},
		Author:   new(string),
		PlatformTp: PlatformTp,
	}
	data, err := json.Marshal(param)
	if err != nil {
		return err
	}
	resData, err := HttpSend(this.addr+"/v1/ft/notify/sms/send", bytes.NewReader(data), "POST", nil)
	if err != nil {
		return err
	}
	res := new(Response)
	if err := json.Unmarshal(resData, res); err != nil {
		return err
	}
	if res.Code != 0 {
		return fmt.Errorf("%d-%s", res.Code, res.Message)
	}
	return nil
}

func HttpSend(url string, body io.Reader, method string, headers map[string]string) ([]byte, error) {
	if len(method) == 0 {
		method = "GET"
	}
	req, err := http.NewRequest(method, url, body)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", "application/json")
	for k, v := range headers {
		fmt.Println(k, v)
		req.Header.Set(k, v)
	}
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		return nil, errors.New(resp.Status)
	}

	content, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	if len(content) == 0 {
		return nil, errors.New("nil resp")
	}
	return content, nil
}
