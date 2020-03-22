package sdk

import (
	"encoding/json"
	"io/ioutil"
	"net/url"

	"errors"
	ypclnt "github.com/yunpian/yunpian-go-sdk/sdk"
)

var GYunPainSdk YunPainSdk
type YunPainSdk struct{
	yunpianClient ypclnt.YunpianClient
	apikey string
}

type ResponseYunPian struct{
	Code int        `json:"code"`
	Msg string       `json:"msg"`
	ShortUrl ShortUrlYunPian    `json:"short_url"`
}

type ShortUrlYunPian struct{
	Sid	        string    `json:"sid"`
	LongUrl	string    `json:"long_url"`
	ShortUrl	string    `json:"short_url"`
	EnterUrl	string   `json:"enter_url"`
	Name	     string   `json:"name"`
	StatExpire	string   `json:"stat_expire"`
}

func (this *YunPainSdk) Init(apikey string) {
	this.apikey = apikey
	this.yunpianClient =  ypclnt.New(apikey)
}

func (this *YunPainSdk) GetShortUrl(longurl string) (string, error){
	param := make(url.Values)
	param.Add("long_url", longurl)
	param.Add(ypclnt.APIKEY, this.apikey)


	resp,err := this.yunpianClient.Post("https://sms.yunpian.com/v2/short_url/shorten.json", param.Encode(), nil, "")
	if err != nil{
		return "",err
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		return "", errors.New(resp.Status)
	}

	content, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	if len(content) == 0 {
		return "", errors.New("nil resp")
	}
	response := new(ResponseYunPian)
	err = json.Unmarshal(content, response)
	if err != nil {
		return "",err
	}
	return response.ShortUrl.ShortUrl, nil
}