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

// PublishList 实现了 video.VideoService 接口
func (vs *VideoServiceImpl) PublishList(ctx context.Context, param *video.PublishListParam) (resp *video.PublishListResp, err error) {
	resp = new(video.PublishListResp)
	defer errProcess(&resp.BaseResp, &err)

	err = param.IsValid()
	if err != nil {
		return
	}

	videoInfos, err := publishList(ctx, param)
	if err != nil {
		return
	}

	resp.VideoList = videoInfos

	return
}

func publishList(ctx context.Context, param *video.PublishListParam) (resp []*common.VideoInfo, err error) {
	videos, err := db.QueryVideoList(ctx, param.LookUserId)
	if err != nil {
		return
	}

	var favoriteList []model.Favorite
	videoIds := modelVideoToVideoIdList(videos)
	if param.LoggedInUser != nil {
		favoriteList, err = db.QueryFavorite(ctx, *param.LoggedInUser, videoIds)
		if err != nil {
			return
		}
	}

	userIds := new([]int64)
	if param.LoggedInUser != nil && *param.LoggedInUser == param.LookUserId {
		*userIds = []int64{param.LookUserId}
	} else {
		*userIds = modelVideoToUserIdList(videos)
	}

	userInfos, err := userrpc.QueryUser(ctx, &user.QueryUserParam{
		LoggedUserId: param.LoggedInUser,
		UserIds:      *userIds,
	})
	if err != nil {
		return
	}

	resp = parseVideoInfoList(videos, userInfos, favoriteList)

	return
}
