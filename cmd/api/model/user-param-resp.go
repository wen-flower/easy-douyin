package model

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
}
