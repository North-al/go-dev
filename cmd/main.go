package main

import "northal.com/internal/server"

func main() {

	err := server.InitHttp()

	if err != nil {
		panic("http init error")
	}

	// r := gin.Default()
	// r.GET("/ping", func(c *gin.Context) {
	// 	c.JSON(200, gin.H{
	// 		"message": "pong",
	// 		"data":    db,
	// 	})
	// })
	// r.Run() // 监听并在 0.0.0.0:8080 上启动服务
}
