package handler

import (
	"context"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/wen-flower/easy-douyin/cmd/api/consts"
	"github.com/wen-flower/easy-douyin/cmd/api/model"
	"github.com/wen-flower/easy-douyin/cmd/api/mw"
	"github.com/wen-flower/easy-douyin/cmd/api/rpc"
	"github.com/wen-flower/easy-douyin/kitex_gen/user"
)

// Login 处理登录接口
// @router /douyin/user/login [POST]
func Login(ctx context.Context, req *app.RequestContext) {
	var err error
	defer errProcess(req, &err)

	var param model.LoginParam
	if err = req.BindAndValidate(&param); err != nil {
		return
	}
	uid, err := rpc.CheckUser(context.Background(), &user.CheckUserParam{
		Username: param.Username,
		Password: param.Password,
	})

	if err != nil {
		return
	}
	req.Set(consts.JwtIdentityKey, uid)
	// 调用 Hertz JWT 中间件的登录方法
	mw.JwtMiddleware.LoginHandler(ctx, req)
}

// Register 处理注册接口
// @router /douyin/user/register [POST]
func Register(ctx context.Context, req *app.RequestContext) {
	var err error
	defer errProcess(req, &err)

	var param model.RegisterParam
	if err = req.BindAndValidate(&req); err != nil {
		return
	}

	var uid *int64
	uid, err = rpc.CreateUser(ctx, &user.CreateUserParam{
		Username: param.Username,
		Password: param.Password,
	})
	if err != nil {
		return
	}
	req.Set(consts.JwtIdentityKey, uid)
	mw.JwtMiddleware.LoginHandler(ctx, req)
}
