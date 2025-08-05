package home_page

import (
	"net/http"

	"virtual_simulation_center/api/internal/database"
	"virtual_simulation_center/api/internal/models/home_page"

	"github.com/gin-gonic/gin"
)

func GetHomeCarousel(c *gin.Context) {
	var carousels []home_page.HomeCarousel

	// 从 'home_carousel' 表中查找所有记录
	result := database.DB.Find(&carousels)

	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve home carousel data"})
		return
	}

	c.JSON(http.StatusOK, carousels)
}
