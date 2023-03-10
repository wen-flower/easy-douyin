package mw

import (
	"context"
	"errors"
	"net/http"
	"strconv"
	"time"

	"github.com/wen-flower/easy-douyin/cmd/api/consts"
	"github.com/wen-flower/easy-douyin/cmd/api/model"
	"github.com/wen-flower/easy-douyin/cmd/api/utils"
	"github.com/wen-flower/easy-douyin/pkg/errno"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/golang-jwt/jwt/v4"
	jwtmw "github.com/hertz-contrib/jwt"
)

var JwtMiddleware *jwtmw.HertzJWTMiddleware

// OptionalJwtMiddlewareFunc 有 token 则处理，没有则不处理
func OptionalJwtMiddlewareFunc() app.HandlerFunc {
	return func(ctx context.Context, req *app.RequestContext) {
		_, err := JwtMiddleware.ParseToken(ctx, req)
		if errors.Is(err, jwtmw.ErrEmptyFormToken) || errors.Is(err, jwtmw.ErrEmptyQueryToken) || errors.Is(err, jwtmw.ErrEmptyAuthHeader) {
			return
		}
		if err != nil {
			msg := JwtMiddleware.HTTPStatusMessageFunc(err, ctx, req)
			JwtMiddleware.Unauthorized(ctx, req, 200, msg)
			req.Abort()
			return
		}
		JwtMiddleware.MiddlewareFunc()(ctx, req)
	}
}

func InitJWT() {
	JwtMiddleware, _ = jwtmw.New(&jwtmw.HertzJWTMiddleware{
		Key:           []byte(consts.JwtSecretKey),
		TokenLookup:   "header: Authorization, query: token, form: token", //cookie: jwt
		TokenHeadName: "Bearer",
		TimeFunc:      time.Now,
		Timeout:       time.Hour,
		MaxRefresh:    time.Hour,
		IdentityKey:   consts.JwtIdentityKey,
		// 设置从 token 提取用户信息的函数
		IdentityHandler: func(ctx context.Context, req *app.RequestContext) interface{} {
			claims := jwtmw.ExtractClaims(ctx, req)
			return claims[consts.JwtIdentityKey]
		},
		// 设置登录时为 token 添加自定义负载信息的函数，如果不传入这个参数，则 token 的 payload 部分默认存储 token 的过期时间和创建时间
		// data 对应 Authenticator 方法返回的第一个参数
		PayloadFunc: func(data interface{}) jwtmw.MapClaims {
			if v, ok := data.(string); ok {
				return jwtmw.MapClaims{
					consts.JwtIdentityKey: v,
				}
			}
			return jwtmw.MapClaims{}
		},
		// 认证用户的登录信息，配合 HertzJWTMiddleware.LoginHandler 使用
		Authenticator: func(ctx context.Context, req *app.RequestContext) (interface{}, error) {
			uid := utils.GetLoggedInUID(req)
			if uid == nil {
				return nil, errno.AuthorizationFailedErr
			}
			return strconv.FormatInt(*uid, 10), nil
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
				if errors.Is(e, jwt.ErrTokenExpired) {
					return "登录过期"
				}
				if errors.Is(e, jwtmw.ErrEmptyFormToken) || errors.Is(e, jwtmw.ErrEmptyQueryToken) || errors.Is(e, jwtmw.ErrEmptyAuthHeader) {
					return "请先登录"
				}
				hlog.CtxErrorf(ctx, "error = %v", t.Error())
				return errno.ServiceErr.Msg()
			}
		},
	})
}
