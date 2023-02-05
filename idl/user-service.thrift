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

service UserService {
    CreateUserResp CreateUser(1: CreateUserParam param)
    CheckUserResp CheckUser(1: CheckUserParam param)
    QueryUserResp QueryUser(1: QueryUserParam param)
}