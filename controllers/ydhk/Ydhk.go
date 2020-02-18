package ydhk

import (
	"LianFaPhone/lfp-marketing-api/api"
	"LianFaPhone/lfp-marketing-api/models"
	. "LianFaPhone/lfp-marketing-api/thirdcard-api/ydhk"
	"fmt"
	"go.uber.org/zap"
	"time"
)

import (
	. "LianFaPhone/lfp-marketing-api/controllers"
	"github.com/kataras/iris"
)
import (
	apibackend "LianFaPhone/lfp-api/errdef"
	. "LianFaPhone/lfp-base/log/zap"
)

type Ydhk struct{
	Controllers
}

func (this * Ydhk) GetProtocal(ctx iris.Context) {
	pCode := ctx.FormValue("province_code")
	token,err := new(ReProtocal).Send(pCode)
	if err != nil {
		ZapLog().With(zap.Error(err)).Error("Retoken send err")
		this.ExceptionSerive(ctx, apibackend.BASERR_INTERNAL_SERVICE_ACCESS_ERROR.Code(), err.Error())
		return
	}
	this.Response(ctx, token)
}

func (this * Ydhk) GetToken(ctx iris.Context) {
	token,err := new(ReToken).Send()
	if err != nil {
		ZapLog().With(zap.Error(err)).Error("Retoken send err")
		this.ExceptionSerive(ctx, apibackend.BASERR_INTERNAL_SERVICE_ACCESS_ERROR.Code(), err.Error())
		return
	}
	this.Response(ctx, &api.FtResYdhkToken{token})
}

func (this * Ydhk) GetsAddr(ctx iris.Context) {
	ll, err := new(ReAddr).Send()
	if err != nil {
		ZapLog().With(zap.Error(err)).Error("Update err")
		this.ExceptionSerive(ctx, apibackend.BASERR_INTERNAL_SERVICE_ACCESS_ERROR.Code(), err.Error())
		return
	}
	this.Response(ctx, ll)
}

func (this * Ydhk) ListNumberPool(ctx iris.Context) {
	param := new(api.FtYdhkNumberPoolList)
	if err := ctx.ReadJSON(param); err != nil {
		this.ExceptionSerive(ctx, apibackend.BASERR_INVALID_PARAMETER.Code(), apibackend.BASERR_INVALID_PARAMETER.Desc())
		ZapLog().Error("param err", zap.Error(err))
		return
	}
	if param.Page >=1 {
		param.Page = param.Page -1
	}

	ll,err := new(ReCardSearch).Send(param.ProviceCode, param.Provice, param.CityCode, param.City, param.Searchkey, param.Page, param.Size)
	if err != nil {
		ZapLog().With(zap.Error(err)).Error("Retoken send err")
		this.ExceptionSerive(ctx, apibackend.BASERR_INTERNAL_SERVICE_ACCESS_ERROR.Code(), err.Error())
		return
	}
	this.Response(ctx, ll)
}

func (this * Ydhk) LockNumber(ctx iris.Context) {
	param := new(api.FtYdhkNumberPoolLock)
	if err := ctx.ReadJSON(param); err != nil {
		this.ExceptionSerive(ctx, apibackend.BASERR_INVALID_PARAMETER.Code(), apibackend.BASERR_INVALID_PARAMETER.Desc())
		ZapLog().Error("param err", zap.Error(err))
		return
	}

	flag,err := new(ReCloseNumber).Send(param.ProviceCode, param.CityCode,  param.Number, param.Token)
	if err != nil {
		ZapLog().With(zap.Error(err)).Error("Retoken send err")
		this.ExceptionSerive(ctx, apibackend.BASERR_CARDMARKET_PHONEPOOL_LOCK_FAIL.Code(), err.Error())
		return
	}
	if !flag {
		this.ExceptionSerive(ctx, apibackend.BASERR_CARDMARKET_PHONEPOOL_LOCK_FAIL.Code(), apibackend.BASERR_CARDMARKET_PHONEPOOL_LOCK_FAIL.OriginDesc())
		return
	}
	this.Response(ctx, nil)
}

func (this * Ydhk) Apply(ctx iris.Context) {

	//FtYdhkApply
	param := new(api.FtYdhkApply)
	if err := ctx.ReadJSON(param); err != nil {
		this.ExceptionSerive(ctx, apibackend.BASERR_INVALID_PARAMETER.Code(), apibackend.BASERR_INVALID_PARAMETER.Desc())
		ZapLog().Error("param err", zap.Error(err))
		return
	}

	errCode, orderId,oaoFlag,err := new(ReOrderSubmit).Send(param.AccessToken, param.Phone,  param.NewPhone, param.LeagalName, param.CertificateNo, param.Address, param.Province, param.City, param.SendProvince, param.SendCity,param.SendDistrict)
	if err != nil {
		ZapLog().With(zap.Error(err)).Error("Retoken send err")
		this.ExceptionSerive(ctx, errCode.Code(), err.Error())
		return
	}

	go func(){
		orderNo := fmt.Sprintf("D%s%d", time.Now().Format("060102030405000"), GIdGener.Gen())
		modelParam := &models.CardOrder{
			OrderNo:  &orderNo,
			TrueName: &param.LeagalName,
			IdCard:   &param.CertificateNo,
			ClassTp:  param.ClassTp,
			//		CountryCode: p.CountryCode,
			Phone:        &param.Phone,
			//Province:     &param.Province,
			ProvinceCode: &param.Province,
			//City:         p.City,
			CityCode:     &param.City,
			//Area:         &param.Area,
			AreaCode:     &param.SendDistrict,
			//Town:         p.Town,
			Address:      &param.Address,
			IP:           param.IP,
			PhoneOSTp:    param.PhoneOSTp,
			ClassIsp: param.ClassISP,
			NewPhone: &param.NewPhone,
			ThirdOrderNo: &orderId,
		}

		modelParam.Valid = new(int)
		*modelParam.Valid = 1

		modelParam.Status = new(int)
		if oaoFlag == true {
			*modelParam.Status =	models.CONST_OrderStatus_New
		}else{
			*modelParam.Status =	models.CONST_OrderStatus_New_UnFinish
		}

		if err := modelParam.Add(); err != nil {
			//		tx.Rollback()
			ZapLog().With(zap.Error(err)).Error("database err")
			this.ExceptionSerive(ctx, apibackend.BASERR_DATABASE_ERROR.Code(), apibackend.BASERR_DATABASE_ERROR.Desc())
			return
		}

	}()

	this.Response(ctx, &api.FtResYdhkApply{orderId,oaoFlag})
}


func (this * Ydhk) UploadPhoto(ctx iris.Context) {
	//ll, err := new(models.CardClass).Gets()
	//if err != nil {
	//	ZapLog().With(zap.Error(err)).Error("Update err")
	//	this.ExceptionSerive(ctx, apibackend.BASERR_DATABASE_ERROR.Code(), apibackend.BASERR_DATABASE_ERROR.Desc())
	//	return
	//}
	//this.Response(ctx, ll)
}

func (this *Ydhk) FtConfirm(ctx iris.Context) {
	param := new(api.FtCardOrderConfirm)

	err := Tools.ShouldBindJSON(ctx, param)
	if err != nil {
		this.ExceptionSerive(ctx, apibackend.BASERR_INVALID_PARAMETER.Code(), apibackend.BASERR_INVALID_PARAMETER.Desc())
		ZapLog().Error("param err", zap.Error(err))
		return
	}
	if (param.SuccFlag != nil) && (*param.SuccFlag !=0) {
		err = new(models.CardOrder).UpdateStatusByOrderNo(param.OrderNo, models.CONST_OrderStatus_New)
		if err != nil {
			ZapLog().With(zap.Error(err)).Error("database err")
			this.ExceptionSerive(ctx, apibackend.BASERR_DATABASE_ERROR.Code(), apibackend.BASERR_DATABASE_ERROR.Desc())
			return
		}
	}


	if param.Log != nil && len(*param.Log) > 1 {
		go func (){
			new(models.CardOrderLog).FtParseAdd(nil, &param.OrderNo, param.Log).Add()
		}()
	}
	this.Response(ctx, nil)
}