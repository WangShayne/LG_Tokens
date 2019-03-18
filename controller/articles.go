package controller

import (
	"github.com/LG_Tokens/common"
	"github.com/LG_Tokens/db"
	"github.com/LG_Tokens/util"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetArticles(requestBody *common.RequestBody, c *gin.Context) {
	p := requestBody.Params.([]interface{})
	plen := len(p)
	if plen == 0 {
		c.JSON(http.StatusOK, util.CreateResponseBody(requestBody, 10003, nil))
		return
	}

	article := db.QueryArticles(p[0].(string))

	c.JSON(http.StatusOK, util.CreateResponseBody(requestBody, 0, article))

}
