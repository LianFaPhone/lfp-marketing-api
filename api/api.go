package api

type(
	CardOrderApply struct{
		VerifyId   *string    `valid:"required" json:"verify_uuid,omitempty"` //短信验证码，防止薅羊毛

		TrueName   *string    `valid:"required" json:"true_name,omitempty"`
		IdCard     *string    `valid:"required" json:"idcard,omitempty"`
		CountryCode *string   `valid:"optional" json:"country_code,omitempty"`
		Phone      *string    `valid:"required" json:"phone,omitempty"`
		Province   *string    `valid:"optional" json:"province,omitempty"` //上海市
		City       *string    `valid:"optional" json:"city,omitempty"`  //上海市
		Area       *string    `valid:"optional" json:"area,omitempty"` //浦东新区
		ProvinceCode   *string    `valid:"optional" json:"province_code,omitempty"` //上海市
		CityCode       *string    `valid:"optional" json:"city_code,omitempty"`  //上海市
		AreaCode       *string    `valid:"optional" json:"area_code,omitempty"` //浦东新区
		Town       *string    `valid:"optional" json:"town,omitempty"`  //东明路街道
		Address    *string    `valid:"optional" json:"address,omitempty"` //小区及门牌号
		ClassTp    *int       `valid:"required" json:"class_tp,omitempty"`
		ClassAlias *string    `valid:"optional" json:"class_alias,omitempty"`
		ClassISP    *int      `valid:"required" json:"class_isp,omitempty"`
		IP         string     `valid:"optional" json:"—"`
		Ips        *int    `valid:"optional" json:"-"`  //屏幕材质？
		PhoneOSTp  *int       `valid:"optional" json:"-"`
		Status     *int       `valid:"optional"  json:"-"`
		Number     *string    `valid:"optional" json:"new_phone,omitempty"`  //选的号码
	}

	ResChinaAddrCode struct{
		Code      string  `json:"code,omitempty"`
		Region    string  `json:"region,omitempty"`
		SubRegions []*ResChinaAddrCode  `json:"sub_regions,omitempty"`
	}

	SmsSend struct{
		Language    *string    `valid:"required" json:"language,omitempty"`
		CountryCode *string    `valid:"optional" json:"country_code,omitempty"`
		Phone       *string    `valid:"required" json:"phone,omitempty"`
		PlayTp      int        `valid:"optional" json:"play_tp,omitempty"`
	}
	SmsVerify struct{
		VerifyUUId    *string    `valid:"required" json:"verify_uuid,omitempty"`
		VerifyCode    *string    `valid:"required" json:"verify_code,omitempty"`
		CountryCode   *string    `valid:"required" json:"country_code,omitempty"`
		Phone         *string    `valid:"required" json:"phone,omitempty"`
	}
	StreetGets struct{
		AreaCode    *string    `valid:"required" json:"area_code,omitempty"`
	}
	FtPhoneNumberList struct{
		UseFlag    int        `valid:"-"  json:"-"`
		Valid      int       `valid:"-"  json:"-"`
		Page       int64       `valid:"optional" json:"page,omitempty"`
		Size       int64       `valid:"optional" json:"size,omitempty"`
	}

	FtPhoneNumberCheck struct{
		Number      *string    		`valid:"required"  json:"number,omitempty"`
	}

)