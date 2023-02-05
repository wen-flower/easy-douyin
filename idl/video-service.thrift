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

service VideoService {
    CreateVideoResp CreateVideo(1: CreateVideoParam param)
    PublishListResp PublishList(1: PublishListParam param)
    VideoFeedResp VideoFeed (1: VideoFeedParam param)
}