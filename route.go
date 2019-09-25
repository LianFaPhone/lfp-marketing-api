package main

import (
	"github.com/kataras/iris"
	"github.com/iris-contrib/middleware/cors"
	"LianFaPhone/lfp-marketing-api/controllers"
	. "LianFaPhone/lfp-marketing-api/controllers/order"
	. "LianFaPhone/lfp-marketing-api/controllers/area"
)

func (this *WebServer) routes()  {
	app := this.mIris
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


	v1 := app.Party("/v1/ft/market", crs/*, interceptor.Interceptor*/)
	{
			//活动，添加，list，更新，警用 全是后端的
		cardParty := v1.Party("/card")
		{
			ac := new(CardOrder)
			{
				cardParty.Post("/apply", ac.Apply)
			}
			pp := new(controllers.PhoneNumberPool)
			{
				cardParty.Post("/phone_number_pool/get", pp.Get)
				cardParty.Post("/phone_number_pool/number-check", pp.NumberCheck)
				cardParty.Post("/phone_number_pool/list", pp.FtList)
			}
			pl := new(controllers.PhoneNumberLevel)
			{
				cardParty.Post("/phone_number_level/gets", pl.FtGets)
			}

		}
		smsParty := v1.Party("/sms")
		{
			smsCtrl := new(controllers.Sms)

			smsParty.Post("/send", smsCtrl.Send)
			smsParty.Post("/verify", smsCtrl.Verify)
		}

		regionCodeParty := v1.Party("/addrcode")
		{
			ac := new(ChinaAddrCode)

			regionCodeParty.Post("/street/gets", ac.GetStreet)
			regionCodeParty.Post("/region/gets", ac.Gets)
		}

	}


}

