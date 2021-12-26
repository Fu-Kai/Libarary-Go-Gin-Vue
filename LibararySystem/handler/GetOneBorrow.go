package handler

import (
	"LibararySystem/db"
	"fmt"
	"github.com/gin-gonic/gin"
)

func GetOneBorrow(c *gin.Context) { // gin框架定义的处理器，都是这种格式的函数
	// 查询数据库中，所有学生
	borrow, book, err := db.GetOneBorrow(c.Param("sid"))
	if err != nil {
		fmt.Println(err)
		respError(c, err)
		return
	}
	c.JSON(200, gin.H{
		"code":       0,
		"borrowData": borrow,
		"bookData":   book,
	})
}
