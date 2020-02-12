package api

type (
	OldCardOrderApply struct {
		TrueName *string `valid:"required" json:"truename,omitempty"`
		IdCard   *string `valid:"required" json:"idcard,omitempty"`
		Phone    *string `valid:"required" json:"phone,omitempty"`
		Address  string  `valid:"optional" json:"address,omitempty"` //小区及门牌号
		Address2 string  `valid:"optional" json:"address,omitempty"` //小区及门牌号
		Channel  string  `valid:"optional" json:"channel,omitempty"` //小区及门牌号

	}

	CardOrderApply struct {
		VerifyId *string `valid:"optional" json:"verify_uuid,omitempty"` //短信验证码，防止薅羊毛

		TrueName     *string `valid:"required" json:"true_name,omitempty"`
		IdCard       *string `valid:"required" json:"idcard,omitempty"`
		CountryCode  *string `valid:"optional" json:"country_code,omitempty"`
		Phone        *string `valid:"required" json:"phone,omitempty"`
		Province     *string `valid:"optional" json:"province,omitempty"`      //上海市
		ProvinceCode *string `valid:"optional" json:"province_code,omitempty"` //上海市
		City         *string `valid:"optional" json:"city,omitempty"`          //上海市
		CityCode     *string `valid:"optional" json:"city_code,omitempty"`     //上海市
		Area         *string `valid:"optional" json:"area,omitempty"`          //浦东新区
		AreaCode     *string `valid:"optional" json:"area_code,omitempty"`     //浦东新区
		Town         *string `valid:"optional" json:"town,omitempty"`          //东明路街道, town street
		TownCode     *string `valid:"optional" json:"town_code,omitempty"`     //东明路街道
		Address      *string `valid:"optional" json:"address,omitempty"`       //小区及门牌号
		ClassBigTp   *int    `valid:"optional" json:"class_big_tp,omitempty"`
		ClassTp      *int    `valid:"optional" json:"class_tp,omitempty"`
		ClassName   *string `valid:"optional" json:"class_name,omitempty"`
		ClassISP     *int    `valid:"optional" json:"class_isp,omitempty"`
		IP           string  `valid:"optional" json:"ip,omitempty"`
		Ips          *int    `valid:"optional" json:"-"` //屏幕材质？
		PhoneOSTp    *int    `valid:"optional" json:"class_isp,omitempty"`
		Status       *int    `valid:"optional"  json:"-"`
		Number       *string `valid:"optional" json:"new_phone,omitempty"` //选的号码
		FinishFlag     *int    `valid:"optional" json:"finish_flag,omitempty"`
		NumberPoolFlag  *int `valid:"optional" json:"numberpool_flag,omitempty"` //选的号码
		SessionId string    `valid:"optional" json:"session_id,omitempty"`
	}

	ResCardOrderApply struct {
		OrderNo   string `valid:"optional" json:"order_no,omitempty" `    //订单号
	}

	FtCardOrderStatus struct {
		OrderNo   string   `valid:"required"  json:"order_no,omitempty"`    //订单号
		Status     *int    `valid:"optional"  json:"status,omitempty"` //订单状态
	}

	ResChinaAddrCode struct {
		Code       string              `json:"code,omitempty"`
		Region     string              `json:"region,omitempty"`
		SubRegions []*ResChinaAddrCode `json:"sub_regions,omitempty"`
	}

	SmsSend struct {
		Language    *string `valid:"optional" json:"language,omitempty"`
		CountryCode *string `valid:"optional" json:"country_code,omitempty"`
		Phone       *string `valid:"required" json:"phone,omitempty"`
		PlayTp      int     `valid:"optional" json:"play_tp,omitempty"`
	}
	SmsVerify struct {
		VerifyUUId  *string `valid:"required" json:"verify_uuid,omitempty"`
		VerifyCode  *string `valid:"required" json:"verify_code,omitempty"`
		CountryCode *string `valid:"optional" json:"country_code,omitempty"`
		Phone       *string `valid:"required" json:"phone,omitempty"`
	}
	StreetGets struct {
		AreaCode *string `valid:"required" json:"area_code,omitempty"`
	}
	FtPhoneNumberList struct {
		UseFlag int   `valid:"-"  json:"-"`
		Valid   int   `valid:"-"  json:"-"`
		Page    int64 `valid:"optional" json:"page,omitempty"`
		Size    int64 `valid:"optional" json:"size,omitempty"`
	}

	FtPhoneNumberCheck struct {
		Number *string `valid:"required"  json:"number,omitempty"`
	}

	FtCardClass struct{
		ISP   *int    `valid:"optional" json:"isp,omitempty"      gorm:"column:isp;type:int(11)"`
		BigTp    *int    `valid:"optional" json:"big_tp,omitempty"       gorm:"column:big_tp;type:int(11)"` //加上type:int(11)后AUTO_INCREMENT无效
		Tp    *int    `valid:"optional" json:"tp,omitempty"       gorm:"column:tp;type:int(11)"` //加上type:int(11)后AUTO_INCREMENT无效
		Name *string `valid:"optional" json:"name,omitempty"     gorm:"column:name;type:varchar(50)"`

	}

	FtPhoneNumberLock struct{
		SessionId  *string   `valid:"required" json:"session_id,omitempty"`
		Number *string `valid:"required" json:"number,omitempty"`
	}
	FtPhoneNumberUnLock struct{
		SessionId  *string   `valid:"required" json:"session_id,omitempty"`
		Number *string `valid:"required" json:"number,omitempty"`
	}
)
