package pusher

//
//import "LianFaPhone/lfp-marketing-api/models"
//import (
//	"encoding/json"
//	"time"
//	"go.uber.org/zap"
//	"bytes"
//	"sync"
//	. "LianFaPhone/bas-base/log/zap"
//	"LianFaPhone/lfp-marketing-api/common"
//	"fmt"
//	"LianFaPhone/lfp-marketing-api/db"
//)
//
//type Response struct {
//	Code    int         `json:"code"`
//	Message string      `json:"message"`
//	Data    interface{} `json:"data,omitempty"`
//}
//
//type PushMsg struct{
//	MsgType  int      `json:"msg_type"`
//	Data    interface{} `json:"data,omitempty"`
//}
//
//const CONST_MAX_WORKER_NUM = 30
//const(
//	CONST_PUSH_EVENT_CARDORDER = 1
//)
//
////1--10 裂变
//const CONST_FISSION_ACTIVITY = 1
//const CONST_FISSION_Red = 2
//const CONST_FISSION_Robber = 3
//
////11-20 是抽奖
//
//var GTasker Tasker
//type Tasker struct{
//	mPushRobCh chan int
//	mRunFlag   bool
//	sync.WaitGroup
//}
//
//func (this *Tasker) Init() {
//	this.mRunFlag = true
//	this.mPushRobCh = make(chan int, 100)
//}
//
//func (this *Tasker) Start() {
//	this.Add(2)
//	//go this.runRobber()
//	go this.run()
//}
//
//func (this *Tasker) Stop() {
//	this.mRunFlag = false
//	close(this.mPushRobCh)
//	this.Wait()
//}
//
//func (this *Tasker) PushEvent(event int) error {
//	if event <= 0 {
//		return nil
//	}
//
//	select{
//	case this.mPushRobCh <- event:
//	default:
//		return nil
//	}
//
//	return nil
//}
//
////func (this *Tasker) runRobber() {
////	defer this.Done()
////	for {
////		select{
////			case r,ok :=<- this.mPushRobCh:
////				if !ok {
////					return
////				}
////				data,err := json.Marshal(r)
////				if err != nil {
////					ZapLog().Error("json Marshal err", zap.Error(err))
////					continue
////				}
////				notifyMsg := new(models.NotifyMsg).ParseAdd(models.CONST_NOTIFY_MSG_TYPE_ROB, string(data))
////				if err := notifyMsg.Add(); err != nil {
////					ZapLog().Error("NotifyMsg Add err", zap.Error(err))
////					continue
////				}
////				if !this.mRunFlag {
////					continue
////				}
////				select{
////				case this.mCh <- notifyMsg:
////				default:
////					ZapLog().Warn("notify Chan is full so lost", zap.String("notifyMsg", string(data)))
////				}
////
////		}
////	}
////}
//
//func (this *Tasker) run() {
//	defer this.Done()
//	defer models.PanicPrint()
//	//delTicker := time.NewTicker(time.Hour * 1)
//	sheetTicker := time.NewTicker(time.Hour * 1)
//	smsTicker := time.NewTicker(time.Second*600)
//	bspTicker := time.NewTicker(time.Second*600)
//	wait := new(sync.WaitGroup)
//	count := 0
//	for {
//		var msg *models.NotifyMsg
//		var ok bool
//		select{
//			//case <- delTicker.C:
//			//	this.delHander()  //定时清理垃圾
//			//	ZapLog().Debug("time to del")
//			//定时计算昨天的数据
//			//case msg, ok = <- this.mPushRobCh:
//			//	//批量推送
//			//	ZapLog().Debug("chan recved")
//			//	if !ok {
//			//		return
//			//	}
//			//	count++
//			//	wait.Add(1)
//			//	go func(){
//			//		defer wait.Done()
//			//		this.chHander(msg)
//			//	}()
//			//	if count >= CONST_MAX_WORKER_NUM {
//			//		wait.Wait()
//			//		count = 0
//			//	}
//			case <-smsTicker.C:
//				ZapLog().Debug("time to ticker")
//				wait.Wait()
//				count = 0
//				this.smsTickHander()
//			case <-bspTicker.C:
//				ZapLog().Debug("time to ticker")
//				wait.Wait()
//				count = 0
//				this.smsTickHander()
//			case <-sheetTicker.C:
//				ZapLog().Debug("time to ticker")
//				wait.Wait()
//				count = 0
//				this.smsTickHander()
//		}
//
//	}
//}
//
////func (this *Tasker) delHander() {
////	defer models.PanicPrint()
////	conds := []*models.SqlPairCondition{
////		&models.SqlPairCondition{"created_at <= ?", time.Now().Unix() - 3*24*3600},
////	}
////	if err := new(models.NotifyMsg).Del(conds); err != nil {
////		ZapLog().Error("Del err", zap.Error(err))
////		return
////	}
////}
//
////短信补发机制，只补发10分钟到2小时之间的订单而且 只补发1次
//func (this *Tasker) smsTickHander() {
//	defer models.PanicPrint()
//	nowTime := time.Now().Unix()
//	conds := []*models.SqlPairCondition{
//		&models.SqlPairCondition{"created_at >= ?", nowTime - 7200},
//		&models.SqlPairCondition{"created_at <= ?", nowTime - 600},
//		&models.SqlPairCondition{"sms_flag = ?", 0},
//	}
//	ZapLog().Debug("start tickHander")
//	//批量推送
//	for ;true; {
//		msgs,err := new(models.CardOrder).GetLimitByCond(CONST_MAX_WORKER_NUM, conds)
//		if err != nil {
//			ZapLog().Error("GetsByLimit err", zap.Error(err))
//			return
//		}
//		if len(msgs) == 0 {
//			ZapLog().Debug("GetsByLimit nil")
//			return
//		}
//		wait := new(sync.WaitGroup)
//		for j:=0; j < len(msgs);j++ {
//			tempJ := j
//			if msgs[tempJ].Content == nil || msgs[tempJ].Id == nil {
//				continue
//			}
//			wait.Add(1)
//			go func(){
//				defer wait.Done()
//				tx := db.GDbMgr.Get().Begin()
//				affectedRows,err := new(models.NotifyMsg).TxSetPushFlag(tx, *msgs[tempJ].Id)
//				if err != nil {
//					ZapLog().Error("TxSetPushFlag err", zap.Error(err))
//					tx.Rollback()
//					return
//				}
//				if affectedRows <= 0 {
//					tx.Rollback()
//					return
//				}
//				flag := this.callRobber(msgs[tempJ])
//				if !flag {
//					ZapLog().Error("callRobber faild")
//					tx.Rollback()
//				}else{
//					tx.Commit()
//				}
//				if err := new(models.NotifyMsg).IncrTryCount(*msgs[tempJ].Id); err != nil {
//					ZapLog().Error("IncrTryCount err", zap.Error(err))
//					return
//				}
//
//				ZapLog().Debug("call ok "+fmt.Sprintf("%d", tempJ))
//			}()
//		}
//		wait.Wait()
//		ZapLog().Debug("wait ok")
//		if len(msgs) < CONST_MAX_WORKER_NUM {
//			ZapLog().Info("GetsByLimit < 50 stop")
//			return
//		}
//	}
//}
//
//func (this *Tasker) chHander(event int) {
//	defer models.PanicPrint()
//	if msg == nil || msg.Content == nil || msg.Id == nil{
//		return
//	}
//	if msg.TryCount != nil && *msg.TryCount >= 3 {
//		return
//	}
//	tx := db.GDbMgr.Get().Begin()
//	affectedRows,err := new(models.NotifyMsg).TxSetPushFlag(tx, *msg.Id)
//	if err != nil {
//		ZapLog().Error("TxSetPushFlag err", zap.Error(err))
//		tx.Rollback()
//		return
//	}
//	if affectedRows <= 0 {
//		tx.Rollback()
//		return
//	}
//	flag := this.callRobber(msg)
//	if !flag {
//		tx.Rollback()
//		ZapLog().Error("callRobber faild")
//		*msg.TryCount = *msg.TryCount+1
//		this.mCh <- msg
//	}else{
//		tx.Commit()
//	}
//	if err := new(models.NotifyMsg).IncrTryCount(*msg.Id); err != nil {
//		ZapLog().Error("IncrTryCount err", zap.Error(err))
//		return
//	}
//}
//
//func (this *Tasker) callRobber(nfy *models.NotifyMsg) bool {
//	robMsg := new(models.Robber)
//	if err := json.Unmarshal([]byte(*nfy.Content), robMsg); err != nil {
//		ZapLog().Error("json Unmarshal err", zap.Error(err))
//		return false
//	}
//
//	acty, err := new(models.Activity).GetByUuidFromCache(*robMsg.ActivityUUId)
//	if err != nil {
//		ZapLog().Error("Activity GetById err", zap.Error(err))
//		return false
//	}
//	if acty == nil {
//		ZapLog().Error("Activity GetById nofind")
//		return true
//	}
//	if acty.Valid == nil || *acty.Valid == 0 {
//		ZapLog().Info("no need push")
//		return true
//	}
//	//spr, err := new(models.Sponsor).GetSponsorByIdFromCache(*acty.SponsorId)
//	//if err != nil {
//	//	ZapLog().Error("Sponsor GetSponsorByIdFromCache err", zap.Error(err))
//	//	return false
//	//}
//	//
//	//if spr.Valid == nil || *spr.Valid == 0 {
//	//	ZapLog().Info("no need push")
//	//	return true
//	//}
//	if acty.NotifyUrl == nil || len(*acty.NotifyUrl) < 2 {
//		ZapLog().Error("NotifyUrl is nil, no need push")
//		return true
//	}
//	robMsg.SponsorAccount = acty.SponsorAccount
//	robMsg.OffAt = acty.OffAt
//	robMsg.ApiKey = new(string)
//	robMsg.Language = acty.Language
//	robMsg.ApiKey = new(string)
//	*robMsg.ApiKey = "70123b81-9754-4fd4-b2cd-ecac5d3947cd"
//	headers := make(map[string] string)
//	headers["Api-Key"] = "70123b81-9754-4fd4-b2cd-ecac5d3947cd"
//	headers["Req-Real-Ip"] = ""
//
//	//pushMag := &PushMsg{
//	//	MsgType:CONST_FISSION_Robber,
//	//	Data: robMsg,
//	//}
//	data,err := json.Marshal(robMsg)
//	if err != nil {
//		ZapLog().Error("json Marshal err", zap.Error(err))
//		return false
//	}
//	ZapLog().Info(" push ", zap.Int("notifymsg_id", *nfy.Id), zap.Int("robber_id", *robMsg.Id), zap.String("content", string(data)))
//	body := bytes.NewReader(data)
//	resData, err := common.HttpSend(*acty.NotifyUrl, body,"POST", headers)
//	if err != nil {
//		ZapLog().Error("HttpSend err", zap.Error(err))
//		return false
//	}
//	//ZapLog().Info("push ",zap.Int("notifymsg_id", *nfy.Id), zap.Int("robber_id", *msg.Id))
//	response := new(Response)
//	if err := json.Unmarshal(resData, response); err != nil {
//		ZapLog().Error("json UnMarshal Response err", zap.Error(err))
//		return false
//	}
//	if response.Code != 0 {
//		ZapLog().Error(" Response code not 0 err:"+response.Message)
//		return false
//	}
//	ZapLog().Info(" push success ", zap.Int("notifymsg_id", *nfy.Id), zap.Int("robber_id", *robMsg.Id), zap.String("content", string(data)))
//	return true
//}
