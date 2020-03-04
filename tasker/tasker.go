package tasker

import (
	. "LianFaPhone/lfp-base/log/zap"
	"LianFaPhone/lfp-marketing-api/config"
	"LianFaPhone/lfp-marketing-api/models"
	"go.uber.org/zap"
	"os/exec"
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

	jthkNewUnfinishNotifyTicker := time.NewTicker(time.Minute * 3)
	//cardOrderUnFinishSms5MinTicker.Stop()
	jthkFailNotifyTicker := time.NewTicker(time.Minute * 60)
	//cardOrderUnFinishSms1HourTicker.Stop()

	ydhkUnFinishSmallCheckTicker := time.NewTicker(time.Minute * 3)
	if !config.GConfig.Task.YdhkUnfinishFlag {
		ydhkUnFinishSmallCheckTicker.Stop()
	}
	ydhkUnFinishCheckTicker := time.NewTicker(time.Minute * 20)
	if !config.GConfig.Task.YdhkUnfinishFlag {
		ydhkUnFinishCheckTicker.Stop()
	}

	ydhkExpressTicker := time.NewTicker(time.Minute * 15)

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

	//oao测试 和 发短信 都有时效性要求，超过时间就无效了。所以性能不够得时候得加机器
	go func() {//15分钟处理不完就失效了
		defer models.PanicPrint()
		for {
			select {
			case <-jthkNewUnfinishNotifyTicker.C:
				this.jthkNewUnFinishNotify()
			}
		}
	}()

	go func() {
		defer models.PanicPrint()
		for {
			select {
			case <-jthkFailNotifyTicker.C:
				 this.jthkFailNotify()
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
			case <- ydhkUnFinishCheckTicker.C:
				this.ydhkOaoWork("ydhk_oao", 3610, true)
			}
		}
	}()

	go func() { //超过 1分钟500笔oao订单后需要调整
		defer models.PanicPrint()
		for {
			select {
			case <-ydhkUnFinishSmallCheckTicker.C:
				this.ydhkOaoWork("ydhk_small_oao", 10*60, false)
			}
		}
	}()

	go func() {
		defer models.PanicPrint()
		for {
			select {
			case <-ydhkExpressTicker.C:
				this.ydhkExpressWork() //24小时
			}
		}
	}()

	go func() {
		defer models.PanicPrint()
		for {
			//return
			hour := time.Now().Hour()
			if hour >=2 && hour <= 5 {
				this.clearWork()
			}
			if hour == 5 {
				cmd := exec.Command("rm", "-f", config.GConfig.Server.FilePath +"/2*.xlsx")
				cmd.Run()
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

