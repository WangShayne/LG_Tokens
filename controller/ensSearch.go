package controller

import (
	"github.com/LG_Tokens/common"
	"github.com/LG_Tokens/db"
	"github.com/LG_Tokens/util"
	"github.com/gin-gonic/gin"
)

// 域名查询
// method: lg_ensSearch
// params: ["域名文字"]

// result: ""
func SearchEns(requestBody *common.RequestBody, c *gin.Context) {

	p := requestBody.Params.([]interface{})

	domainName := p[0].(string)

	ensResp, errCode := db.QueryENS(domainName)

	c.JSON(200, util.CreateResponseBody(requestBody, errCode, ensResp))

}

func SearchEnsByKey(requestBody *common.RequestBody, c *gin.Context) {

	p := requestBody.Params.([]interface{})

	key := p[0].(string)

	ensByKey, errCode := db.QueryENSByKey(key)

	c.JSON(200, util.CreateResponseBody(requestBody, errCode, ensByKey))

}
