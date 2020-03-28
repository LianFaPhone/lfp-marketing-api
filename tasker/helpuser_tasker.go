package tasker

import (
	. "LianFaPhone/lfp-base/log/zap"
	"LianFaPhone/lfp-marketing-api/config"
	"LianFaPhone/lfp-marketing-api/models"
	"LianFaPhone/lfp-marketing-api/thirdcard-api/ydjthk"
	"fmt"
	"go.uber.org/zap"
	"strconv"
	"strings"
	"time"
)

func (this *Tasker) jtydhkHelpUserWork() {
	defer models.PanicPrint()

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

	//orderTable := new(models.CardOrder).TableName()
	//OrderRetryTable :=  new(models.CardOrderRetry).TableName()
	//jsonStr := fmt.Sprintf("left join %s on %s.order_no = %s.order_no", OrderRetryTable, OrderRetryTable, orderTable)


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

	time.Sleep(time.Second*3)
	startId := int64(0)
	maxLimit := 100

	for ;true; {
		conds := []*models.SqlPairCondition{
			&models.SqlPairCondition{"id > ?", startId},
			&models.SqlPairCondition{"created_at >= ?", time.Now().Unix() - 52*3600},
			&models.SqlPairCondition{"status = ?", models.CONST_OrderStatus_HelpUser_Apply_Doing},
			&models.SqlPairCondition{"third_order_at <= ?", time.Now().Unix() - 4*3600},
		}
		if len(partnerIds) > 0 {
			conds = append(conds, &models.SqlPairCondition{"partner_id in (?)", partnerIds})
		}

		orderArr, err := new(models.CardOrder).GetLimitByCond2(maxLimit, conds, nil)
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
			if *orderArr[i].Id > startId {
				startId = *orderArr[i].Id
			}
			if orderArr[i] == nil || orderArr[i].Status == nil{
				continue
			}

			mp := &models.CardOrder{
				Id: orderArr[i].Id,
				Status : new(int),
				ThirdOrderAt: new(int64),
			}
			*mp.ThirdOrderAt = time.Now().Unix()

			if orderArr[i].PartnerGoodsCode == nil || orderArr[i].Province == nil || orderArr[i].City == nil || orderArr[i].Area == nil || orderArr[i].Address == nil || orderArr[i].Phone == nil || orderArr[i].TrueName == nil || orderArr[i].IdCard == nil || orderArr[i].PartnerGoodsCode == nil{
				*mp.Status = models.CONST_OrderStatus_Fail_Already_Retry
				mp.Update()
				new(models.CardOrderLog).FtParseAdd2(orderArr[i].Id, orderArr[i].OrderNo,"帮助用户下单失败|表中数据缺失").Add()
				continue
			}

			existFlag,err := HaveHistoryOrderSame(*orderArr[i].IdCard, *orderArr[i].Phone, *orderArr[i].PartnerGoodsCode)
			if err != nil  {
				ZapLog().Error("HaveHistoryOrderSame err", zap.Error(err))
			}else if existFlag {
				*mp.Status = models.CONST_OrderStatus_Fail_Already_Retry
				mp.Update()
				new(models.CardOrderLog).FtParseAdd2(orderArr[i].Id, orderArr[i].OrderNo, "帮助用户下单失败|关联订单已完成").Add()
				continue
			}
			queryValues,err := GetUrlParam(*orderArr[i].PartnerGoodsCode)
			if err != nil {
				*mp.Status = models.CONST_OrderStatus_Fail_Already_Retry
				mp.Update()
				new(models.CardOrderLog).FtParseAdd2(orderArr[i].Id, orderArr[i].OrderNo, "帮助用户下单失败|额外参数缺失").Add()
				continue
			}
			channelId := queryValues.Get("channelId")
			productId := queryValues.Get("productId")
			isOao,_ := strconv.ParseBool(queryValues.Get("isOao"))

			if len(channelId) <= 0 || len(productId) <= 0 {
				*mp.Status = models.CONST_OrderStatus_Fail_Already_Retry
				mp.Update()
				new(models.CardOrderLog).FtParseAdd2(orderArr[i].Id, orderArr[i].OrderNo,  "帮助用户下单失败|额外参数缺失").Add()
				continue
			}
			time.Sleep(time.Millisecond*100)
			token,err := new(ydjthk.ReToken).Send(isOao, channelId)
			if err != nil {
				*mp.Status = GethelpUserStatus(*orderArr[i].OrderNo, *orderArr[i].Status)
				mp.Update()
				new(models.CardOrderLog).FtParseAdd2(orderArr[i].Id, orderArr[i].OrderNo, "帮助用户下单失败|请求token失败，"+err.Error()).Add()
				continue
			}
			province,ok := ProvinceMap[*orderArr[i].Province]
			if !ok {
				*mp.Status = models.CONST_OrderStatus_Fail_Already_Retry
				mp.Update()
				new(models.CardOrderLog).FtParseAdd2(orderArr[i].Id, orderArr[i].OrderNo,  "帮助用户下单失败|省份匹配不上").Add()
				continue
			}
			city, ok := province.CityMap[*orderArr[i].City]
			if !ok {
				*mp.Status = models.CONST_OrderStatus_Fail_Already_Retry
				mp.Update()
				new(models.CardOrderLog).FtParseAdd2(orderArr[i].Id, orderArr[i].OrderNo, "帮助用户下单失败|城市匹配不上").Add()
				continue
			}
			area,ok := city.AreaMap[*orderArr[i].Area]
			if !ok {
				*mp.Status = models.CONST_OrderStatus_Fail_Already_Retry
				mp.Update()
				new(models.CardOrderLog).FtParseAdd2(orderArr[i].Id, orderArr[i].OrderNo, "帮助用户下单失败|区（县）匹配不上").Add()
				continue
			}
			time.Sleep(time.Millisecond*100)
			numbers,err := new(ydjthk.ReCardSearch).Send(isOao, province.ProvinceId, province.ProvinceName, city.CityId, city.CityName, "", 1, 10)
			if err != nil {
				*mp.Status = GethelpUserStatus(*orderArr[i].OrderNo, *orderArr[i].Status)
				mp.Update()
				new(models.CardOrderLog).FtParseAdd2(orderArr[i].Id, orderArr[i].OrderNo, "帮助用户下单失败|获取新号码失败，"+err.Error()).Add()
				continue
			}
			time.Sleep(time.Millisecond*100)
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
				*mp.Status = GethelpUserStatus(*orderArr[i].OrderNo, *orderArr[i].Status)
				mp.Update()
				new(models.CardOrderLog).FtParseAdd2(orderArr[i].Id, orderArr[i].OrderNo,  "帮助用户下单失败|无法锁定新号码").Add()
				continue
			}
			time.Sleep(time.Millisecond*1000)
			_, thirdOrderNo,oaoFlag,orderErr := new(ydjthk.ReOrderSubmit).Parse(channelId, productId, nil).Send(isOao, token,  *orderArr[i].Phone, chooseNumber, *orderArr[i].TrueName, *orderArr[i].IdCard, *orderArr[i].Address, province.ProvinceId, city.CityId, province.ProvinceId, city.CityId, area.AreaId)
			if orderErr != nil {
				ZapLog().Error("helpuser ReOrderSubmit err", zap.Error(orderErr), zap.String("channelId", channelId), zap.String("productId", productId), zap.String("other", *orderArr[i].Phone+ chooseNumber+ *orderArr[i].TrueName+ *orderArr[i].IdCard+ *orderArr[i].Address+ *orderArr[i].Province+ *orderArr[i].City+ province.ProvinceId+ city.CityId+ area.AreaId))
				*mp.Status = GenHelpUserFailStatus(GethelpUserStatus(*orderArr[i].OrderNo, *orderArr[i].Status), orderErr.Error())
				mp.Update()
				new(models.CardOrderLog).FtParseAdd2(orderArr[i].Id, orderArr[i].OrderNo, "帮助用户下单失败|下单失败，"+orderErr.Error()).Add()
				continue
			}

			log := "帮助用户下单成功|新订单已完成"
			if !oaoFlag {
				//*mp.Status = models.CONST_OrderStatus_New_UnFinish// 这行无效

				newUrl,err := new(ydjthk.ReIdCheckUrl).Send(isOao, channelId, thirdOrderNo, chooseNumber, token)
				if err != nil {
					ZapLog().Error("ReIdCheckUrl send err", zap.Error(err))
					log = "帮助用户下单成功|获取上传照片网址失败，"+err.Error()
					//*mp.Status = GethelpUserStatus(*orderArr[i].OrderNo, *orderArr[i].Status)
				}else{
					//sendUnFinishNotify(newUrl, orderArr[i])
					log = "帮助用户下单成功|新订单未完成，等待上传照片"
					new(models.CardOrderUrl).FtParseAdd(orderArr[i].Id, orderArr[i].OrderNo, &newUrl).Add()
				}
				*mp.Status = GethelpUserStatus(*orderArr[i].OrderNo, *orderArr[i].Status)
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
			if err:=mp.MaxIdByOrderNo(*orderArr[i].OrderNo); err != nil {
				ZapLog().Error("helpuser maxIdByorderNo error", zap.Error(err),zap.String("no",  *orderArr[i].OrderNo))
			} // 这一步是关键，让它重新被检测
			time.Sleep(time.Second * 1)
		}

		if len(orderArr) < maxLimit {
			break
		}
		time.Sleep(time.Second * 3)
	}

}

func GethelpUserStatus(orderNo string, oldstatus int) int {
	count, err := new(models.CardOrderRetry).Incr(orderNo)
	if err != nil {
		ZapLog().Error("new(models.CardOrderRetry).Incr err", zap.Error(err))
		return oldstatus
	}
	if count >= 6 {
		return models.CONST_OrderStatus_Fail_Already_Retry
	}
	return oldstatus
}

func GenHelpUserFailStatus(oldStatus int, errMsg string) int {
	if oldStatus == models.CONST_OrderStatus_Fail_Already_Retry {
		return models.CONST_OrderStatus_Fail_Already_Retry
	}

	 if strings.Contains(errMsg, "系统错误") {
			return oldStatus
	 }else if strings.Contains(errMsg, "订购的号码不存在") {
			return oldStatus
	 }else if strings.Contains(errMsg, "号码已被占用") {
			return oldStatus
	 }else if strings.Contains(errMsg, "系统忙") {
			return oldStatus
	 }
	return models.CONST_OrderStatus_Fail_Already_Retry
}

func HaveHistoryOrderSame(idcard, phone, partnerGoodsCode string) (bool,error){
	param := &models.CardOrder{
		IdCard: &idcard,
		Phone: &phone,
		PartnerGoodsCode: &partnerGoodsCode,
	}
	condStr := fmt.Sprintf("( status = %d or status = %d ) and created_at >= %d", models.CONST_OrderStatus_New, models.CONST_OrderStatus_Already_Delivered, time.Now().Unix() - 10*3600)
	existFlag,err := param.Exist(condStr)
	return existFlag,err
}
