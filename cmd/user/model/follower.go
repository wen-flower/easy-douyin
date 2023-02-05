package model

import "time"

type Follower struct {
	ID           int64     `gorm:"column:id;primary_key"` //主键ID
	UID          int64     `gorm:"column:uid"`            //用户ID
	FollowerUser int64     `gorm:"column:follower_user"`  //粉丝ID
	Status       int       `gorm:"column:status"`         //是否关注（1：关注，0：取消关注）
	CreatedAt    time.Time `gorm:"column:created_at"`     //第一次关注时间
	UpdatedAt    time.Time `gorm:"column:updated_at"`     //最后一次更新时间
}

// TableName 结构体对应的数据库表名
func (f *Follower) TableName() string {
	return "tb_follower"
}

// 结构体字段对应的数据表列名常量
const (
	FollowerID           = "id"
	FollowerUID          = "uid"
	FollowerFollowerUser = "follower_user"
	FollowerStatus       = "status"
	FollowerCreatedAt    = "created_at"
	FollowerUpdatedAt    = "updated_at"
)
