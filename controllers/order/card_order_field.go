package order

import (
	"LianFaPhone/lfp-marketing-api/models"
	"github.com/kataras/iris"
)


func (this *CardOrder) BkFields(ctx iris.Context) {
	this.Response(ctx, models.CardOrderFieldArr)
}
