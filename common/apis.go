package common

// apis 列表

//  请求结构体

type RequestBody struct {
	ID      float64
	Jsonrpc string
	Method  string
	Params  interface{}
}

type ErrorResponse struct {
	Code int64  `json:"code"`
	Msg  string `json:"msg"`
}
