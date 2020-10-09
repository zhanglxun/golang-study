package main

import "github.com/gin-gonic/gin"

func main() {
	r := gin.Default()
	r.GET("/user/search", func(context *gin.Context) {
		username := context.DefaultQuery("username", "jarvis")
		address := context.Query("address")

		context.JSON(200, gin.H{
			"message":  "pong",
			"username": username,
			"address":  address,
		})
	})
	r.Run(":8081")

	/*
		http://localhost:8081/user/search?address=shenzhen
		http://localhost:8081/user/search?address=shenzhen&username=king
	*/
}
