package common

type Ens struct {
	ID         interface{} `json:"id"`
	DomainName string      `json:"domainName"`
	PubKey     string      `json:"pubKey"`
}

type EnsRespByKey struct {
	DomainName []string `json:"domainName"`
	PubKey     string   `json:"pubKey"`
}
