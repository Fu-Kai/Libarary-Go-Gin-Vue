package handler

import (
	"LibararySystem/db"
	"github.com/gin-gonic/gin"
)

func DeleteStudent(c *gin.Context) { // gin框架定义的处理器，都是这种格式的函数
	// 查询数据库中，所有学生
	err := db.DeleteStudent(c.Param("sid"))
	if err != nil {
		respError(c, err)
		return
	}
	respOK(c, "删除 学生id:"+c.Param("sid"))
}
