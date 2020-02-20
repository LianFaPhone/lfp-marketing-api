package ydhk

import (
	"LianFaPhone/lfp-marketing-api/common"
	"bytes"
	"encoding/json"
	"fmt"
)

type(
	ReIdCheckUrl struct{
		MsgType  string  `json:"MsgType"`
		Version string    `json:"Version"`
		OrderId string     `json:"orderId"`
		ChannelId string  `json:"channelId"`
		NewPhone string  `json:"mobileId"`
		Token   string   `json:"accessToken"`
		ModuleName string   `json:"moduleName"`

	//{"MsgType":"OnlineCertPageParamReq","Version":"1.0.0","orderId":"44200220221249110034","channelId":"C10000032277","mobileId":"19813737336","accessToken":"LIVEHK_ACCESSTOKEN_C10000032277_31A90E24EE4BE12D76111D514F552516_26","moduleName":"LIVEHK"}
	}

	ResIdCheckUrl struct{
		// {"hRet":"0","retMsg":"成功","allProvinceInfo":
		Ret   string    `json:"hRet"`
		Msg string   `json:"retMsg"`
		Url string  `json:"urlparams"`
	//{
	//	"hRet": "0",
	//	"retMsg": "成功",
	//	"urlparams": "https://activate.online-cmcc.cn/edcreg/saleCardMiniAppointment/signTheCheckAPP?requestSource=200993&transactionID=20099320200220221710385066&signature=j3jqZy9Lw2ztL1HfHy4NJEcf33HL7EwDdYTzP9tKnWl9IvztP7u2tUGl0Qt18Jt49MJJEqibrGKS0IxeEZkj75w1NBzahaOmGpikfCY91NSqRZA6nXccm5gKxMn%2BbTewwI8ZKhuGfmbDCjzLmsfyK0n0sz%2B0C8kIi6Smvz2JDuIzD1jLiNb7lSFxR3yTWMp1GhDEqWn4sbVOOyyGlbzdoKGZDjGDFVO9LRsn4HFzA9mCgr%2FXOouiYkZFEQ33AsHp4Qqt%2BG4b6uIxBfkg%2BuqByADVag8%2B6zUT%2FX81WWw5kpBgblAczhAwdb9MLEZup16pd3BZ4YtYIo9jTtxjm4ot3w%3D%3D&billId=93a1d850c354c91e2a8fedbe5f9b3b75&channelId=48ea9573ce2d9d57296ab192303ca869&busiCode=PREORDAIN_CHECK",
	//	"MsgType": "OnlineCertPageParamResp",
	//	"Version": "1.0.0"
	//}
	}


)

func (this *ReIdCheckUrl) Send(orderId, newPhone, token string) (string, error) {
	this.MsgType = "OnlineCertPageParamReq"
	this.Version = Const_Version
	this.ChannelId = Const_ChannelId
	this.ModuleName = "LIVEHK"
	this.OrderId = orderId
	this.NewPhone = newPhone
	this.Token = token

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

	resData, err := common.HttpSend(Const_Url+"/rwx/rwkweb/rwkCommon/queryIdCheckUrl", bytes.NewReader(reqData),"POST", heads)
	if err != nil {
		return "", err
	}
	res := new(ResIdCheckUrl)
	if err = json.Unmarshal(resData, res); err != nil {
		return  "", err
	}
	if res.Ret != "0" {
		return "", fmt.Errorf("%s",res.Msg)
	}

	return res.Url, nil
}
