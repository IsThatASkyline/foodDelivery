package api

import (
	"context"
	"errors"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

// App основная структура для приложения
type App struct {
	engine *gin.Engine
}

func NewApp() *App {
	engine := gin.Default()
	engine.ContextWithFallback = true

	return &App{
		engine: engine,
	}
}

type Module interface {
	RegisterRoutes(*gin.RouterGroup)
}

func (a *App) SetupRoutes(modules ...Module) {
	api := a.engine.Group("/api")

	for _, module := range modules {
		module.RegisterRoutes(api)
	}
}

func (a *App) StartServe(ctx context.Context) {
	server := &http.Server{
		Addr:           ":8080",
		Handler:        a.engine,
		ReadTimeout:    time.Second * 60,
		WriteTimeout:   time.Second * 60,
		MaxHeaderBytes: 1 << 20,
	}

	// Запускаем сервер
	go func() {
		if err := server.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			log.Fatalf("Service running error: %s", err.Error())
		}
	}()

	<-ctx.Done()
	log.Println("Server stop start")

	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	if err := server.Shutdown(ctx); err != nil {
		log.Printf("Server forced to shutdown: %s", err.Error())
	}
	log.Println("Server stopped")
}
