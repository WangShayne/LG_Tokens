package controller

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/elgs/gojq"
	"github.com/spf13/viper"

	common "github.com/LG_Tokens/common"
	db "github.com/LG_Tokens/db"

	util "github.com/LG_Tokens/util"
	"github.com/gin-gonic/gin"
	"github.com/mikemintang/go-curl"
)

// address: 0xcb97e65f07da24d46bcdd078ebebd7c6e6e3d750

// 换算法币汇率
// method: lg_rate
// {"jsonrpc":"2.0","method":"LG_rate","params":[""],"id":44}

//  {
//     "id": 44,
//     "jsonrpc": "2.0",
//     "result": {
//         "symbol": "ETH",
//         "price_usd": "146.654290138",
//         "price_btc": "0.03741033",
//         "price_cny": "990.3270904439"
//     }
// }

func GetRate(requestBody *common.RequestBody, c *gin.Context) {
	var ETHPriceAPI = viper.GetString("ETHPriceAPI")

	req := curl.NewRequest()
	res, err := req.
		SetUrl(ETHPriceAPI).
		Get()

	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusOK, util.CreateResponseBody(requestBody, 10001, nil))
	} else {
		if res.IsOk() {
			var ethRes common.ETHPrice
			var LGRate common.ETHRate
			var usdToCny common.USDToCNYRate

			parser, err := gojq.NewStringQuery(res.Body)
			if err != nil {
				fmt.Println(err)
				c.JSON(http.StatusOK, util.CreateResponseBody(requestBody, 10001, nil))
			}
			symbol, err := parser.Query("[0].symbol")
			priceUsd, err := parser.Query("[0].price_usd")
			priceBtc, err := parser.Query("[0].price_btc")
			ethRes = common.ETHPrice{
				symbol.(string),
				priceUsd.(string),
				priceBtc.(string),
			}

			usdToCny = db.QueryRate()

			fBuyPri, _ := strconv.ParseFloat(usdToCny.FBuyPri, 64)
			priceUsdFloat, _ := strconv.ParseFloat(ethRes.PriceUsd, 64)
			priceCNY := (fBuyPri / 100) * priceUsdFloat
			priceCNYStr := strconv.FormatFloat(priceCNY, 'f', 10, 64)

			LGRate = common.ETHRate{
				ethRes,
				priceCNYStr,
			}
			c.JSON(http.StatusOK, util.CreateResponseBody(requestBody, 00000, &LGRate))

		} else {
			fmt.Println(res.Raw)
		}
	}
}

func GetRateByOutApi() {

	URL := viper.GetString("RateURL")
	req := curl.NewRequest()
	res, err := req.
		SetUrl(URL).
		Get()

	if err != nil {
		fmt.Println(err)
	} else {
		if res.IsOk() {
			var usdToCny *common.USDToCNYRate

			parser, err := gojq.NewStringQuery(res.Body)
			if err != nil {
				fmt.Println(err)
			}
			name, err := parser.Query("result.[0].data1.name")
			fBuyPri, err := parser.Query("result.[0].data1.fBuyPri")
			date, err := parser.Query("result.[0].data1.date")
			time, err := parser.Query("result.[0].data1.time")
			dateTime := date.(string) + " " + time.(string)
			usdToCny = &common.USDToCNYRate{
				name.(string),
				fBuyPri.(string),
				dateTime,
			}

			db.InsertRate(usdToCny)
		} else {
			fmt.Println(res.Raw)
		}
	}
}
