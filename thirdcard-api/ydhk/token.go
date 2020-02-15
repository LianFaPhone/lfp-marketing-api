package ydhk

import (
"LianFaPhone/lfp-common"

"bytes"
"encoding/json"
"fmt"
)

type (
	ReToken struct {
		MsgType  string  `json:"MsgType"`
		Version string    `json:"Version"`
		ChannelId string   `json:"channelId"`
	}
	ResToken struct {
		Ret   string    `json:"hRet"`
		Msg string   `json:"retMsg"`
		AccessToken string `json:"accessToken"`
	}
)

func (this *ReToken) Send() (string, error) {
	this.MsgType = "LivehkGetDataReq"
	this.Version = Const_Version
	this.ChannelId = Const_ChannelId

	heads := map[string]string{
		"Accept": "application/json, text/plain, */*",
		"Host": Const_Host,
		"Origin": Const_Url,
		"Referer": Const_Url+"/rwx/rwkvue/young/",
	}

	reqData,err := json.Marshal(this)
	if err != nil {
		return "", err
	}

	resData, err := common.HttpSend(Const_Url+"/rwx/rwkweb/livehk/getLivehkData", bytes.NewReader(reqData),"POST", heads)
	if err != nil {
		return "", err
	}
	res := new(ResToken)
	if err = json.Unmarshal(resData, res); err != nil {
		return  "", err
	}
	if res.Ret != "0" {
		return "", fmt.Errorf("%s-%s",res.Ret, res.Msg)
	}
	if len(res.AccessToken) <= 5 {
		return "", fmt.Errorf("get wrong token [%s]",res.AccessToken)
	}
	return res.AccessToken, nil
}

