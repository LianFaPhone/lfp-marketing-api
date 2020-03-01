package tasker

import (
	. "LianFaPhone/lfp-base/log/zap"
	"LianFaPhone/lfp-marketing-api/config"
	"LianFaPhone/lfp-marketing-api/models"
	"LianFaPhone/lfp-marketing-api/thirdcard-api/ydhk"
	"go.uber.org/zap"
	"time"
)

func (this *Tasker) ydhkOaoWork() {
	defer models.PanicPrint()

	recoder, err := new(models.IdRecorder).GetByName("ydhk_oao")
	if err != nil {
		ZapLog().Error("IdRecorder GetByName err", zap.Error(err))
		return
	}

	if recoder == nil {
		if recoder, err = new(models.IdRecorder).Add("ydhk_oao", 0); err != nil {
			ZapLog().Error("IdRecorder Add card_order_ips err", zap.Error(err))
			return
		}
	}

	startId := int64(0)
	if recoder.IdTag != nil {
		startId = *recoder.IdTag
	}

	for true {
		conds := []*models.SqlPairCondition{
			&models.SqlPairCondition{"id > ?", startId},
			//&models.SqlPairCondition{"class_big_tp = ?", 5},
			&models.SqlPairCondition{"created_at >= ?", time.Now().Unix() - 15*60},
			//条件还得处理下
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
				log:= "信息不全"
				new(models.CardOrderLog).FtParseAdd(nil, orderArr[i].OrderNo, &log).Add()
				continue
			}

			mp := &models.CardOrder{
				Id: orderArr[i].Id,
			}

			yidongArr,err := new(ydhk.ReOrderSerach).Send(*orderArr[i].Phone, *orderArr[i].IdCard)
			if err != nil {
				log:= err.Error()
				new(models.CardOrderLog).FtParseAdd(nil, orderArr[i].OrderNo, &log).Add()
				continue
			}

			oaoFlag := false
			for j:=0; j< len(yidongArr); j++ {
				if yidongArr[j].Number == nil {
					continue
				}
				if * yidongArr[j].Number == *orderArr[i].NewPhone {
					oaoFlag = true
					break
				}
			}

			if !oaoFlag {
				log:= "oao未发现"
				new(models.CardOrderLog).FtParseAdd(nil, orderArr[i].OrderNo, &log).Add()
				continue
			}


			mp.Status = new(int)
			*mp.Status = models.CONST_OrderStatus_New


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

	if len(config.GConfig.Jthk.ParterCode) <= 0 {
		return
	}
	parter,err := new(models.PdPartner).GetByCode(config.GConfig.Jthk.ParterCode)
	if err != nil {
		//日志
		return
	}
	if parter == nil {
		return
	}

	for true {
		conds := []*models.SqlPairCondition{
			&models.SqlPairCondition{"id > ?", startId},
			//&models.SqlPairCondition{"class_big_tp = ?", 5},
			&models.SqlPairCondition{"created_at >= ?", time.Now().Unix() - 24*3600},
			//条件还得处理下
			&models.SqlPairCondition{"partner_id = ?", parter.Id},
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
			if orderArr[i].ExpressNo != nil && len(*orderArr[i].ExpressNo) > 0{
				haveExpreeFlag = true
				continue
			}

			if orderArr[i].Phone == nil || orderArr[i].NewPhone == nil || orderArr[i].IdCard == nil {
				log:= "信息不全"
				new(models.CardOrderLog).FtParseAdd(nil, orderArr[i].OrderNo, &log).Add()
				continue
			}

			mp := &models.CardOrder{
				Id: orderArr[i].Id,
			}

			yidongArr,err := new(ydhk.ReOrderSerach).Send(*orderArr[i].Phone, *orderArr[i].IdCard)
			if err != nil {
				log:= err.Error()
				new(models.CardOrderLog).FtParseAdd(nil, orderArr[i].OrderNo, &log).Add()
				continue
			}

			var chooseOne *ydhk.OrderInfo
			for j:=0; j< len(yidongArr); j++ {
				if yidongArr[j].Number == nil {
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



			mp.Status = new(int)
			*mp.Status = models.CONST_OrderStatus_Already_Delivered


			if err = mp.Update(); err != nil {
				ZapLog().Error("CardOrder Update err", zap.Error(err))
				return
			}
			time.Sleep(time.Second * 1)
		}

		if ! haveExpreeFlag {
			return
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


