package order


import (
	apibackend "LianFaPhone/lfp-api/errdef"
	. "LianFaPhone/lfp-base/log/zap"
	"LianFaPhone/lfp-marketing-api/api"
	"LianFaPhone/lfp-marketing-api/common"
	"LianFaPhone/lfp-marketing-api/config"
	"github.com/go-redis/redis"

	. "LianFaPhone/lfp-marketing-api/controllers"
	"LianFaPhone/lfp-marketing-api/db"
	"LianFaPhone/lfp-marketing-api/models"
	"fmt"
	"github.com/kataras/iris"
	"go.uber.org/zap"
	"time"
)


func (this *CardOrder) BkFileCreate(ctx iris.Context) {
	param := new(api.BkCardOrderList)

	err := Tools.ShouldBindJSON(ctx, param)
	if err != nil {
		this.ExceptionSerive(ctx, apibackend.BASERR_INVALID_PARAMETER.Code(), apibackend.BASERR_INVALID_PARAMETER.Desc())
		ZapLog().Error("param err", zap.Error(err))
		return
	}

	rd := common.RandomDigit(5)
	timeStr := time.Now().Format("2006-01-02-15-04-05")
	fileName := timeStr+"-"+rd +".xlsx"
	excel,err := common.NewExcel("sheet1", config.GConfig.Server.FilePath + "/" +fileName)
	if err != nil {
		this.ExceptionSerive(ctx, apibackend.BASERR_DATABASE_ERROR.Code(), apibackend.BASERR_DATABASE_ERROR.Desc())
		ZapLog().Error("create file err", zap.Error(err))
		return
	}
	_,err = db.GRedis.GetConn().Set("filecreate_"+fileName, "", time.Duration(time.Second*7200)).Result()
	if err != nil {
		ZapLog().With(zap.Error(err)).Error("redis set err")
		this.ExceptionSerive(ctx, apibackend.BASERR_DATABASE_ERROR.Code(), apibackend.BASERR_DATABASE_ERROR.Desc())

		return
	}
	this.Response(ctx, fileName)

	go func() {
		ZapLog().Info("BkFileCreate start "+fileName)
		if param.Page <= 0 {
			param.Page = 1
		}
		if param.Size <= 0 {
			param.Size = 400000
		}
		condPair := make([]*models.SqlPairCondition, 0, 15)
		if param.LikeStr != nil && len(*param.LikeStr) > 0 {
			condPair = append(condPair, &models.SqlPairCondition{"true_name like ?", "%" + *param.LikeStr + "%"})
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
		userId,_ := ctx.URLParamInt64("limit_userid")
		if userId > 0{
			limitGoodsArr,err :=new(models.PdPartnerGoods).GetsByUserId(userId)
			if err != nil {
				ZapLog().With(zap.Error(err)).Error("GetsGoodsByUserId err")
				this.ExceptionSerive(ctx, apibackend.BASERR_DATABASE_ERROR.Code(), apibackend.BASERR_DATABASE_ERROR.Desc())
				return
			}
			if len(limitGoodsArr) <= 0 {
				this.Response(ctx, nil)
			}
			for i:=0; i < len(limitGoodsArr); i++ {
				condPair = append(condPair, &models.SqlPairCondition{"partner_goods_code = ?", limitGoodsArr[i].Code})
			}
		}

		headers :=[] string{
			"序号","订单号","套餐名","订单状态","姓名","身份证","手机号","省","市","区县","镇街道","详细地址","新手机号","ICCID","归属地","快递","快递单号","发货时间","照片上传","黑名单","下单时间","来源",
		}

		if err := excel.AddHeader(headers); err != nil {
			ZapLog().Error("excel AddHeader err", zap.Error(err))
			return
		}
		total := param.Size
		newSize := int64(1000)
		AllPage := total/newSize
		condStr := ""
		modelParam := new(models.CardOrder).BkParseList(param)
		ZapLog().Info("BkFileCreate "+fileName, zap.Int64("pageall", AllPage))
		//needFields := GetSelectFields(ctx.URLParam("fields"))
		for page:= int64(0); page < AllPage ;page ++ {
			var results []*models.CardOrder
			if page == 0 {
				results, err = modelParam.GetsWithConds(param.Page, newSize, nil, condPair, condStr)
				if err != nil {
					ZapLog().With(zap.Error(err)).Error("Verify err")
					this.ExceptionSerive(ctx, apibackend.BASERR_DATABASE_ERROR.Code(), apibackend.BASERR_DATABASE_ERROR.Desc())
					return
				}
			}else{
				results, err = modelParam.GetsWithConds(1, newSize, nil, condPair, condStr)
				if err != nil {
					ZapLog().With(zap.Error(err)).Error("Verify err")
					this.ExceptionSerive(ctx, apibackend.BASERR_DATABASE_ERROR.Code(), apibackend.BASERR_DATABASE_ERROR.Desc())
					return
				}
			}
			if len(results) == 0 {
				break
			}


			coArr := &results
			minId := int64(0)
			for i := 0; i < len(*coArr); i++ {
				temp := (*coArr)[i]
				if i == len(*coArr) - 1 {
					minId = *temp.Id
				}
				appendData := make([]string, len(headers), len(headers))
				if temp.Id != nil {
					appendData[0] = fmt.Sprintf("%d", *temp.Id)
				}
				if temp.OrderNo != nil {
					appendData[1] = *temp.OrderNo
				}

				if temp.PartnerGoodsCode != nil {
					//ZapLog().Info("cardclass 1")
					cc,err := new(models.PdPartnerGoods).GetByCodeFromCache(*temp.PartnerGoodsCode)
					//ZapLog().Info("cardclass 2", zap.Error(err), zap.Any("cc",cc))
					if err == nil && cc != nil{
						temp.PartnerGoodsName = cc.Name
					}
				}
				if temp.PartnerGoodsName != nil {
					appendData[2] = *temp.PartnerGoodsName
				}

				if temp.Status != nil {
					m := models.OrderStatusMap[*temp.Status]
					temp.StatusName = &m
					if temp.StatusName != nil {
						appendData[3] = *temp.StatusName
					}
				}
				if temp.TrueName != nil {
					appendData[4] = *temp.TrueName
				}
				if temp.IdCard != nil {
					appendData[5] = *temp.IdCard
				}
				if temp.Phone != nil {
					appendData[6] = *temp.Phone
				}
				if temp.Province != nil {
					appendData[7] = *temp.Province
				}
				if temp.City != nil {
					appendData[8] = *temp.City
				}
				if temp.Area != nil {
					appendData[9] = *temp.Area
				}
				if temp.Town != nil {
					appendData[10] = *temp.Town
				}
				if temp.Address != nil {
					appendData[11] = *temp.Address
				}
				if temp.NewPhone != nil {
					appendData[12] = *temp.NewPhone
				}
				if temp.ICCID != nil {
					appendData[13] = *temp.ICCID
				}
				if temp.Guishudi != nil {
					appendData[14] = *temp.Guishudi
				}

				if temp.Express != nil {
					appendData[15] = *temp.Express
				}
				if temp.ExpressNo != nil {
					appendData[16] = *temp.ExpressNo
				}
				if temp.DeliverAt != nil {
					str := time.Unix(*temp.DeliverAt, 0).Format("2006-01-02 15:04")
					appendData[17] = str
				}
				if temp.IdCardPicFlag != nil {
					str := "否"
					if *temp.IdCardPicFlag == 1 {
						str = "是"
					}
					appendData[18] = str
				}
				if temp.CreatedAt != nil {
					str := time.Unix(*temp.CreatedAt, 0).Format("2006-01-02 15:04")//"2006-01-02 15:04-05"
					appendData[20] = str
				}

				if param.BlackSwitch != nil && *param.BlackSwitch == 1 {
					this.CheckBlack(temp)
				}
				if temp.IsBacklist != nil && *temp.IsBacklist == 1 {
					appendData[19] = "是"
				}
				if temp.AdTp != nil {
					switch *temp.AdTp {
					case models.CONST_ADTRACK_Tp_KuaiShou:
						appendData[21] = "快手"
						break
					case models.CONST_ADTRACK_Tp_DouYin:
						appendData[21] = "抖音"
						break
					default:
						appendData[21] = ""
						break
					}
				}
				if err := excel.AppendCache(appendData); err != nil {
					ZapLog().With(zap.Error(err)).Error("excel.AppendCache err")
					return
				}
				//if page %100 == 0 {
				//	if err:= excel.Sync(); err != nil {
				//		ZapLog().With(zap.Error(err)).Error("excel.Sync err")
				//		return
				//	}
				//}
			}
			if len(results) < int(newSize) {
				break
			}
			//if err:= excel.Sync(); err != nil {
			//	ZapLog().With(zap.Error(err)).Error("excel.Sync err")
			//	return
			//}
			condStr = fmt.Sprintf("id < %d", minId)
		}
		if err:= excel.Sync(); err != nil {
			ZapLog().With(zap.Error(err)).Error("excel.Sync err")
			return
		}
		ZapLog().Info("BkFileCreate success "+fileName)
		url := "http://file.lfcxwifi.com" + "/" + fileName
		_,err = db.GRedis.GetConn().Set("filecreate_"+fileName, url, time.Duration(time.Second*7200)).Result()
		if err != nil {
			ZapLog().With(zap.Error(err)).Error("redis set err")
		}
	}()
}

func (this *CardOrder) BkFileGet(ctx iris.Context) {
	fileName := ctx.URLParam("filename")

	if len(fileName) <= 0 {
		this.ExceptionSerive(ctx, apibackend.BASERR_INVALID_PARAMETER.Code(), apibackend.BASERR_INVALID_PARAMETER.Desc())
		ZapLog().Error("param err")
		return
	}

	url,err := db.GRedis.GetConn().Get("filecreate_"+fileName).Result()
	if err == redis.Nil {
		this.ExceptionSerive(ctx, apibackend.BASERR_OBJECT_NOT_FOUND.Code(), apibackend.BASERR_OBJECT_NOT_FOUND.Desc())
		ZapLog().Error("nofind")
		return
	}
	if err != nil {
		ZapLog().With(zap.Error(err)).Error("redis set err")
		this.ExceptionSerive(ctx, apibackend.BASERR_DATABASE_ERROR.Code(), apibackend.BASERR_DATABASE_ERROR.Desc())
		return
	}

	this.Response(ctx, url)
}

func (this *CardOrder) BkExport(ctx iris.Context) {
	this.bkSubList(ctx, nil)
}
