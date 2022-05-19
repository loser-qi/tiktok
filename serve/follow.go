package serve

import (
	"gorm.io/gorm"
	"tiktok/util"
)

func HasFollow(followerUserId, followedUserId int64) bool {
	var cnt int64
	util.Db.Table("is_follow").Where("follower_user_id = ? and followed_user_id = ?", followerUserId, followedUserId).Count(&cnt)
	return cnt != 0
}

func GetFollow(followerUserId, followedUserId int64) bool {
	if followedUserId == followerUserId {
		return true
	}
	if !HasFollow(followerUserId, followedUserId) {
		return false
	}
	var actionType int
	util.Db.Table("is_follow").Select("action_type").Where("follower_user_id = ? and followed_user_id = ?", followerUserId, followedUserId).Take(&actionType)
	return actionType == 1
}

func GetFollowList(userId int64) []User {
	var userIdList []int64
	var userList []User
	util.Db.Table("is_follow").Select("followed_user_id").Where("follower_user_id = ? and action_type = 1", userId).Find(&userIdList)
	userList = make([]User, len(userIdList))
	for i, followUserId := range userIdList {
		user, _ := GetUserById(followUserId)
		userList[i] = *user
	}
	return userList
}

func GetFollowerList(userId int64) []User {
	var userIdList []int64
	var userList []User
	util.Db.Table("is_follow").Select("follower_user_id").Where("followed_user_id = ? and action_type = 1", userId).Find(&userIdList)
	userList = make([]User, len(userIdList))
	for i, followUserId := range userIdList {
		user, _ := GetUserById(followUserId)
		userList[i] = *user
	}
	return userList
}

func SaveFollow(followerUserId, followedUserId int64, actionType int) error {
	var resErr error = nil
	if !HasFollow(followerUserId, followedUserId) {
		resErr = util.Db.Transaction(func(tx *gorm.DB) error {
			err1 := tx.Table("is_follow").Create(map[string]interface{}{
				"follower_user_id": followerUserId,
				"followed_user_id": followedUserId,
				"action_type":      actionType,
			})
			if err1.Error != nil {
				return err1.Error
			}
			if err2 := UpdateFollowCount(followerUserId, actionType, tx); err2 != nil {
				return err2
			}
			if err3 := UpdateFollowerCount(followedUserId, actionType, tx); err3 != nil {
				return err3
			}
			return nil
		})
	} else {
		resErr = util.Db.Transaction(func(tx *gorm.DB) error {
			if err := tx.Table("is_follow").Where("follower_user_id = ? and followed_user_id = ?", followerUserId, followedUserId).Update("action_type", actionType); err.Error != nil {
				return err.Error
			}
			if err := UpdateFollowCount(followerUserId, actionType, tx); err != nil {
				return err
			}
			if err := UpdateFollowerCount(followedUserId, actionType, tx); err != nil {
				return err
			}
			return nil
		})
	}
	return resErr
}
