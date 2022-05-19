package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"tiktok/serve"
	"tiktok/util"
)

func Favorite(c *gin.Context) {
	tokenStr := c.Query("token")
	videoIdStr := c.Query("video_id")
	actionTypeStr := c.Query("action_type")
	userId, err := util.DecodeToken(tokenStr)
	if err != nil {
		c.JSON(http.StatusOK, Resp{
			StatusCode: 1,
			StatusMsg:  err.Error(),
		})
		return
	}
	videoId, _ := strconv.ParseInt(videoIdStr, 10, 64)
	actionType, _ := strconv.Atoi(actionTypeStr)
	err = serve.SaveFavorite(userId, videoId, actionType)
	if err != nil {
		c.JSON(http.StatusOK, Resp{
			StatusCode: 2,
			StatusMsg:  err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, Resp{
		StatusCode: 0,
	})
}

func FavoriteList(c *gin.Context) {
	tokenStr := c.Query("token")
	userId, err := util.DecodeToken(tokenStr)
	if err != nil {
		c.JSON(http.StatusOK, Resp{
			StatusCode: 1,
			StatusMsg:  err.Error(),
		})
		return
	}
	videoIdList := serve.GetVideoIdList(userId)
	videoResp := make([]VideoResp, len(videoIdList))
	for i, videoId := range videoIdList {
		video, _ := serve.GetVideoById(videoId)
		videoResp[i] = *VideoConv(video, true, serve.GetFollow(userId, video.UserId))
	}
	c.JSON(http.StatusOK, FavoriteListResp{
		Resp: Resp{
			StatusCode: 0,
		},
		VideoList: videoResp,
	})
}
