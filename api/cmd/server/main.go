// cmd/server/main.go
package main

import (
	"api/internal/database"                // 注意这里的导入路径
	models "api/internal/models/home_page" // 注意这里的导入路径，确保它指向新的模型文件
	"api/internal/routes"
	"api/pkg/config"
	"log"
)

func main() {
	// 1. 加载配置
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("Fatal: could not load config: %v", err)
	}

	log.Printf("DEBUG: The DSN string read from config is: [%s]", cfg.Database.DSN)
	// 2. 连接数据库
	database.Connect(cfg.Database.DSN)

	// 3. 自动迁移数据库模型 (可选，在开发环境很方便)
	// 这会确保数据库中的表结构和你的 models 定义一致
	err = database.DB.AutoMigrate(&models.HomePageCarousel{})
	if err != nil {
		log.Fatalf("Fatal: could not auto migrate database: %v", err)
	}

	// 4. 设置并获取路由引擎
	r := routes.SetupRouter()

	// 5. 启动服务器
	log.Printf("Server starting on port %s", cfg.Server.Port)
	if err := r.RunTLS(cfg.Server.Port, "./testcrt/server.pem", "./testcrt/server.key"); err != nil {
		log.Fatalf("Fatal: Failed to run server: %v", err)
	}

}
