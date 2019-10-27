package aliyun_api

import "LianFaPhone/lfp-marketing-api/common"
import (
	"encoding/json"
	"fmt"
)

type (
	ReqIdCardInfo struct {
	}

	ResIdCardInfo struct {
		Ret   int         `json:"ret,omitempty"`
		Msg   string      `json:"msg,omitempty"`
		LogId string      `json:"log_id,omitempty"`
		Data  *IdCardInfo `json:"data,omitempty"`
	}

	IdCardInfo struct {
		Area     string `json:"area,omitempty"`
		Number   string `json:"number,omitempty"`
		Province string `json:"province,omitempty"`
		Addrcode string `json:"addrcode,omitempty"`
		City     string `json:"city,omitempty"`
		Sex      string `json:"sex,omitempty"`
		Length   int    `json:"length,omitempty"`
		Birth    string `json:"birth,omitempty"`
		Region   string `json:"region,omitempty"`
		Age      int    `json:"age,omitempty"`
	}
)

func (this *ReqIdCardInfo) GetByNumber(no string) (*IdCardInfo, error) {
	data, err := common.HttpSend("https://api10.aliyun.venuscn.com/id-card/query?number="+no, nil, "GET", map[string]string{"Authorization": "APPCODE"})
	if err != nil {
		return nil, err
	}

	res := new(ResIdCardInfo)
	if err = json.Unmarshal(data, res); err != nil {
		return nil, err
	}
	if res.Ret != 200 {
		return nil, fmt.Errorf("%d-%s", res.Ret, res.Msg)
	}
	return res.Data, nil

}

func (this *ReqIdCardInfo) CheckNumber(no, name string) (bool, error) {
	_, err := this.GetByNumber(no)
	if err != nil {
		return false, err
	}

	return false, nil
}
