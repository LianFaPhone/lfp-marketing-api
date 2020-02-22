package tasker

import (
	. "LianFaPhone/lfp-base/log/zap"
	"LianFaPhone/lfp-marketing-api/config"
	"LianFaPhone/lfp-marketing-api/models"
	"go.uber.org/zap"
	"time"
)

/*
日期统计 每1小时一次，每次查下昨天的是否计算完，然后再算今天的.先全量计算，以后增量计算
"2006-01-02 15:04:05"
*/

var GTasker Tasker

type Tasker struct {
}

func (this *Tasker) Init() error {

	return nil
}

func (this *Tasker) Start() {
	if !config.GConfig.Task.Task_flag {
		return
	}
	go this.run()
}

func (this *Tasker) run() {
	//dateSheetTicker := time.NewTicker(time.Hour *2)
	//ipsTicker := time.NewTicker(time.Minute *10)
	activeCodeTicker := time.NewTicker(time.Hour * 10)
	activeCodeTicker.Stop()

	cardOrderSheetTicker := time.NewTicker(time.Second *  time.Duration(config.GConfig.Task.SheetTicker))
	if !config.GConfig.Task.SheetFlag {
		cardOrderSheetTicker.Stop()
	}
	cardOrderIpsTicker := time.NewTicker(time.Second * time.Duration(config.GConfig.Task.IpsTicker))
	if !config.GConfig.Task.IpsFlag {
		cardOrderIpsTicker.Stop()
	}

	cardOrderHistoryTicker := time.NewTicker(time.Hour * 6)
	cardOrderHistoryTicker.Stop()

	cardOrderUnFinishSms5MinTicker := time.NewTicker(time.Minute * 3)
	cardOrderUnFinishSms5MinTicker.Stop()
	cardOrderUnFinishSms1HourTicker := time.NewTicker(time.Minute * 60)
	cardOrderUnFinishSms1HourTicker.Stop()

	ydhkUnFinishCheckTicker := time.NewTicker(time.Minute * 3)
	if !config.GConfig.Task.YdhkUnfinishFlag {
		ydhkUnFinishCheckTicker.Stop()
	}

	go func() {
		defer models.PanicPrint()
		return
		for {
			busyflag:= true
			hour := time.Now().Hour()
			if hour >=1 && hour <= 5 {
				busyflag = false
			}
			select {
			case <-cardOrderSheetTicker.C:
				this.sheetWork(busyflag)
			case <-activeCodeTicker.C:
				//this.activeCodeWork()
			case <-cardOrderHistoryTicker.C:
				//go this.orderHistory()

			}
		}
	}()

	go func() {
		defer models.PanicPrint()
		for {
			select {
			case <-cardOrderUnFinishSms5MinTicker.C:
				//this.newUnFinishSmsWork5min()
			case <-cardOrderUnFinishSms1HourTicker.C:
				//this.newUnFinishSmsWork5hour()
			}
		}
	}()

	go func() {
		defer models.PanicPrint()
		for {
			return
			select {
			case <-cardOrderIpsTicker.C:
				busyflag:= true
				hour := time.Now().Hour()
				if hour >=1 && hour <= 5 {
					busyflag = false
				}
				this.ipsWork(busyflag)
			}
		}
	}()

	go func() {
		defer models.PanicPrint()
		for {
			select {
			case <-ydhkUnFinishCheckTicker.C:
				//this.ydhkWork()
			}
		}
	}()

	go func() {
		defer models.PanicPrint()
		for {
			return
			hour := time.Now().Hour()
			if hour >=2 && hour <= 5 {
				this.clearWork()
			}
			time.Sleep(time.Hour)
		}
	}()
}

func (this *Tasker) activeCodeWork() {
	defer models.PanicPrint()

	t := time.Now().Unix() - 3600*24*7
	conds := []*models.SqlPairCondition{
		&models.SqlPairCondition{"created_at <= ?", t},
	}

	if err := new(models.ActiveCode).Del(conds); err != nil {
		ZapLog().Error("ActiveCode del err", zap.Error(err))
		return
	}
}

