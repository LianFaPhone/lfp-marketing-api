package area

import (
	apibackend "LianFaPhone/lfp-api/errdef"
	. "LianFaPhone/lfp-base/log/zap"
	. "LianFaPhone/lfp-marketing-api/controllers"
	"LianFaPhone/lfp-marketing-api/models"
	"github.com/kataras/iris"
	"go.uber.org/zap"
)

type BsProvice struct {
	Controllers
}

func (this *BsProvice) Gets(ctx iris.Context) {
	ll, err := new(models.BsProvice).Gets()
	if err != nil {
		ZapLog().With(zap.Error(err)).Error("Gets err")
		this.ExceptionSerive(ctx, apibackend.BASERR_DATABASE_ERROR.Code(), apibackend.BASERR_DATABASE_ERROR.Desc())
		return
	}
	this.Response(ctx, ll)
}
