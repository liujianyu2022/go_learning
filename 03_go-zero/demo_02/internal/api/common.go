package api

type Error struct {
	Code string `json:"code"`
	Msg  string `json:"msg"`
}

func NewError(code string, msg string) *Error {
	return &Error{
		Code: code,
		Msg:  msg,
	}
}

var (
	DBError         = NewError("Error001", "数据库错误")
	AlreadyRegister = NewError("Error002", "用户已注册")
)

type Result struct {
	Code string `json:"code"`
	Msg  string `json:"msg"`
	Data any    `json:"data,omitempty"`
}

func Success(data any) *Result {
	return &Result{
		Code: "200",
		Msg:  "success",
		Data: data,
	}
}

func Fail(err *Error) *Result {
	return &Result{
		Code: err.Code,
		Msg:  err.Msg,
	}
}
