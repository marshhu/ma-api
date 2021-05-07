package api

// ApiResult 接口返回体
type ApiResult struct {
	Code    int         `json:"code"`
	Message string      `json:"message,omitempty"`
	Data    interface{} `json:"data,omitempty"`
	Stack   string      `json:"stack,omitempty"`
}

func NewApiResult(code int, msg string, data interface{}) *ApiResult {
	return &ApiResult{
		Code:    code,
		Message: msg,
		Data:    data,
	}
}

const (
	SuccessCode  int = 0
	FailedCode int = -1
)

func Success(data interface{}) *ApiResult {
	return NewApiResult(SuccessCode, "", data)
}

func Failed(msg string) *ApiResult {
	return NewApiResult(FailedCode, msg, nil)
}

func FailedWithCode(code int,msg string) *ApiResult {
	return NewApiResult(code, msg, nil)
}