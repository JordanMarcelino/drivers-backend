package main

import (
	"github.com/JordanMarcelino/drivers-backend/internal/config"
	"github.com/JordanMarcelino/drivers-backend/internal/httpserver"
	"github.com/JordanMarcelino/drivers-backend/internal/pkg/logger"
)

func main() {
	cfg := config.InitConfig(".")

	logger.SetLogrusLogger(cfg)

	httpserver.StartGinHttpServer(cfg)
}
