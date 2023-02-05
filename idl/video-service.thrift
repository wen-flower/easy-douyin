namespace go video

include "common.thrift"

struct CreateVideoParam {
    1: i64 video_id // 视频 ID
    2: string title // 视频标题
    3: i64 user_id // 用户 ID
}

struct CreateVideoResp {
    1: common.BaseResp base_resp
}

struct PublishListParam {
    1: i64 look_user_id // 要查询的用户ID
    2: optional i64 logged_in_user // 登录用户的ID
}

struct PublishListResp {
    1: common.BaseResp base_resp
    2: list<common.VideoInfo> video_list
}

struct VideoFeedParam {
    1: i64 latest_time // 限制返回视频的最新投稿时间戳，精确到秒，不填表示当前时间
    2: optional i64 logged_in_user // 登录用户的ID
}

struct VideoFeedResp {
    1: common.BaseResp base_resp
    2: list<common.VideoInfo> video_list
}

struct FavoriteVideoParam{
    1: i64 logged_user_id // 登录用户的 ID
    2: i64 video_id // 要点赞的视频
    3: bool action // 操作：关注（true） ，取关（false）
}  

struct FavoriteVideoResp{
    1: common.BaseResp base_resp
}  

struct FavoriteListParam{
    1: optional i64 logged_user_id // 登录用户的 ID
    2: i64 look_user_id // 要查询用户的 ID
}  

struct FavoriteListResp{
    1: common.BaseResp base_resp
    2: list<common.VideoInfo> video_list
}  

struct CommentVideoParam{
    1: i64 logged_user_id // 登录用户的 ID
    2: i64 video_id // 要评论的视频 ID
    3: string comment_text // 评论的内容
}  

struct CommentVideoResp{
    1: common.BaseResp base_resp
    2: common.CommentInfo comment
} 

struct DeleteCommentParam{
    1: i64 logged_user_id // 登录用户的 ID
    2: i64 video_id // 要删除评论的视频 ID
    3: i64 comment_id // 要删除的评论ID
}  

struct DeleteCommentResp{
    1: common.BaseResp base_resp
}  

struct CommentListParam{
    1: optional i64 logged_user_id // 登录用户的 ID
    2: i64 video_id // 要查看评论的视频 ID
}

struct CommentListResp{
    1: common.BaseResp base_resp
    2: list<common.CommentInfo> comment_list
}

service VideoService {
    CreateVideoResp CreateVideo(1: CreateVideoParam param)
    PublishListResp PublishList(1: PublishListParam param)
    VideoFeedResp VideoFeed(1: VideoFeedParam param)

    FavoriteVideoResp FavoriteVideo(1: FavoriteVideoParam param)
    FavoriteListResp FavoriteList(1: FavoriteListParam param)

    CommentVideoResp CommentVideo(1: CommentVideoParam param)
    DeleteCommentResp DeleteComment(1: DeleteCommentParam param)
    CommentListResp CommentList(1: CommentListParam param)
}