package controller

import (
	"github.com/LG_Tokens/common"
	"github.com/LG_Tokens/db"
	"github.com/LG_Tokens/util"
	"github.com/gin-gonic/gin"
	"net/http"
)

// 域名
// method: lg_ensRegist
// params: ["域名文字"]

// result: true, false

// 域名注册接口  成功返回true  失败返回fasle
func RegisterEns(requestBody *common.RequestBody, c *gin.Context) {

	p := requestBody.Params.([]interface{})
	plen := len(p)
	if plen == 0 || plen == 1 {
		c.JSON(http.StatusOK, util.CreateResponseBody(requestBody, 10003, nil))
		return
	}

	domainName := p[0].(string)
	pubKey := p[1].(string)

	ensResp, errCode := db.InsertENS(domainName, pubKey)
	c.JSON(200, util.CreateResponseBody(requestBody, errCode, ensResp))
}
