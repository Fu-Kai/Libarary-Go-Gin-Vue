package handler

import (
	"LibararySystem/db"
	"github.com/gin-gonic/gin"
)

func GetAllBorrows(c *gin.Context) { // gin框架定义的处理器，都是这种格式的函数
	// 查询数据库中，所有学生
	borrows, err := db.GetAllBorrows()
	if err != nil {
		respError(c, err)
		return
	}
	respOK(c, borrows)
}
