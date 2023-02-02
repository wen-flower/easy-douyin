package model

import (
	"gorm.io/gorm"
	"time"
)

type User struct {
	ID            int64          `gorm:"column:id;primary_key"` //主键ID
	UID           int64          `gorm:"column:uid"`            //用户ID
	Username      string         `gorm:"column:username"`       //用户名
	Password      string         `gorm:"column:password"`       //用户密码
	FollowCount   int64          `gorm:"column:follow_count"`   //关注总数
	FollowerCount int64          `gorm:"column:follower_count"` //粉丝总数
	CreatedAt     time.Time      `gorm:"column:created_at"`     //注册时间
	UpdatedAt     time.Time      `gorm:"column:updated_at"`     //最后一次更新时间
	DeletedAt     gorm.DeletedAt `gorm:"column:deleted_at"`     //逻辑删除的时间
}

// TableName 结构体对应的数据库表名
func (u *User) TableName() string {
	return "tb_user"
}

// 结构体字段对应的数据表列名常量
const (
	UserID            = "id"
	UserUID           = "uid"
	UserUsername      = "username"
	UserPassword      = "password"
	UserFollowCount   = "follow_count"
	UserFollowerCount = "follower_count"
	UserCreatedAt     = "created_at"
	UserUpdatedAt     = "updated_at"
	UserDeletedAt     = "deleted_at"
)
