package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"tiktok/serve"
	"tiktok/util"
	"time"
)

func Comment(c *gin.Context) {
	tokenStr := c.Query("token")
	videoIdStr := c.Query("video_id")
	actionTypeStr := c.Query("action_type")
	commentText := c.DefaultQuery("comment_text", "")
	commentIdStr := c.DefaultQuery("comment_id", "")
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
	if actionType == 1 {
		err = serve.SaveComment(userId, videoId, commentText, time.Now())
	} else if actionType == 2 {
		commentId, _ := strconv.ParseInt(commentIdStr, 10, 64)
		err = serve.DelCommentById(commentId)
	}
	if err != nil {
		c.JSON(http.StatusOK, Resp{
			StatusCode: 2,
			StatusMsg:  err.Error(),
		})
	}
	c.JSON(http.StatusOK, Resp{
		StatusCode: 0,
	})
}

func CommentList(c *gin.Context) {
	tokenStr := c.Query("token")
	videoIdStr := c.Query("video_id")
	userId, err := util.DecodeToken(tokenStr)
	videoId, _ := strconv.ParseInt(videoIdStr, 10, 64)
	commentList := serve.GetCommentListByVideoId(videoId)
	commentResp := make([]CommentResp, len(commentList))
	if err != nil {
		for i, comment := range commentList {
			commentResp[i] = *CommentConv(&comment, false)
		}
	} else {
		for i, comment := range commentList {
			commentResp[i] = *CommentConv(&comment, serve.GetFollow(userId, comment.UserId))
		}
	}
	c.JSON(http.StatusOK, CommentListResp{
		Resp: Resp{
			StatusCode: 0,
		},
		CommentList: commentResp,
	})
}
