package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	r := gin.Default()

	r.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "值：%v", "你好,Golang")
	})

	r.GET("/news", func(c *gin.Context) {
		c.String(200, "我是新闻页面")
	})

	r.POST("/add", func(c *gin.Context) {
		c.String(200, "添加数据")
	})

	r.PUT("/edit", func(c *gin.Context) {
		c.String(200, "编辑数据")
	})

	r.PUT("/delete", func(c *gin.Context) {
		c.String(200, "删除数据")
	})

	r.GET("/json", func(c *gin.Context) {
		c.JSON(200, map[string]interface{}{
			"success": true,
			"msg":     "你好Golang",
		})
	})

	r.Run()
}
