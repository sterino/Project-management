package http

import (
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"log"
	"project-management/internal/api/handler"
	"project-management/internal/api/routes"
)

type Server struct {
	engine *gin.Engine
}

func NewServer(taskHandler *handler.TaskHandler, userHandler *handler.UserHandler, projectHandler *handler.ProjectHandler) *Server {
	router := gin.Default()

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	routes.InitRoutes(router.Group("/api"), taskHandler, userHandler, projectHandler)

	return &Server{router}
}

func (s *Server) Run(infoLog *log.Logger, errorLog *log.Logger) {
	infoLog.Printf("starting server on: 8080")
	err := s.engine.Run(":8080")
	errorLog.Fatal(err)
}
