package main

import (
	"github.com/kataras/iris"
	"github.com/iris-contrib/middleware/cors"
	"LianFaPhone/lfp-marketing-api/controllers"
	. "LianFaPhone/lfp-marketing-api/controllers/order"
)

func (this *WebServer) bkroutes()  {
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
			simPy.Post("/update", ac.BkUpdate)
			simPy.Post("/status-sets", ac.BkUpdatesStatus)
			simPy.Post("/extra-import", ac.BkOrderExtraInport)

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
		}
		statusPy := v1bk.Party("/orderstatus")
		{
			ac := new(OrderStatus)

			statusPy.Post("/gets", ac.Gets)
		}

	}
}