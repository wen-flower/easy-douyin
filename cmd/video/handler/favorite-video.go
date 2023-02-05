package handler

import (
	"context"

	"github.com/wen-flower/easy-douyin/cmd/video/dal/db"
	"github.com/wen-flower/easy-douyin/kitex_gen/video"
)

// FollowUser 实现了 user.UserService 接口
func (vs *VideoServiceImpl) FavoriteVideo(ctx context.Context, param *video.FavoriteVideoParam) (resp *video.FavoriteVideoResp, err error) {
	resp = new(video.FavoriteVideoResp)
	defer errProcess(&resp.BaseResp, &err)

	err = param.IsValid()
	if err != nil {
		return
	}

	err = followUser(ctx, param)

	return
}

func followUser(ctx context.Context, param *video.FavoriteVideoParam) error {
	if param.Action {
		return db.FavoriteVideo(ctx, param.LoggedUserId, param.VideoId)
	}
	return db.CancelFavoriteVideo(ctx, param.LoggedUserId, param.VideoId)
}
