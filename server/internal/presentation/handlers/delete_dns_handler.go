package handlers

import (
	"errors"
	"log"
	"net"
	"net/http"
	"yadroTestAssignment/server/internal/application/contracts"
	"yadroTestAssignment/server/internal/application/service"
	"yadroTestAssignment/server/internal/presentation"

	"github.com/gin-gonic/gin"
)

func NewDeleteDNSHandler(DNSServer contracts.DNSServer) gin.HandlerFunc {
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

		err := DNSServer.DeleteDNS(dns)
		if err != nil {
			if errors.Is(err, service.ErrFileIsNotCreated) {
				ctx.AbortWithStatusJSON(http.StatusNotFound, presentation.FileNotFound)
				log.Println("file is not created")
				return
			}
			if errors.Is(err, service.ErrDNSNotFound) {
				ctx.AbortWithStatusJSON(http.StatusNotFound, presentation.DNSNotFound)
				log.Println("dns not found:", dns)
				return
			}

			ctx.AbortWithStatusJSON(http.StatusInternalServerError, presentation.InternalError)
			log.Println("internal server error:", err)
			return
		}

		ctx.Status(http.StatusOK)
	}
}
