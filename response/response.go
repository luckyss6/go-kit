package response

const (
	SUCCESS_CODE = 1000 + iota
	ERROR_CODE
	DATABASE_ERROR
)

var MessageMap = map[int]string{
	SUCCESS_CODE:   "success",
	ERROR_CODE:     "error",
	DATABASE_ERROR: "database error",
}

type Resposne struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

func NewResponse(code int, msg string, data interface{}) Resposne {
	return Resposne{Code: code, Msg: msg, Data: data}
}

func (r *Resposne) Success(data interface{}) {
	*r = NewResponse(SUCCESS_CODE, MessageMap[SUCCESS_CODE], data)
}

func (r *Resposne) Error(code int, msg string) {
	*r = NewResponse(code, msg, nil)
}
