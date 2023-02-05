package model

import "time"

type Video struct {
	ID            int64     `gorm:"column:id;primary_key"` //主键ID
	Vid           int64     `gorm:"column:vid"`            //视频ID
	UID           int64     `gorm:"column:uid"`            //用户ID
	Title         string    `gorm:"column:title"`          //视频标题
	FavoriteCount int64     `gorm:"column:favorite_count"` //视频点赞数
	CommentCount  int64     `gorm:"column:comment_count"`  //视频评论数
	CreatedAt     time.Time `gorm:"column:created_at"`     //视频发布时间
	UpdatedAt     time.Time `gorm:"column:updated_at"`     //最后一次更新时间
	DeletedAt     time.Time `gorm:"column:deleted_at"`     //逻辑删除的时间
}

// TableName 结构体对应的数据库表名
func (v *Video) TableName() string {
	return "tb_video"
}

// 结构体字段对应的数据表列名常量
const (
	VideoID            = "id"
	VideoVid           = "vid"
	VideoUID           = "uid"
	VideoTitle         = "title"
	VideoFavoriteCount = "favorite_count"
	VideoCommentCount  = "comment_count"
	VideoCreatedAt     = "created_at"
	VideoUpdatedAt     = "updated_at"
	VideoDeletedAt     = "deleted_at"
)
