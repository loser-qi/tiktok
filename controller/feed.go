package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"tiktok/serve"
	"tiktok/util"
)

func Feed(c *gin.Context) {
	tokenStr := c.Query("token")
	userId, err := util.DecodeToken(tokenStr)
	videoList := serve.GetVideoListByLimit(30)
	videoResp := make([]VideoResp, len(videoList))
	if err != nil {
		for i, video := range videoList {
			videoResp[i] = *VideoConv(&video, false, false)
		}
	} else {
		for i, video := range videoList {
			videoResp[i] = *VideoConv(&video, serve.GetFavorite(userId, video.Id), serve.GetFollow(userId, video.UserId))
		}
	}
	c.JSON(http.StatusOK, FeedResp{
		Resp: Resp{
			StatusCode: 0,
		},
		VideoList: videoResp,
	})
}
