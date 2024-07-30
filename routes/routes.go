package routes

import (
	"todo-list-api/handlers"
	"todo-list-api/middleware"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	api := gin.Default()
	api.POST("/register", handlers.Register)
	api.POST("/login", handlers.Login)
	authorized := api.Group("/")
	authorized.Use(middleware.AuthMiddleware())
	{
		authorized.GET("/tasks", handlers.GetTask)
		authorized.POST("/tasks", handlers.CreateTask)
		authorized.PUT("/tasks", handlers.UpdateTask)
		authorized.DELETE("/tasks/:id", handlers.DeleteTask)

	}
	return api
}
