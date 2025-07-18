// internal/routes/routes.go
package routes

import (
	"api/internal/handlers" // 注意这里的导入路径

	"github.com/gin-gonic/gin"
)

// SetupRouter 配置所有路由
func SetupRouter() *gin.Engine {
	r := gin.Default()

	// 这里可以设置全局中间件，比如 CORS
	// r.Use(cors.Default())

	// 设置静态文件服务
	// URL /static/* 会映射到本地的 ./public/ 目录
	r.Static("/static", "./public")

	// API 路由组
	api := r.Group("/api")
	{
		// GET /api/slides
		api.GET("/slides", handlers.GetSlides)
		// 在这里添加其他 API 路由，比如用户相关的
		// api.POST("/users", handlers.CreateUser)
	}
	return r
}