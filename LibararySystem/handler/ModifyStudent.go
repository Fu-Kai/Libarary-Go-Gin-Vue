package handler

import (
	"LibararySystem/db"
	"fmt"
	"github.com/gin-gonic/gin"
)

func ModifyStudent(c *gin.Context) { // gin框架定义的处理器，都是这种格式的函数
	student := db.Student_fukai{}
	if err := c.ShouldBind(&student); err != nil {
		fmt.Println(student)
		respError(c, err)
		return
	}
	fmt.Println(student)
	if err := db.ModifyStudent(student, c.Param("sid")); err != nil {
		respError(c, err)
		return
	}

	respOK(c, "已修改")
}
