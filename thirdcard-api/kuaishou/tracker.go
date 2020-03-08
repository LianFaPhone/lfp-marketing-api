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
		ErrMsg *string `json:"error_msg"`
		HostName *string `json:"host-name"`
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
	//ZapLog().Info("kuaishou res",zap.String("res", string(resData)))
	if res.Ret != 1 {
		if res.ErrMsg != nil {
			return fmt.Errorf("%d-%s", res.Ret, *res.ErrMsg)
		}
		return fmt.Errorf("%d-%s", res.Ret, string(resData))
	}
	return nil
}

