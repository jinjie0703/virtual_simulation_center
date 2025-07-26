package routes

import (
	"api/internal/handlers/home_page"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	// 中间件必须在路由前面注册
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:5173"}, // 开发环境，允许你的 Vue 前端地址访问
		AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Accept", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))
	r.Static("/static", "./public") // 静态文件目录

	api := r.Group("/api")
	{
		api.GET("/HomeCarousel", home_page.GetHomeCarousel)
		api.GET("/PartnerCarousel", home_page.GetCarouselByType) // 更新处理器
		api.GET("/FeatureSection", home_page.GetFeatureSection)
		api.GET("/UpdatesSectionNews", home_page.GetUpdatesSectionNews)
		api.GET("/UpdatesSectionCompetitions", home_page.GetUpdatesSectionCompetitions)
	}
	return r
}
