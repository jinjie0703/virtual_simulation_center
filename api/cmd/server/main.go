// cmd/server/main.go
package main

import (
	"log"
	"virtual_simulation_center/api/internal/database"
	"virtual_simulation_center/api/internal/routes"
	"virtual_simulation_center/api/pkg/config"
)

func main() {
	// 1. 加载配置 (这一步保持不变)
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("配置加载失败: %v", err)
	}

	// 2. 从配置中获取当前激活环境的DSN
	//    这是最主要的改动！
	activeDSN, err := cfg.GetDSN()
	if err != nil {
		log.Fatalf("获取数据库DSN失败: %v", err)
	}

	// (可选) 打印调试信息，确认我们获取到了正确的DSN
	log.Printf("当前激活的数据库环境: [%s]", cfg.Database.Active)
	log.Printf("将要使用的数据库连接字符串(DSN): [%s]", activeDSN)

	// 3. 使用获取到的DSN连接数据库
	database.Connect(activeDSN)

	// 4. 设置并获取路由引擎
	r := routes.SetupRouter()

	// 5. 启动服务器
	log.Printf("服务开启的端口为 %s", cfg.Server.Port)

	// 注意：对于 Docker 部署，我们必须启动一个普通的 HTTP 服务器。
	// Azure App Service 平台会在外部为我们处理 HTTPS。
	log.Printf("服务将在 %s 端口上启动一个 HTTP 服务器", cfg.Server.Port)
	if err := r.Run(cfg.Server.Port); err != nil {
    log.Fatalf("Fatal: HTTP 服务启动失败: %v", err)
	}
	// 注意：您使用的是RunTLS，这意味着您需要配置HTTPS。
	// 如果只是本地开发，有时会用 r.Run(cfg.Server.Port) 来启动HTTP服务。
	// 这里我保持您原有的RunTLS不变。
	// if err := r.RunTLS(cfg.Server.Port, "./api/testcrt/localhost+2.pem", "./api/testcrt/localhost+2-key.pem"); err != nil {
	// 	log.Fatalf("Fatal: 服务启动失败: %v", err)
	// }
}