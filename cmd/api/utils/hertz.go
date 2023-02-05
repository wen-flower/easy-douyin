package utils

import (
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/wen-flower/easy-douyin/cmd/api/consts"
	"net/http"
	"strconv"
)

// GetLoggedInUID 获取登录用户的 UID，如果没有登录则返回 nil
func GetLoggedInUID(req *app.RequestContext) *int64 {
	// 存的是字符串，int64会丢失精度
	uidStr := req.GetString(consts.JwtIdentityKey)
	if uidStr == "" {
		return nil
	}
	uid, err := strconv.ParseInt(uidStr, 10, 64)
	if err != nil {
		return nil
	}
	return &uid
}

// SetLoggedInUID 设置登录用户的 UID
func SetLoggedInUID(req *app.RequestContext, uid int64) {
    req.Set(consts.JwtIdentityKey, strconv.FormatInt(uid, 10))
}

// SendJson 返回响应请求的 JSON 数据
func SendJson(req *app.RequestContext, data interface{}) {
	req.JSON(http.StatusOK, data)
}
