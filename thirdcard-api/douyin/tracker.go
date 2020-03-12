package douyin

import (
	"LianFaPhone/lfp-marketing-api/common"
	"fmt"
	"net/url"
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

func (this *ReTracker) Send(link, source string, eventTime,event_type int64) ( error) {

	link = url.QueryEscape(link)

	url:=fmt.Sprintf("https://ad.toutiao.com/track/activate/?link=%s&source=%s&conv_time=%d&event_type=%d", link, source, eventTime, event_type)

	_, err := common.HttpSend(url, nil,"GET", nil)
	if err != nil {
		return  err
	}
	//res := new(ResTracker)
	//if err = json.Unmarshal(resData, res); err != nil {
	//	return   err
	//}
	////ZapLog().Info("douyin res",zap.String("res", string(resData)))
	//if res.Ret != 1 {
	//	if res.ErrMsg != nil {
	//		return fmt.Errorf("%d-%s", res.Ret, *res.ErrMsg)
	//	}
	//	return fmt.Errorf("%d-%s", res.Ret, string(resData))
	//}
	return nil
}


