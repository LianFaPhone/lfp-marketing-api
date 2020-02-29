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
	param := new(api.BkPartnerGoodsAdd)

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

	partner,err := new(models.PdPartner).GetById(*param.PartnerId)
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
	param := new(api.BkPartnerGoods)

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
	className := ctx.FormValue("code")
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
	if cc == nil || (cc.Valid!=nil && *cc.Valid == 0) {
		this.ExceptionSerive(ctx, apibackend.BASERR_OBJECT_NOT_FOUND.Code(), apibackend.BASERR_OBJECT_NOT_FOUND.Desc())
		return
	}

	partner ,err := new(models.PdPartner).GetByIdFromCache(*cc.PartnerId)
	if err != nil {
		ZapLog().With(zap.Error(err)).Error("Update err")
		this.ExceptionSerive(ctx, apibackend.BASERR_DATABASE_ERROR.Code(), apibackend.BASERR_DATABASE_ERROR.Desc())
		return
	}
	if partner == nil || (partner.Valid!=nil && *partner.Valid == 0) {
		this.ExceptionSerive(ctx, apibackend.BASERR_OBJECT_NOT_FOUND.Code(), apibackend.BASERR_OBJECT_NOT_FOUND.Desc())
		return
	}

	newCC := &api.FtResPdPartnerGoodsGet{
		Code   : cc.Code,
		UrlParam : cc.UrlParam,
		ImgUrl :cc.ImgUrl,
		NoExpAddr: partner.NoExpAddr,
		MinAge: partner.MinAge,
		MaxAge: partner.MaxAge,
		SmsFlag: cc.SmsFlag,
		IdcardDispplay: cc.IdcardDispplay,

	}

	this.Response(ctx, newCC)
}


func (this *PdPartnerGoods) Update(ctx iris.Context) {
	//设置套餐，图片，上传文件
	param := new(api.BkPartnerGoods)

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
	param := new(api.BkPartnerGoodsList)

	err := controllers.Tools.ShouldBindJSON(ctx, param)
	if err != nil {
		this.ExceptionSerive(ctx, apibackend.BASERR_INVALID_PARAMETER.Code(), apibackend.BASERR_INVALID_PARAMETER.Desc())
		ZapLog().Error("param err", zap.Error(err))
		return
	}

	results, err := new(models.PdPartnerGoods).ParseList(param).ListWithConds(param.Page, param.Size, nil, nil)
	if err != nil {
		ZapLog().With(zap.Error(err)).Error("Update err")
		this.ExceptionSerive(ctx, apibackend.BASERR_DATABASE_ERROR.Code(), apibackend.BASERR_DATABASE_ERROR.Desc())
		return
	}
	coArr := results.List.(*[]*models.PdPartnerGoods)
	for i := 0; i < len(*coArr); i++ {
		temp := (*coArr)[i]
		if temp.PartnerId == nil {
			continue
		}
		pt,err := new(models.PdPartner).GetByIdFromCache(*temp.PartnerId)
		if err != nil {
			continue
		}
		if pt == nil {
			continue
		}
		temp.PartnerName = pt.Name

	}
	this.Response(ctx, results)
}

func (this *PdPartnerGoods) UpdateStatus(ctx iris.Context) {
	//设置套餐，图片，上传文件
	param := new(api.BkPartnerGoodsStatusUpdate)

	err := controllers.Tools.ShouldBindJSON(ctx, param)
	if err != nil {
		this.ExceptionSerive(ctx, apibackend.BASERR_INVALID_PARAMETER.Code(), apibackend.BASERR_INVALID_PARAMETER.Desc())
		ZapLog().Error("param err", zap.Error(err))
		return
	}

	ll, err := new(models.PdPartnerGoods).UpdateStatus(param.Id, param.Valid)
	if err != nil {
		ZapLog().With(zap.Error(err)).Error("Update err")
		this.ExceptionSerive(ctx, apibackend.BASERR_DATABASE_ERROR.Code(), apibackend.BASERR_DATABASE_ERROR.Desc())
		return
	}
	this.Response(ctx, ll)
}