package handler

import (
	"context"
	"strconv"

	"github.com/jaevor/go-nanoid"
	"github.com/wen-flower/easy-douyin/cmd/user/cfg"
	"github.com/wen-flower/easy-douyin/cmd/user/dal/db"
	"github.com/wen-flower/easy-douyin/cmd/user/model"
	"github.com/wen-flower/easy-douyin/kitex_gen/user"
	"github.com/wen-flower/easy-douyin/pkg/errno"
)

// CreateUser 实现了 user.UserService 接口
func (us *UserServiceImpl) CreateUser(ctx context.Context, param *user.CreateUserParam) (resp *user.CreateUserResp, err error) {
	resp = new(user.CreateUserResp)
	defer errProcess(&resp.BaseResp, &err)

	err = param.IsValid()
	if err != nil {
		return
	}

	uid, err := createUser(ctx, param)
	if err != nil {
		return
	}

	resp.UserId = &uid

	return
}

func createUser(ctx context.Context, param *user.CreateUserParam) (int64, error) {
	uidGenerator, err := nanoid.CustomASCII("0123456789", cfg.UserIdLength)
	if err != nil {
		return 0, err
	}

	exists, err := db.ExistsUserByUsername(ctx, param.Username)
	if err != nil {
		return 0, err
	}
	if exists {
		return 0, errno.UsernameAlreadyExistErr
	}

	uid, err := strconv.ParseInt(uidGenerator(), 10, 64)
	if err != nil {
		return 0, nil
	}

	_user := &model.User{
		UID:      uid,
		Username: param.Username,
		Password: passwordDigest(param.Password),
	}

	err = db.CreateUser(ctx, _user)

	return uid, err
}
