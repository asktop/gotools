<h1>API签名验证规则</h1>

[TOC]

# 1. 申请获取 `APP-KEY` 和 `APP-SECRET`

# 2. API签名验证Header参数
**API签名验证所需 Request Header 参数如下：**
| 参数 | 类型 | 必须 | 描述 | 取值 |
| --- | --- | --- | --- | --- |
| APP-KEY | string | true | 申请的APP-KEY | 例：3e5832293dc9a119aeee163a024b79f1 |
| APP-SIGNATURE | string | true | API签名 | 例：jO9vANFp4ZqrjdVxKoumGt1z/aM= |
| APP-TIMESTAMP | string | true | 签名时间戳 | 单位：毫秒，例：1533805471865 |

# 3. API签名生成规则
## 3.1 总规则
签名前所需要准备的数据为 `HTTP_METHOD` + `HTTP_REQUEST_URL` + `TIMESTAMP` + `POST_BODY`，数据按此规则拼接成签名信息。然后，进行第一次 `Base64` 编码，然后对第一次 `Base64` 编码的结果使用秘钥 `APP-SECRET` 进行 `HMAC-SHA1` 签名，最后对签名二进制结果进行第二次 `Base64` 编码，得到最终签名字符串 `APP-SIGNATURE`。

## 3.2 各部分数据详细规则
1.`HTTP_METHOD`  
`GET`, `POST`, `DELETE`, `PUT` 需要大写。  

2.`HTTP_REQUEST_URL`  
拼接签名信息时需要对 URI 中的参数，按照==按照字母表升序排序==！
```
示例：
如果请求的 URL 为：
https://api.m.cc/v2/orders?c=value1&b=value2&a=value3

则签名时排序后的 URL 为：
https://api.m.cc/v2/orders?a=value3&b=value2&c=value1
```

3.`TIMESTAMP`  
访问 API 时的 UNIX EPOCH 时间戳（毫秒，11位），需要和服务器之间的时间差少于 30 秒。

4.`POST_BODY`  
如果是 POST 请求，POST 请求数据也需要被签名，签名规则如下：  
所有请求的 key ==按照字母表升序排序==，然后进行 url 参数化，并使用 & 连接。  
并且，请求 `Content-Type` 必须是 `application/json` 。
```
示例：
pa=value3&pb=value2&pc=value1
```

# 4. API签名生成示例
## 4.1 请求数据
1.Request Header 参数  
| 参数 | 取值 |
| --- | --- |
| APP-KEY | 3e5832293dc9a119aeee163a024b79f1 |
| APP-TIMESTAMP | 1533805471865 |

2.请求 `POST` `https://api.m.cc/v2/orders`  

3.POST请求参数
```
{
  "type": "limit",
  "side": "buy",
  "amount": "100.0",
  "price": "100.0",
  "symbol": "btcusdt"
}
```

## 4.2 签名生成
1.拼接签名信息如下：  
```
POSThttps://api.m.cc/v2/orders1533805471865amount=100.0&price=100.0&side=buy&symbol=btcusdt&type=limit
```

2.进行第一次 Base64 编码，得到：
```
UE9TVGh0dHBzOi8vYXBpLm0uY2MvdjIvb3JkZXJzMTUzMzgwNTQ3MTg2NWFtb3VudD0xMDAuMCZwcmljZT0xMDAuMCZzaWRlPWJ1eSZzeW1ib2w9YnRjdXNkdCZ0eXBlPWxpbWl0
```

3.拷贝在申请 `APP-KEY` 时获得的秘钥 `APP-SECRET` （以 `a13444ca8eef5637358915eeb16f30d35ead9b36` 为例）， 对第一次 `Base64` 编码的结果使用秘钥 `APP-SECRET` 进行 `HMAC-SHA1` 签名，并对签名二进制结果进行第二次 `Base64` 编码，得到最终签名 `APP-SIGNATURE` ：
```
jO9vANFp4ZqrjdVxKoumGt1z/aM=
```
