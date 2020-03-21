package ydjthk

import (
"LianFaPhone/lfp-marketing-api/common"
	"LianFaPhone/lfp-marketing-api/config"
	"bytes"
"encoding/json"
"fmt"
)

type(
	ReProtocal struct{
		MsgType  string  `json:"MsgType"`
		Version string    `json:"Version"`
		ProviceCode string   `json:"provcode"`
	}

	ResProtocal struct{
		// {"hRet":"0","retMsg":"成功","allProvinceInfo":
		Ret   string    `json:"hRet"`
		Msg string   `json:"retMsg"`
		ResTest string `json:"restext"`
	}
)

func (this *ReProtocal) Send(isOao bool, ProviceCode string) (string, error) {
	this.MsgType = "LiveHKCardNaprotocolReq"
	this.Version = Const_Version
	this.ProviceCode = ProviceCode

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

	resData, err := common.HttpSend(config.GConfig.Jthk.Url+"/rwx/rwkweb/livehk/livehkProtocol/getCardNaprotocol", bytes.NewReader(reqData),"POST", heads)
	if err != nil {
		return "", err
	}
	res := new(ResProtocal)
	if err = json.Unmarshal(resData, res); err != nil {
		return  "", err
	}
	if !(res.Ret == "0" || res.Ret == "000") {
		return "", fmt.Errorf("%s-%s",res.Ret, res.Msg)
	}

	return res.ResTest, nil
}
