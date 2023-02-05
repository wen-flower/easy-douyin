package handler

import (
	"context"

	"github.com/wen-flower/easy-douyin/cmd/video/dal/db"
	"github.com/wen-flower/easy-douyin/cmd/video/model"
	"github.com/wen-flower/easy-douyin/kitex_gen/video"
)

// CreateVideo 实现了 video.VideoService
func (vs *VideoServiceImpl) CreateVideo(ctx context.Context, param *video.CreateVideoParam) (resp *video.CreateVideoResp, err error) {
	resp = new(video.CreateVideoResp)
	defer errProcess(&resp.BaseResp, &err)

	err = param.IsValid()
	if err != nil {
		return
	}

	err = createVideo(ctx, param)

	return
}

func createVideo(ctx context.Context, param *video.CreateVideoParam) error {
	return db.CreateVideo(ctx, &model.Video{
		Vid:   param.VideoId,
		UID:   param.UserId,
		Title: param.Title,
	})
}
