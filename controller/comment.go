package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"tiktok/serve"
	"tiktok/util"
	"time"
)

// Comment 评论操作的控制层实现
func Comment(c *gin.Context) {
	// 获取参数
	tokenStr := c.Query("token")
	videoIdStr := c.Query("video_id")
	actionTypeStr := c.Query("action_type")
	commentText := c.DefaultQuery("comment_text", "")
	commentIdStr := c.DefaultQuery("comment_id", "")

	// 解析token
	userId, err := util.DecodeToken(tokenStr)
	if err != nil {
		c.JSON(http.StatusOK, Resp{
			StatusCode: 1,
			StatusMsg:  err.Error(),
		})
		return
	}

	// 根据actionType进行相应评论操作
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

// CommentList 视频评论列表
func CommentList(c *gin.Context) {
	// 获取参数
	tokenStr := c.Query("token")
	videoIdStr := c.Query("video_id")

	// 解析token
	userId, err := util.DecodeToken(tokenStr)

	// 获取评论列表
	videoId, _ := strconv.ParseInt(videoIdStr, 10, 64)
	commentList := serve.GetCommentListByVideoId(videoId)

	// 根据登录状态添加关注关系
	commentResp := make([]CommentResp, len(commentList))
	if err != nil {
		for i, comment := range commentList {
			commentResp[i] = *CommentConv(&comment, false)
		}
	} else {
		// 如果已登录，则需要查询评论user和当前user的关注关系
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
