package home_page

import (
	"net/http"

	"virtual_simulation_center/api/internal/database"
	"virtual_simulation_center/api/internal/models/home_page"

	"github.com/gin-gonic/gin"
)

// GetFeatureSection 用于获取特色功能区数据
func GetFeatureSection(c *gin.Context) {
	var features []home_page.FeatureSection

	// 明确指定表名，以确保查询的准确性
	result := database.DB.Table(home_page.FeatureSection{}.TableName()).Find(&features)

	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "获取数据失败"})
		return
	}

	c.JSON(http.StatusOK, features)
}
