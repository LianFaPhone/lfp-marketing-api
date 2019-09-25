package order

import(
	. "LianFaPhone/lfp-marketing-api/controllers"
	"github.com/kataras/iris"
	"LianFaPhone/lfp-marketing-api/models"
)


type OrderStatus struct{
	Controllers
}

func (this * OrderStatus) Gets(ctx iris.Context){
	this.Response(ctx, models.OrderStatusArr)
}
