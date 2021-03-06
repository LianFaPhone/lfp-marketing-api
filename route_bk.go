package main

import (
	"LianFaPhone/lfp-marketing-api/controllers"
	"LianFaPhone/lfp-marketing-api/controllers/area"
	"LianFaPhone/lfp-marketing-api/controllers/black"
	"LianFaPhone/lfp-marketing-api/controllers/class"
	"LianFaPhone/lfp-marketing-api/controllers/jthk"
	. "LianFaPhone/lfp-marketing-api/controllers/order"
	"LianFaPhone/lfp-marketing-api/controllers/phonepool"
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
		ydjthkPy := v1bk.Party("/ydjthk")
		{
			bsP := new(jthk.Ydhk)
			ydjthkPy.Any("/ticket/set", bsP.SetTicket)
		}
		areaPy := v1bk.Party("/area")
		{
			bsP := new(area.BsProvice)
			areaPy.Get("/bsprovince/gets", bsP.Gets)
			bsC := new(area.BsCity)
			areaPy.Get("/bscity/gets", bsC.Gets)
		}
		phoneNumberPy := v1bk.Party("/numberpool")
		{
			ac := new(phonepool.PhoneNumberPool)

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
			ac := new(phonepool.PhoneNumberLevel)

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
			cardOrderLogPy.Post("/gets", ac.BkGets)
		}
		//CardIdcardPic
		cardOrderPicPy := v1bk.Party("/card_order_idcardpic")
		{
			ac := new(CardOrderIdCardPic)

			cardOrderPicPy.Post("/get", ac.BkGet)
		}
		simPy := v1bk.Party("/card_order")
		{
			ac := new(CardOrder)

			simPy.Post("/get", ac.BkGet)
			simPy.Post("/list", ac.BkList)
			simPy.Post("/list-all", ac.BkListAll)
			simPy.Post("/list-new", ac.BkListNew)
			simPy.Post("/list-new-unfinish", ac.BkListNewunfinish)
			simPy.Post("/list-export", ac.BkListExport)
			simPy.Post("/export", ac.BkExport)
			simPy.Post("/file/create", ac.BkFileCreate)
			simPy.Get("/file/get", ac.BkFileGet)
			simPy.Post("/list-deliver", ac.BkListDeliver)
			simPy.Post("/list-waitdone", ac.BkListWaitdone)
			simPy.Post("/list-alreadydone", ac.BkListAlreadydone)
			simPy.Post("/list-recyclebin", ac.BkListRecyclebin)
			simPy.Post("/list-unmatch", ac.BkListUnmatch)
			simPy.Post("/list-activated", ac.BkListActivated)
			simPy.Post("/update", ac.BkUpdate)
			simPy.Post("/status-sets", ac.BkUpdatesStatus)
			simPy.Post("/express-import", ac.BkOrderExpressInport)
			simPy.Post("/idcard-check", ac.BkIdCardCheck)
			simPy.Post("/active-import", ac.BkOrderActiveInport)
			simPy.Post("/new-import", ac.BkOrderNewInport)
			simPy.Post("/retry-apply", ac.BkRetryApply)
			simPy.Post("/field/gets", ac.BkFields)
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
		classBigTpPy := v1bk.Party("/pdparter")
		{
			ac := new(class.PdPartner)

			classBigTpPy.Any("/gets", ac.Gets)
			classBigTpPy.Post("/add", ac.Add)
			classBigTpPy.Any("/get", ac.Get)
			classBigTpPy.Post("/update", ac.Update)
			classBigTpPy.Post("/list", ac.List)
			classBigTpPy.Post("/status-update", ac.UpdateStatus)
		}
		classPy := v1bk.Party("/pdpartergoods")
		{
			ac := new(class.PdPartnerGoods)

			classPy.Any("/gets", ac.Gets)
			classPy.Post("/add", ac.Add)
			classPy.Any("/get", ac.Get)
			classPy.Post("/update", ac.Update)
			classPy.Post("/list", ac.List)
			classPy.Post("/status-update", ac.UpdateStatus)
		}
		statusPy := v1bk.Party("/orderstatus")
		{
			ac := new(OrderStatus)

			statusPy.Post("/gets", ac.Gets)
		}
		adTpPy := v1bk.Party("/ad-tp")
		{
			ac := new(controllers.AdTp)

			adTpPy.Post("/gets", ac.Gets)
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
