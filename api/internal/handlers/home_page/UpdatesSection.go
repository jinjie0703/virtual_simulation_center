package home_page

import (
	"net/http"

	"api/internal/database"
	"api/internal/models/home_page"

	"github.com/gin-gonic/gin"
)

// GetUpdatesSection 用于获取动态更新区数据
func GetUpdatesSection(c *gin.Context) {
	var updates []home_page.UpdatesSection

	// 从 'updates_section' 表中查找所有记录
	// 为了让最新的动态显示在最前面，我们按日期降序排序
	result := database.DB.Order("date desc").Find(&updates)

	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve updates section data"})
		return
	}

	c.JSON(http.StatusOK, updates)
}
