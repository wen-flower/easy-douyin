package handler

import (
	user "github.com/wen-flower/easy-douyin/kitex_gen/user"
)

// UserServiceImpl implements the last service interface defined in the IDL.
type UserServiceImpl struct{}

// 检查 UserServiceImpl 是否实现了 user.UserService 接口
var _ user.UserService = (*UserServiceImpl)(nil)
