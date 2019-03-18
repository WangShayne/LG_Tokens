package controller

import (
	"fmt"
	"github.com/LG_Tokens/common"
	"github.com/LG_Tokens/db"
	"github.com/LG_Tokens/util"
	"github.com/elgs/gojq"
	"github.com/gin-gonic/gin"
	"github.com/mikemintang/go-curl"
	"github.com/spf13/viper"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"
	"time"
)

var ethAddress = "0x0000000000000000000000000000000000000000"

var ch chan *common.TokenPrice

func GetTokensPrice(requestBody *common.RequestBody, c *gin.Context) {
	var priceSlice []*common.TokenPrice

	p := requestBody.Params.([]interface{})
	plen := len(p)

	if len(p) == 0 {
		ch = make(chan *common.TokenPrice, 1)
		go GetTokenPriceByAddress(ethAddress)
	} else {
		ch = make(chan *common.TokenPrice, plen)
		for i := 0; i < plen; i++ {
			time.Sleep(100 * time.Millisecond)
			go GetTokenPriceByAddress(p[i].(string))
		}
	}

	for i := 0; i < plen; i++ {
		v := <-ch
		priceSlice = append(priceSlice, v)
	}
	c.JSON(http.StatusOK, util.CreateResponseBody(requestBody, 00000, priceSlice))
}

func GetTokenPriceByAddress(address string) {

	var tkp *common.TokenPrice

	// eth币价格
	if address == ethAddress {

		var ETHPriceAPI = viper.GetString("ETHPriceAPI")

		req := curl.NewRequest()
		res, err := req.
			SetUrl(ETHPriceAPI).
			Get()
		if err != nil {
			fmt.Println(err)
			tkp = crateNilTokenPrice(address)
			ch <- tkp
			return
		} else {
			if res.IsOk() {
				parser, err := gojq.NewStringQuery(res.Body)
				if err != nil {
					fmt.Println(err)
					tkp = crateNilTokenPrice(address)
					ch <- tkp
					return
				}

				symbol, err := parser.Query("[0].symbol")
				priceUsd, err := parser.Query("[0].price_usd")
				priceBtc, err := parser.Query("[0].price_btc")

				var usdToCny common.USDToCNYRate
				usdToCny = db.QueryRate()

				fBuyPri, _ := strconv.ParseFloat(usdToCny.FBuyPri, 64)
				priceUsdFloat, _ := strconv.ParseFloat(priceUsd.(string), 64)
				priceCNY := (fBuyPri / 100) * priceUsdFloat
				priceCNYStr := strconv.FormatFloat(priceCNY, 'f', 10, 64)

				tkp = &common.TokenPrice{
					ethAddress,
					symbol.(string),
					priceUsd.(string),
					priceBtc.(string),
					priceCNYStr,
				}

			} else {
				fmt.Println(res.Raw)
			}
		}

	} else {
		// 其他代币价格

		fmt.Println(address)
		var TokenPriceURL = viper.GetString("TokenPriceURL")
		resp, err := http.PostForm(TokenPriceURL, url.Values{"address": {address}})
		if err != nil {
			fmt.Println(err)
			fmt.Println("resp 出错")
			tkp = crateNilTokenPrice(address)
			ch <- tkp
			return
		}

		defer resp.Body.Close()
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			fmt.Println(err)
			fmt.Println("body 出错")
			tkp = crateNilTokenPrice(address)
			ch <- tkp
			return
		} else {
			parser, err := gojq.NewStringQuery(string(body))
			if err != nil {
				fmt.Println(err)
				tkp = crateNilTokenPrice(address)
				ch <- tkp
				return
			}

			symbol, err := parser.Query("data.symbol")
			priceUsd, err := parser.Query("data.price_usd")
			priceBtc, err := parser.Query("data.price_btc")
			priceCny, err := parser.Query("data.price_cny")

			priceUsdStr := strconv.FormatFloat(priceUsd.(float64), 'f', 10, 64)
			priceBtcStr := strconv.FormatFloat(priceBtc.(float64), 'f', 10, 64)
			priceCnyStr := strconv.FormatFloat(priceCny.(float64), 'f', 10, 64)

			tkp = &common.TokenPrice{
				address,
				symbol.(string),
				priceUsdStr,
				priceBtcStr,
				priceCnyStr,
			}

		}
	}
	ch <- tkp
	return
}

func crateNilTokenPrice(address string) *common.TokenPrice {
	return &common.TokenPrice{
		address,
		"",
		"",
		"",
		"",
	}
}
