package idcard_api

import (
	"LianFaPhone/lfp-marketing-api/common"
	"net/url"
	"strings"
	"encoding/json"
	"fmt"
	"LianFaPhone/lfp-marketing-api/config"
)

type ReqIdCardCheck struct{

}

type ResIdCardCheck struct{
	Name string        `json:"name"`
	IdNo string        `json:"idNo"`
	RespMessage string `json:"respMessage"`
	RespCode    string `json:"respCode"`
	Province   string  `json:"province"`
	City    string     `json:"city"`
	County string      `json:"county"`
	Birthday string    `json:"birthday"`
	Sex string         `json:"sex"`
	Age string         `json:"age"`
}

func (this * ReqIdCardCheck) Check(idNo, name string) (bool, string, error) {

	res, err := this.Send(idNo, name)
	if err != nil {
		return false, "", err
	}

	if res.RespCode == "0000" {
		return true, "",nil
	}
	if res.RespCode == "0008" || res.RespCode == "0007" || res.RespCode == "0004" {
		return false, "",nil
	}
	return false, res.RespCode, fmt.Errorf("%s-%s", res.RespCode, res.RespMessage)
}

func (this * ReqIdCardCheck) Send(idNo, name string) (*ResIdCardCheck,error) {
	data := make(url.Values)
	data["idNo"] = []string{idNo}
	data["name"] = []string{name}

	body := strings.NewReader(data.Encode())

	resData,err := common.HttpSend("https://idenauthen.market.alicloudapi.com/idenAuthentication", body, "POST", map[string]string{"Authorization":"APPCODE "+config.GConfig.IdCardCheck.AppCode, "Content-Type":"application/x-www-form-urlencoded"})

	if err != nil {
		return nil,err
	}
	res := new(ResIdCardCheck)
	if err := json.Unmarshal(resData, res); err != nil {
		return nil,err
	}
	return res,nil
}
