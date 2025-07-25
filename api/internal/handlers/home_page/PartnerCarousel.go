package home_page

import (
	"net/http"

	"api/internal/database"
	"api/internal/models/home_page"

	"github.com/gin-gonic/gin"
)

// GetPartnerCarousel 用于获取合作伙伴轮播数据
func GetPartnerCarousel(c *gin.Context) {
	var partners []home_page.PartnerCarousel

	// 从 'partner_carousel' 表中查找所有记录
	result := database.DB.Find(&partners)

	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve partner carousel data"})
		return
	}

	c.JSON(http.StatusOK, partners)
}
