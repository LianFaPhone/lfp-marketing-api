package tasker

import (
	. "LianFaPhone/lfp-base/log/zap"
	"LianFaPhone/lfp-marketing-api/config"
	"LianFaPhone/lfp-marketing-api/models"
	"LianFaPhone/lfp-marketing-api/thirdcard-api/ding"
	"LianFaPhone/lfp-marketing-api/thirdcard-api/ydjthk"
	"fmt"
	"go.uber.org/zap"


		"strings"
	"time"
)


func (this *Tasker) ydjthkOaoWork(idRecorderName string, delayTime int64, SetFailFlag bool) {
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
			&models.SqlPairCondition{"status = ?", models.CONST_OrderStatus_New_UnFinish},
			&models.SqlPairCondition{"third_order_at <= ?", time.Now().Unix() - delayTime},
			//条件还得处理下
		}

		if len(partnerIds) > 0 {
			conds = append(conds, &models.SqlPairCondition{"partner_id in (?)", partnerIds})
		}

		orderArr, err := new(models.CardOrder).GetLimitByCond2(10, conds, nil)
		if err != nil {
			ZapLog().Error("CardOrder GetLimitByCond err", zap.Error(err))
			return
		}
		if orderArr == nil || len(orderArr) <= 0 {
			//fmt.Println("nil data")
			return
		}

		endTime := time.Now().Add(time.Hour * 5).Format("2006-01-02")
		startTime := time.Now().Add(-time.Hour * 24*7).Format("2006-01-02")

		//ZapLog().Sugar().Infof("jthktasker %d", len(orderArr))

		//记录id, 倒叙
		for i := len(orderArr) - 1; i >= 0; i-- {
			if *orderArr[i].Id > startId {
				startId = *orderArr[i].Id
			}
			if orderArr[i] == nil || orderArr[i].Status == nil || orderArr[i].ThirdOrderNo == nil || len(*orderArr[i].ThirdOrderNo) <= 1{
				continue
			}
			new(models.CardOrderLog).FtParseAdd2(nil, orderArr[i].OrderNo, fmt.Sprintf("oao检测：开始+%v",SetFailFlag)).Add()

			if  *orderArr[i].Status != models.CONST_OrderStatus_New_UnFinish {
				continue
			}

			mp := &models.CardOrder{
				Id: orderArr[i].Id,
				Status: new(int),
			}

			if orderArr[i].Phone == nil || orderArr[i].NewPhone == nil || orderArr[i].IdCard == nil {
				log:= "OAO检测：信息不全"
				new(models.CardOrderLog).FtParseAdd(nil, orderArr[i].OrderNo, &log).Add()
				*mp.Status = models.CONST_OrderStatus_Fail_Retry
				if err = mp.Update(); err != nil {
					ZapLog().Error("CardOrder Update err", zap.Error(err))
				}
				continue
			}

			resOrderShortSerach,err := new(ydjthk.ReYgOrderSerach).Send(*orderArr[i].ThirdOrderNo, startTime, endTime);
			if err != nil {
				log:= "OAO检测：网络错误，"+err.Error()
				new(models.CardOrderLog).FtParseAdd(nil, orderArr[i].OrderNo, &log).Add()
				continue
			}
			if resOrderShortSerach.Success != true {
				log:= fmt.Sprintf("OAO检测：%v-%v", resOrderShortSerach.Success, resOrderShortSerach.Message)
				new(models.CardOrderLog).FtParseAdd(nil, orderArr[i].OrderNo, &log).Add()
				if !SetFailFlag {
					continue
				}
				continue
				*mp.Status = models.CONST_OrderStatus_HelpUser_Apply_Doing
			}else{
				if resOrderShortSerach.Total == 0 {
					log:= fmt.Sprintf("OAO检测：oao未发现,%v", SetFailFlag)
					new(models.CardOrderLog).FtParseAdd(nil, orderArr[i].OrderNo, &log).Add()
					if !SetFailFlag {
						continue
					}
					*mp.Status = models.CONST_OrderStatus_HelpUser_Apply_Doing
				}else{
					new(models.CardOrderLog).FtParseAdd2(nil, orderArr[i].OrderNo, "OAO检测|成功找到订单").Add()
					*mp.Status = models.CONST_OrderStatus_New
				}
			}

			if err = mp.Update(); err != nil {
				ZapLog().Error("CardOrder Update err", zap.Error(err))
				return
			}
			time.Sleep(time.Millisecond * 10)
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
		if !SetFailFlag {
			time.Sleep(time.Millisecond * 100)
		}else{
			time.Sleep(time.Second * 2)
		}
	}

}


func (this *Tasker) ydjthkExpressWork(idRecordName string, delay int64) {
	defer models.PanicPrint()

	recoder, err := new(models.IdRecorder).GetByName(idRecordName)
	if err != nil {
		ZapLog().Error("IdRecorder GetByName err", zap.Error(err))
		return
	}

	if recoder == nil {
		if recoder, err = new(models.IdRecorder).Add(idRecordName, 0); err != nil {
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
			&models.SqlPairCondition{"status = ?", models.CONST_OrderStatus_New},
			&models.SqlPairCondition{"third_order_at <= ?", time.Now().Unix() - delay},
			//条件还得处理下
			//&models.SqlPairCondition{"partner_id in ?", parter.Id},
		}
		if len(partnerIds) > 0 {
			conds = append(conds, &models.SqlPairCondition{"partner_id in (?)", partnerIds})
		}

		orderArr, err := new(models.CardOrder).GetLimitByCond2(10, conds, nil)
		if err != nil {
			ZapLog().Error("CardOrder GetLimitByCond err", zap.Error(err))
			return
		}
		if orderArr == nil || len(orderArr) <= 0 {
			//fmt.Println("nil data")
			return
		}

		endTime := time.Now().Add(time.Hour * 10 ).Format("2006-01-02")
		startTime := time.Now().Add(-time.Hour * 24 * 8).Format("2006-01-02")


		//ZapLog().Sugar().Infof("jthktasker express %d", len(orderArr))

		//记录id, 倒叙
		haveExpreeFlag := false
		for i := len(orderArr) - 1; i >= 0; i-- {
			if *orderArr[i].OrderNo == "D12004101203170004664" {
				ZapLog().Info("D12004101203170004664 find", zap.Any("status", orderArr[i].Status))
			}
			if *orderArr[i].Id > startId {
				startId = *orderArr[i].Id
			}
			if orderArr[i] == nil || orderArr[i].Status == nil || orderArr[i].ThirdOrderNo == nil|| len(*orderArr[i].ThirdOrderNo) <=1{
				continue
			}
			if  *orderArr[i].Status != models.CONST_OrderStatus_New {
				continue
			}
			new(models.CardOrderLog).FtParseAdd2(nil, orderArr[i].OrderNo, "快递查询|开始").Add()

			if orderArr[i].ExpressNo != nil && len(*orderArr[i].ExpressNo) >= 2 && orderArr[i].Express !=nil && len(*orderArr[i].Express) >= 1{
				haveExpreeFlag = true
				new(models.CardOrderLog).FtParseAdd2(nil, orderArr[i].OrderNo, "快递查询|信息库中已存在").Add()
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

			resOrderSearch,err := new(ydjthk.ReYgOrderSerach).Send(*orderArr[i].ThirdOrderNo, startTime, endTime);
			if err != nil {
				log:= "快递查询：网络问题，"+err.Error()
				new(models.CardOrderLog).FtParseAdd(nil, orderArr[i].OrderNo, &log).Add()
				continue
			}
			if resOrderSearch.Success != true {
				log:= fmt.Sprintf("快递查询：%v-%v", resOrderSearch.Success, resOrderSearch.Message)
				new(models.CardOrderLog).FtParseAdd(nil, orderArr[i].OrderNo, &log).Add()
				continue
			}

			if resOrderSearch.Total == 0 {
				log:= "快递查询：未查到相关信息"
				new(models.CardOrderLog).FtParseAdd(nil, orderArr[i].OrderNo, &log).Add()
				continue
			}

			if resOrderSearch.Datas == nil || len(resOrderSearch.Datas) <= 0 {
				log:= "快递查询：未查到相关信息"
				new(models.CardOrderLog).FtParseAdd(nil, orderArr[i].OrderNo, &log).Add()
				continue
			}

			if resOrderSearch.Datas[0].ShipmentNo == nil || len(*resOrderSearch.Datas[0].ShipmentNo) < 2 {
				log:= "快递查询：未查到相关信息"
				new(models.CardOrderLog).FtParseAdd(nil, orderArr[i].OrderNo, &log).Add()
				continue
			}
			haveExpreeFlag = true

			mp.Express = resOrderSearch.Datas[0].ShipmentCompany
			mp.ExpressNo = resOrderSearch.Datas[0].ShipmentNo

			mp.Status = new(int)
			*mp.Status = models.CONST_OrderStatus_Already_Delivered
			mp.DeliverAt = new(int64)
			*mp.DeliverAt = time.Now().Unix()
			new(models.CardOrderLog).FtParseAdd2(nil, orderArr[i].OrderNo, "快递查询|成功找到快递信息").Add()

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
		time.Sleep(time.Second * 2)
	}

}

func (this *Tasker) ydjthkActivedWork(idRecordName string, delay int64) {
	defer models.PanicPrint()

	recoder, err := new(models.IdRecorder).GetByName(idRecordName)
	if err != nil {
		ZapLog().Error("IdRecorder GetByName err", zap.Error(err))
		return
	}

	if recoder == nil {
		if recoder, err = new(models.IdRecorder).Add(idRecordName, 0); err != nil {
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

	firstTicketErrFlag := true

	for true {
		conds := []*models.SqlPairCondition{
			&models.SqlPairCondition{"id > ?", startId},
			&models.SqlPairCondition{"status = ?", models.CONST_OrderStatus_Already_Delivered},
			//&models.SqlPairCondition{"third_order_at >= ?", time.Now().Unix() - 15*24*3600},
			&models.SqlPairCondition{"third_order_at <= ?", time.Now().Unix() - delay},
			//条件还得处理下
			//&models.SqlPairCondition{"partner_id in ?", parter.Id},
		}
		if len(partnerIds) > 0 {
			conds = append(conds, &models.SqlPairCondition{"partner_id in (?)", partnerIds})
		}

		orderArr, err := new(models.CardOrder).GetLimitByCond2(10, conds, nil)
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
			if orderArr[i] == nil || orderArr[i].Status == nil || orderArr[i].ThirdOrderNo == nil || len(*orderArr[i].ThirdOrderNo) <= 2 {
				continue
			}
			if  *orderArr[i].Status != models.CONST_OrderStatus_Already_Delivered {
				continue
			}


			new(models.CardOrderLog).FtParseAdd2(nil, orderArr[i].OrderNo, "激活查询|开始").Add()

			if orderArr[i].ThirdOrderAt == nil || *(orderArr[i].ThirdOrderAt) <= 1  {
				orderArr[i].ThirdOrderAt = orderArr[i].CreatedAt
			}

			if orderArr[i].ThirdOrderAt == nil || *(orderArr[i].ThirdOrderAt) <= 1  {
				log:= "激活查询：信息不全"
				new(models.CardOrderLog).FtParseAdd(nil, orderArr[i].OrderNo, &log).Add()
				continue
			}

			uTime := time.Unix(*orderArr[i].ThirdOrderAt, 0)
			endTime := uTime.Add(time.Hour * 24 * 8 ).Format("2006-01-02")
			startTime := uTime.Add(time.Hour * 2).Format("2006-01-02")

			mp := &models.CardOrder{
				Id: orderArr[i].Id,
			}

			resOrderSearch,err := new(ydjthk.ReYgOrderSerach).Send(*orderArr[i].ThirdOrderNo, startTime, endTime);
			if err != nil {
				log:= "激活查询："+err.Error()
				if firstTicketErrFlag && (strings.Contains(err.Error(), "没有权限") || strings.Contains(err.Error(), "无权访问") ){
					new(ding.ReDing).Send("移动花卡ticket权限过期，请重新设置")
					firstTicketErrFlag = false
				}
				new(models.CardOrderLog).FtParseAdd(nil, orderArr[i].OrderNo, &log).Add()
				continue
			}
			if resOrderSearch.Success != true {
				log:= fmt.Sprintf("激活查询：%v-%v", resOrderSearch.Success, resOrderSearch.Message)
				new(models.CardOrderLog).FtParseAdd(nil, orderArr[i].OrderNo, &log).Add()
				continue
			}

			if resOrderSearch.Total == 0 {
				log:= "激活查询：未查到相关信息"
				new(models.CardOrderLog).FtParseAdd(nil, orderArr[i].OrderNo, &log).Add()
				continue
			}

			if resOrderSearch.Datas == nil || len(resOrderSearch.Datas) <= 0 {
				log:= "激活查询：未查到相关信息"
				new(models.CardOrderLog).FtParseAdd(nil, orderArr[i].OrderNo, &log).Add()
				continue
			}
			if resOrderSearch.Datas[0].Status == nil {
				continue
			}
			if *resOrderSearch.Datas[0].Status != ydjthk.Yg_Status_Already_Activated {
				continue
			}


			mp.Status = new(int)
			*mp.Status = models.CONST_OrderStatus_Already_Activated
			if resOrderSearch.Datas[0].ActiveTime != nil {
				tt,_ := time.ParseInLocation("2006-01-02 15:04:05", *resOrderSearch.Datas[0].ActiveTime, time.Local)
				mp.ActiveAt = new(int64)
				*mp.ActiveAt = tt.Unix()
			}


			new(models.CardOrderLog).FtParseAdd2(nil, orderArr[i].OrderNo, "激活查询|成功找到激活信息").Add()

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
		time.Sleep(time.Second * 2)
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
