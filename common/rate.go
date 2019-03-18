package common

// type ETHPriceResponse struct {
// 	_ []ETHPrice
// }

// eth价格
type ETHPrice struct {
	Symbol   string `json:"symbol"`
	PriceUsd string `json:"price_usd"`
	PriceBtc string `json:"price_btc"`
}

// 人民币美元汇率
type USDToCNYRate struct {
	Name       string
	FBuyPri    string // 买入价格
	UpdateTime string // 更新时间
}

// eth返回汇率
type ETHRate struct {
	ETHPrice
	PriceCNY string `json:"price_cny"`
}



