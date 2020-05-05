package ydjthk

import (
	"LianFaPhone/lfp-marketing-api/common"
	"LianFaPhone/lfp-marketing-api/config"
	"net/http"
	"time"

	//. "LianFaPhone/lfp-tools/autoorder-search-yidonghuaka/config"
	"encoding/json"
	"fmt"
	"net/url"
)

type (
	ReYgOrderSerach struct {
	}
	ResYgOrderSerach struct {
		Total   int    `json:"total"`
		Page    int   `json:"page"`
		Records    int   `json:"records"`
		Success  bool    `json:"success"`
		Message string  `json:"message"`
		Datas  []*YgOrderInfo `json:"rows"`
	}
	YgOrderInfo struct{
		OrderId *string  `json:"orderid"`     //订单号
		OrderTime *string `json:"actiontime"`  //下单时间
		GoodsTitle *string `json:"goodstitle"`           //:"移动花卡-宝藏版",
		Mobilephone *string      `json:"mobilephone"`      //:"16637207824",
		ShipmentNo   *string      `json:"shipmentno"`    //快递订单号:"JDVB02166598201",
		ShipmentCompany  *string    `json:"shipmentcompany"`  //:null,快递公司
		ShipmentCompanyCode *string `json:"shipmentcompanycode"`
		//"status":"AC",
		Status *string `json:"status"`
		//"name":null,
		//"address":null,
		OaoModel *int   `json:"oaomodel"`
		OaoModel2 *string   `json:"oaomodel2"`
		Status2 *string `json:"status2"`
		ActiveTime *string `json:"activtime"`  //下单时间
		//"shipList":null
	}


)
//idcard后6位
func (this *ReYgOrderSerach) Send(thridOrderNo string, startData, endDate string) (*ResYgOrderSerach, error) {


	formBody := make(url.Values)
	formBody.Add("params['orderid']", thridOrderNo)
	formBody.Add("params['startTime']", startData)
	formBody.Add("params['endTime']", endDate)
	config.GConfig.Lock.Lock()
	ticket := config.GConfig.Jthk.CkTicket
	config.GConfig.Lock.Unlock()


	ckArr := make([]*http.Cookie, 0, 3)
	cookie1 := &http.Cookie{Name: "AUTH_TICKET",Value: ticket, HttpOnly: true, Expires: time.Now().Add(time.Hour*24*365*100)}
	cookie2 := &http.Cookie{Name: "JSESSIONID",Value: ticket, HttpOnly: true, Expires: time.Now().Add(time.Hour*24*365*100)}
	cookie3 := &http.Cookie{Name: "admin-domain",Value: "%2Fadmin%2F", HttpOnly: true, Expires: time.Now().Add(time.Hour*24*365*100)}
	ckArr=append(ckArr, cookie1)
	ckArr=append(ckArr, cookie2)
	ckArr=append(ckArr, cookie3)

	//heads := map[string]string{
	////	"Content-Type": "application/x-www-form-urlencoded; charset=UTF-8",
	////	"Accept": "*/*",
	////	//"Host": config.GConfig.Jthk.Host,
	////	"X-Requested-With":"XMLHttpRequest",
	////	"Origin": config.GConfig.Jthk.SearchUrl,
	////	"Referer": config.GConfig.Jthk.SearchUrl+"/rwkgzh/views/youthCard/order/orderList.jsp?mobilephone="+phone+"&certificateNo="+idcard,
	//}

	//ZapLog().Sugar().Infof("url[%s][%v][%v]", config.GConfig.Jthk.SearchUrl,heads, formBody )

	url:="https://yg.cmicrwx.cn/opesp-portal/fcyrCorderTra/fcyrCorderTraReport.ajax?"

	resData, err := common.HttpSend3(url+formBody.Encode(), nil,"GET", nil, ckArr)
	if err != nil {
		return nil, err
	}
	if resData == nil || len(resData) == 0 {
		return nil, fmt.Errorf("nil msg")
	}
	res := new(ResYgOrderSerach)
	if err = json.Unmarshal(resData, res); err != nil {

		return  nil, err
	}
	//if res.Ret !=200 {
	//	return nil, fmt.Errorf("%d-%s",res.Ret, res.Msg)
	//}
	return res, nil
}

const(
	Yg_Status_Already_Activated = "8"  //已激活
	Yg_Status_Deliver = "5"  //已发货
	Yg_Status_New     = "4"  //备货
	Yg_Status_Init    = "0"  //初始状态
)
