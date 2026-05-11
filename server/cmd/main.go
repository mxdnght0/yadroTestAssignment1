package main

import (
	"log/slog"
	"os"
	"yadroTestAssignment/server/internal/application/service"
	"yadroTestAssignment/server/internal/config"
	"yadroTestAssignment/server/internal/infrastructure/file"
	"yadroTestAssignment/server/internal/logger"
	"yadroTestAssignment/server/internal/presentation/router"
)

func main() {
	log := logger.New()

	cfg := config.NewConfig()
	storage := file.NewStorage(cfg)
	svc := service.NewDNS(storage)
	r := router.NewRouter(svc, log)

	log.Info("starting server", slog.String("port", cfg.GetPort()))
	if err := r.Run(":" + cfg.GetPort()); err != nil {
		log.Error("server failed", slog.Any("error", err))
		os.Exit(1)
	}
}
