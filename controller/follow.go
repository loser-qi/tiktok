package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"tiktok/serve"
	"tiktok/util"
)

// Relation 关系操作
func Relation(c *gin.Context) {
	// 获取参数
	tokenStr := c.Query("token")
	toUserIdStr := c.Query("to_user_id")
	actionTypeStr := c.Query("action_type")

	// 解析token
	userId, err := util.DecodeToken(tokenStr)
	if err != nil {
		c.JSON(http.StatusOK, Resp{
			StatusCode: 1,
			StatusMsg:  "user doesn't exist",
		})
		return
	}

	// 存储关注关系
	toUserId, _ := strconv.ParseInt(toUserIdStr, 10, 64)
	actionType, _ := strconv.Atoi(actionTypeStr)
	err = serve.SaveFollow(userId, toUserId, actionType)
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

// FollowList 用户关注列表
func FollowList(c *gin.Context) {
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

	// 获取所有关注
	userList := serve.GetFollowList(userId)
	userResp := make([]UserResp, len(userList))
	for i, user := range userList {
		userResp[i] = *UserConv(&user, true)
	}
	c.JSON(http.StatusOK, UserListResp{
		Resp: Resp{
			StatusCode: 0,
		},
		UserList: userResp,
	})
}

// FollowerList 用户粉丝列表
func FollowerList(c *gin.Context) {
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

	// 获取所有粉丝
	userList := serve.GetFollowerList(userId)
	userResp := make([]UserResp, len(userList))
	for i, user := range userList {
		userResp[i] = *UserConv(&user, true)
	}
	c.JSON(http.StatusOK, UserListResp{
		Resp: Resp{
			StatusCode: 0,
		},
		UserList: userResp,
	})
}
