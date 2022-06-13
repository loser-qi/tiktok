package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"tiktok/serve"
	"tiktok/util"
)

// Favorite 赞操作
func Favorite(c *gin.Context) {
	// 获取参数
	tokenStr := c.Query("token")
	videoIdStr := c.Query("video_id")
	actionTypeStr := c.Query("action_type")

	// 解析token
	userId, err := util.DecodeToken(tokenStr)
	if err != nil {
		c.JSON(http.StatusOK, Resp{
			StatusCode: 1,
			StatusMsg:  err.Error(),
		})
		return
	}

	// 存储点赞关系
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

// FavoriteList 点赞列表
func FavoriteList(c *gin.Context) {
	// 获取参数
	tokenStr := c.Query("token")

	// 解析token
	userId, err := util.DecodeToken(tokenStr)
	if err != nil {
		c.JSON(http.StatusOK, Resp{
			StatusCode: 1,
			StatusMsg:  err.Error(),
		})
		return
	}

	// 获取所有点赞的videoId
	videoIdList := serve.GetFavoriteVideoIdList(userId)

	// 根据videoId获取video
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
