package controllers

import(
	"github.com/kataras/iris"
	"LianFaPhone/lfp-marketing-api/models"
)


type ClassTp struct{
	Controllers
}

func (this * ClassTp) Gets(ctx iris.Context){
	this.Response(ctx, models.ClassTpArr)
}