namespace go user

include "common.thrift"

struct CreateUserParam {
    1: string username // 用户名
    2: string password // 密码
}

struct CreateUserResp {
    1: common.BaseResp base_resp
    2: optional i64 user_id // 用户ID
}

struct CheckUserParam {
    1: string username // 用户名
    2: string password // 密码
}

struct CheckUserResp {
    1: common.BaseResp base_resp
    2: optional i64 user_id // 用户ID
}

struct QueryUserParam {
    1: i64 logged_user_id // 登录用户的 ID
    2: list<i64> user_ids // 用户 ID 列表
}

struct QueryUserResp {
    1: common.BaseResp base_resp
    2: list<common.UserInfo> user_list // 用户信息列表
}

struct FollowUserParam {
    1: i64 logged_user_id // 登录用户的 ID
    2: i64 follow_user_id // 要关注的用户 ID
    3: bool action // 操作：关注（true） ，取关（false）
}

struct FollowUserResp {
    1: common.BaseResp base_resp
}

struct FollowListParam {
    1: i64 logged_user_id // 登录用户的 ID
    2: i64 look_user_id // 要查询用户的 ID
}

struct FollowListResp {
    1: common.BaseResp base_resp
    2: list<common.UserInfo> user_list // 关注用户的信息列表
}

struct FollowerListParam {
    1: i64 logged_user_id // 登录用户的 ID
    2: i64 look_user_id // 要查询用户的 ID
}

struct FollowerListResp {
    1: common.BaseResp base_resp
    2: list<common.UserInfo> user_list // 粉丝的信息列表
}

service UserService {
    CreateUserResp CreateUser(1: CreateUserParam param)
    CheckUserResp CheckUser(1: CheckUserParam param)
    QueryUserResp QueryUser(1: QueryUserParam param)

    FollowUserResp FollowUser(1: FollowUserParam param)
    FollowListResp FollowList(1: FollowListParam param)
    FollowerListResp FollowerList(1: FollowerListParam param)
}