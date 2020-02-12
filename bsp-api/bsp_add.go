package bsp_api

import (
	"LianFaPhone/lfp-marketing-api/models"
	"time"
)

type (
	ReBspAdd struct {
		Request
		Content struct {
			ServiceOrderOutOrderId           string `json:"content,omitempty"`
			ServiceOrderSubmitTime           string `json:"serviceOrderSubmitTime,omitempty"`
			ServiceOrderCusName              string `json:"serviceOrderCusName,omitempty"`
			ServiceOrderCusCardNo            string `json:"serviceOrderCusCardNo,omitempty"`
			ServiceOrderCusAccPhone          string `json:"serviceOrderCusAccPhone,omitempty"`
			ServiceOrderCusContactPhone      string `json:"serviceOrderCusContactPhone,omitempty"`
			ServiceOrderIccid                string `json:"serviceOrderIccid,omitempty"`
			ServiceOrderPhoneCellcore        string `json:"serviceOrderPhoneCellcore,omitempty"`
			ServiceOrderReceiverName         string `json:"serviceOrderReceiverName,omitempty"`
			ServiceOrderReceiverProvCode     string `json:"serviceOrderReceiverProvCode,omitempty"`
			ServiceOrderReceiverCityCode     string `json:"serviceOrderReceiverCityCode,omitempty"`
			ServiceOrderReceiverDistrictCode string `json:"serviceOrderReceiverDistrictCode,omitempty"`
			ServiceOrderReceiverAddress      string `json:"serviceOrderReceiverAddress,omitempty"`
			ServiceOrderType                 string `json:"serviceOrderType,omitempty"`
			ServiceOrderSource               string `json:"serviceOrderSource,omitempty"`
			ServiceOrderCpsRefereePeople     string `json:"serviceOrderCpsRefereePeople,omitempty"`
			ServiceOrderRefereePeople        string `json:"serviceOrderRefereePeople,omitempty"`
			ServiceOrderPayMethod            string `json:"serviceOrderPayMethod,omitempty"`
			ServiceOrderPayStatus            string `json:"serviceOrderPayStatus,omitempty"`
			ServiceOrderPayTranid            string `json:"serviceOrderPayTranid,omitempty"`
			ServiceOrderPayOrderid           string `json:"serviceOrderPayOrderid,omitempty"`
			ServiceOrderCardPic1             string `json:"serviceOrderCardPic1,omitempty"`
			ServiceOrderActivationReferee    string `json:"serviceOrderActivationReferee,omitempty"`
		} `json:"content,omitempty"`
	}
)

func (this *ReBspAdd) Add(p *models.CardOrder) error {
	this.Method = "syn.orderinfo.SynJdServiceOrder"
	this.AccessToken = ""
	this.Version = "1.0"
	this.Content.ServiceOrderOutOrderId = p.OrderNo
	this.Content.ServiceOrderSubmitTime = time.Now().In(time.FixedZone("UTC", 8*3600)).Format("2006-01-02 15:04:05")
	this.Content.ServiceOrderCusName = *p.TrueName
	this.Content.ServiceOrderCusCardNo = *p.IdCard

	return nil
}
