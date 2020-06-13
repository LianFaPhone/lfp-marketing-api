package order

import (
	apibackend "LianFaPhone/lfp-api/errdef"
	. "LianFaPhone/lfp-base/log/zap"
	"LianFaPhone/lfp-marketing-api/api"
	. "LianFaPhone/lfp-marketing-api/controllers"
	"LianFaPhone/lfp-marketing-api/models"
	"fmt"
	"github.com/kataras/iris"
	"go.uber.org/zap"
	"time"
)

type CardOrder struct {
	//IdGener
	Controllers
}

func (this *CardOrder) BkList(ctx iris.Context) {
	this.bkSubList(ctx, nil)
}

func (this *CardOrder) BkListAll(ctx iris.Context) {
	this.bkSubList(ctx, nil)
}

func (this *CardOrder) BkListNew(ctx iris.Context) {
	status := models.CONST_OrderStatus_New
	this.bkSubList(ctx, &status)
}

func (this *CardOrder) BkListNewunfinish(ctx iris.Context) {
	status := models.CONST_OrderStatus_New_UnFinish
	this.bkSubList(ctx, &status)
}

func (this *CardOrder) BkListExport(ctx iris.Context) {
	status := models.CONST_OrderStatus_Already_Export
	this.bkSubList(ctx, &status)
}

func (this *CardOrder) BkListDeliver(ctx iris.Context) {
	status := models.CONST_OrderStatus_Already_Delivered
	this.bkSubList(ctx, &status)
}

func (this *CardOrder) BkListWaitdone(ctx iris.Context) {
	status := models.CONST_OrderStatus_Wait_Done
	this.bkSubList(ctx, &status)
}

func (this *CardOrder) BkListAlreadydone(ctx iris.Context) {
	status := models.CONST_OrderStatus_Already_Done
	this.bkSubList(ctx, &status)
}

func (this *CardOrder) BkListRecyclebin(ctx iris.Context) {
	status := models.CONST_OrderStatus_Recyclebin
	this.bkSubList(ctx, &status)
}

func (this *CardOrder) BkListUnmatch(ctx iris.Context) {
	status := models.CONST_OrderStatus_UnMatch
	this.bkSubList(ctx, &status)
}

func (this *CardOrder) BkListActivated(ctx iris.Context) {
	status := models.CONST_OrderStatus_Already_Activated
	this.bkSubList(ctx, &status)
}

func (this *CardOrder) bkSubList(ctx iris.Context, status *int) {
	param := new(api.BkCardOrderList)

	err := Tools.ShouldBindJSON(ctx, param)
	if err != nil {
		this.ExceptionSerive(ctx, apibackend.BASERR_INVALID_PARAMETER.Code(), apibackend.BASERR_INVALID_PARAMETER.Desc())
		ZapLog().Error("param err", zap.Error(err))
		return
	}
	if param.Status == nil {
		param.Status = status
	}

	condPair := make([]*models.SqlPairCondition, 0, 15)
	if param.LikeStr != nil && len(*param.LikeStr) > 0 {
		condPair = append(condPair, &models.SqlPairCondition{"true_name like ?", "%" + *param.LikeStr + "%"})
		condPair = append(condPair, &models.SqlPairCondition{"order_no like ?", "%" + *param.LikeStr + "%"})
		condPair = append(condPair, &models.SqlPairCondition{"phone like ?", "%" + *param.LikeStr + "%"})
		condPair = append(condPair, &models.SqlPairCondition{"new_phone like ?", "%" + *param.LikeStr + "%"})
	}
	if param.StartCreatedAt != nil {
		condPair = append(condPair, &models.SqlPairCondition{"created_at >= ?", param.StartCreatedAt})
	}
	if param.EndCreatedAt != nil {
		condPair = append(condPair, &models.SqlPairCondition{"created_at <= ?", param.EndCreatedAt})
	}
	if param.StartDeliverAt != nil {
		condPair = append(condPair, &models.SqlPairCondition{"deliver_at >= ?", param.StartDeliverAt})
	}
	if param.EndDeliverAt != nil {
		condPair = append(condPair, &models.SqlPairCondition{"deliver_at <= ?", param.EndDeliverAt})
	}

	if param.StartActiveAt != nil {
		condPair = append(condPair, &models.SqlPairCondition{"active_at >= ?", param.StartActiveAt})
	}
	if param.EndActiveAt != nil {
		condPair = append(condPair, &models.SqlPairCondition{"active_at <= ?", param.EndActiveAt})
	}
	//extend := ctx.Values().GetString("extend")
	//if len(extend) > 0 {
	//	arr := strings.Split(extend, ",")
	//	condPair = append(condPair, &models.SqlPairCondition{"class_tp in (?)", arr})
	//}
	if param.UploadFlag != nil {
		if *param.UploadFlag == 0 {
			condPair = append(condPair, &models.SqlPairCondition{"dataurl1 == ?", "null"})
		} else if *param.UploadFlag == 1 {
			condPair = append(condPair, &models.SqlPairCondition{"dataurl1 != ?", "null"})
		}
	}

	limitGoodsArr,err := GetsGoodsByUserId(ctx.URLParam("limit_userid"))
	if err != nil {
		ZapLog().With(zap.Error(err)).Error("GetsGoodsByUserId err")
		this.ExceptionSerive(ctx, apibackend.BASERR_DATABASE_ERROR.Code(), apibackend.BASERR_DATABASE_ERROR.Desc())
		return
	}
	for i:=0; i < len(limitGoodsArr); i++ {
		condPair = append(condPair, &models.SqlPairCondition{"partner_goods_code = ?", limitGoodsArr[i].Code})
	}

	results, err := new(models.CardOrder).BkParseList(param).ListWithConds(param.Page, param.Size, nil, condPair)
	if err != nil {
		ZapLog().With(zap.Error(err)).Error("Verify err")
		this.ExceptionSerive(ctx, apibackend.BASERR_DATABASE_ERROR.Code(), apibackend.BASERR_DATABASE_ERROR.Desc())
		return
	}
	coArr := results.List.(*[]*models.CardOrder)
	for i := 0; i < len(*coArr); i++ {
		temp := (*coArr)[i]
		this.AttachName(temp)

		if param.BlackSwitch == nil || *param.BlackSwitch == 0 {
			this.CheckBlack(temp)
		}

	}
	this.Response(ctx, results)
}

func (this *CardOrder) BkUpdate(ctx iris.Context) {
	param := new(api.BkCardOrder)

	err := Tools.ShouldBindJSON(ctx, param)
	if err != nil {
		this.ExceptionSerive(ctx, apibackend.BASERR_INVALID_PARAMETER.Code(), apibackend.BASERR_INVALID_PARAMETER.Desc())
		ZapLog().Error("param err", zap.Error(err))
		return
	}

	err = new(models.CardOrder).BkParse(param).UpdateByOrderNo()
	if err != nil {
		ZapLog().With(zap.Error(err)).Error("Update err")
		this.ExceptionSerive(ctx, apibackend.BASERR_DATABASE_ERROR.Code(), apibackend.BASERR_DATABASE_ERROR.Desc())
		return
	}
	this.Response(ctx, nil)

}

func (this *CardOrder) BkGet(ctx iris.Context) {
	param := new(api.BkCardOrder)

	err := Tools.ShouldBindJSON(ctx, param)
	if err != nil {
		this.ExceptionSerive(ctx, apibackend.BASERR_INVALID_PARAMETER.Code(), apibackend.BASERR_INVALID_PARAMETER.Desc())
		ZapLog().Error("param err", zap.Error(err))
		return
	}

	res, err := new(models.CardOrder).BkParse(param).GetByOrderNo(*param.OrderNo, nil)
	if err != nil {
		ZapLog().With(zap.Error(err)).Error("Update err")
		this.ExceptionSerive(ctx, apibackend.BASERR_DATABASE_ERROR.Code(), apibackend.BASERR_DATABASE_ERROR.Desc())
		return
	}
	if res != nil {
		this.AttachName(res)

		res.CardOrderLog, _ = new(models.CardOrderLog).GetsByOrderNo(*param.OrderNo)
		res.CardIdcardPic, _ = new(models.CardIdcardPic).GetByOrderNo(param.OrderNo)

		if res.IdCard != nil && res.PartnerGoodsCode != nil{
			relateOrders, _ := new(models.CardOrder).GetsByIdcardAndPartnerGoods(*res.IdCard, *res.PartnerGoodsCode, nil, 50, &models.SqlPairCondition{fmt.Sprintf("created_at >= ? and id != %d", *res.Id), time.Now().Unix() - 30*24*3600})
			for i:=0; i< len(relateOrders);i++{
				this.AttachName(relateOrders[i])
			}
			res.RelateCardOrder = relateOrders
		}
	}
	this.Response(ctx, res)

}

func (this *CardOrder) BkUpdatesStatus(ctx iris.Context) {
	param := new(api.BkCardOrderStatusUpdates)

	err := Tools.ShouldBindJSON(ctx, param)
	if err != nil {
		this.ExceptionSerive(ctx, apibackend.BASERR_INVALID_PARAMETER.Code(), apibackend.BASERR_INVALID_PARAMETER.Desc())
		ZapLog().Error("param err", zap.Error(err))
		return
	}

	err = new(models.CardOrder).UpdatesStatusByOrderNo(param.OrderNo, *param.Status)
	if err != nil {
		ZapLog().With(zap.Error(err)).Error("Update err")
		this.ExceptionSerive(ctx, apibackend.BASERR_DATABASE_ERROR.Code(), apibackend.BASERR_DATABASE_ERROR.Desc())
		return
	}
	this.Response(ctx, nil)

}

/*********************************/
func (this *CardOrder) CheckBlack(temp *models.CardOrder)  {
	var areaBlack *models.BlacklistAreaCacheStt
	areaBlackFlag := false
	if temp.IsBacklist == nil && temp.Phone != nil {
		flag, err := new(models.BlacklistPhone).ExistByPhone(*temp.Phone)
		if err == nil {

		}
		if flag {
			temp.IsBacklist = new(int)
			*temp.IsBacklist = 1
			return
		}
	}
	if temp.IsBacklist == nil && temp.IdCard != nil {
		flag, err := new(models.BlacklistIdcard).ExistByIdCard(*temp.IdCard)
		if err == nil {

		}
		if flag {
			temp.IsBacklist = new(int)
			*temp.IsBacklist = 1
			return
		}
	}
	if !areaBlackFlag {
		areaBlackFlag = true
		var err error
		areaBlack, err = new(models.BlacklistArea).GetFromCache("")
		if err != nil {

		}
	}

	if areaBlack == nil {
		return
	}
	if temp.ProvinceCode != nil {
		_, ok := areaBlack.ProviceM[*temp.ProvinceCode]
		if ok {
			temp.IsBacklist = new(int)
			*temp.IsBacklist = 1
			return
		}
	}
	if temp.CityCode != nil {
		_, ok := areaBlack.CityM[*temp.CityCode]
		if ok {
			temp.IsBacklist = new(int)
			*temp.IsBacklist = 1
			return
		}
	}
	if temp.AreaCode != nil {
		_, ok := areaBlack.AreaM[*temp.AreaCode]
		if ok {
			temp.IsBacklist = new(int)
			*temp.IsBacklist = 1
			return
		}
	}
}

func (this *CardOrder) AttachName(temp *models.CardOrder)  {
	if temp == nil {
		return
	}
	if temp.AdTp != nil {
		v,ok := models.AdTpMap[*temp.AdTp]
		if ok {
			temp.AdTpName = &v
		}
	}

	if temp.PartnerGoodsCode != nil {
		cc,err := new(models.PdPartnerGoods).GetByCodeFromCache(*temp.PartnerGoodsCode)
		if err == nil && cc != nil{
			temp.PartnerGoodsName = cc.Name
		}
	}
	if temp.PhoneOSTp != nil {
		m, _ := models.PhoneOsMap[*temp.PhoneOSTp]
		if m != nil {
			temp.PhoneOSName = &m.Name
		}
	}
	if temp.Status != nil {
		m := models.OrderStatusMap[*temp.Status]
		temp.StatusName = &m
	}
}