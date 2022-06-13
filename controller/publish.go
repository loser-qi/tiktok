package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"path/filepath"
	"tiktok/serve"
	"tiktok/util"
	"time"
)

// Publish 视频投稿
func Publish(c *gin.Context) {
	// 获取参数
	tokenStr := c.PostForm("token")
	title := c.PostForm("title")
	data, err := c.FormFile("data")
	if err != nil {
		c.JSON(http.StatusOK, Resp{
			StatusCode: 1,
			StatusMsg:  err.Error(),
		})
		return
	}

	// 解析token
	userId, err := util.DecodeToken(tokenStr)
	if err != nil {
		c.JSON(http.StatusOK, Resp{
			StatusCode: 2,
			StatusMsg:  "user doesn't exist",
		})
		return
	}

	// 存储上传的视频至public文件夹
	playFilename := filepath.Base(data.Filename)
	playFilename = fmt.Sprintf("%d_%s_%s", userId, time.Now().Format("20060102150405"), playFilename)
	playFilepath := filepath.Join("./public/", playFilename)
	if err = c.SaveUploadedFile(data, playFilepath); err != nil {
		c.JSON(http.StatusOK, Resp{
			StatusCode: 3,
			StatusMsg:  err.Error(),
		})
		return
	}

	// 制作封面图，获取封面图文件名
	coverFilename := util.MakeCover(playFilename, "./public/")

	// 将视频相关信息存入数据库
	serve.SaveVideo(userId, title, playFilename, coverFilename, time.Now())
	c.JSON(http.StatusOK, Resp{
		StatusCode: 0,
		StatusMsg:  playFilename + " uploaded successfully",
	})
}

// PublishList 发布列表
func PublishList(c *gin.Context) {
	// 获取参数
	tokenStr := c.Query("token")

	// 解析token
	userId, err := util.DecodeToken(tokenStr)
	if err != nil {
		c.JSON(http.StatusOK, Resp{
			StatusCode: 1,
			StatusMsg:  "user doesn't exist",
		})
		return
	}

	// 根据userId获取视频列表
	videoList := serve.GetVideoListByUserId(userId)
	videoResp := make([]VideoResp, len(videoList))
	for i, video := range videoList {
		videoResp[i] = *VideoConv(&video, serve.GetFavorite(userId, video.Id), true)
	}
	c.JSON(http.StatusOK, PublishListResp{
		Resp: Resp{
			StatusCode: 0,
		},
		VideoList: videoResp,
	})
}
