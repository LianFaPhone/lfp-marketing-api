package tasker

import (
	. "LianFaPhone/lfp-base/log/zap"
	"LianFaPhone/lfp-marketing-api/config"
	"LianFaPhone/lfp-marketing-api/models"
	"LianFaPhone/lfp-marketing-api/thirdcard-api/ydhk"
	"fmt"
	"go.uber.org/zap"
	"strings"
	"time"
)


func (this *Tasker) ydhkOaoWork(idRecorderName string, delayTime int64, SetFailFlag bool) {
	defer models.PanicPrint()

	recoder, err := new(models.IdRecorder).GetByName(idRecorderName)
	if err != nil {
		ZapLog().Error("IdRecorder GetByName err", zap.Error(err))
		return
	}

	if recoder == nil {
		if recoder, err = new(models.IdRecorder).Add(idRecorderName, 0); err != nil {
			ZapLog().Error("IdRecorder Add ydhk_oao err", zap.Error(err))
			return
		}
	}

	startId := int64(0)
	if recoder.IdTag != nil {
		startId = *recoder.IdTag
	}

	if len(config.GConfig.Jthk.ParterCode) <= 0 || len(config.GConfig.Jthk.ParterCodeArr) <= 0 {
		return
	}

	partnerIds := GetJtHkPartnerIds()
	if len(partnerIds) <= 0 {
		return
	}

	for true {
		conds := []*models.SqlPairCondition{
			&models.SqlPairCondition{"id > ?", startId},
			//&models.SqlPairCondition{"class_big_tp = ?", 5},
			&models.SqlPairCondition{"created_at <= ?", time.Now().Unix() - delayTime},
			//条件还得处理下
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

		//ZapLog().Sugar().Infof("jthktasker %d", len(orderArr))

		//记录id, 倒叙
		for i := len(orderArr) - 1; i >= 0; i-- {
			if *orderArr[i].Id > startId {
				startId = *orderArr[i].Id
			}
			if orderArr[i] == nil || orderArr[i].Status == nil{
				continue
			}
			if  *orderArr[i].Status != models.CONST_OrderStatus_New_UnFinish {
				continue
			}
			if orderArr[i].Phone == nil || orderArr[i].NewPhone == nil || orderArr[i].IdCard == nil {
				log:= "OAO检测：信息不全"
				new(models.CardOrderLog).FtParseAdd(nil, orderArr[i].OrderNo, &log).Add()
				continue
			}

			mp := &models.CardOrder{
				Id: orderArr[i].Id,
				Status: new(int),
			}

			resOrderShortSerach,err := new(ydhk.ReOrderShortSerach).Send(*orderArr[i].Phone, *orderArr[i].IdCard);
			if err != nil {
				log:= "OAO检测："+err.Error()
				new(models.CardOrderLog).FtParseAdd(nil, orderArr[i].OrderNo, &log).Add()
				continue
			}
			if resOrderShortSerach.Ret != 200 {
				log:= fmt.Sprintf("OAO检测：%d-%s", resOrderShortSerach.Ret, resOrderShortSerach.Msg)
				new(models.CardOrderLog).FtParseAdd(nil, orderArr[i].OrderNo, &log).Add()
				if !SetFailFlag {
					continue
				}
				*mp.Status = models.CONST_OrderStatus_Fail
				if err = mp.Update(); err != nil {
					ZapLog().Error("CardOrder Update err", zap.Error(err))
				}
				continue
			}

			resOrderSearch,err := new(ydhk.ReOrderSerach).Send(*orderArr[i].Phone, *orderArr[i].IdCard)
			if err != nil {
				log:= "OAO检测："+err.Error()
				new(models.CardOrderLog).FtParseAdd(nil, orderArr[i].OrderNo, &log).Add()
				continue
			}
			if resOrderSearch.Ret != 200 {
				log:= fmt.Sprintf("OAO检测：%d-%s", resOrderSearch.Ret, resOrderSearch.Msg)
				new(models.CardOrderLog).FtParseAdd(nil, orderArr[i].OrderNo, &log).Add()
				if !SetFailFlag {
					continue
				}
				*mp.Status = models.CONST_OrderStatus_Fail
				if err = mp.Update(); err != nil {
					ZapLog().Error("CardOrder Update err", zap.Error(err))
				}
				continue
			}

			var chooseOne *ydhk.OrderInfo
			yidongArr := resOrderSearch.Datas
			for j:=0; j< len(yidongArr); j++ {
				if yidongArr[j].Number == nil || orderArr[i].NewPhone == nil{
					continue
				}
				if * yidongArr[j].Number == *orderArr[i].NewPhone {
					chooseOne = yidongArr[j]
					break
				}
			}

			if chooseOne == nil {
				log:= "OAO检测：oao未发现"
				new(models.CardOrderLog).FtParseAdd(nil, orderArr[i].OrderNo, &log).Add()
				if !SetFailFlag {
					continue
				}
				*mp.Status = models.CONST_OrderStatus_Fail
			}else{
				*mp.Status = models.CONST_OrderStatus_New
			}



			if err = mp.Update(); err != nil {
				ZapLog().Error("CardOrder Update err", zap.Error(err))
				return
			}
			time.Sleep(time.Second * 1)
		}

		err = recoder.Update(startId)
		if err != nil {
			ZapLog().Error("IdRecorder Update err", zap.Error(err))
			return
		}
		//fmt.Println("startId= ", startId)
		if len(orderArr) < 10 {
			break
		}
		time.Sleep(time.Second * 5)
	}

}


func (this *Tasker) ydhkExpressWork() {
	defer models.PanicPrint()

	recoder, err := new(models.IdRecorder).GetByName("ydhk_express")
	if err != nil {
		ZapLog().Error("IdRecorder GetByName err", zap.Error(err))
		return
	}

	if recoder == nil {
		if recoder, err = new(models.IdRecorder).Add("ydhk_express", 0); err != nil {
			ZapLog().Error("IdRecorder Add card_order_ips err", zap.Error(err))
			return
		}
	}

	startId := int64(0)
	if recoder.IdTag != nil {
		startId = *recoder.IdTag
	}

	if len(config.GConfig.Jthk.ParterCode) <= 0 || len(config.GConfig.Jthk.ParterCodeArr) <= 0{
		return
	}
	partnerIds := GetJtHkPartnerIds()

	if len(partnerIds) <= 0 {
		return
	}

	for true {
		conds := []*models.SqlPairCondition{
			&models.SqlPairCondition{"id > ?", startId},
			//&models.SqlPairCondition{"class_big_tp = ?", 5},
			&models.SqlPairCondition{"created_at <= ?", time.Now().Unix() - 24*3600},
			//条件还得处理下
			//&models.SqlPairCondition{"partner_id in ?", parter.Id},
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

		//ZapLog().Sugar().Infof("jthktasker express %d", len(orderArr))

		//记录id, 倒叙
		haveExpreeFlag := false
		for i := len(orderArr) - 1; i >= 0; i-- {
			if *orderArr[i].Id > startId {
				startId = *orderArr[i].Id
			}
			if orderArr[i] == nil || orderArr[i].Status == nil {
				continue
			}
			if  *orderArr[i].Status != models.CONST_OrderStatus_New {
				continue
			}
			if orderArr[i].ExpressNo != nil && len(*orderArr[i].ExpressNo) > 0 && orderArr[i].Express !=nil && len(*orderArr[i].Express) > 0{
				haveExpreeFlag = true
				continue
			}

			if orderArr[i].Phone == nil || orderArr[i].NewPhone == nil || orderArr[i].IdCard == nil {
				log:= "快递查询：信息不全"
				new(models.CardOrderLog).FtParseAdd(nil, orderArr[i].OrderNo, &log).Add()
				continue
			}

			mp := &models.CardOrder{
				Id: orderArr[i].Id,
			}

			resOrderSearch,err := new(ydhk.ReOrderSerach).Send(*orderArr[i].Phone, *orderArr[i].IdCard)
			if err != nil {
				log:= "快递查询："+err.Error()
				new(models.CardOrderLog).FtParseAdd(nil, orderArr[i].OrderNo, &log).Add()
				continue
			}
			if resOrderSearch.Ret != 200 {
				log:= fmt.Sprintf("快递查询：%d-%s", resOrderSearch.Ret, resOrderSearch.Msg)
				new(models.CardOrderLog).FtParseAdd(nil, orderArr[i].OrderNo, &log).Add()
				continue
			}

			var chooseOne *ydhk.OrderInfo
			yidongArr := resOrderSearch.Datas
			for j:=0; j< len(yidongArr); j++ {
				if yidongArr[j].Number == nil || orderArr[i].NewPhone == nil {
					continue
				}
				if * yidongArr[j].Number == *orderArr[i].NewPhone {
					chooseOne = yidongArr[j]
					break
				}
			}
			if chooseOne == nil {
				continue
			}
			haveExpreeFlag = true

			mp.Express = chooseOne.ShipmentCompany
			mp.ExpressNo = chooseOne.ShipmentNo

			if chooseOne.Status != nil{
				*chooseOne.Status = strings.ToUpper(*chooseOne.Status)
				if !strings.HasPrefix(*chooseOne.Status, "S")  { // 不成功
					log:= "快递查询：状态错误-"+*chooseOne.Status
					new(models.CardOrderLog).FtParseAdd(nil, orderArr[i].OrderNo, &log).Add()
					//continue
				}
			}

			if (mp.Express == nil) && (chooseOne.ShipmentNo !=nil) && (len(chooseOne.Tid) > 0) && (chooseOne.ShipmentCompanyCode != nil) {
				orderDetail,err := new(ydhk.ReOrderDetailSerach).Send(*orderArr[i].Phone, *orderArr[i].IdCard, chooseOne.Tid, *chooseOne.ShipmentCompanyCode, *chooseOne.ShipmentNo)
				if err != nil {
					log:= "快递详情查询:"+err.Error()
					new(models.CardOrderLog).FtParseAdd(nil, orderArr[i].OrderNo, &log).Add()
				}
				if orderDetail != nil {
					mp.Express = orderDetail.ShipmentCompany
				}
			}

			mp.Status = new(int)
			*mp.Status = models.CONST_OrderStatus_Already_Delivered


			if err = mp.Update(); err != nil {
				ZapLog().Error("CardOrder Update err", zap.Error(err))
				return
			}
			time.Sleep(time.Second * 1)
		}

		if ! haveExpreeFlag {
			//return
		}

		err = recoder.Update(startId)
		if err != nil {
			ZapLog().Error("IdRecorder Update err", zap.Error(err))
			return
		}
		//fmt.Println("startId= ", startId)
		if len(orderArr) < 10 {
			break
		}
		time.Sleep(time.Second * 5)
	}

}


func GetJtHkPartnerIds() []*int64{
	if len(config.GConfig.Jthk.ParterCode) <= 0 || len(config.GConfig.Jthk.ParterCodeArr) <= 0 {
		return nil
	}

	partnerIds := make([]*int64, 0)
	for i:=0; i< len(config.GConfig.Jthk.ParterCodeArr);i++ {
		if len(config.GConfig.Jthk.ParterCodeArr[i]) <= 0 {
			continue
		}
		parter,err := new(models.PdPartner).GetByCode(config.GConfig.Jthk.ParterCodeArr[i])
		if err != nil {
			ZapLog().Error("GetByCOde err", zap.Error(err))
			continue
		}
		if parter == nil {
			continue
		}
		partnerIds = append(partnerIds, parter.Id)
	}

	return partnerIds
}
