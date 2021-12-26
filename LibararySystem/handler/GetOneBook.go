package handler

import (
	"LibararySystem/db"
	"github.com/gin-gonic/gin"
)

func GetOneBook(c *gin.Context) { // gin框架定义的处理器，都是这种格式的函数
	// 查询数据库中，所有学生
	book, err := db.GetOneBook(c.Param("bid"))
	if err != nil {
		respError(c, err)
		return
	}
	respOK(c, book)
}
