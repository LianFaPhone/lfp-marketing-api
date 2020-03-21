package jthk

import (
	"LianFaPhone/lfp-marketing-api/models"
	"strings"
	apibackend "LianFaPhone/lfp-api/errdef"
)

func (this *Ydhk) ParseFailStatus(errMsg string) int {
	if strings.Contains(errMsg, "欠费号码") {
		return models.CONST_OrderStatus_Fail_Retry
	}else if strings.Contains(errMsg, "已超时") {
		return models.CONST_OrderStatus_Fail_Retry
	}else if strings.Contains(errMsg, "系统错误") {
		return models.CONST_OrderStatus_Fail_Retry
	}else if strings.Contains(errMsg, "订购的号码不存在") {
		return models.CONST_OrderStatus_Fail_Retry
	}else if strings.Contains(errMsg, "号码已被占用") {
		return models.CONST_OrderStatus_Fail_Retry
	}else if strings.Contains(errMsg, "系统忙") {
		return models.CONST_OrderStatus_Fail_Retry
	}else if strings.Contains(errMsg, "无法申请新号卡") {
		return models.CONST_OrderStatus_Fail_Retry
	}else if strings.Contains(errMsg, "锁定号码失败") {
		return models.CONST_OrderStatus_Fail_Retry
	}else if strings.Contains(errMsg, "获取新号码失败"){
		return models.CONST_OrderStatus_Fail_Retry
	}
	return models.CONST_OrderStatus_Fail
}

func (this *Ydhk) ParseExcetionCode(oldCode apibackend.EnumBasErr, errMsg string)  apibackend.EnumBasErr {
	if strings.Contains(errMsg, "系统错误") {
		return apibackend.BASERR_CARDMARKET_PHONECARD_APPLY_HelpUser
	}else  if strings.Contains(errMsg, "系统忙") {
		return apibackend.BASERR_CARDMARKET_PHONECARD_APPLY_HelpUser
	}else if strings.Contains(errMsg, "获取新号码失败"){
		return apibackend.BASERR_CARDMARKET_PHONECARD_APPLY_HelpUser
	} else if strings.Contains(errMsg, "无法申请新号卡") {
		return apibackend.BASERR_CARDMARKET_PHONECARD_FastAPPLY_Dxnbhk
	}else if strings.Contains(errMsg, "欠费号码") {
		return apibackend.BASERR_CARDMARKET_PHONECARD_FastAPPLY_Dxnbhk
	}
	return oldCode
}
