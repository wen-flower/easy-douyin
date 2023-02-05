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
    1: i64 follow_count (to.tag = 'json:"follow_count"')
    /**
     * 粉丝总数
     */
    2: i64 follower_count (to.tag = 'json:"follower_count"')
    /**
     * 用户ID
     */
    3: i64 id (to.tag = 'json:"id"')
    /**
     * 关注状态，true-已关注，false-未关注
     */
    4: bool followed  (to.tag = 'json:"is_follow"')
    /**
     * 用户名称
     */
    5: string name (to.tag = 'json:"name"')
}