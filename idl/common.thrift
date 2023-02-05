namespace go common

struct BaseResp {
    1: i16 code // 状态码
    2: string msg // 状态码描述
}

/**
 * 用户信息
 */
struct UserInfo {
    /**
     * 关注总数
     */
    1: i64 follow_count (to.tag = 'json:"follow_count,string"')
    /**
     * 粉丝总数
     */
    2: i64 follower_count (to.tag = 'json:"follower_count,string"')
    /**
     * 用户ID
     */
    3: i64 id (to.tag = 'json:"id,string"')
    /**
     * 关注状态，true-已关注，false-未关注
     */
    4: bool followed  (to.tag = 'json:"is_follow"')
    /**
     * 用户名称
     */
    5: string name (to.tag = 'json:"name"')
}

/**
 * 视频信息
 */
struct VideoInfo {
    /**
     * 视频作者信息
     */
    UserInfo author (go.tag = 'json:"author"')
    /**
     * 视频的评论总数
     */
    i64 comment_count (go.tag = 'json:"comment_count,string"')
    /**
     * 视频封面地址
     */
    string cover_url (go.tag = 'json:"cover_url"')
    /**
     * 视频的点赞总数
     */
    i64 favorite_count (go.tag = 'json:"favorite_count,string"')
    /**
     * 视频唯一标识
     */
    i64 id (go.tag = 'json:"id,string"')
    /**
     * 点赞状态，true-已点赞，false-未点赞
     */
    bool favorited (go.tag = 'json:"is_favorite"')
    /**
     * 视频播放地址
     */
    string play_url (go.tag = 'json:"play_url"')
    /**
     * 视频标题
     */
    string title (go.tag = 'json:"title"')
}