package response

// Body 统一 API 返回结构
type Body struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data,omitempty"`
}

func OK(data interface{}) *Body {
	return &Body{Code: 0, Msg: "ok", Data: data}
}

func Fail(code int, msg string) *Body {
	return &Body{Code: code, Msg: msg}
}
