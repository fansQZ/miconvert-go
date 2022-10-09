// Package constants
// @Author: fzw
// @Create: 2022/10/8
// @Description: package about constants
package result

// ResultCode
// @Description: response code and message return to the frontend
//
type ResultCode struct {
	code    int    //response code
	message string //response message
}

func (r ResultCode) GetCode() int {
	return r.code
}

func (r ResultCode) GetMessage() string {
	return r.message
}

var (
	CUSTOM_SIMPLE_ERROR_MESSAGE = &ResultCode{code: 1, message: "自定义错误"}
	ERR                         = &ResultCode{code: 500, message: "server error"}
	OK                          = &ResultCode{code: 200, message: "response succeeded"}
)
