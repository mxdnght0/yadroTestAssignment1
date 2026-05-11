package main

import (
	"log"
	"yadroTestAssignment/server/internal/application/service"
	"yadroTestAssignment/server/internal/config"
	"yadroTestAssignment/server/internal/infrastructure/file"
	"yadroTestAssignment/server/internal/presentation/router"
)

func main() {
	cfg := config.NewConfig()
	storage := file.NewStorage(cfg)
	svc := service.NewDNS(storage)
	r := router.NewRouter(svc)

	log.Printf("Starting server at :%s", cfg.GetPort())
	if err := r.Run(":" + cfg.GetPort()); err != nil {
		log.Fatalf("Server failed: %v", err)
	}
}
