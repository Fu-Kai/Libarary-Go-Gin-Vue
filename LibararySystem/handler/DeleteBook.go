package handler

import (
	"LibararySystem/db"
	"github.com/gin-gonic/gin"
)

func DeleteBook(c *gin.Context) { // gin框架定义的处理器，都是这种格式的函数
	// 查询数据库中，所有学生
	err := db.DeleteBook(c.Param("bid"))
	if err != nil {
		respError(c, err)
		return
	}
	respOK(c, "删除 书本id:"+c.Param("bid"))
}
