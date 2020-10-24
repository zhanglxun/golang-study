package router

import (
	"github.com/gin-gonic/gin"
	"golangStudy/3.apiServer/handler/sd"
	"golangStudy/3.apiServer/handler/user"
	"golangStudy/3.apiServer/router/middleware"
	"net/http"
)

// Load loads the middlewares, routes, handlers.
func Load(g *gin.Engine, mw ...gin.HandlerFunc) *gin.Engine {
	// Middlewares.
	g.Use(gin.Recovery())
	g.Use(middleware.NoCache)
	g.Use(middleware.Options)
	g.Use(middleware.Secure)
	g.Use(mw...)
	// 404 Handler.
	g.NoRoute(func(c *gin.Context) {
		c.String(http.StatusNotFound, "The incorrect API route.")
	})

	/**
	请求body 需要用此方式
	{
	    "username":"admin",
	    "password":"admin"
	}
	*/
	u := g.Group("/v1/user")
	{
		u.POST("", user.Create)
		u.POST("/:username", user.UserInfo)
	}

	// The health check handlers
	svcd := g.Group("/sd")
	{
		svcd.GET("/health", sd.HealthCheck)
		svcd.GET("/disk", sd.DiskCheck)
		svcd.GET("/cpu", sd.CPUCheck)
		svcd.GET("/ram", sd.RAMCheck)
	}

	return g
}
