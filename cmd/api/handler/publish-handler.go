package handler

import (
	"context"
	"io"
	"mime/multipart"
	"strconv"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/h2non/filetype"
	"github.com/jaevor/go-nanoid"
	"github.com/wen-flower/easy-douyin/cmd/api/model"
	"github.com/wen-flower/easy-douyin/cmd/api/utils"
	"github.com/wen-flower/easy-douyin/kitex_gen/video"
	"github.com/wen-flower/easy-douyin/pkg/cos"
	"github.com/wen-flower/easy-douyin/pkg/errno"
	"github.com/wen-flower/easy-douyin/pkg/rpc/videorpc"
)

// PublishList 发布列表接口
// @router /douyin/publish/list [GET]
func PublishList(ctx context.Context, req *app.RequestContext) {
	var err error
	defer errProcess(req, &err)

	var param model.PublishListParam
	if err = req.BindAndValidate(&param); err != nil {
		return
	}

	videoInfos, err := videorpc.PublishList(ctx, &video.PublishListParam{
		LookUserId:   param.UserId,
		LoggedInUser: utils.GetLoggedInUID(req),
	})

	resp := new(model.PublishListResp)
	resp.Ok()
	resp.VideoList = videoInfos

	utils.SendJson(req, resp)
}

var vidGenerator, _ = nanoid.CustomASCII("0123456789", 18)

// PublishVideo 投稿视频
// @router /douyin/publish/action [POST]
func PublishVideo(ctx context.Context, req *app.RequestContext) {
	var err error
	defer errProcess(req, &err)

	var param model.PublishVideoParam
	if err = req.BindAndValidate(&param); err != nil {
		return
	}

	file, err := param.Data.Open()
	defer func(file multipart.File) {
		err := file.Close()
		if err != nil {
			hlog.Errorf("err = %v", err)
		}
	}(file)
	if err != nil {
		return
	}

	head := make([]byte, 261)
	_, err = file.Read(head)
	if err != nil {
		return
	}
	if !filetype.IsVideo(head) {
		err = errno.AuthorizationFailedErr
		return
	}

	_, err = file.Seek(0, io.SeekStart)
	if err != nil {
		return
	}

	vid := vidGenerator()
	err = cos.VideoUpload(vid, file) // tp.MIME.Value
	if err != nil {
		hlog.Errorf("upload error = %v", err)
		return
	}

	vidInt64, err := strconv.ParseInt(vid, 10, 64)
	if err != nil {
		return
	}
	err = videorpc.CreateVideo(ctx, &video.CreateVideoParam{
		VideoId: vidInt64,
		Title:   param.Title,
		UserId:  *utils.GetLoggedInUID(req),
	})
	if err != nil {
		return
	}

	resp := model.PublishVideoResp{}
	resp.Ok()
	utils.SendJson(req, resp)
}
