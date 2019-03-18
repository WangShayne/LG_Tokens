package common

type Ens struct {
	ID         interface{} `json:"id"`
	DomainName string      `json:"domainName"`
	PubKey     string      `json:"pubKey"`
}

type EnsResp struct {
	IsOK bool   `json:"isOK"`
	Msg  string `json:"msg"`
	*Ens
}

func (ensResp *EnsResp) CreateEnsResp(isok bool, msg string, ens *Ens) *EnsResp {

	ensResp = &EnsResp{
		isok,
		msg,
		ens,
	}
	return ensResp
}
