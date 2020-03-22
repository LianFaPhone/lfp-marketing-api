package dxnbhk

import (
	"LianFaPhone/lfp-marketing-api/common"
	"LianFaPhone/lfp-marketing-api/config"
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"net/url"
	"strings"
	"time"
)

type ReOrderSubmit struct{
	//timeStamp  string
	//services   string
	ContactNumber string
	UserName    string
	IdNumber     string
	AgentMark   string
	ProductId   string
	Province    string
	City        string
	Area        string
	Address     string
}

type ResOrderSubmit struct{
	Success   bool  `json:"success"`
	Msg        string  `json:"msg"`
	SelectNumber    string  `json:"selectNumber"`
	OrderNumber     string  `json:"orderNumber"`
	UploadUrl      string  `json:"uploadUrl"`
}

func (this *ReOrderSubmit) Send() (*ResOrderSubmit, error) {
	//预先定义好顺序，不是比程序计算更快啊
	//address>agentMark>Area>City>ContactNumber>IdNumber>ProductId>Province>services>timeStamp>UserName
	if len(this.AgentMark) <= 0 {
		return nil,fmt.Errorf("nil AgentMark")
	}
	timeStr :=  time.Now().Format("20060102150405")
	md5Str := this.Address+this.AgentMark+this.Area+this.City+this.ContactNumber+this.IdNumber+this.ProductId+this.Province+ "submit"+timeStr+this.UserName+"NB@Zt.123"
	md5Str = strings.ToUpper(md5V(md5Str))

	formBody := make(url.Values)
	formBody.Set("address", this.Address)
	formBody.Set("agentMark", this.AgentMark)
	formBody.Set("area", this.Area)
	formBody.Set("city", this.City)
	formBody.Set("contactNumber", this.ContactNumber)
	formBody.Set("idNumber", this.IdNumber)
	formBody.Set("productId", this.ProductId)
	formBody.Set("province", this.Province)
	formBody.Set("services", "submit")
	formBody.Set("timeStamp",timeStr)
	formBody.Set("userName", this.UserName)
	formBody.Set("sign", md5Str)

	header := make(map[string]string)
	header["Content-Type"] = "application/x-www-form-urlencoded"
	resBytes,err := common.HttpFormSend(config.GConfig.Dxnbhk.Url, formBody, "POST", header)
	if err != nil {
		return nil,err
	}

	res := new(ResOrderSubmit)
	if err := json.Unmarshal(resBytes, res); err != nil {
		return nil,err
	}
	return res,nil
}

func md5V(str string) string  {
	h := md5.New()
	h.Write([]byte(str))
	return hex.EncodeToString(h.Sum(nil))
}