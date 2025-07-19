
package home_page

import (
	"net/http"

	"api/internal/database"
	"api/internal/models/home_page" // Go 会自动找到新的 HomePageCarousel 模型

	"github.com/gin-gonic/gin"
)

// 我们可以把函数名也改得更贴切一些
func GetHomePageCarousel(c *gin.Context) {
	// 变量类型从 []models.Slide 改为 []models.HomePageCarousel
	var carousels []models.HomeCarousel

	// GORM 查询现在会作用于 home_page_carousel 表
	result := database.DB.Find(&carousels)

	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve carousel data"})
		return
	}

	// 返回查询到的新数据
	c.JSON(http.StatusOK, carousels)
}