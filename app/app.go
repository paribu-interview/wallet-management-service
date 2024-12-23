package app

import (
	"context"
	"errors"
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/safayildirim/wallet-management-service/internal/wallet"
	"github.com/safayildirim/wallet-management-service/pkg/config"
	"github.com/safayildirim/wallet-management-service/pkg/db"
	"gorm.io/gorm"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

type Handler interface {
	RegisterRoutes(e *echo.Group)
}

type App struct {
	Config   config.Config
	DB       *gorm.DB
	Server   *echo.Echo
	Handlers []Handler
}

func New() *App {
	// Load configuration
	cfg := config.New()

	// Initialize GORM DB connection
	dbInstance, err := db.NewConnection(cfg.Postgres)
	if err != nil {
		panic(err)
	}

	// Create Echo instance
	server := echo.New()

	server.HTTPErrorHandler = func(err error, c echo.Context) {
		code := http.StatusInternalServerError
		var he *echo.HTTPError
		if errors.As(err, &he) {
			code = he.Code
		}
		c.JSON(code, err)
	}

	// Configure middleware
	server.Use(middleware.Logger())
	server.Use(middleware.Recover())

	var handlers []Handler

	walletRepository := wallet.NewRepository(dbInstance)
	walletService := wallet.NewService(walletRepository)
	walletHandler := wallet.NewHandler(walletService)

	handlers = append(handlers, walletHandler)

	return &App{Config: *cfg, DB: dbInstance, Server: server, Handlers: handlers}
}

func (a *App) Run() error {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	route := a.Server.Group("/api")

	for _, handler := range a.Handlers {
		handler.RegisterRoutes(route)
	}

	// Start the server in a goroutine
	go func() {
		port := fmt.Sprintf(":%d", a.Config.Http.Port)
		log.Println("Starting server on :", port)
		if err := a.Server.Start(port); err != nil && !errors.Is(err,
			http.ErrServerClosed) {
			log.Fatalf("Could not start server: %v", err)
		}
	}()

	// Wait for interrupt signal to gracefully shut down the server
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM)
	<-quit
	log.Println("Shutting down server...")

	// Gracefully shut down the Echo server with a timeout
	ctx, cancel = context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := a.Server.Shutdown(ctx); err != nil {
		log.Fatalf("Server forced to shutdown: %v", err)
	}

	// Close database connection
	sqlDB, err := a.DB.DB()
	if err == nil {
		sqlDB.Close()
	}

	log.Println("Server exited gracefully")

	return nil
}
