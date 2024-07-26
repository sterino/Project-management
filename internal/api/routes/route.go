package routes

import (
	"github.com/gin-gonic/gin"
	"project-management/internal/api/handler"
)

func InitRoutes(router *gin.RouterGroup, taskHandler *handler.TaskHandler, userHandler *handler.UserHandler, projectHandler *handler.ProjectHandler) {
	tasks := router.Group("/tasks")
	{
		tasks.GET("/")
		tasks.POST("/")
		tasks.GET("/:id")
		tasks.PUT("/:id")
		tasks.DELETE("/:id")
	}

	users := router.Group("/users")
	{
		users.GET("/")
		users.POST("/")
		users.GET("/:id")
		users.PUT("/:id")
		users.DELETE("/:id")
		users.GET("/:id/tasks")
	}
	projects := router.Group("/projects")
	{
		projects.GET("/", projectHandler.ListProjects)
		projects.POST("/", projectHandler.CreateProject)
		projects.GET("/:id")
		projects.PUT("/:id")
		projects.DELETE("/:id")
	}
}
