package model

import "time"

type Favorite struct {
	ID        int64     `gorm:"column:id;primary_key"` //主键ID
	Vid       int64     `gorm:"column:vid"`            //视频ID
	UID       int64     `gorm:"column:uid"`            //用户ID
	Status    int       `gorm:"column:status"`         //是否点赞（1：已点赞，0：未点赞）
	CreatedAt time.Time `gorm:"column:created_at"`     //第一次点赞时间
	UpdatedAt time.Time `gorm:"column:updated_at"`     //最后一次更新时间
}

// TableName 结构体对应的数据库表名
func (f *Favorite) TableName() string {
	return "tb_favorite"
}

// 结构体字段对应的数据表列名常量
const (
	FavoriteID        = "id"
	FavoriteVid       = "vid"
	FavoriteUID       = "uid"
	FavoriteStatus    = "status"
	FavoriteCreatedAt = "created_at"
	FavoriteUpdatedAt = "updated_at"
)
