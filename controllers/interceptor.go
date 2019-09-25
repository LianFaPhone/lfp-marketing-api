package controllers

//import (
//	"github.com/kataras/iris"
//	"strings"
//	"LianFaPhone/bas-api/apibackend"
//)
//
//var PathWhiteList map[string] bool = map[string] bool{
//	//"/v1/ft/marketing/dispatch": true,
//	//"/v1/ft/marketing/pack": true,
//	"/v1/ft/fissionshare/sms/send":  true,
//	"/v1/ft/fissionshare/sms/verify":  true,
//	"/v1/ft/fissionshare/red/get": true,
//	"/v1/ft/fissionshare/activity/get": true,
//	"/v1/ft/fissionshare/robber/get": true,
//	"/v1/ft/fissionshare/robber/add": true,
//	"/v1/ft/fissionshare/robber/get/ten": true,
//
//	"/v1/ft/fissionshare/robber/search": true,       //这几个 要求apikey
//	"/v1/ft/fissionshare/robber/searchs": true,
//	"/v1/ft/fissionshare/robber/sets-transferflag": true,
//	"/v1/ft/fissionshare/robber/set-transferflag": true,
//
//	"/v1/ft/fissionshare/page/get": true,
//	"/v1/ft/fissionshare/share-info/get": true,
//	"/v1/ft/fissionshare/page-share-info/get": true,
//
//}
//
//type ApiKeyClaims struct{
//	ApiKey      string
//	ActyUuids   []*string
//}
//
//type Interceptor struct{
//	Controllers
//}
//
//func (this *Interceptor) Interceptor(ctx iris.Context) {
//	 _,ok := PathWhiteList[ctx.Path()];
//	if ok || strings.HasPrefix(ctx.Path(), "/v1/bk") {
//		ctx.Next()
//		return
//	}
//
//	apiKey := ctx.GetHeader("Api-Key")
//	if len(apiKey) == 0 {
//		this.ExceptionSerive(ctx, apibackend.BASERR_ACTIVITY_FISSIONSHARE_ILLEGAL_APIKEY.Code(), apibackend.BASERR_ACTIVITY_FISSIONSHARE_ILLEGAL_APIKEY.OriginDesc())
//		ZapLog().Error( "apikey is null err "+ctx.Path())
//		return
//	}
//	if apiKey != "70123b81-9754-4fd4-b2cd-ecac5d3947cd" {
//		this.ExceptionSerive(ctx, apibackend.BASERR_ACTIVITY_FISSIONSHARE_ILLEGAL_APIKEY.Code(), apibackend.BASERR_ACTIVITY_FISSIONSHARE_ILLEGAL_APIKEY.OriginDesc())
//		ZapLog().Error( "apikey is not match "+ctx.Path())
//		return
//	}
//	ctx.Next()
//	return
//
//}
//
//func (this *Interceptor) MarketInterceptor(ctx iris.Context) {
//	_,ok := PathWhiteList[ctx.Path()];
//	if ok || strings.HasPrefix(ctx.Path(), "/v1/bk") {
//		ctx.Next()
//		return
//	}
//
//	apiKey := ctx.GetHeader("Api-Key")
//	if len(apiKey) == 0 {
//		this.ExceptionSerive(ctx, apibackend.BASERR_ACTIVITY_FISSIONSHARE_ILLEGAL_APIKEY.Code(), apibackend.BASERR_ACTIVITY_FISSIONSHARE_ILLEGAL_APIKEY.OriginDesc())
//		ZapLog().Error( "apikey is null err "+ctx.Path())
//		return
//	}
//	if apiKey != "70123b81-9754-4fd4-b2cd-ecac5d3947cd" {
//		this.ExceptionSerive(ctx, apibackend.BASERR_ACTIVITY_FISSIONSHARE_ILLEGAL_APIKEY.Code(), apibackend.BASERR_ACTIVITY_FISSIONSHARE_ILLEGAL_APIKEY.OriginDesc())
//		ZapLog().Error( "apikey is not match "+ctx.Path())
//		return
//	}
//	ctx.Next()
//	return
//
//}

//apikey 是否与对应activity_uuid 匹配   业务层发
// 不过参数不存在activity_uuid，生成apikey对应的账号信息 业务层