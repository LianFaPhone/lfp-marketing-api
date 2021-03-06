package ydjthk

import (
	"LianFaPhone/lfp-marketing-api/common"
	"LianFaPhone/lfp-marketing-api/config"
	"bytes"
	"encoding/json"
	"fmt"
)

type(
	ReAddr struct{
		MsgType  string  `json:"MsgType"`
		Version string    `json:"Version"`
		CardProductId string   `json:"cardProductId"`
	}

	ResAddr struct{
		// {"hRet":"0","retMsg":"成功","allProvinceInfo":
		Ret   string    `json:"hRet"`
		Msg string   `json:"retMsg"`
		AllProvinceInfo []Provice `json:"allProvinceInfo"`
	}

 Provice struct {
ProvinceId string   `json:"provinceId"`
ProvinceName string     `json:"provinceName"`
CityList    []*City     `json:"cityList"`
CityMap     map[string] *City   `json:"-"`
}

 City struct{
AreaList   []*Area     `json:"area"`
CityId string      `json:"cityId"`
CityName string    `json:"cityName"`
AreaMap  map[string]*Area     `json:"-"`
}

 Area struct{
AreaName string       `json:"areaName"`
AreaId   string       `json:"areaId"`
}
)

func (this *ReAddr) Send(isOao bool) ([]Provice, error) {
	this.MsgType = "GetProvCityInfoReq"
	this.Version = Const_Version
	this.CardProductId = ""

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
		return nil, err
	}

	resData, err := common.HttpSend(config.GConfig.Jthk.Url+"/rwx/rwkweb/rwkCommon/getAllProvInfoTotal", bytes.NewReader(reqData),"POST", heads)
	if err != nil {
		return nil, err
	}
	res := new(ResAddr)
	if err = json.Unmarshal(resData, res); err != nil {
		return  nil, err
	}
	if res.Ret != "0" {
		return nil, fmt.Errorf("%s-%s",res.Ret, res.Msg)
	}

	return res.AllProvinceInfo, nil
}