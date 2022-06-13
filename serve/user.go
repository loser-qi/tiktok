package serve

import (
	"errors"
	"gorm.io/gorm"
	"tiktok/util"
)

// HasUser 根据username查询是否存在user
func HasUser(username string) bool {
	var cnt int64
	util.Db.Model(&User{}).Where("username like ?", username).Count(&cnt)
	return cnt != 0
}

// GetUserById 根据id获取user
func GetUserById(id int64) (*User, error) {
	user := &User{}
	util.Db.Where("id = ?", id).Take(user)
	if user.Id == 0 {
		return nil, errors.New("user not exist")
	}
	return user, nil
}

// GetUserByUsernameAndPassword 根据username和password获取user
func GetUserByUsernameAndPassword(username, password string) (*User, error) {
	if !HasUser(username) {
		return nil, errors.New("user not exist")
	}
	user := &User{}
	util.Db.Where("username like ? and password like ?", username, password).Take(user)
	if user.Id == 0 {
		return nil, errors.New("password is not correct")
	}
	return user, nil
}

// SaveUser 存储user到数据库
func SaveUser(username, password string) int64 {
	user := &User{
		Username: username,
		Password: password,
	}
	util.Db.Create(user)
	return user.Id
}

// UpdateFollowCount 更改关注数
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

// UpdateFollowerCount 更改粉丝数
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
