package controllers

import (
	"LianFaPhone/lfp-marketing-api/models"
	"github.com/kataras/iris"
)

type PhoneOs struct {
	Controllers
}

func (this *PhoneOs) Gets(ctx iris.Context) {
	this.Response(ctx, models.PhoneOsArr)
}
