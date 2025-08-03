package information_center

import (
	"api/internal/database"
	"api/internal/models/information_center"
	"net/http"

	"github.com/gin-gonic/gin"
)

// GetCompetitionsHandler 处理获取竞赛列表的请求
func GetCompetitionsHandler(c *gin.Context) {
	competitionsList, err := database.GetAllCompetitions()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve competitions from database"})
		return
	}

	c.JSON(http.StatusOK, competitionsList)
}
