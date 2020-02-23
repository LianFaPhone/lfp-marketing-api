package order

import (
	apibackend "LianFaPhone/lfp-api/errdef"
	. "LianFaPhone/lfp-base/log/zap"
	"LianFaPhone/lfp-marketing-api/api"
	"LianFaPhone/lfp-marketing-api/common"
	"LianFaPhone/lfp-marketing-api/models"
	. "LianFaPhone/lfp-marketing-api/controllers"
	"strings"
	"time"

	//"LianFaPhone/lfp-marketing-api/sdk"
	"LianFaPhone/lfp-marketing-api/tasker"
	//"fmt"
	"github.com/kataras/iris"
	"go.uber.org/zap"
	//"strings"
	//"time"
)

//还缺个发送短信的功能
func (this *CardOrder) BkOrderExtraInport(ctx iris.Context) {
	params := make([]*api.BkCardOrderExtraImprot, 0)

	err := ctx.ReadJSON(&params)
	//err := Tools.ShouldBindJSON(ctx, param)
	if err != nil {
		this.ExceptionSerive(ctx, apibackend.BASERR_INVALID_PARAMETER.Code(), apibackend.BASERR_INVALID_PARAMETER.Desc())
		ZapLog().Error("param err", zap.Error(err))
		return
	}

	res := new(api.BkResCardOrderExtraImprot)
	notifyTp := models.CONST_OrderNotifyTp_Express
	for i := 0; i < len(params); i++ {
		aff, err := new(models.CardOrder).BkParseExtraImport(params[i]).UpdatesByOrderNo()
		if err != nil {
			ZapLog().With(zap.Error(err)).Error("Update err")
			this.ExceptionSerive(ctx, apibackend.BASERR_DATABASE_ERROR.Code(), apibackend.BASERR_DATABASE_ERROR.Desc())
			return
		}
		if aff > 0 {
			res.SuccCount += 1
		}
		if err = new(models.CardOrderNotify).Add(params[i].OrderNo, nil, &notifyTp); err != nil {
			ZapLog().With(zap.Error(err)).Error("CardOrderNotify.Add err")
		}
	}
	if res.SuccCount > 0 {
		tasker.GNotifyTasker.Push()
	}

	res.FailCount = len(params) - res.SuccCount

	this.Response(ctx, res)
}

func (this *CardOrder) BkOrderActiveInport(ctx iris.Context) {
	params := make([]*api.BkCardOrderActiveImprot, 0)

	err := ctx.ReadJSON(&params)
	//err := Tools.ShouldBindJSON(ctx, param)
	if err != nil {
		this.ExceptionSerive(ctx, apibackend.BASERR_INVALID_PARAMETER.Code(), apibackend.BASERR_INVALID_PARAMETER.Desc())
		ZapLog().Error("param err", zap.Error(err))
		return
	}

	res := new(api.BkResCardOrderExtraImprot)
	//	notifyTp := models.CONST_OrderNotifyTp_Express
	newPhones := make([]*string, 0, 50)
	for i := 0; i < len(params); i++ {
		newPhones = append(newPhones, params[i].NewPhone)
		if i !=0 && (i%50 == 0 || i == len(params) -1){
			aff,err := new(models.CardOrder).UpdatesStatusByNewphone(newPhones, models.CONST_OrderStatus_Already_Activated)
			if err != nil {
				ZapLog().With(zap.Error(err)).Error("Update err")
				this.ExceptionSerive(ctx, apibackend.BASERR_DATABASE_ERROR.Code(), apibackend.BASERR_DATABASE_ERROR.Desc())
				return
			}
			if aff > 0 {
				res.SuccCount += int(aff)
			}
			newPhones = make([]*string, 0, 50)
		}
	}
	if len(newPhones) > 0 {
		aff,err := new(models.CardOrder).UpdatesStatusByNewphone(newPhones, models.CONST_OrderStatus_Already_Activated)
		if err != nil {
			ZapLog().With(zap.Error(err)).Error("Update err")
			this.ExceptionSerive(ctx, apibackend.BASERR_DATABASE_ERROR.Code(), apibackend.BASERR_DATABASE_ERROR.Desc())
			return
		}
		if aff > 0 {
			res.SuccCount += int(aff)
		}
	}

	res.FailCount = len(params) - res.SuccCount

	this.Response(ctx, res)
}

func (this *CardOrder) BkOrderNewInport(ctx iris.Context) {
	params := make([]*api.BkCardOrderNewImport, 0)

	err := ctx.ReadJSON(&params)
	//err := Tools.ShouldBindJSON(ctx, param)
	if err != nil {
		this.ExceptionSerive(ctx, apibackend.BASERR_INVALID_PARAMETER.Code(), apibackend.BASERR_INVALID_PARAMETER.Desc())
		ZapLog().Error("param err", zap.Error(err))
		return
	}

	res := new(api.BkResCardOrderExtraImprot)
	//	notifyTp := models.CONST_OrderNotifyTp_Express
	for i := 0; i < len(params); i++ {
		//同一个套餐，同一个身份证三个月内出现过就用同一个
		modelParam := &models.CardOrder{
			IdCard: params[i].IdCard,
			Phone:  params[i].Phone,
		}
		order,err := modelParam.GetByIdcardAndPhone(nil)
		if err != nil {

		}
		if order != nil {
			continue
		}


	}

	res.FailCount = len(params) - res.SuccCount

	this.Response(ctx, res)
}

func (this *CardOrder) BkOrderExpressInport(ctx iris.Context) {
	files,err := ParseFiles(ctx)
	if err != nil {
		ZapLog().Error( "ParseFiles err", zap.Error(err), zap.Any("headers", ctx.Request().Header))
		this.ExceptionSerive(ctx, apibackend.BASERR_INVALID_PARAMETER.Code(), "ParseMultipartForm_ERRORS")
		return
	}

	res := new(api.BkResCardOrderExtraImprot)
	for i := 0; i < len(files); i++ {
		//filename := files[i].Filename

		file, err := files[i].Open()
		if err != nil {
			ZapLog().Error("FileOpen  err", zap.Error(err))
			this.ExceptionSerive(ctx, apibackend.BASERR_SYSTEM_INTERNAL_ERROR.Code(), "FileOpen_ERRORS")
			return
		}
		defer file.Close()

		if files[i].Size <= 0 {
			this.ExceptionSerive(ctx, apibackend.BASERR_INVALID_PARAMETER.Code(), "nil data")
			return
		}
		heads,sheet,err := common.ReadFromReader(file, files[i].Size, "")
		if err != nil {
			ZapLog().Error( "ReadFromReader err", zap.Error(err))
			this.ExceptionSerive(ctx, apibackend.BASERR_INVALID_PARAMETER.Code(), "readfromReader fail")
			return
		}
		orderIndex := -1
		expressIndex := -1
		iccidIndex := -1
		expressNoIndex := -1
		newPhoneIndex := -1
		guishudiIndex := -1
		for j:=0; j < len(heads);j++ {
			head := strings.Replace(heads[j], " ", "", -1)
			if strings.Contains(head, "订单号") {
				orderIndex = j
			}else if strings.Contains(head, "快递公司") {
				expressIndex = j
			}else if strings.Contains(head, "快递单号") {
				expressNoIndex = j
			}else if strings.Contains(head, "ICCID") {
				iccidIndex = j
			}else if strings.Contains(head, "新号码") {
				newPhoneIndex = j
			}else if strings.Contains(head, "归属地") {
				guishudiIndex = j
			}
		}

		deliverAt := time.Now().Unix()
		for i:=0; i < len(sheet.Rows); i++ {
			if i == 0 {
				continue
			}
			row := sheet.Rows[i]
			param := new(api.BkCardOrderExtraImprot)

			if orderIndex >= 0 &&(len(row.Cells) -1 >= orderIndex){
				str := row.Cells[orderIndex].String()
				str = strings.Replace(str, " ", "", -1)
				if len(str) < 5 {
					res.FailCount +=1
					continue
				}
				param.OrderNo = &str
			}
			if expressIndex >= 0  &&(len(row.Cells) -1 >= expressIndex){
				str := row.Cells[expressIndex].String()
				str = strings.Replace(str, " ", "", -1)
				if len(str) > 1 {
					param.Express = &str
				}
			}
			if expressNoIndex >= 0 &&(len(row.Cells) -1 >= expressNoIndex) {
				str := row.Cells[expressNoIndex].String()
				str = strings.Replace(str, " ", "", -1)
				if len(str) > 2 {
					param.ExpressNo = &str
				}
			}
			if guishudiIndex >= 0  &&(len(row.Cells) -1 >= guishudiIndex){
				str := row.Cells[guishudiIndex].String()
				str = strings.Replace(str, " ", "", -1)
				if len(str) > 1 {
					param.Guishudi = &str
				}
			}
			if iccidIndex >= 0  &&(len(row.Cells) -1 >= iccidIndex){
				str := row.Cells[iccidIndex].String()
				str = strings.Replace(str, " ", "", -1)
				if len(str) > 2 {
					param.ICCID = &str
				}
			}
			if newPhoneIndex >= 0  &&(len(row.Cells) -1 >= newPhoneIndex){
				str := row.Cells[newPhoneIndex].String()
				str = strings.Replace(str, " ", "", -1)
				if len(str) > 2 {
					param.NewPhone = &str
				}
			}
			if param.OrderNo == nil {
				res.FailCount +=1
				continue
			}
			param.DeliverAt = &deliverAt
			aff, err := new(models.CardOrder).BkParseExtraImport(param).UpdatesByOrderNo()
			if err != nil {
				ZapLog().With(zap.Error(err)).Error("Update err")
				//this.ExceptionSerive(ctx, apibackend.BASERR_DATABASE_ERROR.Code(), apibackend.BASERR_DATABASE_ERROR.Desc())
				res.FailCount +=1
				//return
				continue
			}else{
				if aff > 0 {
					res.SuccCount += 1
				}
			}
		}
	}

	this.Response(ctx, res)
}
