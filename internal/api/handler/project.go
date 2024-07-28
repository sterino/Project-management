package handler

import (
	"errors"
	"github.com/gin-gonic/gin"
	"net/http"
	"project-management/internal/domain/project"
	interfaces "project-management/internal/service/interface"
	"project-management/pkg/response"
)

type ProjectHandler struct {
	projectService interfaces.ProjectService
}

func NewProjectHandler(service interfaces.ProjectService) *ProjectHandler {
	return &ProjectHandler{
		projectService: service,
	}
}

// CreateProject godoc
// @Summary Create a new project
// @Description Create a new project with the input payload
// @Tags projects
// @Accept json
// @Produce json
// @Param project body project.Request true "Project Request"
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

	if err := req.Validate(); err != nil {
		errRes := response.ClientResponse(http.StatusBadRequest, "fields provided are wrong", nil, err.Error())
		c.JSON(http.StatusBadRequest, errRes)
		return
	}

	res, err := ph.projectService.CreateProject(c.Request.Context(), req)
	if err != nil {
		if errors.Is(err, project.ErrorInvalidDate) {
			errRes := response.ClientResponse(http.StatusBadRequest, "fields provided are wrong", nil, err.Error())
			c.JSON(http.StatusBadRequest, errRes)
			return
		}
		errRes := response.ClientResponse(http.StatusInternalServerError, "failed to create project", nil, err.Error())
		c.JSON(http.StatusInternalServerError, errRes)
		return
	}
	successRes := response.ClientResponse(http.StatusCreated, "the project was successfully created", res, nil)
	c.JSON(http.StatusCreated, successRes)
}

// ListProjects godoc
// @Summary List all projects
// @Description Get a list of all projects
// @Tags projects
// @Produce json
// @Success 200 {object} response.Response
// @Failure 500 {object} response.Response
// @Failure 404 {object} response.Response
// @Router /projects [get]
func (ph *ProjectHandler) ListProjects(c *gin.Context) {
	res, err := ph.projectService.ListProjects(c.Request.Context())
	if err != nil {
		if errors.Is(err, project.ErrorNotFound) {
			errRes := response.ClientResponse(http.StatusNotFound, "projects not found", nil, err.Error())
			c.JSON(http.StatusNotFound, errRes)
			return
		}
		errRes := response.ClientResponse(http.StatusInternalServerError, "failed to list projects", nil, err.Error())
		c.JSON(http.StatusInternalServerError, errRes)
		return
	}

	successRes := response.ClientResponse(http.StatusOK, "the projects list", res, nil)
	c.JSON(http.StatusOK, successRes)
}

// GetProject godoc
// @Summary Get a project by ID
// @Description Get details of a project by its ID
// @Tags projects
// @Produce json
// @Param id path string true "Project ID"
// @Success 200 {object} response.Response
// @Failure 404 {object} response.Response
// @Failure 500 {object} response.Response
// @Router /projects/{id} [get]
func (ph *ProjectHandler) GetProject(c *gin.Context) {
	id := c.Param("id")
	res, err := ph.projectService.GetProject(c.Request.Context(), id)
	if err != nil {
		if errors.Is(err, project.ErrorNotFound) {
			errRes := response.ClientResponse(http.StatusNotFound, "project not found", nil, err.Error())
			c.JSON(http.StatusNotFound, errRes)
			return
		}
		errRes := response.ClientResponse(http.StatusInternalServerError, "failed to get project", nil, err.Error())
		c.JSON(http.StatusInternalServerError, errRes)
		return
	}
	successRes := response.ClientResponse(http.StatusOK, "the project details", res, nil)
	c.JSON(http.StatusOK, successRes)
}

// UpdateProject godoc
// @Summary Update a project by ID
// @Description Update details of a project by its ID
// @Tags projects
// @Accept json
// @Produce json
// @Param id path string true "Project ID"
// @Param project body project.Request true "Project Request"
// @Success 200 {object} response.Response
// @Failure 400 {object} response.Response
// @Failure 404 {object} response.Response
// @Failure 500 {object} response.Response
// @Router /projects/{id} [put]
func (ph *ProjectHandler) UpdateProject(c *gin.Context) {
	id := c.Param("id")
	req := project.Request{}
	if err := c.BindJSON(&req); err != nil {
		errRes := response.ClientResponse(http.StatusBadRequest, "fields provided are wrong", nil, err.Error())
		c.JSON(http.StatusBadRequest, errRes)
		return
	}

	if err := req.Validate(); err != nil {
		errRes := response.ClientResponse(http.StatusBadRequest, "fields provided are wrong", nil, err.Error())
		c.JSON(http.StatusBadRequest, errRes)
		return
	}

	err := ph.projectService.UpdateProject(c.Request.Context(), id, req)
	if err != nil {
		if errors.Is(err, project.ErrorNotFound) {
			errRes := response.ClientResponse(http.StatusNotFound, "project not found", nil, err.Error())
			c.JSON(http.StatusNotFound, errRes)
			return
		}
		errRes := response.ClientResponse(http.StatusInternalServerError, "failed to update project", nil, err.Error())
		c.JSON(http.StatusInternalServerError, errRes)
		return
	}
	successRes := response.ClientResponse(http.StatusOK, "the project was successfully updated", nil, nil)
	c.JSON(http.StatusOK, successRes)
}

// DeleteProject godoc
// @Summary Delete a project by ID
// @Description Delete a project by its ID
// @Tags projects
// @Produce json
// @Param id path string true "Project ID"
// @Success 200 {object} response.Response
// @Failure 404 {object} response.Response
// @Failure 500 {object} response.Response
// @Router /projects/{id} [delete]
func (ph *ProjectHandler) DeleteProject(c *gin.Context) {
	id := c.Param("id")
	err := ph.projectService.DeleteProject(c.Request.Context(), id)
	if err != nil {
		if errors.Is(err, project.ErrorNotFound) {
			errRes := response.ClientResponse(http.StatusNotFound, "project not found", nil, err.Error())
			c.JSON(http.StatusNotFound, errRes)
			return
		}
		errRes := response.ClientResponse(http.StatusInternalServerError, "failed to delete project", nil, err.Error())
		c.JSON(http.StatusInternalServerError, errRes)
		return
	}
	successRes := response.ClientResponse(http.StatusOK, "the project was successfully deleted", nil, nil)
	c.JSON(http.StatusOK, successRes)
}

// SearchProjects godoc
// @Summary Search projects by title or manager ID
// @Description Search projects by title or manager ID
// @Tags projects
// @Produce json
// @Param filter query string true "Query"
// @Param val query string true "Value"
// @Success 200 {object} response.Response
// @Failure 400 {object} response.Response
// @Failure 404 {object} response.Response
// @Failure 500 {object} response.Response
// @Router /projects/search [get]
func (ph *ProjectHandler) SearchProjects(c *gin.Context) {
	filter := c.Query("filter")
	val := c.Query("val")

	if filter == "" || val == "" {
		errRes := response.ClientResponse(http.StatusBadRequest, "fields provided are wrong", nil, "filter and value are required")
		c.JSON(http.StatusBadRequest, errRes)
		return
	}

	res, err := ph.projectService.SearchProject(c.Request.Context(), filter, val)
	if err != nil {
		if errors.Is(err, project.ErrorInvalidSearch) {
			errRes := response.ClientResponse(http.StatusBadRequest, "fields provided are wrong", nil, err.Error())
			c.JSON(http.StatusBadRequest, errRes)
			return
		}
		if errors.Is(err, project.ErrorNotFound) {
			errRes := response.ClientResponse(http.StatusNotFound, "projects not found", nil, err.Error())
			c.JSON(http.StatusNotFound, errRes)
			return
		}

		errRes := response.ClientResponse(http.StatusInternalServerError, "failed to search projects", nil, err.Error())
		c.JSON(http.StatusInternalServerError, errRes)
		return
	}
	successRes := response.ClientResponse(http.StatusOK, "the projects list", res, nil)
	c.JSON(http.StatusOK, successRes)
}

// GetProjectTasks godoc
// @Summary Get tasks of a project by project ID
// @Description Get tasks of a project by project ID
// @Tags projects
// @Produce json
// @Param id path string true "Project ID"
// @Success 200 {object} response.Response
// @Failure 404 {object} response.Response
// @Failure 500 {object} response.Response
// @Router /projects/{id}/tasks [get]
func (ph *ProjectHandler) GetProjectTasks(c *gin.Context) {
	id := c.Param("id")
	res, err := ph.projectService.GetProjectTasks(c.Request.Context(), id)
	if err != nil {
		if errors.Is(err, project.ErrorNotFound) {
			errRes := response.ClientResponse(http.StatusNotFound, "project not found", nil, err.Error())
			c.JSON(http.StatusNotFound, errRes)
			return
		}
		errRes := response.ClientResponse(http.StatusInternalServerError, "failed to get project tasks", nil, err.Error())
		c.JSON(http.StatusInternalServerError, errRes)
		return
	}
	successRes := response.ClientResponse(http.StatusOK, "the project tasks list", res, nil)
	c.JSON(http.StatusOK, successRes)
}
