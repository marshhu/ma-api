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

func OK(data interface{}) *ApiResult {
	return NewApiResult(0, "", data)
}

func Error(code int, msg string) *ApiResult {
	return NewApiResult(code, msg, nil)
}
