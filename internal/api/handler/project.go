package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
	_ "project-management/docs"
	"project-management/internal/domain/project"
	"project-management/internal/service"
	"project-management/pkg/response"
)

type ProjectHandler struct {
	projectService *service.ProjectService
}

func NewProjectHandler(service *service.ProjectService) *ProjectHandler {
	return &ProjectHandler{
		projectService: service,
	}
}

// CreateProject godoc
// @Summary Create project
// @Tags project
// @Description Create project
// @Accept json
// @Produce json
// @Param request body project.Request true "Project details"
// @Success 201 {object} response.Response
// @Failure 400 {object} response.Response
// @Failure 500 {object} response.Response
// @Router /projects [post]
func (ph *ProjectHandler) CreateProject(c *gin.Context) {
	req := project.Request{}
	if err := c.BindJSON(&req); err != nil {
		errRes := response.ClientResponse(http.StatusBadRequest, "fields provided are wrong", nil, err.Error())
		c.JSON(http.StatusBadRequest, errRes)
		return
	}
	res, err := ph.projectService.CreateProject(c.Request.Context(), req)
	if err != nil {
		errRes := response.ClientResponse(http.StatusInternalServerError, "failed to create project", nil, err.Error())
		c.JSON(http.StatusInternalServerError, errRes)
		return
	}
	successRes := response.ClientResponse(http.StatusCreated, "the project was successfully created", res, nil)
	c.JSON(http.StatusCreated, successRes)
}

// ListProjects godoc
// @Summary List projects
// @Tags project
// @Description List projects
// @Success 200 {object} response.Response
// @Failure 500 {object} response.Response
// @Router /projects [get]
func (ph *ProjectHandler) ListProjects(c *gin.Context) {
	res, err := ph.projectService.ListProjects(c.Request.Context())
	if err != nil {
		errRes := response.ClientResponse(http.StatusInternalServerError, "failed to list projects", nil, err.Error())
		c.JSON(http.StatusInternalServerError, errRes)
		return
	}

	successRes := response.ClientResponse(http.StatusOK, "the projects list", res, nil)
	c.JSON(http.StatusOK, successRes)
}

func (ph *ProjectHandler) GetProject(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{})
}

func (ph *ProjectHandler) DeleteProject(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{})
}

func (ph *ProjectHandler) UpdateProject(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{})
}
