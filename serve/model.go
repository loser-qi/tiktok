package serve

import "time"

// User 用户模型
type User struct {
	Id            int64   `gorm:"column:id"`
	Username      string  `gorm:"column:username"`
	Password      string  `gorm:"column:password"`
	FollowCount   int64   `gorm:"column:follow_count"`
	FollowerCount int64   `gorm:"column:follower_count"`
	VideoList     []Video `gorm:"foreignKey:user_id;references:id;"`
}

func (User) TableName() string {
	return "user"
}

// Video 视频模型
type Video struct {
	Id            int64     `gorm:"column:id"`
	UserId        int64     `gorm:"column:user_id"`
	Title         string    `gorm:"column:title"`
	PlayPath      string    `gorm:"column:play_path"`
	CoverPath     string    `gorm:"column:cover_path"`
	FavoriteCount int64     `gorm:"column:favorite_count"`
	CommentCount  int64     `gorm:"column:comment_count"`
	CreateTime    time.Time `gorm:"column:create_time"`
	CommentList   []Comment `gorm:"foreignKey:video_id;references:id;"`
}

func (Video) TableName() string {
	return "video"
}

// Comment 评论模型
type Comment struct {
	Id         int64     `gorm:"column:id"`
	UserId     int64     `gorm:"column:user_id"`
	VideoId    int64     `gorm:"column:video_id"`
	Text       string    `gorm:"column:text"`
	CreateTime time.Time `gorm:"column:create_time"`
}

func (Comment) TableName() string {
	return "comment"
}
