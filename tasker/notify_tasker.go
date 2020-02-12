package tasker

import (
	. "LianFaPhone/lfp-base/log/zap"
	"LianFaPhone/lfp-marketing-api/config"
	"LianFaPhone/lfp-marketing-api/models"
	"LianFaPhone/lfp-marketing-api/sdk"
	"go.uber.org/zap"
	"time"
)

var GNotifyTasker NotifyTasker

type NotifyTasker struct {
	mCH           chan bool
	mNotifyTicker *time.Ticker
	mDelTicker    *time.Ticker
}

func (this *NotifyTasker) Init() {
	this.mCH = make(chan bool, 2)
	this.mNotifyTicker = time.NewTicker(time.Minute * 5)
	this.mDelTicker = time.NewTicker(time.Minute * 10)
}

func (this *NotifyTasker) Start() {
	return
	go func() {
		for true {
			this.run()
		}
	}()
}

func (this *NotifyTasker) Push() {
	select {
	case this.mCH <- true:
	default:

	}
}

func (this *NotifyTasker) run() {
	defer models.PanicPrint()
	select {
	case <-this.mCH:
		this.notifyWork()
	case <-this.mNotifyTicker.C:
		this.notifyWork()
	case <-this.mDelTicker.C:
		this.clear()
	}

}

func (this *NotifyTasker) clear() {
	conds := []*models.SqlPairCondition{
		&models.SqlPairCondition{"created_at <= ?", time.Now().Unix() - 3600*24*30},
	}

	if err := new(models.CardOrderNotify).Del(conds); err != nil {
		ZapLog().Error("CardOrderNotify Del err", zap.Error(err))
	}
}

func (this *NotifyTasker) notifyWork() {
	defer models.PanicPrint()
	conds := []*models.SqlPairCondition{
		&models.SqlPairCondition{"try_count < ?", 2},
		&models.SqlPairCondition{"last_at <= ?", time.Now().Unix() - 600},
		&models.SqlPairCondition{"push_flag = ?", 0},
	}

	//批量推送
	for true {
		msgs, err := new(models.CardOrderNotify).GetsByLimit(conds, 20)
		if err != nil {
			ZapLog().Error("GetsByLimit err", zap.Error(err))
			return
		}
		if len(msgs) == 0 {
			ZapLog().Debug("GetsByLimit nil")
			return
		}
		for j := 0; j < len(msgs); j++ {

			if msgs[j].Tp == nil || msgs[j].OrderNo == nil {
				continue
			}
			order, err := new(models.CardOrder).GetByOrderNo(*msgs[j].OrderNo)
			if err != nil {
				ZapLog().Error("GetByOrderNo err", zap.Error(err))
				return
			}
			if order == nil {
				continue
			}
			if order.Valid != nil && *order.Valid == 0 {
				continue
			}
			if config.GConfig.BasNotify.SwitchFlag {
				switch *msgs[j].Tp {
				case models.CONST_OrderNotifyTp_Express:
					param := []string{*order.Express, *order.ExpressNo}
					if err := sdk.GNotifySdk.SendSms(param, *order.Phone, "express_no", 0); err != nil {
						ZapLog().Error("GNotifySdk.SendSms err", zap.Error(err))
						if err := new(models.CardOrderNotify).IncrTryCount(*msgs[j].Id); err != nil {
							ZapLog().Error("IncrTryCount err", zap.Error(err))
						}
						continue
					}

				default:

				}
			}

			if err := new(models.CardOrderNotify).SetPushFlag(*msgs[j].Id); err != nil {
				ZapLog().Error("IncrTryCount err", zap.Error(err))
			}

		}

	}

}
