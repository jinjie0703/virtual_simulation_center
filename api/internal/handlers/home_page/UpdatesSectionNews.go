package home_page

import (
	"net/http"

	"api/internal/database"
	"api/internal/models/home_page"

	"github.com/gin-gonic/gin"
)

// GetUpdatesSection 用于获取动态更新区数据
func GetUpdatesSectionNews(c *gin.Context) {
	var updates []home_page.UpdatesSectionNews

	// 从 'update_section_competitions' 表中查找分类为 '学术荣誉' 的记录
	// 为了让最新的动态显示在最前面，我们按日期降序排序
	result := database.DB.Table("update_section_competitions").Where("category = ?", "学术荣誉").Order("date desc").Find(&updates)

	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve updates section data"})
		return
	}

	c.JSON(http.StatusOK, updates)
}
