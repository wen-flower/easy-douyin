package router

import (
	"github.com/cloudwego/hertz/pkg/app/server"
	"github.com/wen-flower/easy-douyin/cmd/api/handler"
	"github.com/wen-flower/easy-douyin/cmd/api/mw"
)

func Register(h *server.Hertz) {
	h.Use(globalMw()...)

	_douyin := h.Group("/douyin")
	{
		_user := _douyin.Group("/user")
		{
			_user.GET("", mw.JwtMiddleware.MiddlewareFunc(), handler.UserInfo)

			_user.POST("/login", handler.Login)
			_user.POST("/register", handler.Register)
		}
	}
}
