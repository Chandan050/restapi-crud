package main

import (
	
"app/controllers"
"github.com/gin-gonic/gin"
	"github.com/swaggo/gin-swagger"
	"github.com/swaggo/files"


)


// @title Swagger Example API
// @version 1.0
// @description This is a sample server.
// @host localhost:8080
// @BasePath /
func main() {
	// connect to database
	controllers.ConnectDatabase()
	router := gin.Default()

	// register routes
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// API routes group
	api := router.Group("/api")
	{
	api.POST("/students", controllers.CreateStudent)
	api.POST("/courses", controllers.CreateCourse)
	api.POST("/scores", controllers.CreateScore)
	api.GET("/score/:id", controllers.GetScore)
	api.PUT("/score/:id/:scoreid", controllers.UpdateScore)
	}
	router.Run(":8080")


}