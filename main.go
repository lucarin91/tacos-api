package main

import (
	"context"
	"fmt"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/dimfeld/httptreemux/v5"
	"github.com/jackc/pgx/v5/pgxpool"

	"github.com/lucarin91/tacos-api/internal/handler"
	"github.com/lucarin91/tacos-api/internal/middleware"
)

type Config struct {
	LogLevel slog.Level
	Port     int
	DbUrl    string
	secret   []byte
}

var cfg = Config{
	LogLevel: slog.LevelDebug,
	Port:     8080,
	DbUrl:    "postgres://postgres:postgres@localhost:5432",
	secret:   []byte("very secret key"),
}

func main() {
	initLog(cfg.LogLevel)

	pool := getDBPool()
	defer pool.Close()

	r := httptreemux.New()
	r = AddHandlers(r, pool)

	server := &http.Server{
		Handler:      r,
		Addr:         fmt.Sprintf(":%d", cfg.Port),
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 5 * time.Second,
	}

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt)

	go func() {
		slog.Info("Server start", "port", cfg.Port)
		err := server.ListenAndServe()
		if err != nil && err != http.ErrServerClosed {
			panic(err)
		}
	}()

	<-stop

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := server.Shutdown(ctx); err != nil {
		slog.Error("Server forced to stop", "error", err)
	}

	slog.Info("Server stopped")
}

func AddHandlers(router *httptreemux.TreeMux, pool *pgxpool.Pool) *httptreemux.TreeMux {
	router.Use(middleware.Log)

	api := router.NewGroup("/v1")

	api.GET("/token", handler.GetToken(cfg.secret))

	api.GET("/ingredients", handler.GetIngredients(pool))

	auth := middleware.Auth(cfg.secret)
	api.GET("/orders", auth(handler.GetOrders(pool)))
	api.GET("/orders/:id", auth(handler.GetOrder(pool)))
	api.POST("/orders/", auth(handler.CreateOrder(pool)))
	api.DELETE("/orders/:id", auth(handler.DeleteOrder(pool)))

	return router
}

func initLog(level slog.Level) {
	handler := slog.NewTextHandler(
		os.Stdout,
		&slog.HandlerOptions{Level: level},
	)
	logger := slog.New(handler)
	slog.SetDefault(logger)
}

func getDBPool() *pgxpool.Pool {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	pool, err := pgxpool.New(ctx, cfg.DbUrl)
	if err != nil {
		panic(err)
	}

	return pool
}
