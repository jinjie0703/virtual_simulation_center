package information_center

import (
	"net/http"
	"api/internal/database"
	"api/internal/models/information_center"

	"github.com/gin-gonic/gin"
)

// GetNewsHandler 处理获取新闻列表的请求
func GetNewsHandler(c *gin.Context) {
	// 调用数据层获取数据
	newsList, err := database.GetAllNews()
	if err != nil {
		// 如果数据库查询出错，返回服务器内部错误
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve news from database"})
		return
	}

	// 成功获取数据，返回200 OK 和新闻列表
	c.JSON(http.StatusOK, newsList)
}