package utils

import (
	"github.com/cloudwego/hertz/pkg/app/server/binding"
	"github.com/wen-flower/easy-douyin/pkg/errno"
)

func InitHertzValidator() {
	CustomBindErrFunc := func(failField, msg string) error {
		return errno.ParamBindingErr
	}

	CustomValidateErrFunc := func(failField, msg string) error {
		return errno.ParamErr.WithMessage(msg)
	}

	binding.SetErrorFactory(CustomBindErrFunc, CustomValidateErrFunc)
}
