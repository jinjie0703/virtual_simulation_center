package information_center

import (
	"net/http"
	"strconv"

	"api/internal/models/information_center"

	"github.com/gin-gonic/gin"
)

// GetProjectsHandler 处理获取项目列表的请求
func GetProjectsHandler(c *gin.Context) {
	projectsList, err := information_center.GetAllProjects()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve projects from database"})
		return
	}

	c.JSON(http.StatusOK, projectsList)
}

// GetProjectByIDHandler handles the request to get a single project by its ID.
func GetProjectByIDHandler(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid project ID"})
		return
	}

	project, err := information_center.GetProjectByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Project not found"})
		return
	}

	c.JSON(http.StatusOK, project)
}
