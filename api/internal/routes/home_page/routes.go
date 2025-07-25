package routes

import (
	// 注意这里的导入路径

	"api/internal/handlers/home_page"
	"time" // 导入 time 包用于配置

	"github.com/gin-gonic/gin"

	"github.com/gin-contrib/cors"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	// 中间件必须在路由前面注册
	r.Use(cors.New(cors.Config{
		// AllowOrigins:     []string{"https://foo.com"}, // 生产环境使用这个，指定你的前端域名
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
		// 我们保持 URL 不变，还是 /api/slides，但让它调用新的处理函数
		// 注意这里从 handlers.GetSlides 变成了 handlers.GetHomePageCarousel
		// api.GET("/slides", handlers.GetHomePageCarousel)

		// 或者，你也可以把 URL 也改得更贴切，这也是一种好方法
		api.GET("/HomeCarousel", home_page.GetHomeCarousel)
		api.GET("/PartnerCarousel", home_page.GetPartnerCarousel)
		api.GET("/FeatureSection", home_page.GetFeatureSection)
		api.GET("/UpdatesSection", home_page.GetUpdatesSection)
	}
	return r
}
