package home_page

import (
	"net/http"

	"api/internal/database"
	"api/internal/models/home_page"

	"github.com/gin-gonic/gin"
)

// GetUpdatesSectionCompetitions 用于获取动态更新区数据
func GetUpdatesSectionCompetitions(c *gin.Context) {
	var updates []home_page.UpdatesSectionCompetitions

	// 从 'updates_section_competitions' 表中查找所有记录
	// 为了让最新的动态显示在最前面，我们按截止日期降序排序
	result := database.DB.Table("updates_section_competitions").Order("deadline desc").Find(&updates)

	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve updates section data"})
		return
	}

	c.JSON(http.StatusOK, updates)
}
