package errno

import (
	"errors"
	"fmt"
)

// Errno 业务错误
type Errno struct {
	// 业务错误状态码
	code int16
	// 业务错误描述
	msg string
}

// Error 实现了
func (e Errno) Error() string {
	return fmt.Sprintf("err_code=%d, err_msg=%v", e.code, e.msg)
}

// New 创建一个 Errno
func New(code int16, msg string) Errno {
	return Errno{
		code: code,
		msg:  msg,
	}
}

// WithMessage Copy 一份 Errno 并修改 ErrMsg 值
func (e Errno) WithMessage(msg string) Errno {
	e.msg = msg
	return e
}

// ConvertErr 将 error 转化为 Errno
func ConvertErr(err error) Errno {
	Err := Errno{}
	if errors.As(err, &Err) {
		return Err
	}
	s := ServiceErr
	return s
}

// getter/setter 方法

func (e Errno) Code() int16 {
	return e.code
}

func (e Errno) SetCode(code int16) {
	e.code = code
}

func (e Errno) Msg() string {
	return e.msg
}

func (e Errno) SetMsg(msg string) {
	e.msg = msg
}
