package ydhk

import (
	"LianFaPhone/lfp-marketing-api/common"
	"LianFaPhone/lfp-marketing-api/config"
	. "LianFaPhone/lfp-base/log/zap"

	//. "LianFaPhone/lfp-tools/autoorder-search-yidonghuaka/config"
	"encoding/json"
	"fmt"
	"net/url"
)

type (
	ReOrderSerach struct {
	}
	ResOrderSerach struct {
		Ret   int    `json:"code"`
		Msg string   `json:"message"`
		Datas  []*OrderInfo `json:"data"`
	}
	OrderInfo struct{
		Tid string  `json:"tid"`     //订单号
		OrderTime string `json:"orderTime"`  //下单时间
		GoodsTitle string `json:"goodsTitle"`           //:"移动花卡-宝藏版",
		GoodsProvince *string `json:"goodsProvince"` // :"371"
		GoodsCity  *string      `json:"goodsCity"`      //:"371",
		Mobilephone *string      `json:"mobilephone"`      //:"16637207824",
		ShipmentNo   *string      `json:"shipmentNo"`    //快递订单号:"JDVB02166598201",
		ShipmentCompany  *string    `json:"shipmentCompany"`  //:null,快递公司
		//"shipmentCompanyCode":"40",
		//"status":"AC",
		//"name":null,
		//"address":null,
		Number   *string       `json:"number"`    //:"17814678474",
		//"shipList":null
	}

	ReOrderDetailSerach struct {

	}
	ResOrderDetailSerach struct {
		Ret   int    `json:"code"`
		Msg string   `json:"message"`
		Data *OrderInfo `json:"data"`
	}
)
//idcard后6位
func (this *ReOrderSerach) Send(phone, idcard string) ([]*OrderInfo, error) {
	if len(idcard) > 9 {
		idcard = idcard[len(idcard)-6:]
	}

	formBody := make(url.Values)
	formBody.Add("mobilephone", phone)
	formBody.Add("certificateNo", idcard)

	heads := map[string]string{
		"Content-Type": "application/x-www-form-urlencoded; charset=UTF-8",
		"Accept": "*/*",
		//"Host": config.GConfig.Jthk.Host,
		"X-Requested-With":"XMLHttpRequest",
		"Origin": config.GConfig.Jthk.SearchUrl,
		"Referer": config.GConfig.Jthk.SearchUrl+"/rwkgzh/views/youthCard/order/orderList.jsp?mobilephone="+phone+"&certificateNo="+idcard,
	}

	ZapLog().Sugar().Infof("url[%s][%v][%v]", config.GConfig.Jthk.SearchUrl,heads, formBody )

	resData, err := common.HttpFormSend(config.GConfig.Jthk.SearchUrl+"/rwkgzh/youth/youthCard/query.tv", formBody,"POST", heads)
	if err != nil {
		return nil, err
	}
	res := new(ResOrderSerach)
	if err = json.Unmarshal(resData, res); err != nil {
		return  nil, err
	}
	if res.Ret !=200 {
		return nil, fmt.Errorf("%d-%s",res.Ret, res.Msg)
	}
	return res.Datas, nil
}

func (this *ReOrderDetailSerach) Send(phone, idcard, tid string) (*OrderInfo, error) {

	formBody := make(url.Values)
	formBody.Add("mobilephone", phone)
	formBody.Add("certificateNo", idcard)
	formBody.Add("tid", tid)

	heads := map[string]string{
		"Content-Type": "application/x-www-form-urlencoded; charset=UTF-8",
		"Accept": "*/*",
		//"Host": config.GConfig.Jthk.Host,
		"Origin": config.GConfig.Jthk.SearchUrl,
		// http://rwx.mmarket.com/rwkgzh/views/youthCard/order/detailNew.jsp?mobilephone=16637207824&certificateNo=220174&tid=SC19280T19121700402909&shipmentCompanyCode=40&shipmentNo=JDVB02166598201
		//"Referer": GConfig.Server.Url+"/rwkgzh/views/youthCard/order/orderList.jsp?mobilephone="+phone+"&certificateNo="+idcard,
	}



	resData, err := common.HttpFormSend(config.GConfig.Jthk.SearchUrl+"/rwkgzh/youth/youthCard/detail.tv", formBody,"POST", heads)
	if err != nil {
		return nil, err
	}
	res := new(ResOrderDetailSerach)
	if err = json.Unmarshal(resData, res); err != nil {
		return  nil, err
	}
	if res.Ret !=200 {
		return nil, fmt.Errorf("%d-%s",res.Ret, res.Msg)
	}
	return res.Data, nil
}
