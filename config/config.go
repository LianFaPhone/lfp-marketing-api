package config

import (
	"fmt"
	"github.com/ulule/limiter"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"strings"
)

var GConfig Config
var GPreConfig PreConfig

func LoadConfig(path string) *Config {
	data, err := ioutil.ReadFile(path)
	if err != nil {
		panic(fmt.Errorf("Read yml config[%s] err[%v]", path, err).Error())
	}

	err = yaml.Unmarshal([]byte(data), &GConfig)
	if err != nil {
		panic(fmt.Errorf("yml.Unmarshal config[%s] err[%v]", path, err).Error())
	}
	PreProcess()
	return &GConfig
}

func PreProcess() error {
	if GConfig.Cache.CardClassByNameMaxKey < 10 {
		GConfig.Cache.CardClassByNameMaxKey = 10	}
	if GConfig.Cache.CardClassByNameTimeout < 300 {
		GConfig.Cache.CardClassByNameTimeout = 3600
	}
	if GConfig.Cache.SponsorActivityMaxKey < 100 {
		GConfig.Cache.SponsorActivityMaxKey = 100
	}
	if GConfig.Cache.SponsorActivityTimeout < 300 {
		GConfig.Cache.SponsorActivityTimeout = 3600
	}
	if GConfig.Cache.SponsorMaxKey < 100 {
		GConfig.Cache.SponsorMaxKey = 100
	}
	if GConfig.Cache.SponsorTimeout < 300 {
		GConfig.Cache.SponsorTimeout = 86400
	}
	if GConfig.Cache.PageMaxKey < 10 {
		GConfig.Cache.PageMaxKey = 10
	}
	if GConfig.Cache.PageTimeout < 1800 {
		GConfig.Cache.PageTimeout = 1800
	}
	if GConfig.Cache.ShareInfoMaxKey < 10 {
		GConfig.Cache.ShareInfoMaxKey = 10
	}
	if GConfig.Cache.ShareInfoTimeout < 1800 {
		GConfig.Cache.ShareInfoTimeout = 1800
	}

	if GConfig.Cache.RobberMaxKey < 10 {
		GConfig.Cache.RobberMaxKey = 1000
	}
	if GConfig.Cache.RobberTimeout < 600 {
		GConfig.Cache.RobberTimeout = 600
	}

	for i := 0; i < len(GConfig.BussinessLimits.PhoneSms); i++ {
		rate, err := limiter.NewRateFromFormatted(GConfig.BussinessLimits.PhoneSms[i])
		if err != nil {
			return err
		}
		GPreConfig.PhoneSmsLimits = append(GPreConfig.PhoneSmsLimits, &rate)
	}
	for i := 0; i < len(GConfig.BussinessLimits.IpSms); i++ {
		rate, err := limiter.NewRateFromFormatted(GConfig.BussinessLimits.IpSms[i])
		if err != nil {
			return err
		}
		GPreConfig.IpSmsLimits = append(GPreConfig.IpSmsLimits, &rate)
	}
	if len(GConfig.Jthk.ParterCode) > 0 {
		GConfig.Jthk.ParterCode = strings.Replace(GConfig.Jthk.ParterCode, " ", "", -1)
		GConfig.Jthk.ParterCode = strings.TrimSuffix(GConfig.Jthk.ParterCode, ",")
		GConfig.Jthk.ParterCodeArr = strings.Split(GConfig.Jthk.ParterCode, ",")
	}
	return nil
}

type PreConfig struct {
	PhoneSmsLimits []*limiter.Rate
	IpSmsLimits    []*limiter.Rate
}

type Config struct {
	Server System `yaml:"system"`
	Redis  Redis  `yaml:"redis"`
	Db     Mysql  `yaml:"mysql"`
	Cache  Cache  `yaml:"cache"`
	User   User   `yaml:"user"`
	//Aws     Aws          `yaml:"aws"`
	BussinessLimits BussinessLimits `yaml:"bussiness_limits"`
	BasNotify       BasNotify       `yaml:"bas_notify"`
	MarketMain      MarketMain      `yaml:"marketing_main"`
	IdCardCheck     IdCardCheck     `yaml:"idcard_check"`
	ChuangLan       ChuangLan       `yaml:"chuanglan"`
	YunPian        YunPian          `yaml:"yunpian"`
	Aliyun          Aliyun         `yaml:"aliyun"`
	Baidu           Baidu          `yaml:"baidu"`
	Task            Task           `yaml:"task"`
	Jthk            Jthk           `yaml:"jthk"`
	Dxnbhk          Dxnbhk          `yaml:"dxnbhk"`
}

type System struct {
	Port    string `yaml:"port"`
	Debug   bool   `yaml:"debug"`
	LogPath string `yaml:"log_path"`
	Monitor string `yaml:"monitor"`
	BkPort  string `yaml:"bk_port"`
	DevId  string `yaml:"dev_id"`
	FilePath string `yaml:"file_path"`
	LfcxHost string `yaml:"lfcx_host"`
}

type Cache struct {
	CardClassByNameMaxKey         int `yaml:"cardclassbyname_max_key"`
	CardClassByNameTimeout        int `yaml:"cardclassbyname_timeout"`
	SponsorActivityMaxKey  int `yaml:"sponsor_activity_max_key"`
	SponsorActivityTimeout int `yaml:"sponsor_activity_timeout"`
	SponsorMaxKey          int `yaml:"sponsor_max_key"`
	SponsorTimeout         int `yaml:"sponsor_timeout"`
	PageMaxKey             int `yaml:"page_max_key"`
	PageTimeout            int `yaml:"page_timeout"`
	ShareInfoMaxKey        int `yaml:"shareinfo_max_key"`
	ShareInfoTimeout       int `yaml:"shareinfo_timeout"`
	RobberMaxKey           int `yaml:"robber_max_key"`
	RobberTimeout          int `yaml:"robber_timeout"`
}

type Redis struct {
	Network     string `yaml:"network"`
	Host        string `yaml:"host"`
	Port        string `yaml:"port"`
	Password    string `yaml:"password"`
	Database    int    `yaml:"database"`
	MaxIdle     int    `yaml:"maxIdle"`
	MaxActive   int    `yaml:"maxActive"`
	IdleTimeout int    `yaml:"idleTimeout"`
	Prefix      string `yaml:"prefix"`
}

type Mysql struct {
	Host          string `yaml:"host"`
	Port          string `yaml:"port"`
	User          string `yaml:"user"`
	Pwd           string `yaml:"password"`
	Dbname        string `yaml:"dbname"`
	Charset       string `yaml:"charset"`
	Max_idle_conn int    `yaml:"maxIdle"`
	Max_open_conn int    `yaml:"maxOpen"`
	Debug         bool   `yaml:"debug"`
	ParseTime     bool   `yaml:"parseTime"`
}

type User struct {
	Name string `yaml:"name"`
	Pwd  string `yaml:"pwd"`
}

type Aliyun struct {
	AccessKeyId     string `yaml:"accessKeyId"`
	AccessKeySecret string `yaml:"accessKeySecret"`
	OssEndpoint     string `yaml:"oss_endpoint"`
	BucketName     string `yaml:"bucket_name"`
	UpfilePath    string  `yaml:"upfile_path"`
	CardclasspicPath    string  `yaml:"cardclasspic_path"`
}

type BussinessLimits struct {
	PhoneSms []string `yaml:"phone_sms"`
	IpSms    []string `yaml:"ip_sms"`
}

type BasNotify struct {
	Addr             string `yaml:"addr"`
	VerifyCodeSmsTmp string `yaml:"verifycode_sms_tmp"`
	SwitchFlag       bool   `yaml:"switch_flag"`
}

type MarketMain struct {
	FissionPackAddr  string `yaml:"fission_pack_addr"`
	LuckDrawPackAddr string `yaml:"luckdraw_pack_addr"`
}

type IdCardCheck struct {
	AppCode string `yaml:"appcode"`
}

type ChuangLan struct {
	AppId           string `yaml:"appId"`
	AppKey          string `yaml:"appKey"`
	IdcardCheck_url string `yaml:"idcardcheck_url"`
	UnnCheck_url    string `yaml:"unncheck_url"`
	Host            string `yaml:"host"`
}

type YunPian struct {
	ApiKey           string `yaml:"apikey"`
}

type Baidu struct {
	DwzToken string  `yaml:"dwz_token"`
	DwzUrl   string  `yaml:"dwz_url"`
}

type Task struct{
	Task_flag  bool  `yaml:"task_flag"`
	SheetFlag bool  `yaml:"sheet_flag"`
	SheetTicker int64  `yaml:"sheet_ticker"`
	IpsFlag bool  `yaml:"ips_flag"`
	IpsTicker int64  `yaml:"ips_ticker"`
	HisFlag bool  `yaml:"his_flag"`
	HisTicker int64  `yaml:"his_ticker"`
	NotifyFlag bool  `yaml:"notify_flag"`
	NotifyTicker int64  `yaml:"notify_ticker"`
	Unfinish5MNotifyFlag bool  `yaml:"unfinish_5m_notify_flag"`
	Unfinish5MNotifyTicker int64  `yaml:"unfinish_5m_notify_ticker"`
	Unfinish5HNotifyFlag bool  `yaml:"unfinish_5h_notify_flag"`
	Unfinish5HNotifyTicker int64  `yaml:"unfinish_5h_notify_ticker"`

	ActiveFlag bool  `yaml:"active_flag"`
	ActiveTicker int64  `yaml:"active_ticker"`
	YdhkUnfinishFlag bool  `yaml:"ydhk_unfinish_flag"`
	YdhkUnfinishTicker int64  `yaml:"ydhk_unfinish_ticker"`
}

type Jthk struct{
	Host string  `yaml:"host"`
	Url  string  `yaml:"url"`
	SearchUrl  string  `yaml:"search_url"`
	//Channel_id_19 string   `yaml:"channel_id_19"`
	//Product_id_19 string    `yaml:"product_id_19"`
	//Channel_id_39 string   `yaml:"channel_id_39"`
	//Product_id_39 string   `yaml:"product_id_39"`
	Referer_path  string   `yaml:"referer_path"`
	//Channel_id_19_oao string   `yaml:"channel_id_19_oao"`
	//Product_id_19_oao string   `yaml:"product_id_19_oao"`
	//Channel_id_39_oao string   `yaml:"channel_id_39_oao"`
	//Product_id_39_oao  string   `yaml:"product_id_39_oao"`
	Referer_path_oao  string   `yaml:"referer_path_oao"`
	ParterCode    string      `yaml:"partner_code"`
	ParterCodeArr []string   `yaml:"-"`
}

type Dxnbhk struct{
	Url string                `yaml:"url"`
	Partners []DxnbhkReseller     `yaml:"reseller"`
	PartnerIndex int     `yaml:"-"`
}

type DxnbhkReseller struct{
	Name string              `yaml:"name"`
	ResellerId string              `yaml:"reseller_id"`
	ProductId string         `yaml:"product_id"`
	MaxNum int                  `yaml:"max_num"`
	PartnerGoodsCode string                  `yaml:"partnergoods_code"`
	Count     int   `yaml:"-"`
}