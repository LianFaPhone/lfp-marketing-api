
====================

#### 获取商品详情
```
Get /v1/ft/market/partnergoods/get?code=2838737
返回
{
    "code":"2838737",       //商品码
    "url_param":"k=v&k2=v2", //接口额外参数
    "img_url":"",    //图片地址
    "no_exp_addr":"上海,宁波",  //不发快递的地址
    "min_age":18,    //最小年龄
    "max_age":60,    //最大年龄
    "sms_flag":0 ,     //0短信不验证，1短信验证
    "idcard_display":0 //1显示身份证框
}
```

#### 下单
```
POST /v1/ft/market/ydhk/apply
{
    "phone":"1234455",
    "new_phone":"122222",
    "true_name":"nihao",
    "idcard":"330",     //身份证

    "province_code":"100",
    "city_code":"200",

    "express_province":"浙江省",
    "express_province_code":"200",
     "express_city":"嘉兴市",
    "express_city_code":"100",
     "express_district":"秀洲区"
    "express_district_code":"10000"
    "express_address":"2nihao",

    "token":"sddd",

     "partner_id":1,   //渠道号
    "partner_goods_code":2,  //商品码
    "isp":1,
    "ip":"1233333",
}
返回
{
    "code":0,
    "message":"success",
    "data":{
        "third_order_no":"420200202", //第三方订单号
        "oao_model": true             //快递能到否
    }
}

```
#### 下单返回解说
code为 1002110 则弹窗显示错误内容。  

code为其它非0 则显示 请稍后重试。  

code为0，且oao_model为true，则弹出一个网页，网页有两个按钮“为我上门激活”和“立即上传照片”。为我上门激活调用新接口(offline-active);立即上传照片，。则调用下一个接口获取url（idcheckurl-get），并重定向到新url。

code为0，且oao_model为false。则调用下一个接口获取url（idcheckurl-get），并重定向到新url。


#### 获取上传图片地址

```
POST /v1/ft/market/ydhk/idcheckurl-get
{
    "third_order_no":"4221",   //apply下单返回得 第三方订单号
    "new_phone":"13455555",    //选得新手机号
    "token"："tokenssss"
}
返回
{
    "code":0,
    "message":"success",
    "data":"https://www.baidu.com"
}
```

#### 为我上门激活
```
POST /v1/ft/market/ydhk/offline-active
{
    "phone":"1234455",
    "new_phone":"122222",
    "true_name":"nihao",
    "idcard":"330",     //身份证

    "province_code":"100",
    "city_code":"200",

    "express_province":"浙江省",
    "express_province_code":"200",
     "express_city":"嘉兴市",
    "express_city_code":"100",
     "express_district":"秀洲区"
    "express_district_code":"10000"
    "express_address":"2nihao",

    "token":"sddd",

     "class_big_tp":1,
    "class_tp":2,
    "class_name":"dwsdd",
    "class_isp":1,
    "ip":"1233333",
    "class_isp,":1
}
返回
{
    "code":0,
    "message":"success",
    "data":{
        "third_order_no":"420200202", //第三方订单号
        "oao_model": true             //快递能到否
    }
}
```