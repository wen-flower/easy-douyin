package model

import (
	"time"

	"gorm.io/gorm"
)

type Comment struct {
	ID        int64          `gorm:"column:id;primary_key"` //主键ID
	Vid       int64          `gorm:"column:vid"`            //视频ID
	UID       int64          `gorm:"column:uid"`            //用户ID
	Content   string         `gorm:"column:content"`        //评论内容
	CreatedAt time.Time      `gorm:"column:created_at"`     //评论时间
	UpdatedAt time.Time      `gorm:"column:updated_at"`     //最后一次更新时间
	DeletedAt gorm.DeletedAt `gorm:"column:deleted_at"`     //逻辑删除的时间
}

// TableName 结构体对应的数据库表名
func (c *Comment) TableName() string {
	return "tb_comment"
}

// 结构体字段对应的数据表列名常量
const (
	CommentID        = "id"
	CommentVid       = "vid"
	CommentUID       = "uid"
	CommentContent   = "content"
	CommentCreatedAt = "created_at"
	CommentUpdatedAt = "updated_at"
	CommentDeletedAt = "deleted_at"
)
