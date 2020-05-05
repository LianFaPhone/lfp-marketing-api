package jthk

import (
	apibackend "LianFaPhone/lfp-api/errdef"
	"LianFaPhone/lfp-marketing-api/config"
	"LianFaPhone/lfp-marketing-api/db"
	"github.com/kataras/iris"
	"time"
)

func (this * Ydhk) SetTicket(ctx iris.Context) {
	ticket := ctx.URLParam("ticket")

	if len(ticket) <= 2 {
		 cmd := db.GRedis.GetConn().Get("ydjthk_ticket")
		if cmd.Err() != nil {
			this.ExceptionSerive(ctx, apibackend.BASERR_DATABASE_ERROR.Code(), "redis get err")
			return
		}

		config.GConfig.Lock.Lock()
		defer config.GConfig.Lock.Unlock()
		config.GConfig.Jthk.CkTicket = cmd.String()
		this.Response(ctx, nil)
		return
	}

	if err := db.GRedis.GetConn().Set("ydjthk_ticket", ticket, 10*24*time.Hour).Err(); err != nil {
		this.ExceptionSerive(ctx, apibackend.BASERR_DATABASE_ERROR.Code(), "redis get err")
		return
	}

	config.GConfig.Lock.Lock()
	defer config.GConfig.Lock.Unlock()

	config.GConfig.Jthk.CkTicket = ticket

	this.Response(ctx, nil)
}

func (this * Ydhk) GetTicket(ctx iris.Context) {

	config.GConfig.Lock.Lock()
	ticket := config.GConfig.Jthk.CkTicket
	config.GConfig.Lock.Unlock()

	this.Response(ctx, ticket)
}
