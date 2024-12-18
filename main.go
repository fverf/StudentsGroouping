package main

import (
	"jwt-authentication-golang/controllers"
	"jwt-authentication-golang/database"
	"jwt-authentication-golang/middlewares"
	"github.com/gin-gonic/gin"
  )
	func main() {
		// Initialize Database
		database.Connect("host=localhost user=postgres password=ivored19 dbname=studentsbd port=5432 sslmode=disable TimeZone=Europe/Moscow")
		database.Migrate()
	  
		// Initialize Router
		router := initRouter()
		router.Run(":8080")
	  }

  func initRouter() *gin.Engine {
	router := gin.Default()
	api := router.Group("/api")
	{
	  api.POST("/token", controllers.GenerateToken)
	  api.POST("/user/register", controllers.CreateUser)
	  secured := api.Group("/secured").Use(middlewares.Auth())
	  {
		secured.GET("/ping", controllers.Ping)
	  }
	}
	return router
  }