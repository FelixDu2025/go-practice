# go-practice 练习步骤

本仓库是**练习壳**，不是可直接照抄的成品示例。业务实现已抽空，请按下列顺序自己写代码；`go build ./...` 应能通过，`go test ./...` 在你补全实现前会失败，这是预期行为。

建议先通读本文件，再打开对应包里的中文 `TODO` 注释。

---

## 步骤 0：环境与目录

1. 安装 Go 1.22+。
2. 复制环境变量示例：`cp .env.example .env`，按需 `export ADDR=:8080`。
3. 熟悉分层：`handler → service → repository`，入口在 `cmd/server`。
4. 常用命令：`make run` / `make test` / `make tidy`。

---

## 步骤 1：读懂并确认入口组装（wire main）

目标文件：`cmd/server/main.go`、`internal/config`

你需要搞清楚：

- 如何 `config.Load()` 读 `ADDR`、`SHUTDOWN_TIMEOUT_SEC`
- 如何把 `repository → service → handler.NewRouter` 串起来
- 如何 `ListenAndServe`，以及 SIGINT/SIGTERM 下的 `Shutdown`

入口骨架已能编译启动。请边读边改注释里的 TODO，确保你理解每一行，而不是机械保留。

自测：`go run ./cmd/server` 能起来；Ctrl+C 能退出。

---

## 步骤 2：健康检查

目标文件：`internal/handler/health.go`

实现 `GET /health`：

- 状态码 `200`
- JSON 形如 `{"status":"ok"}`（可复用同包的 `writeJSON`）

自测：

```bash
curl -i http://localhost:8080/health
```

相关测试：`TestHealth`（实现前会失败）。

---

## 步骤 3：内存 Todo CRUD（由内向外）

### 3.1 Repository

目标文件：`internal/repository/todo_memory.go`

接口 `TodoRepository` 已给出。请实现 `MemoryTodoRepository`：

- 用 `map` 存数据，注意并发（`sync.RWMutex`）
- `Create`：生成 id、写入时间戳
- `Get` / `Update` / `Delete`：不存在时返回 `ErrNotFound`
- `List`：返回当前全部条目

提示：结构体里预留了 `now`、`newID` 函数字段，便于单测注入；先让 `go test ./internal/repository` 变绿。

### 3.2 Service

目标文件：`internal/service/todo.go`

- 校验：id / title trim 后不能为空 → `ErrInvalidInput`
- `Create`：默认 `Done=false`
- `Update`：按 `UpdateTodoInput` 的指针字段做局部更新（区分「未传」与 `false`）
- 其余委托 repository；把 `ErrNotFound` 透传出去

先让 `go test ./internal/service` 变绿。

### 3.3 Handler + 路由

目标文件：`internal/handler/todo.go`、`routes.go`、`response.go`

- 解析 JSON body / `r.PathValue("id")`
- 把 service 错误映射为 400 / 404 / 500
- 创建成功 `201`，删除成功 `204`
- 路由表已在 `NewRouter` 中列出，请对照补全 handler 行为

端到端自测示例：

```bash
curl -s -X POST http://localhost:8080/api/todos \
  -H 'Content-Type: application/json' \
  -d '{"title":"学习分层","content":"先写 repository"}'

curl -s http://localhost:8080/api/todos
```

相关测试：`TestTodoHandlerCRUDFlow`。

---

## 步骤 4：请求日志中间件

目标文件：`internal/middleware/logging.go`、`handler/routes.go`

实现 `RequestLogger`：

- 包装 `ResponseWriter` 以记录真实状态码
- 输出 method、path、status、耗时（可用 `log/slog`）
- 在 `NewRouter` 里包裹整个 mux

自测：发几次请求，看标准输出是否有结构化日志。

---

## 步骤 5：测试驱动收尾

1. 跑全量：`make test`（或 `go test ./...`），全部通过才算完成。
2. 不要为了「变绿」而删断言；应补实现。
3. 可选加练：把内存仓库换成 SQLite，只改 `repository` 实现与组装处。

---

## 刻意不提供的内容

- 完整 CRUD / 中间件源码答案
- 可直接粘贴的「标准实现」

卡住时：重读本文件对应步骤 + 源码里的中文 TODO，对照失败的测试名推断期望行为。
