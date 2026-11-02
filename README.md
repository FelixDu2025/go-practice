# go-practice

面向初学者的 Go 后端分层练习项目（`github.com/FelixDu2025/go-practice`）。

使用标准库 `net/http`（Go 1.22+ ServeMux），零第三方依赖，结构清晰，方便练习：

- 分层：`handler → service → repository`
- 健康检查
- Todo / 笔记的内存 CRUD
- 请求日志中间件
- 优雅关闭（SIGINT / SIGTERM）

## 目录结构

```text
cmd/server/                 # 程序入口：组装依赖、启动 HTTP、优雅关闭
internal/
  config/                   # 从环境变量读取配置
  handler/                  # HTTP 适配层（路由与 JSON）
  service/                  # 业务规则
  repository/               # 数据访问（当前为内存实现）
  middleware/               # 请求日志等中间件
  model/                    # 领域模型与请求 DTO
.env.example                # 环境变量示例
Makefile                    # run / test / tidy
Dockerfile                  # 可选容器构建
```

## 环境要求

- Go 1.22+

## 快速开始

```bash
# 可选：复制环境变量
cp .env.example .env
# 本项目直接读进程环境变量；如需加载 .env，可自行 export 或使用 direnv 等工具
export ADDR=:8080

make run
# 或：go run ./cmd/server
```

健康检查：

```bash
curl http://localhost:8080/health
```

## API 一览

| 方法 | 路径 | 说明 |
|------|------|------|
| GET | `/health` | 健康检查 |
| GET | `/api/todos` | 列出全部待办 |
| POST | `/api/todos` | 创建待办 |
| GET | `/api/todos/{id}` | 获取单条 |
| PUT | `/api/todos/{id}` | 更新（部分字段） |
| DELETE | `/api/todos/{id}` | 删除 |

### 创建示例

```bash
curl -s -X POST http://localhost:8080/api/todos \
  -H 'Content-Type: application/json' \
  -d '{"title":"学习 Go 分层","content":"先跑通内存 CRUD"}'
```

### 更新示例

```bash
curl -s -X PUT http://localhost:8080/api/todos/<id> \
  -H 'Content-Type: application/json' \
  -d '{"done":true}'
```

## Makefile

```bash
make run    # 启动服务
make test   # 跑单元测试
make tidy   # go mod tidy
```

## 可选：Docker

```bash
docker build -t go-practice .
docker run --rm -p 8080:8080 go-practice
```

## 学习建议

1. 先读 `cmd/server/main.go`，看依赖如何组装。
2. 从 `handler` 跟到 `service`、`repository`，理解每一层职责。
3. 尝试把内存仓库换成 SQLite / PostgreSQL（只改 `repository` 实现）。
4. 扩展字段校验、分页、或换成 `chi` 路由库做对比练习。

## English (brief)

Beginner-friendly layered Go HTTP API with in-memory todo CRUD, request logging, and graceful shutdown. Zero third-party deps (stdlib `net/http`). Run with `make run`, test with `make test`.
