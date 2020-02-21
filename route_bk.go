package main

import (
	"LianFaPhone/lfp-marketing-api/controllers"
	"LianFaPhone/lfp-marketing-api/controllers/black"
	. "LianFaPhone/lfp-marketing-api/controllers/order"
	"LianFaPhone/lfp-marketing-api/controllers/sheet"
	"github.com/iris-contrib/middleware/cors"
	"github.com/kataras/iris"
)

func (this *WebServer) bkroutes() {
	app := this.mBkIris
	crs := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowedMethods:   []string{"POST", "GET", "OPTIONS", "PUT", "DELETE"},
		AllowedHeaders:   []string{"Access-Control-Allow-Origin","Authorization", "X-Requested-With", "X_Requested_With", "Content-Type", "Access-Token", "Accept-Language", "Api-Key", "Req-Real-Ip"},
		AllowCredentials: true,
	})

	app.Any("/", func(ctx iris.Context) {
		ctx.JSON(
			map[string]interface{}{
				"code": 0,
			})
	})
	//interceptor := new(controllers.Interceptor)

	v1bk := app.Party("/v1/bk/market", crs)
	{
		phoneNumberPy := v1bk.Party("/numberpool")
		{
			ac := new(controllers.PhoneNumberPool)

			phoneNumberPy.Post("/adds", ac.Adds)
			phoneNumberPy.Post("/update", ac.Update)
			phoneNumberPy.Any("/get", ac.Get)
			phoneNumberPy.Post("/list", ac.BkList)

			phoneNumberPy.Post("/lock", ac.BkLock)
			phoneNumberPy.Post("/unlock", ac.BkUnLock)
			phoneNumberPy.Post("/use", ac.BkUse)
			phoneNumberPy.Post("/unuse", ac.BkUnUse)

		}
		phoneNumberLvPy := v1bk.Party("/numberpool_level")
		{
			ac := new(controllers.PhoneNumberLevel)

			phoneNumberLvPy.Post("/add", ac.Add)
			phoneNumberLvPy.Post("/update", ac.Update)
			phoneNumberLvPy.Post("/list", ac.List)
		}
		osphonePy := v1bk.Party("/osphone")
		{
			ac := new(controllers.PhoneOs)

			osphonePy.Any("/gets", ac.Gets)
		}
		cardOrderLogPy := v1bk.Party("/card_order_log")
		{
			ac := new(CardOrderLog)

			cardOrderLogPy.Post("/list", ac.BkList)
		}
		simPy := v1bk.Party("/card_order")
		{
			ac := new(CardOrder)

			simPy.Post("/list", ac.BkList)
			simPy.Post("/list-all", ac.BkList)
			simPy.Post("/list-new", ac.BkList)
			simPy.Post("/list-new-unfinish", ac.BkList)
			simPy.Post("/list-export", ac.BkList)
			simPy.Post("/list-deliver", ac.BkList)
			simPy.Post("/list-waitdone", ac.BkList)
			simPy.Post("/list-alreadydone", ac.BkList)
			simPy.Post("/list-recyclebin", ac.BkList)
			simPy.Post("/list-unmatch", ac.BkList)
			simPy.Post("/list-activated", ac.BkList)
			simPy.Post("/update", ac.BkUpdate)
			simPy.Post("/status-sets", ac.BkUpdatesStatus)
			simPy.Post("/express-import", ac.BkOrderExtraInport)
			simPy.Post("/idcard-check", ac.BkIdCardCheck)
			simPy.Post("/active-import", ac.BkOrderActiveInport)
			simPy.Post("/new-import", ac.BkOrderNewInport)
		}

		classsheetPy := v1bk.Party("/classsheet")
		{
			ac := new(sheet.CardClassSheet)

			classsheetPy.Post("/list", ac.BkList)
		}
		areasheetPy := v1bk.Party("/areasheet")
		{
			ac := new(sheet.CardAreaSheet)

			areasheetPy.Post("/list", ac.BkList)
		}
		datesheetPy := v1bk.Party("/monthsheet")
		{
			ac := new(sheet.CardDateSheet)

			datesheetPy.Post("/list", ac.BkList)
		}
		classBigTpPy := v1bk.Party("/cardclassbigtp")
		{
			ac := new(controllers.ClassBigTp)

			classBigTpPy.Any("/gets", ac.Gets)
			classBigTpPy.Post("/add", ac.Add)
			classBigTpPy.Any("/get", ac.Get)
			classBigTpPy.Post("/update", ac.Update)
			classBigTpPy.Post("/list", ac.List)
		}
		classPy := v1bk.Party("/cardclasstp")
		{
			ac := new(controllers.ClassTp)

			classPy.Any("/gets", ac.Gets)
			classPy.Post("/add", ac.Add)
			classPy.Any("/get", ac.Get)
			classPy.Post("/update", ac.Update)
			classPy.Post("/list", ac.List)
		}
		statusPy := v1bk.Party("/orderstatus")
		{
			ac := new(OrderStatus)

			statusPy.Post("/gets", ac.Gets)
		}
		photoPy := v1bk.Party("/photo")
		{
			ac := new(controllers.UploadFile)

			photoPy.Any("/download", ac.PhotoDownload)
		}
		ossmarketPy := v1bk.Party("/oss-market")
		{
			ac := new(controllers.UploadFile)

			ossmarketPy.Any("/upload", ac.UpSso)
			ossmarketPy.Any("/cardclasspic/upload", ac.UpSso)
		}
		blacklistPhonePy := v1bk.Party("/blacklist/phone")
		{
			cc := new(black.BlacklistPhone)
			blacklistPhonePy.Post("/add", cc.Add)
			blacklistPhonePy.Post("/update", cc.Update)
			blacklistPhonePy.Post("/list", cc.List)
		}
		blacklistIdcardPy := v1bk.Party("/blacklist/idcard")
		{
			cc := new(black.BlacklistIdCard)
			blacklistIdcardPy.Post("/add", cc.Add)
			blacklistIdcardPy.Post("/update", cc.Update)
			blacklistIdcardPy.Post("/list", cc.List)
		}
		blacklistAreaPy := v1bk.Party("/blacklist/area")
		{
			cc := new(black.BlacklistArea)
			blacklistAreaPy.Post("/add", cc.Add)
			blacklistAreaPy.Post("/update", cc.Update)
			blacklistAreaPy.Post("/list", cc.List)
		}
		qaPy := v1bk.Party("/qa")
		{
			cc := new(controllers.QaCtrler)
			qaPy.Post("/add", cc.Add)
			qaPy.Post("/update", cc.Update)
			qaPy.Post("/list", cc.List)
		}
	}
}
