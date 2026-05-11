package handlers

import (
	"errors"
	"log/slog"
	"net"
	"net/http"
	"yadroTestAssignment/server/internal/application/contracts"
	"yadroTestAssignment/server/internal/application/service"
	"yadroTestAssignment/server/internal/presentation"

	"github.com/gin-gonic/gin"
)

func NewDeleteDNSHandler(svc contracts.DNSServer, log *slog.Logger) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		dns := ctx.Query("dns")
		if dns == "" {
			ctx.AbortWithStatusJSON(http.StatusBadRequest, presentation.DNSNotFound)
			log.Warn("dns param missing")
			return
		}

		if net.ParseIP(dns) == nil {
			ctx.AbortWithStatusJSON(http.StatusBadRequest, presentation.DNSInvalid)
			log.Warn("invalid dns address", slog.String("dns", dns))
			return
		}

		if err := svc.DeleteDNS(dns); err != nil {
			if errors.Is(err, service.ErrFileIsNotCreated) {
				ctx.AbortWithStatusJSON(http.StatusNotFound, presentation.FileNotFound)
				log.Warn("resolv.conf not found")
				return
			}
			if errors.Is(err, service.ErrDNSNotFound) {
				ctx.AbortWithStatusJSON(http.StatusNotFound, presentation.DNSNotFound)
				log.Warn("dns not found", slog.String("dns", dns))
				return
			}
			ctx.AbortWithStatusJSON(http.StatusInternalServerError, presentation.InternalError)
			log.Error("failed to delete dns", slog.String("dns", dns), slog.Any("error", err))
			return
		}

		log.Info("dns deleted", slog.String("dns", dns))
		ctx.Status(http.StatusOK)
	}
}
