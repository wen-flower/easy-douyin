package db

import (
	"context"
	"errors"
	"github.com/wen-flower/easy-douyin/cmd/user/model"
	"github.com/wen-flower/easy-douyin/pkg/msql"
	"gorm.io/gorm"
)

// CreateUser 创建用户信息
func CreateUser(ctx context.Context, user *model.User) error {
	return DB.WithContext(ctx).Create(user).Error
}

// QueryUserIdAndPasswordByUsername 根据用户名查询用户信息
func QueryUserIdAndPasswordByUsername(ctx context.Context, username string) (*model.User, error) {
	var user model.User
	e := DB.WithContext(ctx).Select(
		model.UserUID,
		model.UserPassword,
	).Where(
		msql.Eq(model.UserUsername), username,
	).First(&user).Error

	if errors.Is(e, gorm.ErrRecordNotFound) {
		return nil, nil
	}

	if e != nil {
		return nil, e
	}

	return &user, nil
}

// ExistsUserByUsername 根据用户名查询用户信息
func ExistsUserByUsername(ctx context.Context, username string) (bool, error) {
	var user model.User
	e := DB.WithContext(ctx).Select(
		model.UserUsername,
	).Where(
		msql.Eq(model.UserUsername), username,
	).First(&user).Error

	if errors.Is(e, gorm.ErrRecordNotFound) {
		return false, nil
	}

	if e != nil {
		return false, e
	}

	return true, nil
}

// QueryUser 查询用户数据
func QueryUser(ctx context.Context, userIds []int64) ([]model.User, error) {
	var users []model.User
	e := DB.WithContext(ctx).Where(msql.In(model.UserUID), userIds).Find(&users).Error

	return users, e
}
