package information_center

import (
	"api/internal/models/information_center"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// GetCompetitionsHandler 处理获取竞赛列表的请求
func GetCompetitionsHandler(c *gin.Context) {
	competitionsList, err := information_center.GetAllCompetitions()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve competitions from database"})
		return
	}

	c.JSON(http.StatusOK, competitionsList)
}

// GetCompetitionByIDHandler handles the request to get a single competition by its ID.
func GetCompetitionByIDHandler(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid competition ID"})
		return
	}

	competition, err := information_center.GetCompetitionByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Competition not found"})
		return
	}

	c.JSON(http.StatusOK, competition)
}
