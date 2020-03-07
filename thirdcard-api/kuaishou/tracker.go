package kuaishou

import (
	"LianFaPhone/lfp-marketing-api/common"
	"encoding/json"
	"fmt"
)

type(
	ReTracker struct{

	}

	ResTracker struct{
		Ret   int    `json:"result"`
	}


)

func (this *ReTracker) Send(callback, eventTp string, eventTime int64) ( error) {

	url:=fmt.Sprintf("http://ad.partner.gifshow.com/track/activate?callback=%s&event_type=%s&event_time=%d", callback, eventTp, eventTime)

	resData, err := common.HttpSend(url, nil,"GET", nil)
	if err != nil {
		return  err
	}
	res := new(ResTracker)
	if err = json.Unmarshal(resData, res); err != nil {
		return   err
	}
	if res.Ret != 1 {
		fmt.Errorf("err%d-%s", res.Ret, string(resData))
	}
	return nil
}

