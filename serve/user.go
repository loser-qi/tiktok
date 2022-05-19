package serve

import (
	"errors"
	"gorm.io/gorm"
	"tiktok/util"
)

func HasUser(username string) bool {
	var cnt int64
	util.Db.Model(&User{}).Where("username=?", username).Count(&cnt)
	return cnt != 0
}

func GetUserById(id int64) (*User, error) {
	user := &User{}
	util.Db.Where("id = ?", id).Take(user)
	if user.Id == 0 {
		return nil, errors.New("user not exist")
	}
	return user, nil
}

func GetUserByUsernameAndPassword(username, password string) (*User, error) {
	if !HasUser(username) {
		return nil, errors.New("user not exist")
	}
	user := &User{}
	util.Db.Where("username = ? and password = ?", username, password).Take(user)
	if user.Id == 0 {
		return nil, errors.New("password is not correct")
	}
	return user, nil
}

func SaveUser(username, password string) int64 {
	user := &User{
		Username: username,
		Password: password,
	}
	util.Db.Create(user)
	return user.Id
}

func UpdateFollowCount(userId int64, actionType int, tx *gorm.DB) error {
	var expr string
	if actionType == 1 {
		expr = "follow_count + 1"
	} else if actionType == 2 {
		expr = "follow_count - 1"
	}
	err := tx.Model(&User{}).Where("id = ?", userId).Update("follow_count", gorm.Expr(expr))
	if err != nil {
		return err.Error
	}
	return nil
}

func UpdateFollowerCount(userId int64, actionType int, tx *gorm.DB) error {
	var expr string
	if actionType == 1 {
		expr = "follower_count + 1"
	} else if actionType == 2 {
		expr = "follower_count - 1"
	}
	err := tx.Model(&User{}).Where("id = ?", userId).Update("follower_count", gorm.Expr(expr))
	if err != nil {
		return err.Error
	}
	return nil
}
