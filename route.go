package main

import (
	"LianFaPhone/lfp-marketing-api/controllers"
	. "LianFaPhone/lfp-marketing-api/controllers/area"
	. "LianFaPhone/lfp-marketing-api/controllers/order"
	"github.com/iris-contrib/middleware/cors"
	"github.com/kataras/iris"
)

func (this *WebServer) routes() {
	app := this.mIris
	crs := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowedMethods:   []string{"POST", "GET", "OPTIONS", "PUT", "DELETE"},
		AllowedHeaders:   []string{"*" /*,"Access-Control-Allow-Origin","Authorization", "X-Requested-With", "X_Requested_With", "Content-Type", "Access-Token", "Accept-Language", "Api-Key", "Req-Real-Ip"*/},
		ExposedHeaders:   []string{"Access-Control-Allow-Origin"},
		AllowCredentials: true,
	})

	app.Any("/", func(ctx iris.Context) {
		ctx.JSON(
			map[string]interface{}{
				"code": 0,
			})
	})
	//interceptor := new(controllers.Interceptor)


	//oldSaveDataPy := app.Party("/saveData", crs /*, interceptor.Interceptor*/)
	//{
	//	oldCO := new(CardOrder)
	//	{
	//		//oldSaveDataPy.Post("/", oldCO.OldApply)
	//	}
	//}

	jsPy := app.Party("/js", crs /*, interceptor.Interceptor*/)
	{
		ac := new(ChinaAddrCode)
		{
			jsPy.Get("/list.json", ac.Gets2)
			jsPy.Get("/town/{param:path}", ac.GetStreet2)
		}
	}

	v1 := app.Party("/v1/ft/market", crs /*, interceptor.Interceptor*/).AllowMethods(iris.MethodOptions)
	{
		//活动，添加，list，更新，警用 全是后端的
		cardClassParty := v1.Party("/cardclass")
		{
			cc := new(controllers.ClassTp)
			{
				cardClassParty.Get("/get", cc.FtGet)
			}
		}

		cardParty := v1.Party("/card")
		{
			ac := new(CardOrder)
			{
				cardParty.Post("/apply", ac.Apply)
				cardParty.Post("/order/confirm", ac.FtConfirm)

				cardParty.Post("/apply-fulfil", ac.Apply)
				cardParty.Any("/verify", ac.FtUpdatePhotoUrls)
			}
			pp := new(controllers.PhoneNumberPool)
			{
				cardParty.Post("/numberpool/get", pp.Get)
				cardParty.Post("/numberpool/number-check", pp.NumberCheck)
				cardParty.Post("/numberpool/list", pp.FtList)
			}
			pl := new(controllers.PhoneNumberLevel)
			{
				cardParty.Post("/numberpool_level/gets", pl.FtGets)
			}

		}
		poolParty := v1.Party("/numberpool")
		{
			pp := new(controllers.PhoneNumberPool)
			{
				poolParty.Post("/get", pp.Get)
				poolParty.Post("/number-check", pp.NumberCheck)
				poolParty.Post("/list", pp.FtList)
				poolParty.Post("/lock", pp.FtLock)
				poolParty.Post("/unlock", pp.FtUnLock)
			}
			pl := new(controllers.PhoneNumberLevel)
			{
				cardParty.Post("/level/gets", pl.FtGets)
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
		//photosupload
		photoParty := v1.Party("/photo")
		{
			ac := new(controllers.UploadFile)

			photoParty.Any("/upload", ac.PhotoUpload)
		}

	}

}
