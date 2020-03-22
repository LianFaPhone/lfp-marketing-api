package tasker

import (
	. "LianFaPhone/lfp-base/log/zap"
	"LianFaPhone/lfp-marketing-api/config"
	"LianFaPhone/lfp-marketing-api/models"
	"LianFaPhone/lfp-marketing-api/sdk"
	"LianFaPhone/lfp-marketing-api/thirdcard-api/ydjthk"
	"fmt"
	"go.uber.org/zap"
	"net/url"
	"strconv"
	"strings"
	"time"
)

//目前只操作集团花卡
func (this *Tasker) jtyhhkThirdRetryWork() {
	defer models.PanicPrint()

	//recoder, err := new(models.IdRecorder).GetByName(idRecorderName)
	//if err != nil {
	//	ZapLog().Error("IdRecorder GetByName err", zap.Error(err))
	//	return
	//}
	//
	//if recoder == nil {
	//	if recoder, err = new(models.IdRecorder).Add(idRecorderName, 0); err != nil {
	//		ZapLog().Error("IdRecorder Add  err", zap.Error(err))
	//		return
	//	}
	//}
	//
	//startId := int64(0)
	//if recoder.IdTag != nil {
	//	startId = *recoder.IdTag
	//}

	if len(config.GConfig.Jthk.ParterCode) <= 0 || len(config.GConfig.Jthk.ParterCodeArr) <= 0 {
		return
	}

	partnerIds := GetJtHkPartnerIds()
	if len(partnerIds) <= 0 {
		return
	}

	ProvinceMap := make(map[string]*ydjthk.Provice)
	provinceArr, err := new(ydjthk.ReAddr).Send(false)
	if err != nil {
		ZapLog().Error("Addr send err", zap.Error(err))
		return
	}

	for i:=0; i < len(provinceArr); i++ {
		p := provinceArr[i]
		ProvinceMap[p.ProvinceName] = &p
		if p.CityMap == nil {
			p.CityMap = make(map[string]*ydjthk.City)
		}
		for j:=0; j< len(p.CityList); j++{
			p.CityMap[p.CityList[j].CityName] = p.CityList[j]
			p.CityList[j].AreaMap = make(map[string]*ydjthk.Area)
			for m:=0; m< len(p.CityList[j].AreaList); m++{
				p.CityList[j].AreaMap[p.CityList[j].AreaList[m].AreaName] = p.CityList[j].AreaList[m]
			}
		}
	}
	time.Sleep(time.Millisecond*2000)
	for q:=0;q<500;q++ {
		conds := []*models.SqlPairCondition{
			//&models.SqlPairCondition{"id > ?", startId},
			//&models.SqlPairCondition{"class_big_tp = ?", 5},
			//&models.SqlPairCondition{"created_at <= ?", time.Now().Unix() - delayTime},
			&models.SqlPairCondition{"status in (?)", []int{models.CONST_OrderStatus_Retry_Apply_Doing, models.CONST_OrderStatus_New_Apply_Doing}},
		}

		if len(partnerIds) > 0 {
			conds = append(conds, &models.SqlPairCondition{"partner_id in (?)", partnerIds})
		}

		orderArr, err := new(models.CardOrder).GetLimitByCond(10, conds, nil)
		if err != nil {
			ZapLog().Error("CardOrder GetLimitByCond err", zap.Error(err))
			return
		}
		if orderArr == nil || len(orderArr) <= 0 {
			//fmt.Println("nil data")
			return
		}

		//记录id, 倒叙
		for i := len(orderArr) - 1; i >= 0; i-- {
			//if *orderArr[i].Id > startId {
			//	startId = *orderArr[i].Id
			//}
			if orderArr[i] == nil || orderArr[i].Status == nil{
				continue
			}

			mp := &models.CardOrder{
				Id: orderArr[i].Id,
				Status: new(int),
			}
			if orderArr[i].PartnerGoodsCode == nil || orderArr[i].Province == nil || orderArr[i].City == nil || orderArr[i].Area == nil || orderArr[i].Address == nil || orderArr[i].Phone == nil || orderArr[i].TrueName == nil || orderArr[i].IdCard == nil{
				*mp.Status = GenFailStatus(orderArr[i], "可重试")
				mp.Update()
				log := "管理员|重试下单失败|表中数据缺失"
				new(models.CardOrderLog).FtParseAdd(orderArr[i].Id, orderArr[i].OrderNo, &log).Add()
				continue
			}

			queryValues,err := GetUrlParam(*orderArr[i].PartnerGoodsCode)
			if err != nil {
				*mp.Status = GenFailStatus(orderArr[i], "可重试")
				mp.Update()
				log := "管理员|重试下单失败|额外参数缺失"
				new(models.CardOrderLog).FtParseAdd(orderArr[i].Id, orderArr[i].OrderNo, &log).Add()
				continue
			}
			channelId := queryValues.Get("channelId")
			productId := queryValues.Get("productId")
			isOao,_ := strconv.ParseBool(queryValues.Get("isOao"))

			if len(channelId) <= 0 || len(productId) <= 0 {
				*mp.Status = GenFailStatus(orderArr[i], "可重试")
				mp.Update()
				log := "管理员|重试下单失败|额外参数缺失"
				new(models.CardOrderLog).FtParseAdd(orderArr[i].Id, orderArr[i].OrderNo, &log).Add()
				continue
			}
			time.Sleep(time.Millisecond*100)
			token,err := new(ydjthk.ReToken).Send(isOao, channelId)
			if err != nil {
				*mp.Status = GenFailStatus(orderArr[i], "可重试")
				mp.Update()
				log := "管理员|重试下单失败|请求token失败，"+err.Error()
				new(models.CardOrderLog).FtParseAdd(orderArr[i].Id, orderArr[i].OrderNo, &log).Add()
				continue
			}
			province,ok := ProvinceMap[*orderArr[i].Province]
			if !ok {
				*mp.Status = GenFailStatus(orderArr[i], "可重试")
				mp.Update()
				log := "管理员|重试下单失败|省份匹配不上"
				new(models.CardOrderLog).FtParseAdd(orderArr[i].Id, orderArr[i].OrderNo, &log).Add()
				continue
			}
			city, ok := province.CityMap[*orderArr[i].City]
			if !ok {
				*mp.Status = GenFailStatus(orderArr[i], "可重试")
				mp.Update()
				log := "管理员|重试下单失败|城市匹配不上"
				new(models.CardOrderLog).FtParseAdd(orderArr[i].Id, orderArr[i].OrderNo, &log).Add()
				continue
			}
			area,ok := city.AreaMap[*orderArr[i].Area]
			if !ok {
				*mp.Status = GenFailStatus(orderArr[i], "可重试")
				mp.Update()
				log := "管理员|重试下单失败|区（县）匹配不上"
				new(models.CardOrderLog).FtParseAdd(orderArr[i].Id, orderArr[i].OrderNo, &log).Add()
				continue
			}
			time.Sleep(time.Millisecond*100)
			numbers,err := new(ydjthk.ReCardSearch).Send(isOao, province.ProvinceId, province.ProvinceName, city.CityId, city.CityName, "", 1, 10)
			if err != nil {
				*mp.Status = GenFailStatus(orderArr[i], "可重试")
				mp.Update()
				log := "管理员|重试下单失败|获取新号码失败，"+err.Error()
				new(models.CardOrderLog).FtParseAdd(orderArr[i].Id, orderArr[i].OrderNo, &log).Add()
				continue
			}

			chooseNumber := ""
			for j:=0; j < len(numbers);j++ {
				flag,_,err := new(ydjthk.ReCloseNumber).Send(isOao, province.ProvinceId,  city.CityId, numbers[j], token)
				if err != nil {
					continue
				}
				if flag {
					chooseNumber = numbers[j]
					break
				}
			}

			if len(chooseNumber) <= 0 {
				*mp.Status = GenFailStatus(orderArr[i], "可重试")
				mp.Update()
				log := "管理员|重试下单失败|无法锁定新号码"
				new(models.CardOrderLog).FtParseAdd(orderArr[i].Id, orderArr[i].OrderNo, &log).Add()
				continue
			}
			time.Sleep(time.Millisecond*1000)
			_, thirdOrderNo,oaoFlag,orderErr := new(ydjthk.ReOrderSubmit).Parse(channelId, productId, nil).Send(isOao, token,  *orderArr[i].Phone, chooseNumber, *orderArr[i].TrueName, *orderArr[i].IdCard, *orderArr[i].Address, *orderArr[i].Province, *orderArr[i].City, province.ProvinceId, city.CityId, area.AreaId)
			if orderErr != nil {
				*mp.Status = GenFailStatus(orderArr[i], orderErr.Error())
				mp.Update()
				log := "管理员|重试下单失败|下单失败，"+orderErr.Error()
				new(models.CardOrderLog).FtParseAdd(orderArr[i].Id, orderArr[i].OrderNo, &log).Add()
				continue
			}

			log := "管理员|重试下单成功|新订单已完成"
			if !oaoFlag {
				*mp.Status = models.CONST_OrderStatus_New_UnFinish

				newUrl,err := new(ydjthk.ReIdCheckUrl).Send(isOao, channelId, thirdOrderNo, chooseNumber, token)
				if err != nil {
					ZapLog().Error("ReIdCheckUrl send err", zap.Error(err))
					log = "管理员|重试下单失败|获取上传照片网址失败，"+err.Error()
					*mp.Status = GenFailStatus(orderArr[i], "可重试")
				}else{
					//sendUnFinishNotify(newUrl, orderArr[i])
					log = "管理员|重试下的那成功|新订单未完成，等待上传照片"
					new(models.CardOrderUrl).FtParseAdd(orderArr[i].Id, orderArr[i].OrderNo, &newUrl).Add()
				}
			}else{
				*mp.Status = models.CONST_OrderStatus_New
			}
			new(models.CardOrderLog).FtParseAdd(orderArr[i].Id, orderArr[i].OrderNo, &log).Add()

			mp.NewPhone = &chooseNumber
			mp.ThirdOrderNo = &thirdOrderNo
			mp.ThirdOrderAt = new(int64)
			*mp.ThirdOrderAt = time.Now().Unix() -2*60
			if err = mp.Update(); err != nil {
				ZapLog().Error("CardOrder Update err", zap.Error(err))
				return
			}
			mp.MaxIdByOrderNo(*orderArr[i].OrderNo) // 这一步是关键，让它重新被检测
			time.Sleep(time.Second * 1)
		}

		if len(orderArr) < 10 {
			break
		}
		time.Sleep(time.Second * 5)
	}

}

func GenFailStatus(temp *models.CardOrder, errMsg string) int {
	if temp.Status == nil {
		return models.CONST_OrderStatus_Fail
	}

	if *temp.Status == models.CONST_OrderStatus_New_Apply_Doing {
		return ParseFailStatus(errMsg)
	}
	if *temp.Status == models.CONST_OrderStatus_Retry_Apply_Doing {
		return models.CONST_OrderStatus_Fail_Already_Retry
	}
	return models.CONST_OrderStatus_Fail
}

func  ParseFailStatus(errMsg string) int {
	if len(errMsg) <= 0 {
		return models.CONST_OrderStatus_Fail
	}
	if strings.Contains(errMsg, "可重试") {
		return models.CONST_OrderStatus_Fail_Retry
	}else if strings.Contains(errMsg, "欠费号码") {
		return models.CONST_OrderStatus_Fail_Retry
	}else if strings.Contains(errMsg, "已超时") {
		return models.CONST_OrderStatus_Fail_Retry
	}else if strings.Contains(errMsg, "系统错误") {
		return models.CONST_OrderStatus_Fail_Retry
	}else if strings.Contains(errMsg, "订购的号码不存在") {
		return models.CONST_OrderStatus_Fail_Retry
	}else if strings.Contains(errMsg, "号码已被占用") {
		return models.CONST_OrderStatus_Fail_Retry
	}else if strings.Contains(errMsg, "系统忙") {
		return models.CONST_OrderStatus_Fail_Retry
	}else if strings.Contains(errMsg, "无法申请新号卡") {
		return models.CONST_OrderStatus_Fail_Retry
	}else if strings.Contains(errMsg, "锁定号码失败") {
		return models.CONST_OrderStatus_Fail_Retry
	}
	return models.CONST_OrderStatus_Fail
}

func sendUnFinishNotify(newUrl string, temp *models.CardOrder) {
	if temp.Phone == nil {
		return
	}
	if len(newUrl) < 3 {
		log:= "短信发送：照片上传提醒，未找到Url信息"
		new(models.CardOrderLog).FtParseAdd(nil, temp.OrderNo, &log).Add()
		return
	}
	if !strings.HasPrefix(newUrl, " ") {
		newUrl = " "+newUrl+" " //可生成短链
	}
	newUrl = strings.Replace(newUrl, "%", "%25", -1)

	log:= "短信发送：照片上传提醒，发送成功"
	platformTp := 2
	if err := sdk.GNotifySdk.SendSms([]string{newUrl}, *temp.Phone, "yd_jthk_new_unfinish",0, &platformTp); err != nil {
		ZapLog().Error("GNotifySdk.SendSms err", zap.Error(err))
		log= "短信发送：照片上传提醒，发送失败;"+err.Error()
	}
	new(models.CardOrderLog).FtParseAdd(nil, temp.OrderNo, &log).Add()
}

func GetUrlParam(partnerGoodsCode string) (url.Values, error) {
	cc, err := new(models.PdPartnerGoods).GetByCodeFromCache(partnerGoodsCode)
	if err != nil {
		return nil,err
	}
	if cc == nil || (cc.Valid!=nil && *cc.Valid == 0) {
		return nil,fmt.Errorf("nofind parterGoods")
	}

	partner ,err := new(models.PdPartner).GetByIdFromCache(*cc.PartnerId)
	if err != nil {
		return nil,err
	}
	if partner == nil || (partner.Valid!=nil && *partner.Valid == 0) {
		return nil,fmt.Errorf("nofind parter")
	}

	UrlParam := new(string)
	if cc.UrlParam != nil && len(*cc.UrlParam) > 0{
		*UrlParam = *cc.UrlParam
		if partner.UrlParam != nil && len(*partner.UrlParam) > 0 {
			*UrlParam = *partner.UrlParam + "&"+*UrlParam
		}
	}else  {
		UrlParam = partner.UrlParam
	}

	if UrlParam == nil {
		return nil,fmt.Errorf("nofind urlParam")
	}

	vv,err := url.ParseQuery(*UrlParam)
	if err != nil {
		return nil,err
	}
	return vv,nil
}

//func (this *Tasker) ydjthkRetryOaoWork(delayTime int64, SetFailFlag bool) {
//	defer models.PanicPrint()
//
//	if len(config.GConfig.Jthk.ParterCode) <= 0 || len(config.GConfig.Jthk.ParterCodeArr) <= 0 {
//		return
//	}
//
//	partnerIds := GetJtHkPartnerIds()
//	if len(partnerIds) <= 0 {
//		return
//	}
//
//	for true {
//		conds := []*models.SqlPairCondition{
//			//&models.SqlPairCondition{"id > ?", startId},
//			&models.SqlPairCondition{"status = ?", models.CONST_OrderStatus_Retry_UnFinish},
//			&models.SqlPairCondition{"third_order_at <= ?", time.Now().Unix() - delayTime},
//			//条件还得处理下
//		}
//
//		if len(partnerIds) > 0 {
//			conds = append(conds, &models.SqlPairCondition{"partner_id in (?)", partnerIds})
//		}
//
//		orderArr, err := new(models.CardOrder).GetLimitByCond(10, conds, nil)
//		if err != nil {
//			ZapLog().Error("CardOrder GetLimitByCond err", zap.Error(err))
//			return
//		}
//		if orderArr == nil || len(orderArr) <= 0 {
//			//fmt.Println("nil data")
//			return
//		}
//
//		//ZapLog().Sugar().Infof("jthktasker %d", len(orderArr))
//
//		//记录id, 倒叙
//		for i := len(orderArr) - 1; i >= 0; i-- {
//			if orderArr[i] == nil || orderArr[i].Status == nil{
//				continue
//			}
//
//			mp := &models.CardOrder{
//				Id: orderArr[i].Id,
//				Status: new(int),
//			}
//
//			if orderArr[i].Phone == nil || orderArr[i].NewPhone == nil || orderArr[i].IdCard == nil {
//				log:= "OAO检测：信息不全"
//				new(models.CardOrderLog).FtParseAdd(nil, orderArr[i].OrderNo, &log).Add()
//				*mp.Status = models.CONST_OrderStatus_Fail_Already_Retry
//				if err = mp.Update(); err != nil {
//					ZapLog().Error("CardOrder Update err", zap.Error(err))
//				}
//				continue
//			}
//
//
//			resOrderShortSerach,err := new(ydjthk.ReOrderShortSerach).Send(*orderArr[i].Phone, *orderArr[i].IdCard);
//			if err != nil {
//				log:= "OAO检测：网络错误，"+err.Error()
//				new(models.CardOrderLog).FtParseAdd(nil, orderArr[i].OrderNo, &log).Add()
//				continue
//			}
//			if resOrderShortSerach.Ret != 200 {
//				log:= fmt.Sprintf("OAO检测：%d-%s", resOrderShortSerach.Ret, resOrderShortSerach.Msg)
//				new(models.CardOrderLog).FtParseAdd(nil, orderArr[i].OrderNo, &log).Add()
//				if !SetFailFlag {
//					continue
//				}
//				*mp.Status = models.CONST_OrderStatus_Fail_Already_Retry
//				if err = mp.Update(); err != nil {
//					ZapLog().Error("CardOrder Update err", zap.Error(err))
//				}
//				continue
//			}
//
//			resOrderSearch,err := new(ydjthk.ReOrderSerach).Send(*orderArr[i].Phone, *orderArr[i].IdCard)
//			if err != nil {
//				log:= "OAO检测：网络错误，"+err.Error()
//				new(models.CardOrderLog).FtParseAdd(nil, orderArr[i].OrderNo, &log).Add()
//				continue
//			}
//			if resOrderSearch.Ret != 200 {
//				log:= fmt.Sprintf("OAO检测：%d-%s", resOrderSearch.Ret, resOrderSearch.Msg)
//				new(models.CardOrderLog).FtParseAdd(nil, orderArr[i].OrderNo, &log).Add()
//				if !SetFailFlag {
//					continue
//				}
//				*mp.Status = models.CONST_OrderStatus_Fail_Already_Retry
//				if err = mp.Update(); err != nil {
//					ZapLog().Error("CardOrder Update err", zap.Error(err))
//				}
//				continue
//			}
//
//			var chooseOne *ydjthk.OrderInfo
//			yidongArr := resOrderSearch.Datas
//			for j:=0; j< len(yidongArr); j++ {
//				if yidongArr[j].Number == nil || orderArr[i].NewPhone == nil{
//					continue
//				}
//				if * yidongArr[j].Number == *orderArr[i].NewPhone {
//					chooseOne = yidongArr[j]
//					break
//				}
//			}
//
//			if chooseOne == nil {
//				log:= "OAO检测：oao未发现"
//				new(models.CardOrderLog).FtParseAdd(nil, orderArr[i].OrderNo, &log).Add()
//				if !SetFailFlag {
//					continue
//				}
//				*mp.Status = models.CONST_OrderStatus_Fail_Already_Retry
//			}else{
//				*mp.Status = models.CONST_OrderStatus_New
//			}
//
//
//
//			if err = mp.Update(); err != nil {
//				ZapLog().Error("CardOrder Update err", zap.Error(err))
//				return
//			}
//			time.Sleep(time.Millisecond * 100)
//		}
//
//		//fmt.Println("startId= ", startId)
//		if len(orderArr) < 10 {
//			break
//		}
//		time.Sleep(time.Second * 1)
//	}
//
//}
//
//func (this *Tasker) ydjthkRetryExpressWork(delay int64) {
//	defer models.PanicPrint()
//
//	if len(config.GConfig.Jthk.ParterCode) <= 0 || len(config.GConfig.Jthk.ParterCodeArr) <= 0{
//		return
//	}
//	partnerIds := GetJtHkPartnerIds()
//
//	if len(partnerIds) <= 0 {
//		return
//	}
//
//	for true {
//		conds := []*models.SqlPairCondition{
//			//&models.SqlPairCondition{"id > ?", startId},
//			&models.SqlPairCondition{"status = ?", models.CONST_OrderStatus_New},
//			&models.SqlPairCondition{"third_order_at <= ?", time.Now().Unix() - delay},
//			//条件还得处理下
//			//&models.SqlPairCondition{"partner_id in ?", parter.Id},
//		}
//		if len(partnerIds) > 0 {
//			conds = append(conds, &models.SqlPairCondition{"partner_id in (?)", partnerIds})
//		}
//
//		orderArr, err := new(models.CardOrder).GetLimitByCond(10, conds, nil)
//		if err != nil {
//			ZapLog().Error("CardOrder GetLimitByCond err", zap.Error(err))
//			return
//		}
//		if orderArr == nil || len(orderArr) <= 0 {
//			//fmt.Println("nil data")
//			return
//		}
//
//		//ZapLog().Sugar().Infof("jthktasker express %d", len(orderArr))
//
//		//记录id, 倒叙
//		haveExpreeFlag := false
//		for i := len(orderArr) - 1; i >= 0; i-- {
//			if orderArr[i] == nil || orderArr[i].Status == nil {
//				continue
//			}
//			if  *orderArr[i].Status != models.CONST_OrderStatus_New {
//				continue
//			}
//			if orderArr[i].ExpressNo != nil && len(*orderArr[i].ExpressNo) >= 2 && orderArr[i].Express !=nil && len(*orderArr[i].Express) >= 1{
//				haveExpreeFlag = true
//				continue
//			}
//
//			if orderArr[i].Phone == nil || orderArr[i].NewPhone == nil || orderArr[i].IdCard == nil {
//				log:= "快递查询：信息不全"
//				new(models.CardOrderLog).FtParseAdd(nil, orderArr[i].OrderNo, &log).Add()
//				continue
//			}
//
//			mp := &models.CardOrder{
//				Id: orderArr[i].Id,
//			}
//
//			resOrderSearch,err := new(ydjthk.ReOrderSerach).Send(*orderArr[i].Phone, *orderArr[i].IdCard)
//			if err != nil {
//				log:= "快递查询：网络问题，"+err.Error()
//				new(models.CardOrderLog).FtParseAdd(nil, orderArr[i].OrderNo, &log).Add()
//				continue
//			}
//			if resOrderSearch.Ret != 200 {
//				log:= fmt.Sprintf("快递查询：%d-%s", resOrderSearch.Ret, resOrderSearch.Msg)
//				new(models.CardOrderLog).FtParseAdd(nil, orderArr[i].OrderNo, &log).Add()
//				continue
//			}
//
//			var chooseOne *ydjthk.OrderInfo
//			yidongArr := resOrderSearch.Datas
//			for j:=0; j< len(yidongArr); j++ {
//				if yidongArr[j].Number == nil || orderArr[i].NewPhone == nil {
//					continue
//				}
//				if * yidongArr[j].Number == *orderArr[i].NewPhone {
//					chooseOne = yidongArr[j]
//					break
//				}
//			}
//			if chooseOne == nil {
//				log:= "快递查询：未查到相关信息"
//				new(models.CardOrderLog).FtParseAdd(nil, orderArr[i].OrderNo, &log).Add()
//				continue
//			}
//			if chooseOne.ShipmentNo == nil || len(*chooseOne.ShipmentNo) < 2 {
//				log:= "快递查询：未查到相关信息"
//				new(models.CardOrderLog).FtParseAdd(nil, orderArr[i].OrderNo, &log).Add()
//				continue
//			}
//			haveExpreeFlag = true
//
//			mp.Express = chooseOne.ShipmentCompany
//			mp.ExpressNo = chooseOne.ShipmentNo
//
//			if chooseOne.Status != nil{
//				*chooseOne.Status = strings.ToUpper(*chooseOne.Status)
//				if !strings.HasPrefix(*chooseOne.Status, "S")  { // 不成功
//					log:= "快递查询：状态错误-"+*chooseOne.Status
//					new(models.CardOrderLog).FtParseAdd(nil, orderArr[i].OrderNo, &log).Add()
//					//continue
//				}
//			}
//
//			if ((mp.Express == nil) || (len(*mp.Express) <= 1)) && (chooseOne.ShipmentNo !=nil) && (len(chooseOne.Tid) > 0) && (chooseOne.ShipmentCompanyCode != nil) {
//				orderDetail,err := new(ydjthk.ReOrderDetailSerach).Send(*orderArr[i].Phone, *orderArr[i].IdCard, chooseOne.Tid, *chooseOne.ShipmentCompanyCode, *chooseOne.ShipmentNo)
//				if err != nil {
//					log:= "快递详情查询:网络问题，"+err.Error()
//					new(models.CardOrderLog).FtParseAdd(nil, orderArr[i].OrderNo, &log).Add()
//				}
//				if orderDetail != nil {
//					mp.Express = orderDetail.ShipmentCompany
//				}
//			}
//
//			mp.Status = new(int)
//			*mp.Status = models.CONST_OrderStatus_Already_Delivered
//			mp.DeliverAt = new(int64)
//			*mp.DeliverAt = time.Now().Unix()
//
//			if err = mp.Update(); err != nil {
//				ZapLog().Error("CardOrder Update err", zap.Error(err))
//				return
//			}
//			time.Sleep(time.Second * 1)
//		}
//
//		if ! haveExpreeFlag {
//			//return
//		}
//
//		if len(orderArr) < 10 {
//			break
//		}
//		time.Sleep(time.Second * 2)
//	}
//
//}
