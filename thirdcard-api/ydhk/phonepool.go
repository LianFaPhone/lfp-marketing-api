package ydhk

import (
	"LianFaPhone/lfp-marketing-api/common"

	//. "LianFaPhone/lfp-tools/autoorder-yidonghuaka/config"
	"bytes"
	"encoding/json"
	"fmt"
)

type (
	ReCardSearch struct {
		MsgType  string  `json:"MsgType"`
		Version string    `json:"Version"`

		ProvCode string   `json:"provCode"`
		Province string   `json:"province"`
		CityCode string   `json:"cityCode"`
		City     string   `json:"city"`

		Selecttype int   `json:"selecttype"`
		Searchkey string   `json:"searchkey"`
		Count int   `json:"count"`
	}

	ResCardSearch struct {
		Ret   string    `json:"hRet"`
		Msg string   `json:"retMsg"`
		Numbers    []string `json:"numbers"`
	}

	ReCloseNumber struct {
		MsgType  string  `json:"MsgType"`
		Version string    `json:"Version"`

		ProvCode string   `json:"provCode"`

		CityCode string   `json:"cityCode"`

		Number     string   `json:"number"`
		AccessToken     string   `json:"accessToken"`
	}

	ResCloseNumber struct {
		Ret   string    `json:"hRet"`
		Msg string   `json:"retMsg"`
		UnLockTime    string `json:"unlockTime"`
	}
)

func (this *ReCardSearch) Send(proCode, proName, cityCode, cityName, Searchkey string,page,size int) ([]string, error) {
	this.MsgType = "LiveHKSelectNumberReq"
	this.Version = Const_Version
	this.Selecttype = page
	this.Count = size
	this.Province = proName
	this.ProvCode = proCode
	this.City = cityName
	this.CityCode = cityCode
	this.Searchkey = Searchkey

	heads := map[string]string{
		"Accept": "application/json, text/plain, */*",
		"Host": Const_Host,
		"Origin": Const_Url,
		"Referer": Const_Url+"/rwx/rwkvue/young/",
	}

	reqData,err := json.Marshal(this)
	if err != nil {
		return nil, err
	}

	resData, err := common.HttpSend(Const_Url+"/rwx/rwkweb/livehk/livehkMobile/selectNumber", bytes.NewReader(reqData),"POST", heads)
	if err != nil {
		return nil, err
	}
	res := new(ResCardSearch)
	if err = json.Unmarshal(resData, res); err != nil {
		return nil, err
	}
	if res.Ret != "0" {
		return nil, fmt.Errorf("%v-%v", res.Ret, res.Msg)
	}
	if len(res.Numbers) == 0 {
		return nil, fmt.Errorf("nil resp")
	}
	return res.Numbers, nil
}

func (this *ReCloseNumber) Send(proCode, cityCode, number, token string) (bool, error) {
	this.MsgType = "LiveHKLockNumberReq"
	this.Version = Const_Version
	this.ProvCode = proCode
	this.CityCode = cityCode
	this.Number  = number
	this.AccessToken = token

	heads := map[string]string{
		"Accept": "application/json, text/plain, */*",
		"Host": Const_Host,
		"Origin": Const_Url,
		"Referer": Const_Url+"/rwx/rwkvue/young/",
	}

	reqData,err := json.Marshal(this)
	if err != nil {
		return false, err
	}

	resData, err := common.HttpSend(Const_Url+"/rwx/rwkweb/livehk/livehkMobile/lockNumber", bytes.NewReader(reqData),"POST", heads)
	if err != nil {
		return false, err
	}
	res := new(ResCloseNumber)
	if err = json.Unmarshal(resData, res); err != nil {
		return false, err
	}
	if res.Ret != "0" {
		return false, fmt.Errorf("%v-%v", res.Ret, res.Msg)
	}

	return true, nil
}
