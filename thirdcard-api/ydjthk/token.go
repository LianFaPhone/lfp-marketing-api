package ydjthk

import (
"LianFaPhone/lfp-marketing-api/common"
	"LianFaPhone/lfp-marketing-api/config"

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

func (this *ReToken) Send(isOao bool, channelId string) (string, error) {
	this.MsgType = "LivehkGetDataReq"
	this.Version = Const_Version
	this.ChannelId = channelId

	heads := map[string]string{
		"Accept": "application/json, text/plain, */*",
		"Host": config.GConfig.Jthk.Host,
		"Origin": config.GConfig.Jthk.Url,
		//"Referer": Const_Url+"/rwx/rwkvue/young/",
	}
	if isOao {
		heads["Referer"] = config.GConfig.Jthk.Url + config.GConfig.Jthk.Referer_path_oao
	}else{
		heads["Referer"] = config.GConfig.Jthk.Url + config.GConfig.Jthk.Referer_path
	}

	reqData,err := json.Marshal(this)
	if err != nil {
		return "", err
	}

	resData, err := common.HttpSend(config.GConfig.Jthk.Url+"/rwx/rwkweb/livehk/getLivehkData", bytes.NewReader(reqData),"POST", heads)
	if err != nil {
		return "", err
	}
	res := new(ResToken)
	if err = json.Unmarshal(resData, res); err != nil {
		return  "", err
	}
	if res.Ret != "0" {
		return "", fmt.Errorf("%s", res.Msg)
	}
	if len(res.AccessToken) <= 5 {
		return "", fmt.Errorf("get wrong token [%s]",res.AccessToken)
	}
	return res.AccessToken, nil
}

