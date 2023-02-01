namespace go user

include "common.thrift"

struct CreateUserRequest {

}

struct CreateUserResponse {

}

service UserService {
    CreateUserResponse CreateUser(1: CreateUserRequest req)
}