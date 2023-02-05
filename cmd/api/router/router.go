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

		_relation := _douyin.Group("/relation", mw.JwtMiddleware.MiddlewareFunc())
		{
			_relation.GET("/follow/list", handler.FollowList)
			_relation.GET("/follower/list", handler.FollowerList)
			_relation.POST("/action", handler.FollowAction)
		}

		_douyin.GET("/feed", handler.VideoFeed)

		_publish := _douyin.Group("/publish", mw.JwtMiddleware.MiddlewareFunc())
		{
			_publish.GET("/list", handler.PublishList)
		}
	}
}
