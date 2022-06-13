package serve

import (
	"errors"
	"gorm.io/gorm"
	"tiktok/util"
	"time"
)

// GetVideoById 根据id获取video
func GetVideoById(id int64) (*Video, error) {
	video := &Video{}
	util.Db.Where("id = ?", id).Take(video)
	if video.Id == 0 {
		return nil, errors.New("video not exist")
	}
	return video, nil
}

// GetVideoListByUserId 根据user获取video列表
func GetVideoListByUserId(userId int64) []Video {
	user := &User{
		Id: userId,
	}
	util.Db.Preload("VideoList").Find(user)
	return user.VideoList
}

// GetVideoListByLimit 按时间倒序获取limit个视频
func GetVideoListByLimit(limit int) []Video {
	var videoList []Video
	util.Db.Order("create_time desc").Limit(limit).Find(&videoList)
	return videoList
}

// SaveVideo 存储视频
func SaveVideo(userId int64, title, playPath, coverPath string, createTime time.Time) {
	video := &Video{
		UserId:     userId,
		Title:      title,
		PlayPath:   playPath,
		CoverPath:  coverPath,
		CreateTime: createTime,
	}
	util.Db.Create(video)
}

// UpdateVideoFavoriteCount 更改视频点赞数
func UpdateVideoFavoriteCount(videoId int64, actionType int, tx *gorm.DB) error {
	var expr string
	if actionType == 1 {
		expr = "favorite_count + 1"
	} else if actionType == 2 {
		expr = "favorite_count - 1"
	}
	err := tx.Model(&Video{}).Where("id = ?", videoId).Update("favorite_count", gorm.Expr(expr))
	if err != nil {
		return err.Error
	}
	return nil
}

// UpdateVideoCommentCount 更改视频评论数
func UpdateVideoCommentCount(videoId int64, actionType int, tx *gorm.DB) error {
	var expr string
	if actionType == 1 {
		expr = "comment_count + 1"
	} else if actionType == 2 {
		expr = "comment_count - 1"
	}
	err := tx.Model(&Video{}).Where("id = ?", videoId).Update("comment_count", gorm.Expr(expr))
	if err != nil {
		return err.Error
	}
	return nil
}
