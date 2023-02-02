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

service UserService {
    CreateUserResp CreateUser(1: CreateUserParam param)
    CheckUserResp CheckUser(1: CheckUserParam param)
}