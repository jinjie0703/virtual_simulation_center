package home_page

import (
	"net/http"

	"api/internal/database"
	"api/internal/models/home_page"

	"github.com/gin-gonic/gin"
)

// GetFeatureSection 用于获取特色功能区数据
func GetFeatureSection(c *gin.Context) {
	var features []home_page.FeatureSection

	// 从 'feature_section' 表中查找所有记录
	result := database.DB.Find(&features)

	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve feature section data"})
		return
	}

	c.JSON(http.StatusOK, features)
}
