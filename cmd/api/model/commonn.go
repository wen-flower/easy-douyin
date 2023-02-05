package model

import "github.com/wen-flower/easy-douyin/pkg/errno"

type (
	BaseResp struct {
		StatusCode int16  `json:"status_code"`
		StatusMsg  string `json:"statue_msg"`
	}
)

func (resp *BaseResp) Ok() {
	resp.StatusCode = errno.Success.Code()
	resp.StatusMsg = errno.Success.Msg()
}
