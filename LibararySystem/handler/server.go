package handler

import (
	"github.com/gin-gonic/gin" // 导入web服务框架
	"net/http"
)

func Start() (err error) {
	// 使用gin框架提供的默认web服务引擎
	r := gin.Default()
	// 放行所有跨域请求
	r.Use(Cors())
	// api接口服务，定义了公开组 /
	public := r.Group("")
	{
		// 定义增改查的接口，并注册到web服务器
		public.POST("/login", Login)
		public.GET("/borrows", GetAllBorrows)
		public.GET("/borrow/:sid", GetOneBorrow)
		public.POST("/delete/borrow/:BRId", DeleteBorrow)
	}

	// api接口服务，定义了路由组 /student
	student := r.Group("student")
	{
		// 定义增改查的接口，并注册到web服务器
		student.GET("", GetAllStudents)
		student.GET("/:sid", GetOneStudents)
		student.POST("/add", AddStudent)
		student.POST("/modify/:sid", ModifyStudent)
		student.POST("/delete/:sid", DeleteStudent)
		student.POST("/borrow", BorrowBook)
		student.POST("/return/:BRId/:sid/:bid", ReturnBook)
	}

	// api接口服务，定义了路由组 /book
	book := r.Group("book")
	{
		// 定义增改查的接口，并注册到web服务器
		book.GET("", GetAllBooks)
		book.GET("/:bid", GetOneBook)
		book.POST("/add", AddBook)
		book.POST("/modify/:bid", ModifyBook)
		book.POST("/delete/:bid", DeleteBook)
	}

	// 启动web服务
	err = r.Run(":7070")
	return err
}

// Cors 跨域进行请求放行
func Cors() gin.HandlerFunc {
	return func(c *gin.Context) {
		method := c.Request.Method

		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Headers", "Content-Type,AccessToken,X-CSRF-Token, Authorization, Token")
		c.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS")
		c.Header("Access-Control-Expose-Headers", "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers, Content-Type")
		c.Header("Access-Control-Allow-Credentials", "true")

		//放行所有OPTIONS方法
		if method == "OPTIONS" {
			c.AbortWithStatus(http.StatusNoContent)
		}
		// 处理请求
		c.Next()
	}
}

// 封装函数，用于向客户端返回正确信息
func respOK(c *gin.Context, data interface{}) {
	c.JSON(200, gin.H{
		"code": 0,
		"data": data,
	})
}

// 封装函数，用于向客户端返回错误消息
func respError(c *gin.Context, msg interface{}) {
	c.JSON(200, gin.H{
		"code":    1,
		"message": msg,
	})
}
