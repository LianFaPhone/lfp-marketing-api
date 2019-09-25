package config

import (
	"fmt"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"github.com/ulule/limiter"
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
	if GConfig.Cache.ActivityMaxKey < 100 {
		GConfig.Cache.ActivityMaxKey = 100
	}
	if GConfig.Cache.ActivityTimeout < 300 {
		GConfig.Cache.ActivityTimeout = 3600
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

	for i:=0; i < len(GConfig.BussinessLimits.PhoneSms); i++ {
		rate, err := limiter.NewRateFromFormatted(GConfig.BussinessLimits.PhoneSms[i])
		if err != nil {
			return err
		}
		GPreConfig.PhoneSmsLimits = append(GPreConfig.PhoneSmsLimits, &rate)
	}
	for i:=0; i < len(GConfig.BussinessLimits.IpSms); i++ {
		rate, err := limiter.NewRateFromFormatted(GConfig.BussinessLimits.IpSms[i])
		if err != nil {
			return err
		}
		GPreConfig.IpSmsLimits = append(GPreConfig.IpSmsLimits, &rate)
	}
	return nil
}

type PreConfig struct{
	PhoneSmsLimits []*limiter.Rate
	IpSmsLimits []*limiter.Rate
}

type Config struct {
	Server  System      `yaml:"system"`
	Redis   Redis       `yaml:"redis"`
	Db      Mysql       `yaml:"mysql"`
	Cache   Cache       `yaml:"cache"`
	User    User        `yaml:"user"`
	//Aws     Aws          `yaml:"aws"`
	BussinessLimits BussinessLimits   `yaml:"bussiness_limits"`
	BasNotify       BasNotify         `yaml:"bas_notify"`
	MarketMain     MarketMain         `yaml:"marketing_main"`
	IdCardCheck    IdCardCheck        `yaml:"idcard_check"`

}

type System struct {
	Port    string      `yaml:"port"`
	Debug   bool        `yaml:"debug"`
	LogPath string      `yaml:"log_path"`
	Monitor string      `yaml:"monitor"`
	BkPort    string      `yaml:"bk_port"`
}

type Cache struct {
	ActivityMaxKey   int      `yaml:"activity_max_key"`
	ActivityTimeout  int      `yaml:"activity_timeout"`
	SponsorActivityMaxKey   int      `yaml:"sponsor_activity_max_key"`
	SponsorActivityTimeout  int      `yaml:"sponsor_activity_timeout"`
	SponsorMaxKey   int      `yaml:"sponsor_max_key"`
	SponsorTimeout  int      `yaml:"sponsor_timeout"`
	PageMaxKey   int     	 `yaml:"page_max_key"`
	PageTimeout   int     	 `yaml:"page_timeout"`
	ShareInfoMaxKey    int   `yaml:"shareinfo_max_key"`
	ShareInfoTimeout   int   `yaml:"shareinfo_timeout"`
	RobberMaxKey          int `yaml:"robber_max_key"`
	RobberTimeout         int   `yaml:"robber_timeout"`
}

type Redis struct {
	Network     string  `yaml:"network"`
	Host        string  `yaml:"host"`
	Port        string  `yaml:"port"`
	Password    string  `yaml:"password"`
	Database    int  `yaml:"database"`
	MaxIdle     int     `yaml:"maxIdle"`
	MaxActive   int     `yaml:"maxActive"`
	IdleTimeout int     `yaml:"idleTimeout"`
	Prefix      string  `yaml:"prefix"`
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

type User struct{
	Name string   `yaml:"name"`
	Pwd  string   `yaml:"pwd"`
}

type Aliyun struct{
	AccessKeyId string   `yaml:"accessKeyId"`
	AccessKeySecret  string    `yaml:"accessKeySecret"`
	OssEndpoint  string    `yaml:"oss_endpoint"`
}

type BussinessLimits struct {
	PhoneSms  []string `yaml:"phone_sms"`
	IpSms  []string `yaml:"ip_sms"`
}

type  BasNotify struct {
	Addr                   string `yaml:"addr"`
	VerifyCodeSmsTmp       string `yaml:"verifycode_sms_tmp"`
}

type MarketMain struct{
	FissionPackAddr       string     `yaml:"fission_pack_addr"`
	LuckDrawPackAddr       string     `yaml:"luckdraw_pack_addr"`
}

type IdCardCheck struct{
	AppCode       string     `yaml:"appcode"`
}