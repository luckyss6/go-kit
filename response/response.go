package response

const (
	SuccessCode = 1000 + iota
	ErrorCode
	DatabaseError
)

var MessageMap = map[int]string{
	SuccessCode:   "success",
	ErrorCode:     "error",
	DatabaseError: "database error",
}

type Resposne struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
	Data any    `json:"data"`
}

func NewResponse(code int, msg string, data any) Resposne {
	return Resposne{Code: code, Msg: msg, Data: data}
}

func Success(data any) Resposne {
	return NewResponse(SuccessCode, MessageMap[SuccessCode], data)
}

func Error(code int, msg string) Resposne {
	return NewResponse(code, msg, nil)
}
