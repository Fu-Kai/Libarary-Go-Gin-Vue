package handler

import (
	"LibararySystem/db"
	"fmt"
	"github.com/gin-gonic/gin"
)

func AddStudent(c *gin.Context) { // gin框架定义的处理器，都是这种格式的函数
	student := db.Student_fukai{}
	if err := c.ShouldBind(&student); err != nil {
		respError(c, err)
		return
	}
	fmt.Println(student)

	// 在数据库中插入一
	if err := db.AddStudent(student); err != nil {
		respError(c, err)
		return
	}

	respOK(c, "已插入"+student.StuName)
}
