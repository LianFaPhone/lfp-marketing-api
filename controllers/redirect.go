package controllers

import (
	"encoding/json"
	"github.com/kataras/iris"
	"io/ioutil"
	"net/http"
	apibackend "LianFaPhone/lfp-api/errdef"
	"LianFaPhone/lfp-marketing-api/config"
	"LianFaPhone/lfp-base/log/zap"
	"go.uber.org/zap"
	"strings"
)

type RedirectController struct {

}

type ResNotifyMsg struct {
	Err            *int             `json:"err,omitempty"`
	ErrMsg         *string          `json:"errmsg,omitempty"`
	TemplateGroupList interface{}      `json:"templategrouplist,omitempty"`
	Templates      interface{}      `json:"template,omitempty"`
	TemplateHistoryList interface{} `json:"templatehistorylist,omitempty"`
}

func (this *ResNotifyMsg) GetErr() int {
	if this.Err == nil {
		return 0
	}
	return *this.Err
}

func (this *ResNotifyMsg) GetErrMsg() string {
	if this.ErrMsg == nil {
		return  ""
	}
	return *this.ErrMsg
}

func (bp *RedirectController) HandlerV1BasNotify(ctx iris.Context) {
	query := ctx.Request().URL.RawQuery
	path := strings.Replace(ctx.Path(), "/bk/fissionshare", "",  -1)
	newUrl := config.GConfig.BasNotify.Addr + path
	if len(query) != 0 {
		newUrl =newUrl + "?" + query
	}
	req, err := http.NewRequest(ctx.Method(), newUrl, ctx.Request().Body)
	if err != nil {
		log.ZapLog().Error("http NewRequest err", zap.Error(err))
		ctx.JSON(&Response{Code: apibackend.BASERR_INTERNAL_CONFIG_ERROR.Code(), Message: "NewRequest_ERROR:" + err.Error()})
		return
	}

	req.Header.Set("Content-Type", "application/json")
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		log.ZapLog().Error("http Do   err",  zap.String("url", newUrl), zap.Error(err))
		ctx.JSON(&Response{Code: apibackend.BASERR_INTERNAL_SERVICE_ACCESS_ERROR.Code(), Message: "BasNotify_HTTP_DO_ERROR:" + err.Error()})
		return
	}
	if resp.StatusCode != 200 {
		log.ZapLog().Error("http Do Response  err", zap.String("url",newUrl ) , zap.String("status", resp.Status))
		ctx.JSON(&Response{Code: apibackend.BASERR_INTERNAL_SERVICE_ACCESS_ERROR.Code(), Message: "BasNotify_HTTP_RESPONSE_ERROR:" + resp.Status})
		return
	}

	content, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.ZapLog().Error("Body readAll  err", zap.Error(err))
		ctx.JSON(&Response{Code: apibackend.BASERR_INTERNAL_SERVICE_ACCESS_ERROR.Code(), Message: "BASNOTIFY_READ_BODY_ERROR:"+ err.Error()})
		return
	}
	defer resp.Body.Close()

	if len(content) == 0 {
		log.ZapLog().Error("  response is null")
		ctx.JSON(&Response{Code: apibackend.BASERR_INTERNAL_SERVICE_ACCESS_ERROR.Code(), Message: "BASNOTIFY_REDIRECT_ERROR:response body content null"})
		return
	}
	if string(content) == "Not Found" {
		log.ZapLog().Error("response is Not Found")
		ctx.JSON(&Response{Code: apibackend.BASERR_INTERNAL_SERVICE_ACCESS_ERROR.Code(), Message: "BASNOTIFY_REDIRECT_ERROR:response is Not Found"})
		return
	}
	log.ZapLog().Debug("BASNOTIFY response content", zap.String("content", string(content)))
	adminRes := new(ResNotifyMsg)
	if err := json.Unmarshal(content, adminRes); err != nil {
		log.ZapLog().Error("Unmarshal  err",  zap.Error(err) )
		ctx.JSON(&Response{Code: apibackend.BASERR_DATA_UNPACK_ERROR.Code(), Message: "BASNOTIFY_REDIRECT_ERROR:response cannot Unmarshal, " + err.Error()})
		return
	}

	if adminRes.GetErr() != 0 {
		log.ZapLog().Error("BASNOTIFY Response.Status.Code err", zap.Int("err", adminRes.GetErr()) , zap.String("errmsg", adminRes.GetErrMsg()))
		ctx.JSON(&Response{Code: adminRes.GetErr(), Message: "BASNOTIFY_REDIRECT_ERROR: " + adminRes.GetErrMsg()})
		return
	}
	if adminRes.TemplateGroupList != nil {
		ctx.JSON(&Response{Code: adminRes.GetErr(), Message: adminRes.GetErrMsg(), Data: adminRes.TemplateGroupList})
		return
	}
	if adminRes.Templates != nil {
		ctx.JSON(&Response{Code: adminRes.GetErr(), Message: adminRes.GetErrMsg(), Data: adminRes.Templates})
		return
	}
	if adminRes.TemplateHistoryList != nil {
		ctx.JSON(&Response{Code: adminRes.GetErr(), Message: adminRes.GetErrMsg(), Data: adminRes.TemplateHistoryList})
		return
	}
	ctx.JSON(&Response{Code: adminRes.GetErr(), Message: adminRes.GetErrMsg()})
	//	l4g.Debug("deal HandleV1Admin username[%s] ok", utils.GetValueUserName(ctx))
	//ctx.Next()
	return
}
