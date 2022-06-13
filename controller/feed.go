package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"tiktok/serve"
	"tiktok/util"
)

// Feed 视频流
func Feed(c *gin.Context) {
	// 获取参数
	tokenStr := c.Query("token")

	//解析token
	userId, err := util.DecodeToken(tokenStr)

	// 按时间倒序获取视频(设置限制为30)
	// 根据登录状态获取点赞关系和关注关系
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
