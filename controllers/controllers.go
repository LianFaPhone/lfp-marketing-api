package controllers

import (
	"LianFaPhone/lfp-marketing-api/common"
	"github.com/kataras/iris/context"
	//"gopkg.exa.center/blockshine-ex/api-article/config"
	"LianFaPhone/lfp-marketing-api/config"
	"LianFaPhone/lfp-marketing-api/db"
)

func Init() error {
	phoneLimiter = common.NewBusLimiter(&db.GRedis, "buslimit_sms_phone_", config.GPreConfig.PhoneSmsLimits)
	ipLimiter = common.NewBusLimiter(&db.GRedis, "buslimit_sms_ip_", config.GPreConfig.IpSmsLimits)
	if err := phoneLimiter.Init(); err != nil {
		return err
	}
	if err := ipLimiter.Init(); err != nil {
		return err
	}
	return nil
}

type (
	Controllers struct {
	}

	Response struct {
		Code    int         `json:"code"`
		Message string      `json:"message"`
		Data    interface{} `json:"data,omitempty"`
	}
)

var (
	Tools *common.Tools
)

func init() {
	Tools = common.New()
}

func (c *Controllers) Response(
	ctx context.Context,
	data interface{}) {

	ctx.JSON(
		Response{
			Code:    0,
			Message: "Success",
			Data:    data,
		})
}

func (c *Controllers) ExceptionSerive(
	ctx context.Context,
	code int,
	message string) {

	ctx.JSON(
		Response{
			Code:    code,
			Message: message,
		})
}

func (c *Controllers) ExceptionSeriveWithData(
	ctx context.Context,
	code int,
	message string,
	data interface{}) {

	ctx.JSON(
		Response{
			Code: code,
			Message: message,
			Data: data,
		})
}
