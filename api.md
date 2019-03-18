
#### 错误代码

- 10001
系统错误

- 10002
method方法不存在或为空

- 10003
提交参数不正确


-----

#### 接口列表

##### 常用代币及模糊搜索 `LG_erc20`

###### Request
```
{
    "jsonrpc": "2.0",
    "method": "LG_erc20",   
    "params": [
    	"LRC"
    ],
    "id": 44
}
```

###### Response

```
{
    "id": 44,
    "jsonrpc": "2.0",
    "result": [
        {
            "code": "loopring",
            "name": "LRC",
            "enName": "Loopring",
            "cnName": "路印协议",
            "logo": "/img/0xef68e7c694f40c8202821edf525de3782458639f.png",
            "address": "0xef68e7c694f40c8202821edf525de3782458639f",
            "isCommon": 1
        }
    ]
}

```
-----
##### 代币余额查询 `LG_getBalance`

###### Request
```
{
    "jsonrpc": "2.0",
    "method": "LG_getBalance",   //接口
    "params": [
        [
            "0xF317f4acfC0D70ccc79A2f24cFBbD7ebc02CFa2E"   // 代币地址
        ],
        "0xeddc1B054649B47b5a2Ab1F1bDeAdBBa35B3C2B5"     //公钥
    ],
    "id": 1
}
```

###### Response

```
{
    "id": 1,
    "jsonrpc": "2.0",
    "result": [
        {
            "address": "0xF317f4acfC0D70ccc79A2f24cFBbD7ebc02CFa2E",    //代币地址
            "balanceOf": "200"      //代币价格
        }
    ]
}

```
-----
##### 代币当前查询 `LG_getTokensPrice`

###### Request
```
{
    "jsonrpc": "2.0",
    "method": "LG_getTokensPrice",          //接口名
    "params": [
    	"0xcb97e65f07da24d46bcdd078ebebd7c6e6e3d750"   // 代币地址
    ],
    "id": 1
}
```

###### Response

```
{
    "id": 1,
    "jsonrpc": "2.0",
    "result": [
        {
            "address": "0xcb97e65f07da24d46bcdd078ebebd7c6e6e3d750",   // 代币地址
            "symbol": "BTM",                    // 代币名
            "price_usd": "0.0857490000",        // 美元价格
            "price_btc": "0.0000224889",        // btc价格
            "price_cny": "0.5767000000"         // 人民币价格
        }
    ]
}

```
-----
##### 域名注册接口 `LG_ensRegister`

###### Request
```
{
    "jsonrpc": "2.0",
    "method": "LG_ensRegister",
    "params": [
    	"www.bai133.com",   // 域名
    	"0xcb97e65f07da24d46bcdd078ebebd7c6e6e3d750" //公钥
    ],
    "id": 1
}
```

###### Response

```
{
    "id": 1,
    "jsonrpc": "2.0",
    "result": {
        "id": 84,
        "domainName": "www.bai133.com",
        "PubKey": "0xcb97e65f07da24d46bcdd078ebebd7c6e6e3d750"
    }
}

```
-----
##### 域名查询公钥接口 `LG_ensSearch`

###### Request
```
{
    "jsonrpc": "2.0",
    "method": "LG_ensSearch",          
    "params": [
    	"www.bai1dd.com"   // 域名
    ],
    "id": 1
}
```

###### Response

```
{
    "id": 1,
    "jsonrpc": "2.0",
    "result": {
        "id": 82,
        "domainName": "www.bai1dd.com",
        "PubKey": "0xcb97e65f07da24d46bcdd078ebebd7c6e6e3d750"
    }
}

```
-----
##### 文章代码查询接口 `LG_article`

###### Request
```
{
    "id" : 1,
    "jsonrpc" : "2.0",
    "method" : "LG_article",
    "params" : ["b"]    // a 用户协议 | b 帮助用心 | c 关于我们
}
```

###### Response

```
{
    "id": 1,
    "jsonrpc": "2.0",
    "result": {
        "name": "帮助中心",
        "codes": "一段帮助代码"
    }
}

```