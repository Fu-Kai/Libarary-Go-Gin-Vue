package handler

import (
	"LibararySystem/db"
	"fmt"
	"github.com/gin-gonic/gin"
)

func BorrowBook(c *gin.Context) { // gin框架定义的处理器，都是这种格式的函数
	borrowTab := db.BorrowTab_fukai{}
	if err := c.ShouldBind(&borrowTab); err != nil {
		fmt.Println(err)
		respError(c, err)
		return
	}
	fmt.Println(borrowTab)

	// 在数据库中插入一
	if err := db.BorrowBook(borrowTab); err != nil {
		fmt.Println(err)
		respError(c, err)
		return
	}

	respOK(c, "已借书")
}
