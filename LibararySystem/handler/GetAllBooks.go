package handler

import (
	"LibararySystem/db"
	"github.com/gin-gonic/gin"
)

func GetAllBooks(c *gin.Context) { // gin框架定义的处理器，都是这种格式的函数
	// 查询数据库中，所有学生
	books, err := db.GetAllBooks()
	if err != nil {
		respError(c, err)
		return
	}
	respOK(c, books)
}
