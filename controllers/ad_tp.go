package controllers

import (
	"LianFaPhone/lfp-marketing-api/models"
	"github.com/kataras/iris"
)

type AdTp struct {
	Controllers
}

func (this *AdTp) Gets(ctx iris.Context) {
	this.Response(ctx, models.AdTpArr)
}

