package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"tiktok/serve"
	"tiktok/util"
)

func Relation(c *gin.Context) {
	tokenStr := c.Query("token")
	userId, err := util.DecodeToken(tokenStr)
	if err != nil {
		c.JSON(http.StatusOK, Resp{
			StatusCode: 1,
			StatusMsg:  "user doesn't exist",
		})
		return
	}
	toUserIdStr := c.Query("to_user_id")
	actionTypeStr := c.Query("action_type")
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

func FollowList(c *gin.Context) {
	tokenStr := c.Query("token")
	userId, err := util.DecodeToken(tokenStr)
	if err != nil {
		c.JSON(http.StatusOK, Resp{
			StatusCode: 1,
			StatusMsg:  "user doesn't exist",
		})
		return
	}
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

func FollowerList(c *gin.Context) {
	tokenStr := c.Query("token")
	userId, err := util.DecodeToken(tokenStr)
	if err != nil {
		c.JSON(http.StatusOK, Resp{
			StatusCode: 1,
			StatusMsg:  "user doesn't exist",
		})
		return
	}
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
