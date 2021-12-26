package handler

import (
	"LibararySystem/db"
	"fmt"
	"github.com/gin-gonic/gin"
)

func AddBook(c *gin.Context) { // gin框架定义的处理器，都是这种格式的函数
	book := db.Book_fukai{}
	if err := c.ShouldBind(&book); err != nil {
		respError(c, err)
		return
	}
	fmt.Println(book)

	// 在数据库中插入一
	if err := db.AddBook(book); err != nil {
		respError(c, err)
		return
	}

	respOK(c, "已插入"+book.BookName)
}
