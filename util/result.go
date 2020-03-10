package util

type Result struct {
	Code int `json:code form:code`
	Msg string `json:msg form:msg`
	Data interface{} `json:data form:data`
}

const(
	Success int = 200
	Fail int = 101 // 失败
)

func NewResult(code int, msg string, data ...interface{}) Result {
	if len(data) > 0 {
		res:= Result{
			Code:code,
			Msg:msg,
			Data:data[0],
		}
		return res
	}else {
		res:= Result{
			Code:code,
			Msg:msg,
		}
		return res
	}
}
