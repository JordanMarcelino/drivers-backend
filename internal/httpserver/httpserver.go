package httpserver

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/JordanMarcelino/drivers-backend/internal/config"
	"github.com/JordanMarcelino/drivers-backend/internal/handler/httphandler"
	"github.com/JordanMarcelino/drivers-backend/internal/httpserver/middleware"
	"github.com/JordanMarcelino/drivers-backend/internal/pkg/database"
	"github.com/JordanMarcelino/drivers-backend/internal/pkg/logger"
	"github.com/JordanMarcelino/drivers-backend/internal/pkg/utils/validationutils"
	"github.com/JordanMarcelino/drivers-backend/internal/repository"
	"github.com/JordanMarcelino/drivers-backend/internal/usecase"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
	"github.com/shopspring/decimal"
)

func initServer(cfg *config.Config) *http.Server {
	db := database.InitStdLib(cfg)

	driverRepo := repository.NewDriverRepository(db)

	driverUseCase := usecase.NewDriverUseCase(driverRepo)

	driverHandler := httphandler.NewDriverHandler(driverUseCase)

	gin.SetMode(cfg.App.Environment)

	r := gin.New()
	r.ContextWithFallback = true

	registerValidators()

	middlewares := []gin.HandlerFunc{
		middleware.ErrorHandler(),
		middleware.Logger(),
		cors.New(cors.Config{
			AllowMethods:     []string{"*"},
			AllowHeaders:     []string{"*"},
			AllowAllOrigins:  true,
			AllowCredentials: true,
		}),
		gin.Recovery(),
	}
	r.Use(middlewares...)

	v1 := r.Group("/v1")
	{
		v1.GET("/salary/driver/list", driverHandler.FindAll)
	}

	srv := &http.Server{
		Addr:    fmt.Sprintf("%s:%d", cfg.HttpServer.Host, cfg.HttpServer.Port),
		Handler: r,
	}

	return srv
}

func StartGinHttpServer(cfg *config.Config) {
	srv := initServer(cfg)

	go func() {
		logger.Log.Info("running server on port :", cfg.HttpServer.Port)
		if err := srv.ListenAndServe(); err != nil {
			if !errors.Is(err, http.ErrServerClosed) {
				logger.Log.Fatal("error while server listen and serve: ", err)
			}
		}
		logger.Log.Info("server is not receiving new requests...")
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

	<-quit

	graceDuration := time.Duration(cfg.HttpServer.GracePeriod) * time.Second

	ctx, cancel := context.WithTimeout(context.Background(), graceDuration)
	defer cancel()

	logger.Log.Info("attempt to shutting down the server...")
	if err := srv.Shutdown(ctx); err != nil {
		logger.Log.Fatal("error shutting down server: ", err)
	}

	logger.Log.Info("http server is shutting down gracefully")
}

func registerValidators() {
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		v.RegisterTagNameFunc(validationutils.FormTagName)
		v.RegisterCustomTypeFunc(validationutils.DecimalType, decimal.Decimal{})
		v.RegisterValidation("shipment_cost_statuses", validationutils.ShipmentCostStatusValidator)
	}
}
