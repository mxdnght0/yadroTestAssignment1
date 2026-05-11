package handlers

import (
	"errors"
	"log"
	"net/http"
	"yadroTestAssignment/server/internal/application/contracts"
	"yadroTestAssignment/server/internal/application/service"
	"yadroTestAssignment/server/internal/presentation"

	"github.com/gin-gonic/gin"
)

func NewGetAllDNSHandler(DNSServer contracts.DNSServer) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		lines, err := DNSServer.GetAllDNS()
		if err != nil {
			if errors.Is(err, service.ErrFileIsNotCreated) {
				ctx.AbortWithStatusJSON(http.StatusNotFound, presentation.DNSNotFound)
				log.Println("file is not created")
				return
			}

			ctx.AbortWithStatusJSON(http.StatusInternalServerError, presentation.InternalError)
			log.Println("internal server error:", err)
			return
		}

		if lines == nil {
			lines = []string{}
		}

		ctx.JSON(http.StatusOK, gin.H{
			"dns_lines": lines,
		})
	}
}
