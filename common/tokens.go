package common

// 代币结构体

type Tokens struct {
	Code     string `json:"code"`     // 代号
	Name     string `json:"name"`     // 名称
	EnName   string `json:"enName"`   // 英文名
	CnName   string `json:"cnName"`   // 中文名
	Logo     string `json:"logo"`     // logo地址
	Address  string `json:"address"`  // 代币地址
	IsCommon int32  `json:"isCommon"` // 是否问常用
}

type TokenPrice struct {
	Address  string `json:"address"`   // 代币地址
	Symbol   string `json:"symbol"`    // 代币称号
	PriceUsd string `json:"price_usd"` // 代币美元价格
	PriceBtc string `json:"price_btc"` // 代币比特币价格
	PriceCNY string `json:"price_cny"` // 代币人民币价格
}
