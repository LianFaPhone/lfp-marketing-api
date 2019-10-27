package main

import (
	"LianFaPhone/lfp-marketing-api/controllers"
	. "LianFaPhone/lfp-marketing-api/controllers/order"
	"github.com/iris-contrib/middleware/cors"
	"github.com/kataras/iris"
)

func (this *WebServer) bkroutes() {
	app := this.mBkIris
	crs := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowedMethods:   []string{"POST", "GET", "OPTIONS", "PUT", "DELETE"},
		AllowedHeaders:   []string{"Authorization", "X-Requested-With", "X_Requested_With", "Content-Type", "Access-Token", "Accept-Language", "Api-Key", "Req-Real-Ip"},
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
		phoneNumberPy := v1bk.Party("/phone_number")
		{
			ac := new(controllers.PhoneNumberPool)

			phoneNumberPy.Post("/adds", ac.Adds)
			phoneNumberPy.Post("/update", ac.Update)
			phoneNumberPy.Post("/get", ac.Get)
			phoneNumberPy.Post("/list", ac.BkList)
		}
		phoneNumberLvPy := v1bk.Party("/phone_number_level")
		{
			ac := new(controllers.PhoneNumberLevel)

			phoneNumberLvPy.Post("/add", ac.Add)
			phoneNumberLvPy.Post("/update", ac.Update)
			phoneNumberLvPy.Post("/list", ac.List)
		}
		osphonePy := v1bk.Party("/osphone")
		{
			ac := new(controllers.PhoneOs)

			osphonePy.Post("/gets", ac.Gets)
		}
		simPy := v1bk.Party("/card_order")
		{
			ac := new(CardOrder)

			simPy.Post("/list", ac.BkList)
			simPy.Post("/list-all", ac.BkList)
			simPy.Post("/list-new", ac.BkList)
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
		datePy := v1bk.Party("/datesheet")
		{
			ac := new(CardDateSheet)

			datePy.Post("/list", ac.BkList)
		}
		areaPy := v1bk.Party("/areasheet")
		{
			ac := new(CardAreaSheet)

			areaPy.Post("/list", ac.BkList)
		}
		classPy := v1bk.Party("/cardclass")
		{
			ac := new(controllers.ClassTp)

			classPy.Post("/gets", ac.Gets)
			classPy.Post("/add", ac.Add)
			classPy.Post("/get", ac.Get)
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
		}

	}
}
