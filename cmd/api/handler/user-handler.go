package handler

import (
	"context"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/wen-flower/easy-douyin/cmd/api/consts"
	"github.com/wen-flower/easy-douyin/cmd/api/model"
	"github.com/wen-flower/easy-douyin/cmd/api/mw"
	"github.com/wen-flower/easy-douyin/cmd/api/utils"
	"github.com/wen-flower/easy-douyin/kitex_gen/user"
	"github.com/wen-flower/easy-douyin/pkg/errno"
	"github.com/wen-flower/easy-douyin/pkg/rpc/userrpc"
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
	uid, err := userrpc.CheckUser(context.Background(), &user.CheckUserParam{
		Username: param.Username,
		Password: param.Password,
	})

	if err != nil {
		return
	}

	utils.SetLoggedInUID(req, *uid)
	// 调用 Hertz JWT 中间件的登录方法
	mw.JwtMiddleware.LoginHandler(ctx, req)
}

// Register 处理注册接口
// @router /douyin/user/register [POST]
func Register(ctx context.Context, req *app.RequestContext) {
	var err error
	defer errProcess(req, &err)

	var param model.RegisterParam
	if err = req.BindAndValidate(&param); err != nil {
		return
	}

	uid, err := userrpc.CreateUser(ctx, &user.CreateUserParam{
		Username: param.Username,
		Password: param.Password,
	})
	if err != nil {
		return
	}
	req.Set(consts.JwtIdentityKey, *uid)
	mw.JwtMiddleware.LoginHandler(ctx, req)
}

// UserInfo 处理获取用户信息接口
// @router /douyin/user [GET]
func UserInfo(ctx context.Context, req *app.RequestContext) {
	var err error
	defer errProcess(req, &err)

	var param model.UserInfoParam
	if err = req.BindAndValidate(&param); err != nil {
		return
	}

	userInfos, err := userrpc.QueryUser(ctx, &user.QueryUserParam{
		LoggedUserId: utils.GetLoggedInUID(req),
		UserIds:      []int64{param.UserId},
	})
	if err != nil {
		return
	}

	if len(userInfos) != 1 {
		err = errno.UserNotExistsErr
		return
	}

	resp := new(model.UserInfoResp)
	resp.Ok()

	resp.User = userInfos[0]

	utils.SendJson(req, resp)
}
