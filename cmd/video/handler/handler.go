package handler

import (
	"github.com/wen-flower/easy-douyin/cmd/video/model"
	"github.com/wen-flower/easy-douyin/kitex_gen/common"
	"github.com/wen-flower/easy-douyin/kitex_gen/video"
	"github.com/wen-flower/easy-douyin/pkg/constant"
	"github.com/wen-flower/easy-douyin/pkg/errno"
)

// VideoServiceImpl implements the last service interface defined in the IDL.
type VideoServiceImpl struct{}

// 检查 VideoServiceImpl 是否实现了 video.VideoService 接口
var _ video.VideoService = (*VideoServiceImpl)(nil)

// 提取出错误的处理流程
func errProcess(baseResp **common.BaseResp, err *error) {
	var resp common.BaseResp
	if *err != nil {
		e := errno.ConvertErr(*err)
		resp.Msg = e.Msg()
		resp.Code = e.Code()
		*err = nil
	} else {
		resp.Code = errno.Success.Code()
		resp.Msg = errno.Success.Msg()
	}
	*baseResp = &resp
}

// 将 model.Video 列表和 common.UserInfo 列表以及 model.Favorite 列表组合为 common.VideoInfo 列表
func parseVideoInfoList(videos []model.Video, userInfos []*common.UserInfo, favoriteList []model.Favorite) []*common.VideoInfo {
	resp := make([]*common.VideoInfo, 0, len(videos))
	userInfoMap := userInfoToMap(userInfos)
	favoriteMap := modelFavoriteToMap(favoriteList)
	for _, _video := range videos {
		resp = append(resp, &common.VideoInfo{
			Author:        userInfoMap[_video.UID],
			CommentCount:  _video.CommentCount,
			CoverUrl:      constant.ParseVideoCoverUrl(_video.Vid),
			FavoriteCount: _video.FavoriteCount,
			Id:            _video.Vid,
			Favorited:     favoriteMap[_video.Vid],
			PlayUrl:       constant.ParseVideoUrl(_video.Vid),
			Title:         _video.Title,
		})
	}
	return resp
}

// 将 model.Comment 列表和 common.UserInfo 列表组合为 common.CommentInfo 列表
func parseCommentInfoList(commentList []model.Comment, userInfos []*common.UserInfo) []*common.CommentInfo {
	resp := make([]*common.CommentInfo, 0, len(commentList))
	userInfoMap := userInfoToMap(userInfos)
	for _, _comment := range commentList {
		resp = append(resp, &common.CommentInfo{
			Content:    _comment.Content,
			CreateDate: _comment.CreatedAt.Format("01-02"),
			Id:         _comment.ID,
			User:       userInfoMap[_comment.UID],
		})
	}
	return resp
}

// 将 model.Favorite 列表转为 map
func modelFavoriteToMap(favoriteList []model.Favorite) map[int64]bool {
	resp := make(map[int64]bool, len(favoriteList))
	for _, favorite := range favoriteList {
		resp[favorite.Vid] = favorite.Status == 1
	}
	return resp
}

// 将 common.UserInfo 列表转为 map
func userInfoToMap(userInfos []*common.UserInfo) map[int64]*common.UserInfo {
	resp := make(map[int64]*common.UserInfo, len(userInfos))
	for _, ui := range userInfos {
		resp[ui.Id] = ui
	}
	return resp
}

// 将 model.Video 列表转为视频ID列表
func modelVideoToVideoIdList(videos []model.Video) []int64 {
	resp := make([]int64, 0, len(videos))
	for _, _video := range videos {
		resp = append(resp, _video.Vid)
	}
	return resp
}

// 将 model.Video 列表转为视频发布者用户ID列表(不重复)
func modelVideoToUserIdList(videos []model.Video) []int64 {
	size := len(videos)
	idMap := make(map[int64]struct{}, size)
	resp := make([]int64, 0, size)
	for _, _video := range videos {
		if _, ok := idMap[_video.UID]; !ok {
			resp = append(resp, _video.UID)
		}
	}
	return resp
}

// 将 model.Favorite 列表转为视频ID列表
func modelFavoriteToVideoIdList(favoriteList []model.Favorite) []int64 {
	resp := make([]int64, 0, len(favoriteList))
	for _, _favorite := range favoriteList {
		resp = append(resp, _favorite.Vid)
	}
	return resp
}

// 将 model.Comment 列表转为评论发布者用户ID列表(不重复)
func modelCommentToUserIdList(commentList []model.Comment) []int64 {
	size := len(commentList)
	idMap := make(map[int64]struct{}, size)
	resp := make([]int64, 0, size)
	for _, _comment := range commentList {
		if _, ok := idMap[_comment.UID]; !ok {
			resp = append(resp, _comment.UID)
		}
	}
	return resp
}
