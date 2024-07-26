package handler

import (
	"github.com/gin-gonic/gin"
	"project-management/internal/service"
)

type UserHandler struct {
	userService *service.UserService
}

func NewUserHandler(service *service.UserService) *UserHandler {
	return &UserHandler{
		userService: service,
	}
}

func (uh *UserHandler) CreateUser(c *gin.Context) {

}

func (uh *UserHandler) ListUser(c *gin.Context) {

}
func (uh *UserHandler) GetUser(c *gin.Context) {

}
func (uh *UserHandler) UpdateUser(c *gin.Context) {

}
func (uh *UserHandler) DeleteUser(c *gin.Context) {

}
