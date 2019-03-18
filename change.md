### 2019-2-14 

1. 数据字段修改

    - en --> enName
    - cn --> cnName

2. bug: address 和 logo 数据相反
3. 代码结构调整
```
    |____routers
    |   |____routers.go
    |____util
    |   |____logger.go
    |   |____viper.go
    |____change.md
    |____config.yaml
    |____controller
    |   |____ensRegist.go
    |   |____rate.go
    |   |____transaction.go
    |   |____article.go
    |   |____ensSearch.go
    |   |____erc20.go
    |____common
    |   |____tokens.go
    |   |____apis.go
    |____static
    |   |____images
    |____db
    |   |____db.go
    |____main.go
    |____sql
    |   |____tokens.sql
```

4. go mod
5. 交叉编译

    windows64 `CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build -o main.exe main.go`

    linux64 `CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o main main.go`


### 2019-2-19

1. 新增法币汇率接口 `LG_rate`
2. 新增错误消息


### 2019-2-26

1. 新增代币余额查询接口 
    -   method `LG_getBalance`
    -   params [ ["address1","address2"] , "pub_key" ]
    
### 2019-2-28

1. 删除汇率查询接口 **~~LG_rate~~**
2. 新增代币价格查询接口 `LG_getTokensPrice`
    
### 2019-3-7

1. 新增域名申请接口 `LG_ensRegister`
2. 新增域名查询公钥接口 `LG_ensSearch`
3. 封装response,增加错误消息

### 2019-3-18

1. 新增文章查询接口 `LG_article`
