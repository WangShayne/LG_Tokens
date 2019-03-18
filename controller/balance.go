package controller

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	. "github.com/LG_Tokens/Logger"
	common "github.com/LG_Tokens/common"
	"github.com/LG_Tokens/util"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

func GetBalance(requestBody *common.RequestBody, c *gin.Context) {
	var addresses []interface{}
	// var pub_key string

	p := requestBody.Params.([]interface{})
	plen := len(p)
	if plen == 0 || plen == 1 {
		c.JSON(http.StatusOK, util.CreateResponseBody(requestBody, 10003, nil))
		return
	}

	addresses = p[0].([]interface{})
	addlen := len(addresses)
	client := &http.Client{}
	WEB3url := viper.GetString("ETHserver.server") + ":" + viper.GetString("ETHserver.WEB3port") + "/getBalance"
	// WEB3url := "http://127.0.0.1:9998/getBalance"
	w3Req := requestBody.Params
	// w3ReqNew := bytes.NewBuffer(w3Req)
	w3ReqJSON, err := json.Marshal(w3Req)
	if err != nil {
		fmt.Println(err.Error())
		c.JSON(http.StatusOK, util.CreateResponseBody(requestBody, 10001, nil))
		return
	}
	w3ReqNew := bytes.NewBuffer([]byte(w3ReqJSON))
	fmt.Println(WEB3url)
	w3request, _ := http.NewRequest("POST", WEB3url, w3ReqNew)
	w3request.Header.Set("Content-type", "application/json")
	w3response, _ := client.Do(w3request)

	balanceList := make([]*common.Balance, addlen)
	if w3response.StatusCode == 200 {
		w3body, _ := ioutil.ReadAll(w3response.Body)
		err := json.Unmarshal([]byte(w3body), &balanceList)
		if err != nil {
			Log.Info(err)
			c.JSON(http.StatusOK, util.CreateResponseBody(requestBody, 10001, nil))
			return
		}
	}

	c.JSON(http.StatusOK, util.CreateResponseBody(requestBody, util.NORMAL_CODE, balanceList))
}
