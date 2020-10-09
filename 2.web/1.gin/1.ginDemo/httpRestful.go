package main

import "github.com/gin-gonic/gin"

func main() {
	r := gin.Default()
	r.GET("/user/info", func(context *gin.Context) {
		context.JSON(200, gin.H{
			"message": "get user info success.",
		})
	})
	r.POST("/user/info", func(context *gin.Context) {
		context.JSON(200, gin.H{
			"message": "create user info success",
		})
	})
	r.PUT("/user/info", func(context *gin.Context) {
		context.JSON(200, gin.H{
			"message": "update user info success",
		})
	})
	r.DELETE("/user/info", func(context *gin.Context) {
		context.JSON(200, gin.H{
			"message": "delete user info success ",
		})

	})
	r.Run()
}
