package class

import (
	"LianFaPhone/lfp-marketing-api/api"
	"LianFaPhone/lfp-marketing-api/common"
	"LianFaPhone/lfp-marketing-api/config"
	"LianFaPhone/lfp-marketing-api/controllers"
	"LianFaPhone/lfp-marketing-api/models"
	"github.com/kataras/iris"
	"go.uber.org/zap"
)
import (
	apibackend "LianFaPhone/lfp-api/errdef"
	. "LianFaPhone/lfp-base/log/zap"
)

type PdPartnerGoods struct {
	controllers.Controllers
}

func (this *PdPartnerGoods) Gets(ctx iris.Context) {

	ll, err := new(models.PdPartnerGoods).Gets()
	if err != nil {
		ZapLog().With(zap.Error(err)).Error("Update err")
		this.ExceptionSerive(ctx, apibackend.BASERR_DATABASE_ERROR.Code(), apibackend.BASERR_DATABASE_ERROR.Desc())
		return
	}
	this.Response(ctx, ll)

}


func (this *PdPartnerGoods) Add(ctx iris.Context) {
	param := new(api.BkCardClassAdd)

	err := controllers.Tools.ShouldBindJSON(ctx, param)
	if err != nil {
		this.ExceptionSerive(ctx, apibackend.BASERR_INVALID_PARAMETER.Code(), apibackend.BASERR_INVALID_PARAMETER.Desc())
		ZapLog().Error("param err", zap.Error(err))
		return
	}

	if param.Code == nil {
		param.Code = new(string)
		*param.Code = common.GetRandomString(10)
	}

	partner,err := new(models.PdPartner).GetById(*param.BigTp)
	if err != nil {
		ZapLog().With(zap.Error(err)).Error("db err")
		this.ExceptionSerive(ctx, apibackend.BASERR_DATABASE_ERROR.Code(), apibackend.BASERR_DATABASE_ERROR.Desc())
		return
	}
	if partner == nil {
		ZapLog().With(zap.Error(err)).Error("nofind partner err")
		this.ExceptionSerive(ctx, apibackend.BASERR_OBJECT_NOT_FOUND.Code(), apibackend.BASERR_OBJECT_NOT_FOUND.Desc())
		return
	}
	param.LongChain = new(string)
	if partner.PrefixPath != nil && len(*partner.PrefixPath) > 0 && partner.Code !=nil {
		*param.LongChain += config.GConfig.Server.LfcxHost+ *partner.PrefixPath +"/"+ *partner.Code +"/"+ *param.Code
	}

	modelParam := new(models.PdPartnerGoods).ParseAdd(param)
	flag,err := modelParam.Unique()
	if err != nil {
		ZapLog().With(zap.Error(err)).Error("db err")
		this.ExceptionSerive(ctx, apibackend.BASERR_DATABASE_ERROR.Code(), apibackend.BASERR_DATABASE_ERROR.Desc())
		return
	}
	if !flag {
		this.ExceptionSerive(ctx, apibackend.BASERR_OBJECT_EXISTS.Code(), apibackend.BASERR_OBJECT_EXISTS.Desc())
		return
	}


	ll, err := modelParam.Add()
	if err != nil {
		ZapLog().With(zap.Error(err)).Error("Update err")
		this.ExceptionSerive(ctx, apibackend.BASERR_DATABASE_ERROR.Code(), apibackend.BASERR_DATABASE_ERROR.Desc())
		return
	}
	this.Response(ctx, ll)

}

func (this *PdPartnerGoods) Get(ctx iris.Context) {
	//设置套餐，图片，上传文件
	param := new(api.BkCardClass)

	err := controllers.Tools.ShouldBindJSON(ctx, param)
	if err != nil {
		this.ExceptionSerive(ctx, apibackend.BASERR_INVALID_PARAMETER.Code(), apibackend.BASERR_INVALID_PARAMETER.Desc())
		ZapLog().Error("param err", zap.Error(err))
		return
	}

	ll, err := new(models.PdPartnerGoods).Parse(param).Get()
	if err != nil {
		ZapLog().With(zap.Error(err)).Error("Update err")
		this.ExceptionSerive(ctx, apibackend.BASERR_DATABASE_ERROR.Code(), apibackend.BASERR_DATABASE_ERROR.Desc())
		return
	}
	this.Response(ctx, ll)
}


func (this *PdPartnerGoods) FtGet(ctx iris.Context) {
	//设置套餐，图片，上传文件
	className := ctx.FormValue("name")
	if len(className) == 0 {
		this.ExceptionSerive(ctx, apibackend.BASERR_INVALID_PARAMETER.Code(), apibackend.BASERR_INVALID_PARAMETER.Desc())
		ZapLog().Error("param err")
		return
	}
	//ZapLog().Sugar().Info("test=", zap.Any("haha", ctx.Request().Header))

	cc, err := new(models.PdPartnerGoods).GetByCodeFromCache(className)
	if err != nil {
		ZapLog().With(zap.Error(err)).Error("Update err")
		this.ExceptionSerive(ctx, apibackend.BASERR_DATABASE_ERROR.Code(), apibackend.BASERR_DATABASE_ERROR.Desc())
		return
	}
	if cc == nil {
		this.ExceptionSerive(ctx, apibackend.BASERR_OBJECT_NOT_FOUND.Code(), apibackend.BASERR_OBJECT_NOT_FOUND.Desc())
		return
	}

	dd := &models.PdPartnerGoods{
		ISP: cc.ISP,
		BigTp: cc.BigTp,
		//Tp: cc.Tp,
		Name: cc.Name,
		ImgUrl: cc.ImgUrl,
		SmsFlag: cc.SmsFlag,
	}

	this.Response(ctx, dd)
}


func (this *PdPartnerGoods) Update(ctx iris.Context) {
	//设置套餐，图片，上传文件
	param := new(api.BkCardClass)

	err := controllers.Tools.ShouldBindJSON(ctx, param)
	if err != nil {
		this.ExceptionSerive(ctx, apibackend.BASERR_INVALID_PARAMETER.Code(), apibackend.BASERR_INVALID_PARAMETER.Desc())
		ZapLog().Error("param err", zap.Error(err))
		return
	}

	ll, err := new(models.PdPartnerGoods).Parse(param).Update()
	if err != nil {
		ZapLog().With(zap.Error(err)).Error("Update err")
		this.ExceptionSerive(ctx, apibackend.BASERR_DATABASE_ERROR.Code(), apibackend.BASERR_DATABASE_ERROR.Desc())
		return
	}
	this.Response(ctx, ll)
}

func (this *PdPartnerGoods) List(ctx iris.Context) {
	//设置套餐，图片，上传文件
	param := new(api.BkCardClassList)

	err := controllers.Tools.ShouldBindJSON(ctx, param)
	if err != nil {
		this.ExceptionSerive(ctx, apibackend.BASERR_INVALID_PARAMETER.Code(), apibackend.BASERR_INVALID_PARAMETER.Desc())
		ZapLog().Error("param err", zap.Error(err))
		return
	}

	ll, err := new(models.PdPartnerGoods).ParseList(param).ListWithConds(param.Page, param.Size, nil, nil)
	if err != nil {
		ZapLog().With(zap.Error(err)).Error("Update err")
		this.ExceptionSerive(ctx, apibackend.BASERR_DATABASE_ERROR.Code(), apibackend.BASERR_DATABASE_ERROR.Desc())
		return
	}
	this.Response(ctx, ll)
}