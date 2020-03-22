package main

import (
	"LianFaPhone/lfp-marketing-api/config"
	//"BastionPay/bas-filetransfer-srv/api"
	"LianFaPhone/lfp-marketing-api/controllers"
	"LianFaPhone/lfp-marketing-api/db"
	"LianFaPhone/lfp-marketing-api/models"
	"math/rand"
	"time"

	"LianFaPhone/lfp-marketing-api/sdk"
	"LianFaPhone/lfp-marketing-api/tasker"
	"os"
)

func Loader() error {
	if err := db.GDbMgr.Init(&db.DbOptions{
		Host:        config.GConfig.Db.Host,
		Port:        config.GConfig.Db.Port,
		User:        config.GConfig.Db.User,
		Pass:        config.GConfig.Db.Pwd,
		DbName:      config.GConfig.Db.Dbname,
		MaxIdleConn: config.GConfig.Db.Max_idle_conn,
		MaxOpenConn: config.GConfig.Db.Max_open_conn,
	}); err != nil {
		return err
	}
	if err := db.GRedis.Init(config.GConfig.Redis.Host, config.GConfig.Redis.Port, config.GConfig.Redis.Password, config.GConfig.Redis.Database); err != nil {
		return err
	}
	models.InitDbTable()
	db.GCache.Init()
	db.GCache.SetBlacklistArea(new(models.BlacklistArea).InnerGetBy)
	db.GCache.SetPdPartnerGoodsByCode(new(models.PdPartnerGoods).InnerGetByCode)
	db.GCache.SetPdPartnerGoodsById(new(models.PdPartnerGoods).InnerGetById)
	db.GCache.SetPdPartnerById(new(models.PdPartner).InnerGetById)
	//db.GCache.SetProvinceByName(new(models.CardClass).InnerGetById)
	//db.GCache.SetShareInfoFunc(new(models.ShareInfo).InnerGetByAcId)
	//db.GCache.SetPageFunc(new(models.Page).InnerGetByAcId)
	//db.GCache.SetRobberFunc(new(models.Robber).InnerGetRedIdAndPhone)
	//db.GCache.SetPageShareInfoFunc(new(models.PageShareInfo).InnerGetByAcId)
	//db.GCache.SetSponsorAkFunc(new(models.Sponsor).InnerGetSponsorByAk)
	//db.GCache.SetFissionApiActyList(new(fisson_api.ReActivity).InnerGet)
	//db.GCache.SetActyYqlFunc()

	//db.GCache.SetUserLevelFunc(new(models.UserLevelSearch).InnerSearch)
	//db.GCache.SetLevelRuleFunc(new(models.RuleListSearch).InnerSearch)
	//pusher.GTasker.Init()
	//pusher.GTasker.Start()
	sdk.GNotifySdk.Init(config.GConfig.BasNotify.Addr)
	tasker.GTasker.Init()
	tasker.GTasker.Start()
	tasker.GNotifyTasker.Init()
	tasker.GNotifyTasker.Start()
	if err := controllers.Init(); err != nil {
		return err
	}
	os.MkdirAll("./photos",os.ModePerm)
	rand.Seed(time.Now().Unix())
	sdk.GYunPainSdk.Init(config.GConfig.YunPian.ApiKey)
	return nil
}

func UnLoader() {
	//pusher.GTasker.Stop()
}
