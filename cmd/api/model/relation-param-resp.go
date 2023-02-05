package model

import "github.com/wen-flower/easy-douyin/kitex_gen/common"

// FollowActionParam 关注操作参数
type FollowActionParam struct {
	// ToUserId 对方用户id
	ToUserId int64 `json:"to_user_id" query:"to_user_id" vd:"@:$ > 0; msg:'用户ID参数错误'"`
	// Action 1-关注，2-取消关注
	Action int16 `json:"action_type" query:"action_type" vd:"@:$ >= 1 && $<=2; msg:'不支持的操作'"`
}

// FollowActionResp 关系操作响应数据
type FollowActionResp struct {
	BaseResp
}

// FollowListParam 查询关注列表参数
type FollowListParam struct {
	// LookUserId 要查询的用户 ID
	LookUserId int64 `json:"user_id" query:"user_id" vd:"@:$ >0; msg:'用户ID参数错误'"`
}

// FollowListResp 查询关注列表响应数据
type FollowListResp struct {
	BaseResp
	UserList []*common.UserInfo `json:"user_list"`
}

// FollowerListParam 查询粉丝列表参数
type FollowerListParam struct {
	// LookUserId 要查询的用户 ID
	LookUserId int64 `json:"user_id" query:"user_id" vd:"@:$ >0; msg:'用户ID参数错误'"`
}

// FollowerListResp 查询粉丝列表响应数据
type FollowerListResp struct {
	BaseResp
	UserList []*common.UserInfo `json:"user_list"`
}
