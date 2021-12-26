package handler

import (
	"LibararySystem/db"
	"github.com/gin-gonic/gin"
)

func DeleteBorrow(c *gin.Context) { // gin框架定义的处理器，都是这种格式的函数
	// 查询数据库中，所有学生
	err := db.DeleteBorrow(c.Param("BRId"))
	if err != nil {
		respError(c, err)
		return
	}
	respOK(c, "删除 BRid:"+c.Param("BRId"))
}
