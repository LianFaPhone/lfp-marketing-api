package tasker

import (
	baidu_api "LianFaPhone/lfp-marketing-api/baidu-api"
	"LianFaPhone/lfp-marketing-api/models"
	"LianFaPhone/lfp-marketing-api/sdk"
	"go.uber.org/zap"
	"time"
)

import (
	. "LianFaPhone/lfp-base/log/zap"
)

func (this *Tasker) newUnFinishSmsWork5hour() {
	defer models.PanicPrint()

	recoder, err := new(models.IdRecorder).GetByName("CardOrderNewUnfinish5H")
	if err != nil {
		ZapLog().Error("IdRecorder GetByName err", zap.Error(err))
		return
	}

	if recoder == nil {
		if recoder, err = new(models.IdRecorder).Add("CardOrderNewUnfinish5H", 0); err != nil {
			ZapLog().Error("IdRecorder Add CardOrderNewUnfinish5H err", zap.Error(err))
			return
		}
	}

	startId := int64(0)
	if recoder.IdTag != nil {
		startId = *recoder.IdTag
	}

	nowTime := time.Now().Unix()
	platformTp :=2

	for true {
		conds := []*models.SqlPairCondition{
			&models.SqlPairCondition{"id > ?", startId},
			&models.SqlPairCondition{"created_at <= ?", nowTime-5*3600},
			&models.SqlPairCondition{"created_at >= ?", nowTime-7*3600},
			&models.SqlPairCondition{"status = ?", models.CONST_OrderStatus_New_UnFinish},
		}

		orderArr, err := new(models.CardOrder).GetLimitByCond(10, conds)
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
			if orderArr[i] == nil {
				continue
			}
			if *orderArr[i].Id > startId {
				startId = *orderArr[i].Id
			}

			cc,err := new(models.CardClass).GetById(*orderArr[i].ClassTp)
			if err != nil {
				ZapLog().Error("Nofind CardClass ", zap.Error(err), zap.Int("classTp", *orderArr[i].ClassTp))
				continue
			}
			if cc == nil || cc.ShortChain == nil ||orderArr[i].Phone == nil{
				continue
			}

			//生成短链，发送短信

			if err := sdk.GNotifySdk.SendSms([]string{*cc.ShortChain}, *orderArr[i].Phone, "new_unfinish",0, &platformTp); err != nil {
				ZapLog().Error("GNotifySdk.SendSms err", zap.Error(err))
				continue
			}

			time.Sleep(time.Millisecond * 1)
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
	}

}

//新未完成订单短信通知
func (this *Tasker) newUnFinishSmsWork5min() {
	defer models.PanicPrint()

	recoder, err := new(models.IdRecorder).GetByName("CardOrderNewUnfinish5M")
	if err != nil {
		ZapLog().Error("IdRecorder GetByName err", zap.Error(err))
		return
	}

	if recoder == nil {
		if recoder, err = new(models.IdRecorder).Add("CardOrderNewUnfinish5M", 0); err != nil {
			ZapLog().Error("IdRecorder Add CardOrderNewUnfinish_5m err", zap.Error(err))
			return
		}
	}

	startId := int64(0)
	if recoder.IdTag != nil {
		startId = *recoder.IdTag
	}

	nowTime := time.Now().Unix()
	platformTp :=2

	for true {
		conds := []*models.SqlPairCondition{
			&models.SqlPairCondition{"id > ?", startId},
			&models.SqlPairCondition{"created_at <= ?", nowTime-5*60},
			&models.SqlPairCondition{"created_at >= ?", nowTime-10*60},
			&models.SqlPairCondition{"status = ?", models.CONST_OrderStatus_New_UnFinish},
		}

		orderArr, err := new(models.CardOrder).GetLimitByCond(10, conds)
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
			if orderArr[i] == nil {
				continue
			}
			if *orderArr[i].Id > startId {
				startId = *orderArr[i].Id
			}

			cc,err := new(models.CardClass).GetById(*orderArr[i].ClassTp)
			if err != nil {
				ZapLog().Error("Nofind CardClass ", zap.Error(err), zap.Int("classTp", *orderArr[i].ClassTp))
				continue
			}
			if cc == nil || cc.ThirdLongChain == nil ||orderArr[i].Phone == nil{
				continue
			}

			//生成短链，发送短信
			reShortChain := &baidu_api.ReDwzCreate{
				Url : *cc.ThirdLongChain, //这个还得改改
				TermOfValidity: "1-year",
			}
			shortChain,err := reShortChain.Send()
			if err != nil {
				ZapLog().Error("baidu_api.ReDwzCreate ",zap.Error(err),  zap.Int("classTp", *orderArr[i].ClassTp))
				continue
			}

			if err := sdk.GNotifySdk.SendSms([]string{shortChain}, *orderArr[i].Phone, "new_unfinish",0, &platformTp); err != nil {
				ZapLog().Error("GNotifySdk.SendSms err", zap.Error(err))
				continue
			}

			time.Sleep(time.Millisecond * 1)
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
	}

}
