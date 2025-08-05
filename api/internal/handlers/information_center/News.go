package information_center

import (
	"net/http"
	"strconv"

	"virtual_simulation_center/api/internal/models/information_center"

	"github.com/gin-gonic/gin"
)

// GetNewsHandler handles the request to get all news.
func GetNewsHandler(c *gin.Context) {
	newsList, err := information_center.GetAllNews()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve news"})
		return
	}
	c.JSON(http.StatusOK, newsList)
}

// GetNewsByIDHandler handles the request to get a single news article by its ID.
func GetNewsByIDHandler(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid news ID"})
		return
	}

	news, err := information_center.GetNewsByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "News not found"})
		return
	}

	c.JSON(http.StatusOK, news)
}
