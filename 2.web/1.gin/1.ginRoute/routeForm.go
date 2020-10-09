package main

import "github.com/gin-gonic/gin"

func searchUser(context *gin.Context) {

	username := context.PostForm("username")
	address := context.PostForm("address")

	context.JSON(200, gin.H{
		"message":  "success",
		"username": username,
		"address":  address,
	})
}

func aboutNew(ctx *gin.Context) {
	ctx.JSON(200, gin.H{
		"message": "about",
	})
}

func main() {
	r := gin.Default()

	r.POST("/user/search", searchUser)
	r.GET("/user/about", aboutNew)

	r.Run()
}
