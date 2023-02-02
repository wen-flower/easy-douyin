package handler

import (
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/wen-flower/easy-douyin/cmd/api/model"
	"github.com/wen-flower/easy-douyin/pkg/errno"
	"net/http"
)

// 提取 API 错误处理流程
func errProcess(req *app.RequestContext, err *error) {
	if *err != nil {
		_errno := errno.ConvertErr(*err)
		req.JSON(http.StatusOK, model.BaseResp{
			StatusCode: _errno.Code(),
			StatusMsg:  _errno.Msg(),
		})
	}
}
