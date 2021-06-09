#1。创建礼品码

```
接口地址 
/giftCode/adminCreatGiftcode 
```
## 请求方式
GET
## 请求示例
```
http://127.0.0.1:8080/giftCode/adminCreatGiftcode?des=description&GN=3&VP=5&GC=10k&CP=llq
```
## 参数  说明

``` 
des 类型string Description描述礼包的字段
```
``` 
GN 类型string GiftNum礼包的可领取数
```
``` 
VP 类型string ValidPeriod礼包的有效时间
```
``` 
GC 类型string GiftContent 礼包的内容
```
``` 
CP 类型string GiftContent CreatePer礼包的创建者
```

```
成功示例 
{
    "GiftCode": "r1czr5u2",
    "condition": "success"
}

错误示列 
{
    "GiftCode": {},
    "condition": "error"
}
```

#2。查询礼品码

```
接口地址 
/giftCode/admininquireGiftCode 
```
## 请求方式
GET
## 请求示例
```
http://127.0.0.1:8080/giftCode/admininquireGiftCode?giftCode=r1czr5u2
```

## 参数  说明

``` 
giftCode 类型string 此字段为需要查询的礼包码
```


```
成功示例 
{
    "condition": "success",
    "data": {
        "AvailableNum": "0",
        "ClaimList": "",
        "CreatTime": "2021-06-09 18:47:07",
        "CreatePer": "llq",
        "Description": "description",
        "GiftContent": "10k",
        "GiftNum": "3",
        "ValidPeriod": "5",
        "giftCode": "r1czr5u2"
    }
}
错误示列 
{
    "condition": "error",
    "giftCode": "GiftCode is error"
}
```

#3。验证礼品码

```
接口地址 
/giftCode/client 
```
## 请求方式
GET
## 请求示例
```
http://127.0.0.1:8080/giftCode/client?giftCode=r1czr5u2&usr=nna
```

## 参数  说明

``` 
giftCode 类型string 此字段为需要验证的礼包码
```
``` 
usr 类型string 此字段用来输入用户名称
```

```
成功示例 
{
    "GiftContent": "10k",
    "condition": "success"
}
错误示列 
{
    "GiftCode": "input usr",
    "condition": "error"
}
```