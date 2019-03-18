package routers

import (
	"fmt"
	"io/ioutil"
	"net/http"

	. "github.com/LG_Tokens/Logger"
	"github.com/LG_Tokens/common"
	ctrl "github.com/LG_Tokens/controller"
	"github.com/LG_Tokens/util"
	"github.com/elgs/gojq"
	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	router := gin.Default()
	router.Static("/img/", "./static/images/")
	router.POST("/", Apis)
	return router
}

func Apis(c *gin.Context) {
	var requestBody common.RequestBody
	reqData, _ := ioutil.ReadAll(c.Request.Body)
	data := string(reqData)
	req, err := gojq.NewStringQuery(data)
	if err != nil {
		Log.Info(err)
		return
	}
	id, err := req.Query("id")
	jsonrpc, err := req.Query("jsonrpc")
	method, err := req.Query("method")
	params, err := req.Query("params")
	requestBody = common.RequestBody{
		id.(float64),
		jsonrpc.(string),
		method.(string),
		params,
	}

	fmt.Println(requestBody.Method)
	switch requestBody.Method {
	// 常用代币及模糊搜索
	case "LG_erc20":

		ctrl.SearchTokens(&requestBody, c)
	// ETH汇率
	case "LG_rate":

		ctrl.GetRate(&requestBody, c)
	// 获取指定块中特定的账户地址余额
	case "LG_getBalance":

		ctrl.GetBalance(&requestBody, c)

	// 获取指定块中特定的账户地址余额
	case "LG_getTokensPrice":

		ctrl.GetTokensPrice(&requestBody, c)

	//  注册域名
	case "LG_ensRegister":

		ctrl.RegisterEns(&requestBody, c)

	//  域名查询公钥
	case "LG_ensSearch":

		ctrl.SearchEns(&requestBody, c)

	default:
		c.JSON(http.StatusOK, util.CreateResponseBody(&requestBody, 10003, nil))
	}

}
