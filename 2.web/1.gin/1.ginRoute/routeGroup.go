package main

import "github.com/gin-gonic/gin"

func login(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "login success",
	})
}
func submit(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "submit success",
	})
}

func read(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "read success",
	})
}

func main() {
	r := gin.Default()

	v1 := r.Group("v1")
	{
		v1.GET("/login", login)
		v1.GET("/submit", submit)
		v1.GET("/read", read)
	}
	r.Run()

}
