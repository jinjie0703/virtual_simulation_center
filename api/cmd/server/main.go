// cmd/server/main.go
package main

import (
	"api/internal/database"
	routes "api/internal/routes/home_page"
	"api/pkg/config"
	"log"
)

func main() {
	// 1. 加载配置
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("配置加载失败: %v", err)
	}

	log.Printf("DEBUG: 从config中读出的dsn是: [%s]", cfg.Database.DSN)
	// 2. 连接数据库
	database.Connect(cfg.Database.DSN)

	// 3. 设置并获取路由引擎
	r := routes.SetupRouter()

	// 4. 启动服务器
	log.Printf("服务开启的端口为 %s", cfg.Server.Port)
	if err := r.RunTLS(cfg.Server.Port, "./testcrt/localhost+2.pem", "./testcrt/localhost+2-key.pem"); err != nil {
		log.Fatalf("Fatal: 服务启动失败: %v", err)
	}

}
