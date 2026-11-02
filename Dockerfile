# 可选：多阶段构建，便于本地或容器中运行练习服务。
FROM golang:1.22-alpine AS builder
WORKDIR /src
COPY go.mod ./
COPY . .
RUN CGO_ENABLED=0 go build -o /out/server ./cmd/server

FROM alpine:3.20
RUN adduser -D -H app
USER app
WORKDIR /app
COPY --from=builder /out/server /app/server
EXPOSE 8080
ENV ADDR=:8080
ENTRYPOINT ["/app/server"]
