package router

import (
	"github.com/gin-gonic/gin"
	"ynbauth/middleware"
	"ynbauth/utils"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	// Public route
	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "Welcome to ynbauth!"})
	})

	// Auth routes
	auth := r.Group("/api")
	{
		auth.POST("/auth-token-obtain", utils.ObtainToken)
		auth.POST("/auth-token-verify", utils.VerifyToken)
		auth.POST("/auth-token-refresh", utils.RefreshToken)
		auth.POST("/new-account", utils. RegisterUser)
	}

	// Protected route
	r.GET("/protected", middleware.AuthMiddleware(), func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "You are authorized!"})
	})

	// Forward requests
	// r.GET("/service-a", middleware.AuthMiddleware(), func(c *gin.Context) {
	// 	resp, err := http.Get("http://localhost:8081") // Microservice A
	// 	if err != nil {
	// 		c.JSON(http.StatusInternalServerError, gin.H{"error": "Service unavailable"})
	// 		return
	// 	}
	// 	defer resp.Body.Close()
	
	// 	c.DataFromReader(resp.StatusCode, resp.ContentLength, resp.Header.Get("Content-Type"), resp.Body, nil)
	// })
	

	return r
}
