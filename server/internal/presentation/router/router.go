package router

import (
	"yadroTestAssignment/server/internal/application/contracts"
	handlers2 "yadroTestAssignment/server/internal/presentation/handlers"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func NewRouter(svc contracts.DNSServer) *gin.Engine {
	r := gin.Default()
	r.Use(cors.New(cors.Config{
		AllowAllOrigins: true,
		AllowMethods:    []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"},
		AllowHeaders:    []string{"Origin", "Authorization", "Content-Type"},
	}))
	r.Use(gin.Recovery())

	r.POST("/dns", handlers2.NewSaveDNSHandler(svc))
	r.DELETE("/dns", handlers2.NewDeleteDNSHandler(svc))
	r.GET("/dns", handlers2.NewGetAllDNSHandler(svc))

	return r
}
