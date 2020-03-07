package api

type (

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

		PartnerId    *int64    `json:"valid:"optional" partner_id,omitempty"`                 //手机卡套餐类型
		PartnerGoodsCode    *string    `valid:"optional" json:"partner_goods_code,omitempty"`                 //手机卡套餐类型
		ClassISP     *int    `valid:"optional" json:"isp,omitempty"`
		IP           string  `valid:"optional" json:"ip,omitempty"`
		Ips          *int    `valid:"optional" json:"-"` //屏幕材质？
		PhoneOSTp    *int    `valid:"optional" json:"class_isp,omitempty"`
		Status       *int    `valid:"optional"  json:"-"`
		Number       *string `valid:"optional" json:"new_phone,omitempty"` //选的号码
		FinishFlag     *int    `valid:"optional" json:"finish_flag,omitempty"`
		NumberPoolFlag  *int `valid:"optional" json:"numberpool_flag,omitempty"` //选的号码
		SessionId string    `valid:"optional" json:"session_id,omitempty"`
		ThirdOrderNo    *string `valid:"optional"  json:"third_order_no,omitempty" ` //订单号
		Log            *string  `valid:"optional"  json:"log,omitempty" `
	}

	ResCardOrderApply struct {
		OrderNo   string `valid:"optional" json:"order_no,omitempty" `    //订单号
	}

	FtCardOrderStatus struct {
		OrderNo   string   `valid:"required"  json:"order_no,omitempty"`    //订单号
		//Status     *int    `valid:"optional"  json:"status,omitempty"` //订单状态
		Log       *string    `valid:"optional"  json:"log,omitempty"`
		SuccFlag  *int   `valid:"optional"  json:"succ_flag,omitempty"`
	}

	FtCardOrderConfirm struct {
		OrderNo   string   `valid:"required"  json:"order_no,omitempty"`    //订单号
		//Status     *int    `valid:"optional"  json:"status,omitempty"` //订单状态
		Log       *string    `valid:"optional"  json:"log,omitempty"`
		SuccFlag  *int   `valid:"optional"  json:"succ_flag,omitempty"`
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


	FtResPdPartnerGoodsGet struct{
		Code     *string  `json:"code,omitempty"     gorm:"column:code;type:varchar(10);unique" `
		UrlParam *string  `json:"url_param,omitempty"     gorm:"column:url_param;type:varchar(200)" `
		ImgUrl *string `json:"img_url,omitempty"     gorm:"column:img_url;type:varchar(250)"`
		HeadImgUrl *string `json:"head_img_url,omitempty"     gorm:"column:head_img_url;type:varchar(250)"`
		TailImgUrl *string `json:"tail_img_url,omitempty"     gorm:"column:tail_img_url;type:varchar(250)"`
		AdTp *int    `json:"ad_tp,omitempty"      gorm:"column:ad_tp;type:int(11)"`

		NoExpAddr   *string      `json:"no_exp_addr,omitempty"     gorm:"column:no_exp_addr;type:varchar(200)"`
		MinAge    *int    `json:"min_age,omitempty"     gorm:"column:min_age;type:int(11)"`
		MaxAge    *int    `json:"max_age,omitempty"     gorm:"column:max_age;type:int(11)"`
		SmsFlag *int    `json:"sms_flag,omitempty"      gorm:"column:sms_flag;type:tinyint(4)"`
		IdcardDispplay *int    `json:"idcard_display,omitempty"      gorm:"column:idcard_display;type:tinyint(3);default 0"`
		BgColor  *string    `json:"bg_color,omitempty"`
		Name     *string   `json:"name,omitempty"`
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

//ydhk
type(
	FtResYdhkToken struct{
		Token  string ` json:"token,omitempty"`
	}

	FtYdhkNumberPoolList struct{
		// "provCode":"551","province":"安徽","cityCode":"558","city":"阜阳市","selecttype":0,"searchkey":"","count":10
		ProviceCode string   `json:"provice_code,omitempty"`
		Provice string        `json:"provice,omitempty"`
		CityCode string       `json:"city_code,omitempty"`
		City string            `json:"city,omitempty"`
		Searchkey  string             `json:"searchkey,omitempty"`
		Page int                `json:"page,omitempty"`
		Size int             `json:"size,omitempty"`
	}

	FtYdhkNumberPoolLock struct{
		ProviceCode string   `json:"province_code,omitempty"`
		CityCode string       `json:"city_code,omitempty"`
		Number  string             `json:"number,omitempty"`
		Token  string             `json:"token,omitempty"`
	}

	FtYdhkApply struct{
		Phone   string  `valid:"required" json:"phone"`
		NewPhone      string  `valid:"required" json:"new_phone"`
		LeagalName    string    `valid:"required" json:"true_name"`
		CertificateNo    string    `valid:"required" json:"idcard"`

		Province    string    `valid:"optional" json:"province_code"`
		City    string        `valid:"optional" json:"city_code"`

		SendProvinceName    string    `valid:"optional" json:"express_province"`
		SendCityName    string    `valid:"optional" json:"express_city"`
		SendDistrictName    string    `valid:"optional" json:"express_district"`

		SendProvince    string    `valid:"optional" json:"express_province_code"`
		SendCity    string    `valid:"optional" json:"express_city_code"`
		SendDistrict    string    `valid:"optional" json:"express_district_code"`
		Address    string     `valid:"optional" json:"express_address"`

		AccessToken     string   `valid:"required" json:"token"`

		PartnerId    *int64    `valid:"optional" json:"-"  `                 //手机卡套餐类型
		PartnerGoodsCode    *string    `valid:"required" json:"partner_goods_code,omitempty"   `                 //手机卡套餐类型

		ClassISP     *int    `valid:"optional" json:"-"`
		IP           *string  `valid:"optional" json:"ip,omitempty"` //从header
		PhoneOSTp    *int    `valid:"optional" json:"device_os_tp,omitempty"` //从header
		AdCallback  *string   `valid:"optional" json:"ad_callback,omitempty"`
		AdTp        *int `valid:"optional" json:"ad_tp,omitempty"`
	}

	FtResYdhkApply struct{
		ThirdOrderId string   `json:"third_order_no"`
		//OrderId string   `json:"order_no"`
		OaoModel  bool    `json:"oao_model"`
	}

	FtIdCheckUrlGet struct {
		OrderId   *string   `valid:"optional"  json:"order_no,omitempty"`
		ThirdOrderNo   string   `valid:"required"  json:"third_order_no,omitempty"`    //订单号
		NewPhone       string    `valid:"optional"  json:"new_phone,omitempty"`
		Token  string    `valid:"optional"  json:"token,omitempty"`
	}
)