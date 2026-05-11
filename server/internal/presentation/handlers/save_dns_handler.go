package handlers

import (
	"log"
	"net"
	"net/http"
	"yadroTestAssignment/server/internal/application/contracts"
	"yadroTestAssignment/server/internal/application/service"
	"yadroTestAssignment/server/internal/presentation"

	"github.com/gin-gonic/gin"
)

func NewSaveDNSHandler(DNSServer contracts.DNSServer) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		dns := ctx.Query("dns")
		if dns == "" {
			ctx.AbortWithStatusJSON(http.StatusBadRequest, presentation.DNSNotFound)
			log.Println("dns not found in query")
			return
		}

		if net.ParseIP(dns) == nil {
			ctx.AbortWithStatusJSON(http.StatusBadRequest, presentation.DNSInvalid)
			log.Println("invalid dns address:", dns)
			return
		}

		err := DNSServer.SaveDNS(dns)
		if err != nil {
			if err == service.ErrDNSAlreadyExists {
				ctx.AbortWithStatusJSON(http.StatusConflict, presentation.DNSAlreadyExists)
				log.Println("dns already exists:", dns)
				return
			}
			ctx.AbortWithStatusJSON(http.StatusInternalServerError, presentation.InternalError)
			log.Println("internal server error:", err)
			return
		}

		ctx.Status(http.StatusCreated)
	}
}
