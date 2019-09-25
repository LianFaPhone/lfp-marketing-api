package controllers

import(
	"github.com/kataras/iris"
	"LianFaPhone/lfp-marketing-api/models"
)



type PhoneOs struct{
	Controllers
}

func (this * PhoneOs) Gets(ctx iris.Context) {
	this.Response(ctx, models.PhoneOsArr)
}