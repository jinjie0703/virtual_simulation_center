package home_page

import (
	"log"
	"net/http"

	"api/internal/database"
	"api/internal/models/home_page"

	"github.com/gin-gonic/gin"
)

func GetUpdatesSectionNews(c *gin.Context) {
	var newsItems []home_page.UpdatesSectionNews

	result := database.DB.Table("updates_section_news").Order("date desc").Find(&newsItems)

	if result.Error != nil {
		log.Printf("数据库错误: %v", result.Error)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "无法检索新闻动态数据"})
		return
	}

	c.JSON(http.StatusOK, newsItems)
}
