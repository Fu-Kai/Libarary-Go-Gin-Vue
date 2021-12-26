package handler

import (
	"LibararySystem/db"
	"github.com/gin-gonic/gin"
)

func ReturnBook(c *gin.Context) { // gin框架定义的处理器，都是这种格式的函数

	// 在数据库中插入一
	if err := db.ReturnBook(c.Param("BRId"), c.Param("sid"), c.Param("bid")); err != nil {
		respError(c, err)
		return
	}

	respOK(c, "已还书")
}
