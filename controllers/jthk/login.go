package jthk

import (
	apibackend "LianFaPhone/lfp-api/errdef"
	"LianFaPhone/lfp-marketing-api/config"
	"github.com/kataras/iris"
)

func (this * Ydhk) SetTicket(ctx iris.Context) {
	ticket := ctx.URLParam("ticket")

	if len(ticket) <= 2 {
		this.ExceptionSerive(ctx, apibackend.BASERR_TOKEN_INVALID.Code(), "ticket invalid")
		return
	}

	config.GConfig.Lock.Lock()
	defer config.GConfig.Lock.Unlock()

	config.GConfig.Jthk.CkTicket = ticket

	this.Response(ctx, nil)
}
