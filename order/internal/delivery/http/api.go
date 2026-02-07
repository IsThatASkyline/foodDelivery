package api

import (
	"context"
	"errors"
	"github.com/IsThatASkyline/foodDelivery/order/internal/delivery/http/v1/handlers"
	clientV1Routes "github.com/IsThatASkyline/foodDelivery/order/internal/delivery/http/v1/routes"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

// App основная структура для приложения
type App struct {
	engine   *gin.Engine
	handlers handlers.Handlers
}

func NewApp() *App {
	engine := gin.Default()
	engine.ContextWithFallback = true

	return &App{
		engine: engine,
	}
}

func (a *App) SetupRoutes() {
	api := a.engine.Group("/api")
	a.setupV1Routes(api)
}

func (a *App) setupV1Routes(router *gin.RouterGroup) {
	apiV1 := router.Group("/v1")

	clientApiV1 := apiV1.Group("")
	clientV1Routes.Setup(clientApiV1, a.handlers)
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
