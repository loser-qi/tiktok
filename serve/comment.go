package serve

import (
	"gorm.io/gorm"
	"tiktok/util"
	"time"
)

// GetCommentListByVideoId 根据videoId获取评论列表
func GetCommentListByVideoId(videoId int64) []Comment {
	video := &Video{
		Id: videoId,
	}
	util.Db.Preload("CommentList", func(db *gorm.DB) *gorm.DB {
		return db.Order("create_time desc")
	}).Find(video)
	return video.CommentList
}

// GetVideoIdByCommentId 根据id获取评论所在视频
func GetVideoIdByCommentId(id int64) int64 {
	var videoId int64
	util.Db.Model(&Comment{}).Select("video_id").Where("id = ?", id).Take(&videoId)
	return videoId
}

// SaveComment 保存评论
func SaveComment(userId, videoId int64, text string, createTime time.Time) error {
	comment := &Comment{
		UserId:     userId,
		VideoId:    videoId,
		Text:       text,
		CreateTime: createTime,
	}
	resErr := util.Db.Transaction(func(tx *gorm.DB) error {
		if err := tx.Create(comment); err.Error != nil {
			return err.Error
		}
		if err := UpdateVideoCommentCount(videoId, 1, tx); err != nil {
			return err
		}
		return nil
	})
	return resErr
}

// DelCommentById 根据id删除评论
func DelCommentById(id int64) error {
	videoId := GetVideoIdByCommentId(id)
	resErr := util.Db.Transaction(func(tx *gorm.DB) error {
		if err := tx.Where("id = ?", id).Delete(&Comment{}); err.Error != nil {
			return err.Error
		}
		if err := UpdateVideoCommentCount(videoId, 2, tx); err != nil {
			return err
		}
		return nil
	})
	return resErr
}
