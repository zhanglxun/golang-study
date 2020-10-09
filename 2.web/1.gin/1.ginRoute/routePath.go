package main

import "github.com/gin-gonic/gin"

func main() {
	r := gin.Default()
	r.GET("/user/search/:username/:address", func(c *gin.Context) {
		username := c.Param("username")
		address := c.Param("address")
		//输出json结果给调用方
		c.JSON(200, gin.H{
			"message":  "pong",
			"username": username,
			"address":  address,
		})
	})

	r.Run(":8080")

	/*
		http://localhost:8080/user/search/jarvis/shenz
	*/
}
