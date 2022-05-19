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

func Publish(c *gin.Context) {
	tokenStr := c.PostForm("token")
	userId, err := util.DecodeToken(tokenStr)
	if err != nil {
		c.JSON(http.StatusOK, Resp{
			StatusCode: 1,
			StatusMsg:  "user doesn't exist",
		})
		return
	}
	title := c.PostForm("title")
	data, err := c.FormFile("data")
	if err != nil {
		c.JSON(http.StatusOK, Resp{
			StatusCode: 2,
			StatusMsg:  err.Error(),
		})
		return
	}
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
	coverFilename := util.MakeCover(playFilename, "./public/")
	serve.SaveVideo(userId, title, playFilename, coverFilename, time.Now())
	c.JSON(http.StatusOK, Resp{
		StatusCode: 0,
		StatusMsg:  playFilename + " uploaded successfully",
	})
}

func PublishList(c *gin.Context) {
	tokenStr := c.Query("token")
	userId, err := util.DecodeToken(tokenStr)
	if err != nil {
		c.JSON(http.StatusOK, Resp{
			StatusCode: 1,
			StatusMsg:  "user doesn't exist",
		})
		return
	}
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
