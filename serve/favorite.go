package serve

import (
	"gorm.io/gorm"
	"tiktok/util"
)

// HasFavorite 数据库中是否存储了user对video的点赞关系
func HasFavorite(userId, videoId int64) bool {
	var cnt int64
	util.Db.Table("is_favorite").Where("user_id = ? and video_id = ?", userId, videoId).Count(&cnt)
	return cnt != 0
}

// GetFavorite 获取user对video的点赞状态
func GetFavorite(userId, videoId int64) bool {
	if !HasFavorite(userId, videoId) {
		return false
	}
	var actionType int
	util.Db.Table("is_favorite").Select("action_type").Where("user_id = ? and video_id = ?", userId, videoId).Take(&actionType)
	return actionType == 1
}

// GetFavoriteVideoIdList 获取user点赞的所有video
func GetFavoriteVideoIdList(userId int64) []int64 {
	var videoIdList []int64
	util.Db.Table("is_favorite").Select("video_id").Where("user_id = ? and action_type = 1", userId).Find(&videoIdList)
	return videoIdList
}

// SaveFavorite 保存点赞关系
func SaveFavorite(userId, videoId int64, actionType int) error {
	var resErr error = nil
	if !HasFavorite(userId, videoId) {
		resErr = util.Db.Transaction(func(tx *gorm.DB) error {
			err1 := tx.Table("is_favorite").Create(map[string]interface{}{
				"user_id":     userId,
				"video_id":    videoId,
				"action_type": actionType,
			})
			if err1.Error != nil {
				return err1.Error
			}
			if err2 := UpdateVideoFavoriteCount(videoId, actionType, tx); err2 != nil {
				return err2
			}
			return nil
		})
	} else {
		resErr = util.Db.Transaction(func(tx *gorm.DB) error {
			if err := tx.Table("is_favorite").Where("user_id = ? and video_id = ?", userId, videoId).Update("action_type", actionType); err.Error != nil {
				return err.Error
			}
			if err := UpdateVideoFavoriteCount(videoId, actionType, tx); err != nil {
				return err
			}
			return nil
		})
	}
	return resErr
}
