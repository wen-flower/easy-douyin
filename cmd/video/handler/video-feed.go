package handler

import (
	"context"

	"github.com/wen-flower/easy-douyin/cmd/video/consts"
	"github.com/wen-flower/easy-douyin/cmd/video/dal/db"
	"github.com/wen-flower/easy-douyin/cmd/video/model"
	"github.com/wen-flower/easy-douyin/kitex_gen/common"
	"github.com/wen-flower/easy-douyin/kitex_gen/user"
	"github.com/wen-flower/easy-douyin/kitex_gen/video"
	"github.com/wen-flower/easy-douyin/pkg/rpc/userrpc"
)

// VideoFeed 实现了 video.VideoService 接口
func (vs *VideoServiceImpl) VideoFeed(ctx context.Context, param *video.VideoFeedParam) (resp *video.VideoFeedResp, err error) {
	resp = new(video.VideoFeedResp)
	defer errProcess(&resp.BaseResp, &err)

	err = param.IsValid()
	if err != nil {
		return
	}

	return
}

func videoFeed(ctx context.Context, param *video.VideoFeedParam) (resp []*common.VideoInfo, err error) {
	videos, err := db.QueryVideoFeed(ctx, param.LatestTime, consts.VideoFeedSize)
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

	userIds := modelVideoToUserIdList(videos)

	userInfos, err := userrpc.QueryUser(ctx, &user.QueryUserParam{
		LoggedUserId: param.LoggedInUser,
		UserIds:      userIds,
	})
	if err != nil {
		return
	}

	resp = parseVideoInfoList(videos, userInfos, favoriteList)

	return
}
