package util

/*
	错误处理
	1 系统级错误
	2 汇率查询接口错误
	3 代币查询接口错误
	4 代币价格查询接口错误
	5 域名注册错误
	6 域名查询错误
*/
const NORMAL_CODE = 0
const (
	// normal

	// 系统级错误
	SERVER_ERR = 10001
	DB_ERR     = 10002
	METHOD_ERR = 10003
	PARAMS_ERR = 10004

	// - 接口级错误 -

	// 汇率查询错误 rate
	RATE_ERR = 20001

	// 代币查询错误 erc20
	ERC20_ERR = 20002

	// 代币价格查询错误 price
	ERC20_PRICE_ERR = 30001

	// 域名注册错误 ens register
	ENS_REGISTER_ERR       = 40001
	ENS_REGISTER_EXIST_ERR = 40002
	// 域名查询错误
	ENS_SEARCH_ERR      = 50001
	ENS_SEARCHBYKEY_ERR = 50002
)

func CreateErrorByCode(code int) (errorBody map[string]interface{}) {
	errorBody = make(map[string]interface{})

	switch code {

	case SERVER_ERR:

		errorBody["code"] = 10001
		errorBody["msg"] = "系统错误,请稍后再试"

	case DB_ERR:
		errorBody["code"] = 10002
		errorBody["msg"] = "数据库连接失败"

	case METHOD_ERR:
		errorBody["code"] = 10003
		errorBody["msg"] = "请检查接口地址"

	case PARAMS_ERR:
		errorBody["code"] = 10004
		errorBody["msg"] = "请检查提交参数"

	case ENS_REGISTER_ERR:
		errorBody["code"] = 40001
		errorBody["msg"] = "注册失败,请稍后再试"

	case ENS_REGISTER_EXIST_ERR:
		errorBody["code"] = 40002
		errorBody["msg"] = "注册失败,域名已存在"
	case ENS_SEARCH_ERR:
		errorBody["code"] = 50001
		errorBody["msg"] = "未查询到相应公钥"

	case ENS_SEARCHBYKEY_ERR:
		errorBody["code"] = 50002
		errorBody["msg"] = "未查询到相应域名"

	default:

		errorBody["code"] = 10001
		errorBody["msg"] = "系统错误,请稍后再试"

	}
	return errorBody
}
