package handler

import (
	"LibararySystem/db"
	"fmt"
	"github.com/gin-gonic/gin"
)

func ModifyBook(c *gin.Context) { // gin框架定义的处理器，都是这种格式的函数
	book := db.Book_fukai{}
	if err := c.ShouldBind(&book); err != nil {
		fmt.Println("nook", book)
		fmt.Println(err)
		respError(c, err)
		return
	}
	fmt.Println("ok", book)
	if err := db.ModifyBook(book, c.Param("bid")); err != nil {
		respError(c, err)
		return
	}

	respOK(c, "已修改")
}
