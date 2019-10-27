package order

import (
	. "LianFaPhone/lfp-marketing-api/controllers"
	"LianFaPhone/lfp-marketing-api/models"
	"github.com/kataras/iris"
)

type OrderStatus struct {
	Controllers
}

func (this *OrderStatus) Gets(ctx iris.Context) {
	this.Response(ctx, models.OrderStatusArr)
}
