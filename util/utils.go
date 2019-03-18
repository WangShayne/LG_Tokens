package util

import (
	"fmt"
	. "github.com/LG_Tokens/common"
)

// func CreateCommonResp(request *RequestBody, result interface{}) (resBody map[string]interface{}) {
// 	resBody = map[string]interface{}{
// 		"id":      request.ID,
// 		"jsonrpc": request.Jsonrpc,
// 		"result":  result,
// 	}
//
// 	return resBody
// }
//
// func CreateErrorResp(request *RequestBody, code int) (resBody map[string]interface{}) {
// 	resBody = map[string]interface{}{
// 		"id":      request.ID,
// 		"jsonrpc": request.Jsonrpc,
// 		"error":   CreateErrorByCode(code),
// 	}
//
// 	return resBody
// }

func CreateResponseBody(request *RequestBody, code int, result interface{}) (resp map[string]interface{}) {
	fmt.Println(request.ID)
	fmt.Println(code)

	if code == NORMAL_CODE {
		resp = map[string]interface{}{
			"id":      request.ID,
			"jsonrpc": request.Jsonrpc,
			"result":  result,
		}

	} else {
		resp = map[string]interface{}{
			"id":      request.ID,
			"jsonrpc": request.Jsonrpc,
			"error":   CreateErrorByCode(code),
		}
	}
	return resp
}
