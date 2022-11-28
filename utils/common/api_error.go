package common

// TODO 同一类code放一起，不同类型的code间隔设置得大一点，这样后面好分辨
const (
	SUCCESS             = 200
	ErrorFailed         = 400
	ErrorParamIllegal   = 6001
	ErrorTokenValidFail = 6002
	ErrorNoPermission   = 6003
	ErrorServerUnknown  = 6004
	ErrorMysqlReadFail  = 6010
	ErrorMysqlWriteFail = 6011
	ErrorShouldBind     = 6013
	ErrorConnectSMTP    = 6014
	ERRORSCNOTACTIVE    = 6020
	ERRORSCPOSTTIMEEXP  = 6022
	ErrorCodeStr        = "400"
)

var ErrorMap = map[int]string{
	SUCCESS:             "请求成功",
	ErrorFailed:         "请求失败",
	ErrorParamIllegal:   "请求参数非法",
	ErrorTokenValidFail: "身份校验失败",
	ErrorNoPermission:   "权限不足",
	ErrorMysqlReadFail:  "数据库读取失败",
	ErrorMysqlWriteFail: "数据库写入失败",
	ErrorServerUnknown:  "服务内部异常未知错误",
	ErrorShouldBind:     "ctx解析请求参数错误",
	ErrorConnectSMTP:    "连接smtp服务器失败",
}
