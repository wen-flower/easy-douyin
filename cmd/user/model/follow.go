package model

import "time"

type Follow struct {
	CreatedAt    time.Time `gorm:"column:created_at"`     //第一次关注时间
	FollowedUser int64     `gorm:"column:followed_user"`  //关注的用户
	ID           int64     `gorm:"column:id;primary_key"` //主键ID
	Status       int       `gorm:"column:status"`         //是否关注（1：关注，0：取消关注）
	UID          int64     `gorm:"column:uid"`            //用户ID
	UpdatedAt    time.Time `gorm:"column:updated_at"`     //最后一次更新时间
}

// TableName 结构体对应的数据库表名
func (f *Follow) TableName() string {
	return "tb_follow"
}

// 结构体字段对应的数据表列名常量
const (
	FollowCreatedAt    = "created_at"
	FollowFollowedUser = "followed_user"
	FollowID           = "id"
	FollowStatus       = "status"
	FollowUID          = "uid"
	FollowUpdatedAt    = "updated_at"
)
