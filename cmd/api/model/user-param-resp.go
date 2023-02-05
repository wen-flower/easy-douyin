package model

import "github.com/wen-flower/easy-douyin/kitex_gen/common"

// LoginParam 登录请求参数
type LoginParam struct {
	// 登录用户名
	Username string `json:"username" query:"username"`
	// 登录密码
	Password string `json:"password" query:"password"`
}

// LoginResp 登录响应数据
type LoginResp struct {
	BaseResp
	UserId *int64  `json:"user_id,string"`
	Token  *string `json:"token"`
}

// RegisterParam 注册请求参数
type RegisterParam struct {
	// 用户名
	Username string `json:"username" query:"username"`
	// 密码
	Password string `json:"password" query:"password"`
}

// RegisterResp 注册响应数据
type RegisterResp struct {
	BaseResp
	UserId *int64  `json:"user_id,string"`
	Token  *string `json:"token"`
}

// UserInfoParam 获取用户信息参数
type UserInfoParam struct {
	UserId int64 `json:"user_id" query:"user_id" vd:"@:$ > 0; msg:'用户ID参数错误'"`
}

// UserInfoResp 获取用户信息响应数据
type UserInfoResp struct {
	BaseResp
	User *common.UserInfo `json:"user"`
}
