package information_center

import (
	"net/http"
	"api/internal/database"
	"api/internal/models/information_center"

	"github.com/gin-gonic/gin"
)

// GetProjectsHandler 处理获取项目列表的请求
func GetProjectsHandler(c *gin.Context) {
	projectsList, err := database.GetAllProjects()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve projects from database"})
		return
	}

	c.JSON(http.StatusOK, projectsList)
}