package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"tiktok/serve"
	"tiktok/util"
)

func UserRegister(c *gin.Context) {
	username := c.Query("username")
	password := c.Query("password")
	if ok := util.CheckUsername(username) && util.CheckPassword(password); !ok {
		c.JSON(http.StatusOK, Resp{
			StatusCode: 1,
			StatusMsg:  "incorrect format",
		})
		return
	}
	if serve.HasUser(username) {
		c.JSON(http.StatusOK, Resp{
			StatusCode: 2,
			StatusMsg:  "user already exist",
		})
		return
	}
	userId := serve.SaveUser(username, password)
	tokenStr := util.EncodeToken(userId)
	c.JSON(http.StatusOK, UserRegisterResp{
		Resp: Resp{
			StatusCode: 0,
		},
		UserId: userId,
		Token:  tokenStr,
	})
}

func UserLogin(c *gin.Context) {
	username := c.Query("username")
	password := c.Query("password")
	if ok := util.CheckUsername(username) && util.CheckPassword(password); !ok {
		c.JSON(http.StatusOK, Resp{
			StatusCode: 1,
			StatusMsg:  "incorrect format",
		})
		return
	}
	user, err := serve.GetUserByUsernameAndPassword(username, password)
	if err != nil {
		c.JSON(http.StatusOK, Resp{
			StatusCode: 2,
			StatusMsg:  err.Error(),
		})
		return
	}
	tokenStr := util.EncodeToken(user.Id)
	c.JSON(http.StatusOK, UserLoginResp{
		Resp: Resp{
			StatusCode: 0,
		},
		UserId: user.Id,
		Token:  tokenStr,
	})
}

func UserInfo(c *gin.Context) {
	tokenStr := c.Query("token")
	userId, err := util.DecodeToken(tokenStr)
	if err != nil {
		c.JSON(http.StatusOK, Resp{
			StatusCode: 1,
			StatusMsg:  "user doesn't exist",
		})
		return
	}
	user, _ := serve.GetUserById(userId)
	c.JSON(http.StatusOK, UserInfoResp{
		Resp: Resp{
			StatusCode: 0,
		},
		UserResp: *UserConv(user, true),
	})
}
