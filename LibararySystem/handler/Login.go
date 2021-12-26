package handler

import (
	"LibararySystem/db"
	"github.com/gin-gonic/gin"
)

func Login(c *gin.Context) { // gin框架定义的处理器，都是这种格式的函数
	loginForm := db.LoginForm_fukai{}
	if err := c.ShouldBind(&loginForm); err != nil {
		respError(c, err)
		return
	}

	// 在数据库查询
	exit, err := db.Login(loginForm)
	if err != nil {
		respError(c, err)
		return
	}
	if exit != 0 {
		c.JSON(200, gin.H{
			"code":  0,
			"data":  "已登录" + loginForm.Acc,
			"StuId": exit,
		})
	}
}
