package errcode

// 统一错误码：0=成功，非0=失败
const (
	CodeOK             = 0
	CodeInvalidParam   = 400
	CodeUnauthorized   = 401
	CodeNotFound       = 404
	CodeInternalError  = 500
	CodeDBError        = 501
	CodeDuplicateEntry = 502
)

var codeMsg = map[int]string{
	CodeOK:             "ok",
	CodeInvalidParam:   "invalid param",
	CodeUnauthorized:   "unauthorized",
	CodeNotFound:       "not found",
	CodeInternalError:  "internal error",
	CodeDBError:        "db error",
	CodeDuplicateEntry: "duplicate entry",
}

func Msg(c int) string {
	if m, ok := codeMsg[c]; ok {
		return m
	}
	return "unknown"
}
