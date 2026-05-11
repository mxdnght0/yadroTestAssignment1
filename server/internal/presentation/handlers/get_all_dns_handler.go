package handlers

import (
	"errors"
	"log/slog"
	"net/http"
	"yadroTestAssignment/server/internal/application/contracts"
	"yadroTestAssignment/server/internal/application/service"
	"yadroTestAssignment/server/internal/presentation"

	"github.com/gin-gonic/gin"
)

func NewGetAllDNSHandler(svc contracts.DNSServer, log *slog.Logger) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		lines, err := svc.GetAllDNS()
		if err != nil {
			if errors.Is(err, service.ErrFileIsNotCreated) {
				ctx.AbortWithStatusJSON(http.StatusNotFound, presentation.FileNotFound)
				log.Warn("resolv.conf not found")
				return
			}
			ctx.AbortWithStatusJSON(http.StatusInternalServerError, presentation.InternalError)
			log.Error("failed to get dns list", slog.Any("error", err))
			return
		}

		if lines == nil {
			lines = []string{}
		}

		log.Info("dns list retrieved", slog.Int("count", len(lines)))
		ctx.JSON(http.StatusOK, gin.H{"dns_lines": lines})
	}
}
