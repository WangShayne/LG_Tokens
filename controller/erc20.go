package controller

import (
	"net/http"

	common "github.com/LG_Tokens/common"
	db "github.com/LG_Tokens/db"

	util "github.com/LG_Tokens/util"
	"github.com/gin-gonic/gin"
)

// 常用的代币列表和模糊搜索的接口
// method: lg_erc20
// params: []

// result: []

func SearchTokens(requestBody *common.RequestBody, c *gin.Context) {
	var keyword string

	p := requestBody.Params.([]interface{})
	plen := len(p)
	// var tokenSlice *[]common.Tokens
	var tokenSlice []common.Tokens
	var errCode int
	if plen == 0 {
		tokenSlice, errCode = db.QueryCommonToken()
	} else {
		for i := 0; i < plen; i++ {
			newA, ok := p[i].(string)
			if ok {
				keyword = newA
			}
		}
		tokenSlice, errCode = db.QueryKeywordToken(keyword)
	}

	c.JSON(http.StatusOK, util.CreateResponseBody(requestBody, errCode, tokenSlice))
}
