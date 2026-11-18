# go-practice（练习仓库）

这是面向初学者的 **Go 后端分层练习壳**（`github.com/FelixDu2025/go-practice`），**不是**已经写完的样例项目。

仓库保留了目录、模型、接口、路由骨架和测试；`handler` / `service` / `repository` / `middleware` 的业务实现已抽成 stub。请你自己完成代码。

**请先打开 [`EXERCISES.md`](./EXERCISES.md)，按步骤练习。**

## 你会练到什么

- 分层：`handler → service → repository`
- `GET /health`
- 内存 Todo CRUD（标准库 `net/http`，Go 1.22+ ServeMux）
- 请求日志中间件
- 配置、优雅关闭与基础测试

零第三方依赖。

## 目录结构

```text
cmd/server/                 # 程序入口：组装依赖、启动 HTTP、优雅关闭
internal/
  config/                   # 从环境变量读取配置
  handler/                  # HTTP 适配层（路由与 JSON）——待你实现
  service/                  # 业务规则——待你实现
  repository/               # 数据访问（内存）——待你实现
  middleware/               # 请求日志等——待你实现
  model/                    # 领域模型与请求 DTO（已给出）
EXERCISES.md                # 有序练习步骤（从这里开始）
.env.example                # 环境变量示例
Makefile                    # run / test / tidy
Dockerfile                  # 可选容器构建
```

## 环境要求

- Go 1.22+

## 快速开始

```bash
cp .env.example .env
export ADDR=:8080

# 当前可编译启动，但业务接口会返回未实现错误，直到你补全代码
make run
```

实现完成后可自测：

```bash
curl http://localhost:8080/health
```

## 目标 API（需自行实现）

| 方法 | 路径 | 说明 |
|------|------|------|
| GET | `/health` | 健康检查 |
| GET | `/api/todos` | 列出全部待办 |
| POST | `/api/todos` | 创建待办 |
| GET | `/api/todos/{id}` | 获取单条 |
| PUT | `/api/todos/{id}` | 更新（部分字段） |
| DELETE | `/api/todos/{id}` | 删除 |

## Makefile

```bash
make run    # 启动服务
make test   # 跑测试（实现前会失败，属预期）
make tidy   # go mod tidy
```

说明：

- `go build ./...`：练习壳应能通过
- `go test ./...`：在补全实现前失败是正常的

## 可选：Docker

实现并确保 `go build ./cmd/server` 通过后：

```bash
docker build -t go-practice .
docker run --rm -p 8080:8080 go-practice
```

## English (brief)

Exercise shell for a layered Go HTTP API. Stubs compile; tests fail until you implement health, in-memory todo CRUD, and request logging. Start with `EXERCISES.md`.
