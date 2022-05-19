package serve

import (
	"errors"
	"gorm.io/gorm"
	"tiktok/util"
	"time"
)

func GetVideoById(id int64) (*Video, error) {
	video := &Video{}
	util.Db.Where("id = ?", id).Take(video)
	if video.Id == 0 {
		return nil, errors.New("video not exist")
	}
	return video, nil
}

func GetVideoListByUserId(userId int64) []Video {
	user := &User{
		Id: userId,
	}
	util.Db.Preload("VideoList").Find(user)
	return user.VideoList
}

func GetVideoListByLimit(limit int) []Video {
	var videoList []Video
	util.Db.Order("create_time desc").Limit(limit).Find(&videoList)
	return videoList
}

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
