package handler

import (
	"encoding/hex"
	"github.com/wen-flower/easy-douyin/cmd/user/consts"
	"github.com/wen-flower/easy-douyin/kitex_gen/common"
	user "github.com/wen-flower/easy-douyin/kitex_gen/user"
	"github.com/wen-flower/easy-douyin/pkg/errno"
	"golang.org/x/crypto/sha3"
)

// UserServiceImpl implements the last service interface defined in the IDL.
type UserServiceImpl struct{}

// 检查 UserServiceImpl 是否实现了 user.UserService 接口
var _ user.UserService = (*UserServiceImpl)(nil)

// 提取出错误的处理流程
func errProcess(baseResp **common.BaseResp, err *error) {
	var resp common.BaseResp
	if *err != nil {
		e := errno.ConvertErr(*err)
		resp.Msg = e.Msg()
		resp.Code = e.Code()
		*err = nil
	} else {
		resp.Code = errno.Success.Code()
		resp.Msg = errno.Success.Msg()
	}
	*baseResp = &resp
}

// 对密码进行摘要
func passwordDigest(password string) string {
	digest := sha3.Sum512([]byte(password + consts.PasswordSecret))
	return hex.EncodeToString(digest[:])
}
