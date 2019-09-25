package bsp_api

type Request struct{
	AccessToken  string `json:"access_token,omitempty"`
	Method       string `json:"method,omitempty"`
	Version      string `json:"version,omitempty"`
}

type Response struct{
	Res_code string    `json:"res_code,omitempty"`
	Res_message string  `json:"res_message,omitempty"`
}

type ReBspIdCardCheck struct{
	Request
	Content struct {
		CartNo string           `json:"cartNo,omitempty"`
		BuyerName string         `json:"buyerName,omitempty"`
		MobilePhone string      `json:"mobilePhone,omitempty"`
		ReceiverName  string           `json:"receiverName,omitempty"`
		ReceiverProv  string           `json:"receiverProv,omitempty"`
		ReceiverCity  string          `json:"receiverCity,omitempty"`
		ReceiverDistrict string      `json:"receiverDistrict,omitempty"`
		ReceiverAdress string     `json:"receiverAdress,omitempty"`
	} `json:"content,omitempty"`

}

type ResBspIdCardCheck struct{
	Response
	Result  struct {
		Code   int       `json:"code,omitempty"`
		Message string  `json:"message,omitempty"`
	}  `json:"result,omitempty"`
}
