package controller

import (
	"tiktok/serve"
	"tiktok/util"
)

type Resp struct {
	StatusCode int32  `json:"status_code"`
	StatusMsg  string `json:"status_msg,omitempty"`
}

type UserResp struct {
	Id            int64  `json:"id,omitempty"`
	Username      string `json:"name,omitempty"`
	FollowCount   int64  `json:"follow_count,omitempty"`
	FollowerCount int64  `json:"follower_count,omitempty"`
	IsFollow      bool   `json:"is_follow,omitempty"`
}

type VideoResp struct {
	Id            int64    `json:"id,omitempty"`
	Author        UserResp `json:"author,omitempty"`
	PlayUrl       string   `json:"play_url,omitempty"`
	CoverUrl      string   `json:"cover_url,omitempty"`
	FavoriteCount int64    `json:"favorite_count,omitempty"`
	CommentCount  int64    `json:"comment_count,omitempty"`
	IsFavorite    bool     `json:"is_favorite,omitempty"`
	Title         string   `json:"title,omitempty"`
}

type CommentResp struct {
	Id         int64    `json:"id,omitempty"`
	User       UserResp `json:"user,omitempty"`
	Content    string   `json:"content,omitempty"`
	CreateDate string   `json:"create_date,omitempty"`
}

// Feed

type FeedResp struct {
	Resp
	VideoList []VideoResp `json:"video_list,omitempty"`
}

// User

type UserRegisterResp struct {
	Resp
	UserId int64  `json:"user_id,omitempty"`
	Token  string `json:"token,omitempty"`
}

type UserLoginResp struct {
	Resp
	UserId int64  `json:"user_id,omitempty"`
	Token  string `json:"token,omitempty"`
}

type UserInfoResp struct {
	Resp
	UserResp `json:"user,omitempty"`
}

// Publish

type PublishListResp struct {
	Resp
	VideoList []VideoResp `json:"video_list,omitempty"`
}

// Favorite

type FavoriteListResp struct {
	Resp
	VideoList []VideoResp `json:"video_list,omitempty"`
}

// Comment

type CommentListResp struct {
	Resp
	CommentList []CommentResp `json:"comment_list,omitempty"`
}

// Follow

type UserListResp struct {
	Resp
	UserList []UserResp `json:"user_list,omitempty"`
}

var (
	staticPrefix = "http://" + util.GetIp() + ":8080/static/"
)

func UserConv(user *serve.User, isFollow bool) *UserResp {
	userResp := &UserResp{
		Id:            user.Id,
		Username:      user.Username,
		FollowCount:   user.FollowCount,
		FollowerCount: user.FollowerCount,
		IsFollow:      isFollow,
	}
	return userResp
}

func VideoConv(video *serve.Video, isFavorite, isFollow bool) *VideoResp {
	user, _ := serve.GetUserById(video.UserId)
	videoResp := &VideoResp{
		Id:            video.Id,
		Author:        *UserConv(user, isFollow),
		PlayUrl:       staticPrefix + video.PlayPath,
		CoverUrl:      staticPrefix + video.CoverPath,
		FavoriteCount: video.FavoriteCount,
		CommentCount:  video.CommentCount,
		IsFavorite:    isFavorite,
		Title:         video.Title,
	}
	return videoResp
}

func CommentConv(comment *serve.Comment, isFollow bool) *CommentResp {
	user, _ := serve.GetUserById(comment.UserId)
	commentResp := &CommentResp{
		Id:         comment.Id,
		User:       *UserConv(user, isFollow),
		Content:    comment.Text,
		CreateDate: comment.CreateTime.Format("01-02"),
	}
	return commentResp
}
