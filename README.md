错误号
=================
```
type EnumBasErr int

const BasErrBegin EnumBasErr = 1000000          //区别老版本
const BasMoreErrBegin EnumBasErr = 1001000  //用户自定义开始

const (
	//系统错误(公用)大家觉得需要自行添加
	BASERR_SUCCESS EnumBasErr = 0 //成功（不可加上起始值）
	//请求相关错误，一般是由于请求格式的输入错误导致。
	BASERR_INVALID_PARAMETER            EnumBasErr = 1000000 + 100 //请求参数无效。    一般用在参数缺失或者格式不对
	BASERR_UNSUPPORTED_METHOD           EnumBasErr = 1000000 + 101 //未知的方法。     请求的url的path未定义或者其它无法识别的方法
	BASERR_URLBODYTOOLARGE_METHOD       EnumBasErr = 1000000 + 102 //请求body过长。   数据长度过长
	BASERR_INVALID_CONTENTLENGTH_METHOD EnumBasErr = 1000000 + 103 //无效的ContentLength。这个是http的header中的
	BASERR_QID_ALREADY_EXISTS           EnumBasErr = 1000000 + 104 //qid已存在。      这个用在异步请求中，标识每次请求，理解成流水号。

	//授权 加密 账户 相关安全错误等, 200不要了
	BASERR_TOKEN_INVALID          EnumBasErr = 1000000 + 201 //无效的访问令牌。  token无效，可与过期混用，貌似目前就admin喝bkadmin用到，单点登录使用。
	BASERR_TOKEN_EXPIRED          EnumBasErr = 1000000 + 202 //访问令牌过期。    同上
	BASERR_INCORRECT_SIGNATURE    EnumBasErr = 1000000 + 203 //无效的签名。      签名无效
	BASERR_UNAUTHORIZED_METHOD    EnumBasErr = 1000000 + 204 //未授权的方法。    该用户方法无权限
	BASERR_UNAUTHORIZED_PARAMETER EnumBasErr = 1000000 + 205 //未授权的参数。    同上
	BASERR_ILLEGAL_DATA           EnumBasErr = 1000000 + 206 //非法数据，       解密失败或者无法解析
	BASERR_TOKEN_PENDING          EnumBasErr = 1000000 + 207 //访问令牌状态未确认。    同上

	BASERR_INCORRECT_ACCOUNT_PWD EnumBasErr = 1000000 + 220 //账号密码不对
	BASERR_INCORRECT_GA_PWD      EnumBasErr = 1000000 + 221 //GA密码错误
	BASERR_INCORRECT_PWD         EnumBasErr = 1000000 + 222 //密码错误
	//
	BASERR_WHITELIST_OUTSIDE EnumBasErr = 1000000 + 251 //白名单之外
	BASERR_BLACKLIST_INSIDE  EnumBasErr = 1000000 + 252 //黑名单内
	BASERR_FROZEN_ACCOUNT    EnumBasErr = 1000000 + 253 //账户未激活        账户处于初始状态
	BASERR_FROZEN_METHOD     EnumBasErr = 1000000 + 254 //功能未激活        部分功能未开启
	BASERR_BLOCK_ACCOUNT     EnumBasErr = 1000000 + 255 //账户被阻，        无法登陆，可能是长期不登录或者注销
	BASERR_DATAFROM_INVALID  EnumBasErr = 1000000 + 256 //数据源不对        数据的来源不合法

	//300-399： 预先定义的服务内部错误（对外显示错误信息）
	BASERR_OBJECT_NOT_FOUND      EnumBasErr = 1000000 + 301 //指定的对象不存在
	BASERR_OBJECT_EXISTS         EnumBasErr = 1000000 + 302 //指定的对象已存在
	BASERR_OBJECT_DATA_NOT_FOUND EnumBasErr = 1000000 + 303 //指定的对象数据不存在
	BASERR_ACCOUNT_NOT_FOUND     EnumBasErr = 1000000 + 304 //账号不存在
	BASERR_OBJECT_DATA_SAME      EnumBasErr = 1000000 + 305 //指定对象数据和上次一致。   这个一般在更新的时候，没必要重复更新。
	BASERR_OBJECT_DATA_NOT_SAME  EnumBasErr = 1000000 + 306 //指定对象数据不相同。   这
	BASERR_OBJECT_DATA_NULL      EnumBasErr = 1000000 + 307 //指定对象数据为空。   这
	BASERR_OBJECT_ZERO           EnumBasErr = 1000000 + 308 //指定对象为0。
	BASERR_OPERATE_FREQUENT      EnumBasErr = 1000000 + 311 //操作太频繁，超过了限制。   操作频率限制
	BASERR_INCORRECT_FORMAT      EnumBasErr = 1000000 + 312 //格式出错。
	BASERR_INCORRECT_PUBKEY      EnumBasErr = 1000000 + 313 //无效的公钥格式。  这个不知是否与上个重复
	BASERR_INCORRECT_FREQUENT    EnumBasErr = 1000000 + 314 //错误太频繁。
	BASERR_NOT_ALLOW_STATE       EnumBasErr = 1000000 + 315 //当前状态不允许
	BASERR_OBJECT_BE_USED        EnumBasErr = 1000000 + 316 //指定的对象已被占用

	//500-999： 预先定义的服务内部错误(不对外显示错误信息，错误信息 统一表示为系统未知错误)
	BASERR_SERVICE_UNKNOWN_ERROR         EnumBasErr = 1000000 + 500 //服务未知错误，     一般500到1000之间的错误。$实在不知道写什么错误就用这个$
	BASERR_INTERNAL_SERVICE_ACCESS_ERROR EnumBasErr = 1000000 + 501 //内部服务访问错误。  调用另一个服务出错, 或者返回的数据 不合理或者有错
	BASERR_INVALID_OPERATION             EnumBasErr = 1000000 + 502 //无效的操作方法。     可能这个方法未开发。 这个绝对是多余了

	BASERR_DATA_PACK_ERROR                 EnumBasErr = 1000000 + 510 //数据打包失败，     json或者pb等打包失败
	BASERR_DATA_UNPACK_ERROR               EnumBasErr = 1000000 + 511 //数据解包失败，     区别无效参数，这个一般用在请求其他服务的返回数据无法解包
	BASERR_DATABASE_ERROR                  EnumBasErr = 1000000 + 512 //数据库操作出错，请重试。 包括redis、mysql等
	BASERR_SERVICE_TEMPORARILY_UNAVAILABLE EnumBasErr = 1000000 + 513 //服务暂不可用      $这个应该也是你困惑的时候该选的$
	BASERR_SEND_TIMEOUT                    EnumBasErr = 1000000 + 514 //发送超时
	BASERR_STATE_NOT_ALLOW                 EnumBasErr = 1000000 + 515 //当前状态不允许

	BASERR_UNKNOWN_BUG EnumBasErr = 1000000 + 555 //未知bug请修复，这个明显是bug

	BASERR_INTERNAL_CONFIG_ERROR EnumBasErr = 1000000 + 556 //恭喜只是配置错误，不是bug。可以联系运维了。

	BASERR_SYSTEM_INTERNAL_ERROR EnumBasErr = 1000000 + 1000 //系统内部错误。      一些内存不足报错啊，什么系统错误啊都可以用这个表示。$这个应该也是你困惑的时候该选的$

	//开发者自定义错误BasMoreErrBegin, 每类应用使用区间50，并且BASERR_ServerName_开头
	//admin和bkadmin 1--49，原则上这些错误的描述信息表现为内部服务错误， 0不可用，与BASERR_SYSTEM_INTERNAL_ERROR 重复
	BASERR_ADMIN_INVALID_VERIFY_STATUS EnumBasErr = 1001000 + 1 //无效的验证状态, 上次已验证的东西超过使用次数，或者 已验证的验证码、手机号、邮箱号等不匹配
	BASERR_ADMIN_INCORRECT_VERIFYCODE  EnumBasErr = 1001000 + 2 //错误的验证码

	//50-99
	BASERR_BASNOTIFY_AWS_ERR                        EnumBasErr = 1001000 + 50
	BASERR_BASNOTIFY_LANCHUANG_ERR                  EnumBasErr = 1001000 + 51
	BASERR_BASNOTIFY_TWL_ERR                        EnumBasErr = 1001000 + 52
	BASERR_BASNOTIFY_TEMPLATE_DEAD                  EnumBasErr = 1001000 + 53
	BASERR_BASNOTIFY_RECIPIENT_EMPTY                EnumBasErr = 1001000 + 54
	BASERR_BASNOTIFY_TEMPLATE_PARSE_FAIL            EnumBasErr = 1001000 + 55
	BASERR_BASNOTIFY_Nexmo_ERR                      EnumBasErr = 1001000 + 56
	BASERR_BASNOTIFY_RongLianYun_ERR                EnumBasErr = 1001000 + 57
	BASERR_BASNOTIFY_DingDing_ERR                   EnumBasErr = 1001000 + 58
	BASERR_BASNOTIFY_DingDing_QunName_NotSet_ERR    EnumBasErr = 1001000 + 59
	BASERR_BASNOTIFY_DingDing_QunName_NotConfig_ERR EnumBasErr = 1001000 + 60
	BASERR_BASNOTIFY_Aliyun_Intl_ERR                EnumBasErr = 1001000 + 61

	//1000-1049 红包裂变

	//1100-1149
	BASERR_CARDMARKET_PHONEPOOL_LOCK_FAIL EnumBasErr = 1001000 + 1100 //号码加锁失败
	BASERR_CARDMARKET_PHONEPOOL_UNLOCK_FAIL EnumBasErr = 1001000 + 1101 //号码解锁失败
	BASERR_CARDMARKET_PHONEPOOL_USE_FAIL EnumBasErr = 1001000 + 1102 //号码占用失败
	BASERR_CARDMARKET_PHONEPOOL_UNUSE_FAIL EnumBasErr = 1001000 + 1103 //号码解占用失败

	BASERR_CARDMARKET_PHONECARD_APPLY_FAID_AND_SHOW EnumBasErr = 1001000 + 1110 //卡申请失败显示弹窗

)

```

0.渠道管理接口
======

#### 渠道添加
```
运营商（常量）
ISP_Dianxin = 1  电信  
ISP_YiDong = 2  移动  
ISP_LiTong = 3  联通

POST /v1/bk/market/pdparter/add
参数header里包含token字段
参数
{
	"isp":1,
	"detail":"详情",
	"name":"移动花卡19",
	"code":"代码",
	"gsd_province":"归属地-省",
	"gsd_city":"归属地-市",
	"gsd_province_code":"归属地省代码",
	"gsd_city_code":"归属地市代码",
	"made_in":"产地",
	"no_exp_addr":"上海市,宁波市", //不发货的城市组
	"min_age":18,  // 最小年龄
	"max_age": 60,  //最大年龄
	"limit_card_count": 3,    //卡最大数量
	"limit_card_period":3600,  //卡最大数量限制周期
	"idcard_five_flag": 1,     //一证五户 开关
	"idcard_five_period":3600,   //一证五户周期
	"repeat_exp_addr_count": 3,   //重复快递地址个数
	"repeat_exp_addr_period":3600,  //重复快递地址周期
	"repeat_phone_count":7,        //重复手机个数
	"repeat_phone_period":3600,     //重复手机周期
	"prefix_path":"",               //路径前缀
	"idcard_display":1,   //   是否显示身份证
	"sms_flag": 0,             //短信验证否
	"stock":100,               //库存
	"production_notes":"nihao啊"   //生产日志


}
返回
{
    "code":0,
    "message":"success",
    "data":{
        "id": 1,
       "isp":1,
    	"detail":"详情",
	    "name":"移动花卡19",
	    "code":"代码",
	    "gsd_province":"归属地-省",
	    "gsd_city":"归属地-市",
	    "gsd_province_code":"归属地省代码",
	    "gsd_city_code":"归属地市代码",
	    "made_in":"产地",
	    "no_exp_addr":"上海市,宁波市", //不发货的城市组
	    "min_age":18,  // 最小年龄
	    "max_age": 60,  //最大年龄
	    "limit_card_count": 3,    //卡最大数量
	    "limit_card_period":3600,  //卡最大数量限制周期
    	"idcard_five_flag": 1,     //一证五户 开关
	    "idcard_five_period":3600,   //一证五户周期
	    "repeat_exp_addr_count": 3,   //重复快递地址个数
	    "repeat_exp_addr_period":3600,  //重复快递地址周期
	    "repeat_phone_count":7,        //重复手机个数
	    "repeat_phone_period":3600,     //重复手机周期
	    "prefix_path":"",               //路径前缀
	    "idcard_display":1,   //   是否显示身份证
    	"sms_flag": 0,             //短信验证否
    	"stock":100,               //库存
    	"production_notes":"nihao啊"   //生产日志
        "valid":1,
        "created_at": 12345, //unix时间戳，秒
        "updated_at":12345   //秒
    }
}
```

#### 渠道更新
```
POST /v1/bk/market/pdparter/update
参数header里包含token字段
参数
{
    "id":1,   //必传
    "isp":2,
	"detail":"详情2",
	"name":"移动花卡39",
	"code":"001",
	"gsd_province":"归属地-省",
	"gsd_city":"归属地-市",
	"gsd_province_code":"归属地省代码",
	"gsd_city_code":"归属地市代码",
	"made_in":"产地",
	"no_exp_addr":"上海市,宁波市",
	"min_age":18,
	"max_age": 60,  
	"limit_card_count": 3,   
	"limit_card_period":3600,  
	"idcard_five_flag": 1,  
	"idcard_five_period":3600,  
	"repeat_exp_addr_count": 3,   
	"repeat_exp_addr_period":3600, 
	"repeat_phone_count":7,
	"repeat_phone_period":3600,    
	"prefix_path":"",              
  	"idcard_display":1, 
	"sms_flag": 1,             
	"stock":100,               
	"production_notes":"nihao啊"   
    
}
返回
{
    "code": 0,
    "message": "Success",
    "data": {
        "id": 1,
        "isp": 2,
        "name": "移动花卡39",
        "code": "001",
        "detail": "详情2",
        "gsd_province": "归属地-省",
        "gsd_city": "归属地-市",
        "gsd_province_code": "归属地省代码",
        "gsd_city_code": "归属地市代码",
        "made_in": "产地",
        "no_exp_addr": "上海市,宁波市",
        "min_age": 18,
        "max_age": 60,
        "limit_card_count": 3,
        "limit_card_period": 3600,
        "idcard_five_flag": 1,
        "idcard_five_period": 3600,
        "repeat_exp_addr_count": 3,
        "repeat_exp_addr_period": 3600,
        "repeat_phone_count": 7,
        "repeat_phone_period": 3600,
        "prefix_path": "",
        "sms_flag": 1,
        "idcard_display": 1,
        "stock": 100,
        "production_notes": "nihao啊",
        "valid": 1,
        "created_at": 1582786424,
        "updated_at": 1582787730
    }
}
```

#### 渠道查询
```
POST /v1/bk/market/pdparter/get
参数header里包含token字段
参数
{
    "id": 1   
}
返回
{
    "code": 0,
    "message": "Success",
    "data": {
        "id": 1,
        "isp": 2,
        "name": "移动花卡39",
        "code": "001",
        "detail": "详情2",
        "gsd_province": "归属地-省",
        "gsd_city": "归属地-市",
        "gsd_province_code": "归属地省代码",
        "gsd_city_code": "归属地市代码",
        "made_in": "产地",
        "no_exp_addr": "上海市,宁波市",
        "min_age": 18,
        "max_age": 60,
        "limit_card_count": 3,
        "limit_card_period": 3600,
        "idcard_five_flag": 1,
        "idcard_five_period": 3600,
        "repeat_exp_addr_count": 3,
        "repeat_exp_addr_period": 3600,
        "repeat_phone_count": 7,
        "repeat_phone_period": 3600,
        "prefix_path": "",
        "sms_flag": 1,
        "idcard_display": 1,
        "stock": 100,
        "production_notes": "nihao啊",
        "valid": 1,
        "created_at": 1582786424,
        "updated_at": 1582787730
    }
}
```

#### 获取所有渠道
```
POST /v1/bk/market/pdparter/gets
参数header里包含token字段
参数
{
    "valid": 1  //可选   
}
返回
{
    "code": 0,
    "message": "Success",
    "data": [
        {
            "id": 1,
            "isp": 2,
            "name": "移动花卡39",
            "code": "001",
            "detail": "详情2",
            "gsd_province": "归属地-省",
            "gsd_city": "归属地-市",
            "gsd_province_code": "归属地省代码",
            "gsd_city_code": "归属地市代码",
            "made_in": "产地",
            "no_exp_addr": "上海市,宁波市",
            "min_age": 18,
            "max_age": 60,
            "limit_card_count": 3,
            "limit_card_period": 3600,
            "idcard_five_flag": 1,
            "idcard_five_period": 3600,
            "repeat_exp_addr_count": 3,
            "repeat_exp_addr_period": 3600,
            "repeat_phone_count": 7,
            "repeat_phone_period": 3600,
            "prefix_path": "",
            "sms_flag": 1,
            "idcard_display": 1,
            "stock": 100,
            "production_notes": "nihao啊",
            "valid": 1,
            "created_at": 1582786424,
            "updated_at": 1582787730
        }
    ]
}
```

#### 套餐分页查询
```
POST /v1/bk/market/pdparter/list
参数header里包含token字段
{

    "isp": 1,
    "valid":1,   //可选
    
	"page":1,    //必选
	"size":100   //必选
}

返回
{
	"total_result":10, //总条数
	"has_next": true,  //是否还有下一页
	"page":1,          //当前页
	"size":10,        //当前返回条数
	"list":[
		{
		"id": 1,
                "isp": 2,
                "name": "移动花卡39",
                "code": "001",
                "detail": "详情2",
                "gsd_province": "归属地-省",
                "gsd_city": "归属地-市",
                "gsd_province_code": "归属地省代码",
                "gsd_city_code": "归属地市代码",
                "made_in": "产地",
                "no_exp_addr": "上海市,宁波市",
                "min_age": 18,
                "max_age": 60,
                "limit_card_count": 3,
                "limit_card_period": 3600,
                "idcard_five_flag": 1,
                "idcard_five_period": 3600,
                "repeat_exp_addr_count": 3,
                "repeat_exp_addr_period": 3600,
                "repeat_phone_count": 7,
                "repeat_phone_period": 3600,
                "prefix_path": "",
                "sms_flag": 1,
                "idcard_display": 1,
                "stock": 100,
                "production_notes": "nihao啊",
                "valid": 1,
                "created_at": 1582786424,
                "updated_at": 1582787730
 		}
	]
}
```

#### 状态更新
```
POST /v1/bk/market/pdparter/status-update
{
    "id":1,
    "valid": 1
}

返回
{
    
}
```


1.渠道商品管理接口
=================  

#### 图片批量上传
```
POST   /v1/bk/market/oss-market/upload
参数header里包含token字段
MultipartForm   名称 file

返回值
[{
	"file":"123.jpg",
	"addr":"www.aaaa/aaa/123.jpg"
}]

```


#### 商品添加  
```
POST /v1/bk/market/pdpartergoods/add
参数header里包含token字段
{
	"partner_id": 1,   //外键
	
	"jd_code":"RtXFST",  //京东码
	"name":"19套餐",    //名称
	"url_param":"",    //额外参数
	"detail":"描述",   //描述
	"img_url":"",      //图片地址
	"short_chain":"短链",
	"long_chain" :"长链",
	"third_long_chain":"第三方链"
 }
返回值
{
    "code": 0,
    "message": "Success",
    "data": {
        "id": 1,
        "partner_id": 1,
        "code": "93bPMf0Ek9",
        "jd_code": "RtXFST",
        "name": "19套餐",
        "url_param": "",
        "detail": "描述",
        "short_chain": "短链",
        "img_url": "",
        "long_chain": "",
        "third_long_chain": "第三方链",
        "valid": 1,
        "created_at": 1582788765,
        "updated_at": 1582788765
    }
}
```
#### 商品更新
```
POST /v1/bk/market/pdpartergoods/update
##参数header里包含token字段
##有些字段可不传
{
    "id":1,    //必传
	"partner_id": 1,
	
	"jd_code":"RtXFST",
	"name":"19套餐2",
	"url_param":"",
	"detail":"描述2",
	"img_url":"",
	"short_chain":"短链",
	"long_chain" :"长链",
	"third_long_chain":"第三方链"
}
返回值
{
    "code": 0,
    "message": "Success",
    "data": {
        "id": 1,
        "partner_id": 1,
        "code": "93bPMf0Ek9",
        "jd_code": "RtXFST",
        "name": "19套餐2",
        "url_param": "",
        "detail": "描述2",
        "short_chain": "短链",
        "img_url": "",
        "long_chain": "长链",
        "third_long_chain": "第三方链",
        "valid": 1,
        "created_at": 1582788765,
        "updated_at": 1582788996
    }
}
```
#### 商品分页查询
```
POST /v1/bk/market/pdpartergoods/list
参数header里包含token字段
{
	"partner_id": 1,
	"page":1,
	"size":10
}

返回
{
	"total_result":10, //总条数
	"has_next": true,  //是否还有下一页
	"page":1,          //当前页
	"size":10,        //当前返回条数
	"list":[
		{
		 "id": 1,
                "partner_id": 1,
                "code": "93bPMf0Ek9",
                "jd_code": "RtXFST",
                "name": "19套餐2",
                "url_param": "",
                "detail": "描述2",
                "short_chain": "短链",
                "img_url": "",
                "long_chain": "长链",
                "third_long_chain": "第三方链",
                "valid": 1,
                "created_at": 1582788765,
                "updated_at": 1582788996
 		}
	]
}
```
#### 商品查询
```
POST /v1/bk/market/pdpartergoods/get
##参数header里包含token字段
##有些字段可不传
{
	"id":1,         //必传
 }
返回值
{
    "code":0,
    "message":"success",
    "data":{
         "id": 1,
                "partner_id": 1,
                "code": "93bPMf0Ek9",
                "jd_code": "RtXFST",
                "name": "19套餐2",
                "url_param": "",
                "detail": "描述2",
                "short_chain": "短链",
                "img_url": "",
                "long_chain": "长链",
                "third_long_chain": "第三方链",
                "valid": 1,
                "created_at": 1582788765,
                "updated_at": 1582788996
    }
}
```
#### 获取所有商品
```
POST /v1/bk/market/pdpartergoods/gets
##参数header里包含token字段
##有些字段可不传
{
	"valid":1,         //可选
 }
返回值
{
    "code":0,
    "message":"success",
    "data":[{
         "id": 1,
                "partner_id": 1,
                "code": "93bPMf0Ek9",
                "jd_code": "RtXFST",
                "name": "19套餐2",
                "url_param": "",
                "detail": "描述2",
                "short_chain": "短链",
                "img_url": "",
                "long_chain": "长链",
                "third_long_chain": "第三方链",
                "valid": 1,
                "created_at": 1582788765,
                "updated_at": 1582788996
    }]
}
```
#### 状态更新
POST /v1/bk/market/pdpartergoods/status-update
{
    "id":1,
    "valid":1
}
返回
{
    "code":0,
    "message":"success"
}
   
   
2.订单管理
=======================

#### 获取设备类型
```
获取手机操作系统列表
POST /v1/bk/market/osphone/gets
返回值
{
    "code": 0,
    "message": "Success",
    "data": [
        {
            "tp": 1,
            "name": "Android"
        },
        {
            "tp": 2,
            "name": "Iphone"
        },
        {
            "tp": 3,
            "name": "Ipad"
        },
        {
            "tp": 4,
            "name": "Other"
        }
    ]
}
```

#### 获取订单所有状态
```
获取订单状态列表
POST /v1/bk/market/orderstatus/gets
返回值
{
    "code": 0,
    "message": "Success",
    "data": [
        {
            "status": 1,
            "name": "新订单已完成"
        },
         {
            "status": 10,
            "name": "新订单未完成"
        },
        {
            "status": 2,
            "name": "已导出"
        },
        {
            "status": 3,
            "name": "已发货"
        },
        {
            "status": 5,
            "name": "待处理"
        },
        {
            "status": 6,
            "name": "已处理"
        },
        {
            "status": 4,
            "name": "回收站"
        },
        {
            "status": 7,
            "name": "未匹配"
        },
        {
            "status": 8,
            "name": "已激活"
        }
    ]
}
```

#### 订单查询
```
POST
/v1/bk/market/card_order/list
/v1/bk/market/card_order/list-all   所有订单
/v1/bk/market/card_order/list-new   新订单（已完成）
/v1/bk/market/card_order/list-new-unfinish   新订单(未完成)
/v1/bk/market/card_order/list-export  已导出订单
/v1/bk/market/card_order/list-deliver  已发货订单
/v1/bk/market/card_order/list-waitdone     待处理订单
/v1/bk/market/card_order/list-alreadydone   已处理订单
/v1/bk/market/card_order/list-recyclebin   回收站
/v1/bk/market/card_order/list-unmatch     未匹配订单
/v1/bk/market/card_order/list-activated   已激活订单

参数header里包含token字段
参数
{
	"true_name":"小明", //可选
	"idcard":"身份证",  //可选
	"phone":"电话",     //可选
	"new_phone":"选择得新号码",  //可选
	"province":"省份",  //可选
	"city":"市",        //可选
	"area":"区县",      //可选
	"partner_id": 1,   //渠道id
	"partner_goods_code":"12345", //渠道商品码
	"isp":1,//运营商
	"device_os_tp":1   //设备类型，1 安卓，2 苹果，3ipad，4 其它， 可选
	"status": 1,   //订单状态，可选
	"min_ips": 3,  //ip重复次数，可选
	"like_str":"1aaaa",  //模糊查询，可选包
	"start_created_at": 123,   //订单开始创建时间，可选
	"end_created_at":456,      //订单结束创建时间，可选
	"start_deliver_at":,       //订单开始发货时间，可选
	"end_deliver_at":, 
	//订单结束发货时间，可选
	"black_switch": 0, //石否查黑名单，0 不查，1查
    "idcard_pic_flag":0, //1已上传身份证照片，0未上传，可选
	"page":1,
	"size":10
}

返回值
{
    "code": 0,
    "message": "Success",
    "data": {
        "total_result": 7, //总个数
        "has_next": true,  //是否有下一页
        "page": 1,   //目前页号
        "size": 2,   //一页个数
        "list": [
            {
                "id": 9,
                "order_no": "D1909241043320004", //订单号，唯一，系统生成
               	"partner_id": 1,   //渠道id
            	"partner_goods_code":"12345", //渠道商品码
            	"isp":1,//运营商
                "class_name": "王卡", //套餐名称
                "status": 1,   //套餐状态
                "status_name": "新订单", //套餐状态名称
                "true_name": "xufengping2345",//姓名
                "idcard": "身份证",
                "country_code": "0086",
                "phone": "15216614713", //手机号
                "province": "北京市", //省
                "province_code": "100", //省代号
                "city": "北京市",     //市
                "city_code": "100100",     //市代码
                "area": "海定区",     //区县
                "area_code":"1001001", //区代码
                "town": "黑庄户",     //镇街道
                "address": "1hao号lou", //详细地址
                "new_phone": "15216617272", //新手机号
                "ip": "127.0.0.1",//ip地址
## #                 "min_ips":1, //ip重复次数
                 "ICCID":"1222",        //ICCID手机号唯一序列号
               "min_ips":1, //ip重复下单数
               "express":"快递公司",
               "express_no":"快递单号",
               "express_remark":"快递备注",
               "deliver_at":1234 , //发货时间，秒
               "guishudi":"归属地",
               "idcard_audit":1,//身份证审核状态
               "is_blacklist": 0, //是否黑名单
               "third_order_no":"第三方订单号",
               "device_os_name":"设备名称",
                "idcard_pic_flag":0,
                "valid": 1, //有效性
                "created_at": 1569293012,//下单时间
                "updated_at": 1569293012//更新时间
            }
        ]
    }
}

```

#### 订单单个更新
```
POST /v1/bk/market/card_order/update
参数header里包含token字段
参数
{
  "order_no":"D1234",  //必选
  "true_name": "hahah", //姓名
  "phone":"12333333",   //手机号
  ""new_phone":"新手机号",
  "idcard":"3302828288", //身份证
  "ICCID":"1222",        //ICCID手机号唯一序列号
  "express":"圆通快递",
  "express_no":"快递单号",
  "express_remark":"快递备注",
  "deliver_at": 1223434,  //快递发货时间
 
  "status":1
  "province":"省",
  "city":"市",
  "area":"区县",
  "town":"镇街道",
   "address":"sjsjaaa"
  
    
  
}
返回值
{
    "code": 0,
    "message": "Success"
}
```

#### 订单状态批量更新
```
POST /v1/bk/market/card_order/status-sets
参数
{
  "order_no":["D123452"], //订单号
    "status": 1           //状态
}
返回值
{
    "code": 0,
    "message": "Success"
}

```

#### 批量导入订单额外信息（快递信息）
```
POST /v1/bk/market/card_order/express-import
参数
multipartform   名称 file
返回值
{
    "code": 0,
    "message": "Success",
    "data": {
        "fail_count": 2,
        "succ_count": 2
    }
}
```

#### 激活导入
```
POST /v1/bk/market/card_order/active-import
参数
[{
    "new_phone":"124455"
},
{
    "new_phone":"124455"
}
]
返回值
{
    "code": 0,
    "message": "Success",
    "data": {
        "fail_count": 2,
        "succ_count": 2
    }
}
```

#### 导出
```
POST /v1/bk/market/card_order/file/create
{

	"status": 1,   //订单状态，可选
	"start_created_at": 123,   //订单开始创建时间，可选
	"end_created_at":456,      //订单结束创建时间，可选
	"black_switch": 0, //石否查黑名单，0 不查，1查
	"page":1,
	"size":10
}

返回值
{
    "code": 0,
    "message": "Success",
    "data": "2020-02-23-09-02-41-47255.xlsx"
}

Get /v1/bk/market/card_order/file/get?filename=2020-02-23-09-02-41-47255.xlsx
返回值
{
     "code": 0,
    "message": "Success",
    "data": "http://file.lfcxwifi.com/2020-02-23-09-02-41-47255.xlsx"
}


```

#### 查询订单的照片信息
```
POST /v1/bk/market/card_order_idcardpic/get
{
    "order_no":"Deeee343444",
}
返回
{
    "order_no":"Deeee343444",
    "pic_url1":"正面照图片地址",
    "pic_url2":"反面照",
    "pic_url3":"免冠照"
}
```

#### 查询订单的日志信息
```
POST /v1/bk/market/card_order_log/list
{
    "order_no":"Deeee343444",
    "page":1,
    "size":10
}
返回
{
    "code": 0,
    "message": "Success",
    "data": {
        "total_result": 7, //总个数
        "has_next": true,  //是否有下一页
        "page": 1,   //目前页号
        "size": 2,   //一页个数
        "list": [
            {
                 "order_no":"Deeee343444",
                 "log":"message", //日志信息
                 "created_at": 1243445,
                 "updated_at":122222
            }
        ]
    }
}
```

3.地址管理
=======================

#### 获取所有省
```
Get  /v1/bk/market/area/bsprovince/gets

返回
{
    "code": 0,
    "message": "Success",
    "data": [
        {
            "id": 31,
            "code": "650000",
            "name": "新疆维吾尔自治区",
            "short_name": "新疆",
            "valid": 1,
            "created_at": 1569938343,
            "updated_at": 1569938343
        },
        {
            "id": 30,
            "code": "640000",
            "name": "宁夏回族自治区",
            "short_name": "宁夏",
            "valid": 1,
            "created_at": 1569938338,
            "updated_at": 1569938338
        }
    ]
}

```
#### 根据省获取市
```
GET  /v1/bk/market/area/bscity/gets?province_code=410000

返回
{
    "code": 0,
    "message": "Success",
    "data": [
        {
            "id": 169,
            "code": "419001",
            "province_code": "410000",
            "name": "济源市",
            "short_name": "济源",
            "valid": 1,
            "created_at": 1569938128,
            "updated_at": 1569938128
        }
    ]
}

```