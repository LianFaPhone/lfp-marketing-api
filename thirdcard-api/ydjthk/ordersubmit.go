package ydjthk

import (
	apibackend "LianFaPhone/lfp-api/errdef"
	"LianFaPhone/lfp-marketing-api/config"

	//. "LianFaPhone/lfp-tools/autoorder-yidonghuaka/config"
	//"LianFaPhone/lfp-base/log/zap"
	"LianFaPhone/lfp-marketing-api/common"
	"bytes"
	"crypto/aes"
	"encoding/base64"
	"encoding/json"
	"fmt"

)

type (
	ReOrderSubmit struct {
		MsgType  string  `json:"MsgType"`
		Version string    `json:"Version"`
		ChannelId string   `json:"channelId"`
		CardProductId string  `json:"cardProductId"`  //加密
		MobilePhone   string  `json:"mobilePhone"`
		MobileId      string  `json:"mobileId"`
		LeagalName    string    `json:"leagalName"`
		CertificateNo    string    `json:"certificateNo"`
		Address    string    `json:"address"`
		Province    string    `json:"province"`
		City    string    `json:"city"`

		SendProvince    string    `json:"sendProvince"`
		SendCity    string    `json:"sendCity"`
		SendDistrict    string    `json:"sendDistrict"`

		AccessToken     string   `json:"accessToken"`
		SellerId     string   `json:"sellerId"`
		SellerMobile     string   `json:"sellerMobile"`
		Ex_field         string   `json:"ex_field"`
		OaoModel    *string    `json:"oaoModel,omitempty"`

	}
	ResOrderSubmit struct {
		Ret   string    `json:"hRet"`
		Msg string   `json:"retMsg"`
		OrderId string   `json:"orderId"`
		OaoModel  bool    `json:"oaoModel"`
	}
)

func (this *ReOrderSubmit) Parse(channelId , productId string, OaoModel *string) *ReOrderSubmit{
	this.ChannelId = channelId
	this.CardProductId = productId
	this.OaoModel = OaoModel
	return this
}

//最后三个是码
func (this *ReOrderSubmit) Send(isOao bool, token, inPhone, newPhone, LegalName,IdCard, address, province, city,  sendprovince, sendcity, sendqu string) (apibackend.EnumBasErr,string,bool, error) {
	//return apibackend.BASERR_SUCCESS, "TestOrder112",true, nil
	this.MsgType = "LiveHKCardTemporaryOrderReq"
	this.Version = Const_Version


	this.CardProductId = base64.StdEncoding.EncodeToString(EcbEncrypt([]byte(this.CardProductId), []byte(token[0:16])))

	this.MobilePhone = base64.StdEncoding.EncodeToString(EcbEncrypt([]byte(inPhone), []byte(token[0:16])))
	this.MobileId = base64.StdEncoding.EncodeToString(EcbEncrypt([]byte(newPhone), []byte(token[0:16])))
	this.LeagalName =  base64.StdEncoding.EncodeToString(EcbEncrypt([]byte(LegalName), []byte(token[0:16])))
	this.CertificateNo = base64.StdEncoding.EncodeToString(EcbEncrypt([]byte(IdCard), []byte(token[0:16])))
	this.Address = base64.StdEncoding.EncodeToString(EcbEncrypt([]byte(address), []byte(token[0:16])))
	this.SendProvince =  base64.StdEncoding.EncodeToString(EcbEncrypt([]byte(sendprovince), []byte(token[0:16])))
	this.SendCity =  base64.StdEncoding.EncodeToString(EcbEncrypt([]byte(sendcity), []byte(token[0:16])))
	this.SendDistrict =  base64.StdEncoding.EncodeToString(EcbEncrypt([]byte(sendqu), []byte(token[0:16])))

	this.Province =  base64.StdEncoding.EncodeToString(EcbEncrypt([]byte(province), []byte(token[0:16])))
	this.City =  base64.StdEncoding.EncodeToString(EcbEncrypt([]byte(city), []byte(token[0:16])))


	this.AccessToken = token

	heads := map[string]string{
		"Accept": "application/json, text/plain, */*",
		"Host": config.GConfig.Jthk.Host,
		"Origin": config.GConfig.Jthk.Url,
		//"Referer": Const_Url+"/rwx/rwkvue/young/",
	}

	path := ""
	if isOao {
		path = "/rwx/rwkweb/livehk/card/delivery"
		heads["Referer"] = config.GConfig.Jthk.Url + config.GConfig.Jthk.Referer_path_oao
	}else{
		path = "/rwx/rwkweb/livehk/card/temporaryorder"
		heads["Referer"] = config.GConfig.Jthk.Url + config.GConfig.Jthk.Referer_path
	}

	reqData,err := json.Marshal(this)
	if err != nil {
		return apibackend.BASERR_DATA_PACK_ERROR,"",false, err
	}

	resData, err := common.HttpSend(config.GConfig.Jthk.Url+path, bytes.NewReader(reqData),"POST", heads)
	if err != nil {
		return apibackend.BASERR_INTERNAL_SERVICE_ACCESS_ERROR, "",false, err
	}
	res := new(ResOrderSubmit)
	if err = json.Unmarshal(resData, res); err != nil {
		return apibackend.BASERR_DATA_UNPACK_ERROR,"",false, err
	}
	if res.Ret != "0" {
		return apibackend.BASERR_CARDMARKET_PHONECARD_APPLY_FAID_AND_SHOW,"",false, fmt.Errorf("%s", res.Msg)
	}

	return apibackend.BASERR_SUCCESS, res.OrderId,res.OaoModel, nil
}

func (this *ReOrderSubmit) OfflineActiveSend( token, inPhone, newPhone, LegalName,IdCard, address, province, city,  sendprovince, sendcity, sendqu string) (apibackend.EnumBasErr,string,bool, error) {
	//return apibackend.BASERR_SUCCESS, "TestOrder112",true, nil
	this.MsgType = "LiveHKCardTemporaryOrderReq"
	this.Version = Const_Version


	this.CardProductId = base64.StdEncoding.EncodeToString(EcbEncrypt([]byte(this.CardProductId), []byte(token[0:16])))

	this.MobilePhone = base64.StdEncoding.EncodeToString(EcbEncrypt([]byte(inPhone), []byte(token[0:16])))
	this.MobileId = base64.StdEncoding.EncodeToString(EcbEncrypt([]byte(newPhone), []byte(token[0:16])))
	this.LeagalName =  base64.StdEncoding.EncodeToString(EcbEncrypt([]byte(LegalName), []byte(token[0:16])))
	this.CertificateNo = base64.StdEncoding.EncodeToString(EcbEncrypt([]byte(IdCard), []byte(token[0:16])))
	this.Address = base64.StdEncoding.EncodeToString(EcbEncrypt([]byte(address), []byte(token[0:16])))
	this.SendProvince =  base64.StdEncoding.EncodeToString(EcbEncrypt([]byte(sendprovince), []byte(token[0:16])))
	this.SendCity =  base64.StdEncoding.EncodeToString(EcbEncrypt([]byte(sendcity), []byte(token[0:16])))
	this.SendDistrict =  base64.StdEncoding.EncodeToString(EcbEncrypt([]byte(sendqu), []byte(token[0:16])))

	this.Province =  base64.StdEncoding.EncodeToString(EcbEncrypt([]byte(province), []byte(token[0:16])))
	this.City =  base64.StdEncoding.EncodeToString(EcbEncrypt([]byte(city), []byte(token[0:16])))


	this.AccessToken = token

	heads := map[string]string{
		"Accept": "application/json, text/plain, */*",
		"Host": config.GConfig.Jthk.Host,
		"Origin": config.GConfig.Jthk.Url,
		//"Referer": Const_Url+"/rwx/rwkvue/young/",
	}


	path := "/rwx/rwkweb/livehk/card/temporaryorder"
	heads["Referer"] = config.GConfig.Jthk.Url + config.GConfig.Jthk.Referer_path_oao


	reqData,err := json.Marshal(this)
	if err != nil {
		return apibackend.BASERR_DATA_PACK_ERROR,"",false, err
	}

	resData, err := common.HttpSend(config.GConfig.Jthk.Url+path, bytes.NewReader(reqData),"POST", heads)
	if err != nil {
		return apibackend.BASERR_INTERNAL_SERVICE_ACCESS_ERROR, "",false, err
	}
	res := new(ResOrderSubmit)
	if err = json.Unmarshal(resData, res); err != nil {
		return apibackend.BASERR_DATA_UNPACK_ERROR,"",false, err
	}
	if res.Ret != "0" {
		return apibackend.BASERR_CARDMARKET_PHONECARD_APPLY_FAID_AND_SHOW,"",false, fmt.Errorf("%s", res.Msg)
	}

	return apibackend.BASERR_SUCCESS, res.OrderId,res.OaoModel, nil
}




func PKCS7Padding(ciphertext []byte, blockSize int) []byte {
	padding := blockSize - len(ciphertext)%blockSize
	padtext := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(ciphertext, padtext...)
}

func PKCS7UnPadding(origData []byte) []byte {
	length := len(origData)
	unpadding := int(origData[length-1])
	return origData[:(length - unpadding)]
}

func EcbDecrypt(data, key []byte) []byte {
	block, _ := aes.NewCipher(key)
	decrypted := make([]byte, len(data))
	size := block.BlockSize()

	for bs, be := 0, size; bs < len(data); bs, be = bs+size, be+size {
		block.Decrypt(decrypted[bs:be], data[bs:be])
	}

	return PKCS7UnPadding(decrypted)
}

func EcbEncrypt(data, key []byte) []byte {
	block, _ := aes.NewCipher(key)
	data = PKCS7Padding(data, block.BlockSize())
	decrypted := make([]byte, len(data))
	size := block.BlockSize()

	for bs, be := 0, size; bs < len(data); bs, be = bs+size, be+size {
		block.Encrypt(decrypted[bs:be], data[bs:be])
	}

	return decrypted
}
