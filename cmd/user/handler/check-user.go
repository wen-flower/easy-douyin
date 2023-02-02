package handler

import (
	"context"
	"github.com/wen-flower/easy-douyin/cmd/user/dal/db"
	"github.com/wen-flower/easy-douyin/kitex_gen/user"
	"github.com/wen-flower/easy-douyin/pkg/errno"
)

// CheckUser 实现了 user.UserService 接口
func (us *UserServiceImpl) CheckUser(ctx context.Context, param *user.CheckUserParam) (resp *user.CheckUserResp, err error) {
	resp = new(user.CheckUserResp)
	defer errProcess(&resp.BaseResp, &err)

	err = param.IsValid()
	if err != nil {
		return
	}

	var uid *int64
	uid, err = checkUser(ctx, param)
	if err != nil {
		return
	}

	resp.UserId = uid

	return
}

func checkUser(ctx context.Context, param *user.CheckUserParam) (*int64, error) {
	_user, err := db.QueryUserIdAndPasswordByUsername(ctx, param.Username)
	if err != nil {
		return nil, err
	}
	if _user == nil || _user.Password != passwordDigest(param.Password) {
		err = errno.AuthorizationFailedErr
		return nil, err
	}

	return &_user.UID, nil
}
