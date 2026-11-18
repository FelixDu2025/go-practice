package main

import (
	"context"
	"errors"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/FelixDu2025/go-practice/internal/config"
	"github.com/FelixDu2025/go-practice/internal/handler"
	"github.com/FelixDu2025/go-practice/internal/repository"
	"github.com/FelixDu2025/go-practice/internal/service"
)

// TODO: 按 EXERCISES.md 步骤 1 理解并完善依赖组装与优雅关闭。
// 当前骨架可以编译启动，但业务层仍是未实现的 stub。
func main() {
	logger := slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelInfo}))
	slog.SetDefault(logger)

	cfg, err := config.Load()
	if err != nil {
		logger.Error("load config", "error", err)
		os.Exit(1)
	}

	// TODO: 组装依赖：MemoryTodoRepository → TodoService → NewRouter
	repo := repository.NewMemoryTodoRepository()
	todoSvc := service.NewTodoService(repo)
	router := handler.NewRouter(todoSvc)

	srv := &http.Server{
		Addr:              cfg.Addr,
		Handler:           router,
		ReadHeaderTimeout: 5 * time.Second,
	}

	go func() {
		logger.Info("server starting", "addr", cfg.Addr)
		if err := srv.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			logger.Error("server failed", "error", err)
			os.Exit(1)
		}
	}()

	// TODO: 监听 SIGINT/SIGTERM，在 ShutdownTimeout 内优雅关闭
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGINT, syscall.SIGTERM)
	sig := <-stop
	logger.Info("shutdown signal received", "signal", sig.String())

	ctx, cancel := context.WithTimeout(context.Background(), cfg.ShutdownTimeout)
	defer cancel()

	if err := srv.Shutdown(ctx); err != nil {
		logger.Error("graceful shutdown failed", "error", err)
		os.Exit(1)
	}
	logger.Info("server stopped")
}
