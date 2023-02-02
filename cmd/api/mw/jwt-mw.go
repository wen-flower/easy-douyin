package mw

import (
	"context"
	"github.com/wen-flower/easy-douyin/cmd/api/consts"
	"github.com/wen-flower/easy-douyin/cmd/api/model"
	"github.com/wen-flower/easy-douyin/cmd/api/utils"
	"github.com/wen-flower/easy-douyin/pkg/errno"
	"net/http"
	"strconv"
	"time"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/hertz-contrib/jwt"
)

var JwtMiddleware *jwt.HertzJWTMiddleware

func InitJWT() {
	JwtMiddleware, _ = jwt.New(&jwt.HertzJWTMiddleware{
		Key:           []byte(consts.JwtSecretKey),
		TokenLookup:   "header: Authorization, query: token", //cookie: jwt
		TokenHeadName: "Bearer",
		TimeFunc:      time.Now,
		Timeout:       time.Hour,
		MaxRefresh:    time.Hour,
		IdentityKey:   consts.JwtIdentityKey,
		// 设置从 token 提取用户信息的函数
		IdentityHandler: func(ctx context.Context, req *app.RequestContext) interface{} {
			claims := jwt.ExtractClaims(ctx, req)
			return claims[consts.JwtIdentityKey]
		},
		// 设置登录时为 token 添加自定义负载信息的函数，如果不传入这个参数，则 token 的 payload 部分默认存储 token 的过期时间和创建时间
		// data 对应 Authenticator 方法返回的第一个参数
		PayloadFunc: func(data interface{}) jwt.MapClaims {
			if v, ok := data.(string); ok {
				return jwt.MapClaims{
					consts.JwtIdentityKey: v,
				}
			}
			return jwt.MapClaims{}
		},
		// 认证用户的登录信息，配合 HertzJWTMiddleware.LoginHandler 使用
		Authenticator: func(ctx context.Context, req *app.RequestContext) (interface{}, error) {
			uid := req.GetInt64(consts.JwtIdentityKey)
			return strconv.FormatInt(uid, 10), nil
		},
		// 设置登录的响应函数
		LoginResponse: func(ctx context.Context, req *app.RequestContext, code int, token string, expire time.Time) {
			req.JSON(http.StatusOK, model.LoginResp{
				BaseResp: model.BaseResp{
					StatusCode: errno.SuccessCode,
					StatusMsg:  "登录成功",
				},
				Token:  &token,
				UserId: utils.GetLoggedInUID(req),
			})
		},
		// 设置 jwt 授权失败后的响应
		Unauthorized: func(ctx context.Context, c *app.RequestContext, code int, message string) {
			c.JSON(http.StatusOK, model.BaseResp{
				StatusCode: errno.AuthorizationFailedErrCode,
				StatusMsg:  message,
			})
		},
		// 一旦 jwt 校验流程产生错误将执行，并返回错误的 message，然后执行 Unauthorized
		HTTPStatusMessageFunc: func(e error, ctx context.Context, c *app.RequestContext) string {
			switch t := e.(type) {
			case errno.Errno:
				return t.Msg()
			default:
				hlog.CtxErrorf(ctx, "error = %v", t.Error())
				return errno.ServiceErr.Msg()
			}
		},
	})
}
