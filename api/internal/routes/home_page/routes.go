package routes

import (
	"virtual_simulation_center/api/internal/database"
	"virtual_simulation_center/api/internal/handlers/home_page"
	"virtual_simulation_center/api/internal/handlers/information_center"
	"virtual_simulation_center/api/internal/handlers/our_team"
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
		api.GET("/PartnerCarousel", home_page.GetCarouselByType)
		api.GET("/FeatureSection", home_page.GetFeatureSection)
		api.GET("/UpdatesSectionNews", home_page.GetUpdatesSectionNews)
		api.GET("/UpdatesSectionCompetitions", home_page.GetUpdatesSectionCompetitions)

		infoGroup := api.Group("/information_center")
		{
			infoGroup.GET("/news", information_center.GetNewsHandler)
			infoGroup.GET("/news/:id", information_center.GetNewsByIDHandler)
			infoGroup.GET("/competitions", information_center.GetCompetitionsHandler)
			infoGroup.GET("/competitions/:id", information_center.GetCompetitionByIDHandler)
			infoGroup.GET("/projects", information_center.GetProjectsHandler)
			infoGroup.GET("/projects/:id", information_center.GetProjectByIDHandler)
		}
		ourTeamGroup := api.Group("/our_team")
		{
			ourTeamGroup.GET("/members", our_team.GetTeamMembers(database.DB))
		}
	}
	return r
}
