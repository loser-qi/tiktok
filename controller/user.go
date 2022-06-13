package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"tiktok/serve"
	"tiktok/util"
)

// UserRegister 用户注册
func UserRegister(c *gin.Context) {
	// 获取参数
	username := c.Query("username")
	password := c.Query("password")

	// 检查是否合法
	if ok := util.CheckUsername(username) && util.CheckPassword(password); !ok {
		c.JSON(http.StatusOK, Resp{
			StatusCode: 1,
			StatusMsg:  "incorrect format",
		})
		return
	}

	// 检查用户是否已存在
	if serve.HasUser(username) {
		c.JSON(http.StatusOK, Resp{
			StatusCode: 2,
			StatusMsg:  "user already exist",
		})
		return
	}

	// 存储用户信息到数据库
	userId := serve.SaveUser(username, password)

	// 生成token
	tokenStr := util.EncodeToken(userId)
	c.JSON(http.StatusOK, UserRegisterResp{
		Resp: Resp{
			StatusCode: 0,
		},
		UserId: userId,
		Token:  tokenStr,
	})
}

// UserLogin 用户登录
func UserLogin(c *gin.Context) {
	// 获取参数
	username := c.Query("username")
	password := c.Query("password")

	// 检查是否合法
	if ok := util.CheckUsername(username) && util.CheckPassword(password); !ok {
		c.JSON(http.StatusOK, Resp{
			StatusCode: 1,
			StatusMsg:  "incorrect format",
		})
		return
	}

	// 从数据库查找用户名和密码与输入一致的user
	user, err := serve.GetUserByUsernameAndPassword(username, password)
	if err != nil {
		c.JSON(http.StatusOK, Resp{
			StatusCode: 2,
			StatusMsg:  err.Error(),
		})
		return
	}

	// 生成token
	tokenStr := util.EncodeToken(user.Id)
	c.JSON(http.StatusOK, UserLoginResp{
		Resp: Resp{
			StatusCode: 0,
		},
		UserId: user.Id,
		Token:  tokenStr,
	})
}

// UserInfo 用户信息
func UserInfo(c *gin.Context) {
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

	// 根据userId获取user
	user, _ := serve.GetUserById(userId)
	c.JSON(http.StatusOK, UserInfoResp{
		Resp: Resp{
			StatusCode: 0,
		},
		UserResp: *UserConv(user, true),
	})
}
