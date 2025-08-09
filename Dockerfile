# ---- Builder Stage ----
# 使用官方的 Go 镜像作为构建环境
FROM golang:1.24-alpine AS builder

# 设置工作目录
WORKDIR /app

# 复制 go.mod 和 go.sum 文件，以便缓存依赖
COPY go.mod go.sum ./
RUN go mod download

# 复制整个 api 文件夹到容器中
COPY api/ ./api/

# 【最终修正】指定正确的编译目标
# Go 编译器会从 ./cmd/server/main.go 开始编译整个项目
RUN CGO_ENABLED=0 GOOS=linux go build -a -o /app/main ./api/cmd/server/

# ---- Final Stage ----
FROM alpine:latest
WORKDIR /root/

# 1. 复制可执行文件
COPY --from=builder /app/main .

# 2. 【关键修改】在 /root/ 下创建一个 "api" 文件夹
RUN mkdir -p api

# 3. 【关键修改】将您本地的 api/public 文件夹，完整复制到容器的 /root/api/public/ 目录下
COPY api/public/ ./api/public/

# 4. 【关键修改】将您本地的 api/configs 文件夹，完整复制到容器的 /root/api/configs/ 目录下
COPY api/configs/ ./api/configs/

EXPOSE 8080
CMD ["./main"]