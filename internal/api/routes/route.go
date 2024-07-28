package routes

import (
	"github.com/gin-gonic/gin"
	"project-management/internal/api/handler"
)

func InitRoutes(router *gin.RouterGroup, taskHandler *handler.TaskHandler, userHandler *handler.UserHandler, projectHandler *handler.ProjectHandler) {
	tasks := router.Group("/tasks")
	{
		tasks.GET("/", taskHandler.ListTasks)
		tasks.POST("/", taskHandler.CreateTask)
		tasks.GET("/:id", taskHandler.GetTask)
		tasks.PUT("/:id", taskHandler.UpdateTask)
		tasks.DELETE("/:id", taskHandler.DeleteTask)
		tasks.GET("/search", taskHandler.SearchTasks)
	}
	users := router.Group("/users")
	{
		users.GET("/", userHandler.ListUsers)
		users.POST("/", userHandler.CreateUser)
		users.GET("/:id", userHandler.GetUser)
		users.PUT("/:id", userHandler.UpdateUser)
		users.DELETE("/:id", userHandler.DeleteUser)
		users.GET("/:id/tasks", userHandler.GetUserTasks)
		users.GET("/search", userHandler.SearchUsers)
	}

	projects := router.Group("/projects")
	{
		projects.GET("/", projectHandler.ListProjects)
		projects.POST("/", projectHandler.CreateProject)
		projects.GET("/:id", projectHandler.GetProject)
		projects.PUT("/:id", projectHandler.UpdateProject)
		projects.DELETE("/:id", projectHandler.DeleteProject)
		projects.GET("/:id/tasks", projectHandler.GetProjectTasks)
		projects.GET("/search", projectHandler.SearchProjects)
	}
}
