package handler

import (
	"context"

	"github.com/wen-flower/easy-douyin/cmd/video/dal/db"
	"github.com/wen-flower/easy-douyin/cmd/video/model"
	"github.com/wen-flower/easy-douyin/kitex_gen/common"
	"github.com/wen-flower/easy-douyin/kitex_gen/user"
	"github.com/wen-flower/easy-douyin/kitex_gen/video"
	"github.com/wen-flower/easy-douyin/pkg/rpc/userrpc"
)

// FavoriteList 实现了 video.VideoService 接口
func (vs *VideoServiceImpl) FavoriteList(ctx context.Context, param *video.FavoriteListParam) (resp *video.FavoriteListResp, err error) {
	resp = new(video.FavoriteListResp)
	defer errProcess(&resp.BaseResp, &err)

	err = param.IsValid()
	if err != nil {
		return
	}

	videoInfos, err := favoriteList(ctx, param)
	if err != nil {
		return
	}

	resp.VideoList = videoInfos

	return
}

func favoriteList(ctx context.Context, param *video.FavoriteListParam) (resp []*common.VideoInfo, err error) {
	favoriteList, err := db.QueryFavoriteList(ctx, param.LookUserId)
	if err != nil {
		return nil, err
	}

	videoIds := modelFavoriteToVideoIdList(favoriteList)

	videos, err := db.QueryVideo(ctx, videoIds)
	if err != nil {
		return
	}

	if param.LoggedUserId != nil {
		favoriteList, err = db.QueryFavorite(ctx, *param.LoggedUserId, videoIds)
		if err != nil {
			return
		}
	} else {
		favoriteList = []model.Favorite{}
	}

	userIds := modelVideoToUserIdList(videos)

	userInfos, err := userrpc.QueryUser(ctx, &user.QueryUserParam{
		LoggedUserId: param.LoggedUserId,
		UserIds:      userIds,
	})
	if err != nil {
		return
	}

	resp = parseVideoInfoList(videos, userInfos, favoriteList)

	return
}
