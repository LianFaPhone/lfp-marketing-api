package baidu_api

import (
	common "LianFaPhone/lfp-marketing-api/common"
	"LianFaPhone/lfp-marketing-api/config"
	"bytes"
	"encoding/json"
	"fmt"
)

type ReDwzCreate struct{
	Url	string `json:"Url"`
	TermOfValidity	string `json:"TermOfValidity"`
}


type ResDwzCreate struct{
	Code int  `json:"Code"`
	ErrMsg  string `json:"ErrMsg"`
	ShortUrl string  `json:"ShortUrl"`
	LongUrl  string `json:"LongUrl"`
}

func (this * ReDwzCreate) Send() (string,error) {
	reBytes,err := json.Marshal(this)
	if err != nil {
		return "",err
	}
	head := map[string]string{"Content-Type":"application/json; charset=UTF-8", "Token":config.GConfig.Baidu.DwzToken}
	resBytes,err := common.HttpSend(config.GConfig.Baidu.DwzUrl+"/admin/v2/create", bytes.NewBuffer(reBytes),"POST", head)
	if err != nil {
		return "",err
	}
	res := new(ResDwzCreate)
	if err = json.Unmarshal(resBytes, res); err != nil {
		return "",err
	}

	if res.Code != 0 {
		return "", fmt.Errorf("%d-%s", res.Code, res.ErrMsg)
	}
	return res.ShortUrl,nil

}