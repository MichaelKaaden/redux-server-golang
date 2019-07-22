package main

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

// setupRouter initializes a router for production and debugging.
func setupRouter() *gin.Engine {
	var router *gin.Engine
	router = gin.Default()

	// configure CORS middleware *BEFORE* setting up any routes
	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"http://localhost:4200"}
	// config.AllowHeaders = []string{"Access-Control-Allow-Origin"}
	config.AllowMethods = []string{"GET", "PUT", "OPTIONS"}
	router.Use(cors.New(config))

	setupRoutes(router)
	return router
}

// setupRouterForTesting initializes a router for testing.
func setupRouterForTesting() *gin.Engine {
	var router *gin.Engine
	gin.SetMode(gin.TestMode)
	router = gin.New()
	setupRoutes(router)
	return router
}

// setupRoutes defines routes and their handlers.
func setupRoutes(router *gin.Engine) {
	router.GET("/counters", GetCounters)
	router.GET("/counters/:id", GetCounter)
	router.PUT("/counters/:id", PutCounter)
	router.PUT("/counters/:id/decrement", DecrementCounter)
	router.PUT("/counters/:id/increment", IncrementCounter)
}

func main() {
	router := setupRouter()
	router.Run(":3000") // listen and serve on 0.0.0.0:8080
}
