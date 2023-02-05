package db

import (
	"context"

	"github.com/wen-flower/easy-douyin/cmd/video/model"
	"github.com/wen-flower/easy-douyin/pkg/msql"
	"gorm.io/gorm"
)

// CreateComment 创建评论
func CreateComment(ctx context.Context, comment *model.Comment) error {
	return DB.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		err := tx.Create(comment).Error
		if err != nil {
			return err
		}

		err = tx.Model(&model.Video{}).Where(msql.Eq(model.VideoVid), comment.Vid).Update(model.VideoCommentCount, msql.Inc(model.VideoCommentCount)).Error

		return err
	})
}

// DeleteComment 删除评论
func DeleteComment(ctx context.Context, uid int64, commentId int64, videoId int64) error {
	return DB.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		res := DB.Where(msql.Eq(model.CommentUID), uid).Delete(&model.Comment{
			ID: commentId,
		})
		err := res.Error
		if err != nil {
			return err
		}
		if res.RowsAffected == 0 {
			return nil
		}

		err = tx.Model(&model.Video{}).Where(msql.Eq(model.VideoVid), videoId).Update(model.VideoCommentCount, msql.Dec(model.VideoCommentCount)).Error

		return err
	})
}

// QueryCommentList 查询评论列表
func QueryCommentList(ctx context.Context, videoId int64) ([]model.Comment, error) {
	var resp []model.Comment
	err := DB.WithContext(ctx).Where(msql.Eq(model.CommentVid), videoId).Find(&resp).Error
	return resp, err
}
