// internal/handlers/slide_handler.go
package handlers

import (
	"net/http"

	"api/internal/database" // 注意这里的导入路径
	"api/internal/models"

	"github.com/gin-gonic/gin"
)

// GetSlides 处理获取所有幻灯片数据的请求
func GetSlides(c *gin.Context) {
	var slides []models.Slide
	result := database.DB.Find(&slides) // 从数据库查询所有 slides
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve slides"})
		return
	}
	c.JSON(http.StatusOK, slides)
}