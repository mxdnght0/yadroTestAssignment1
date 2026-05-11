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

func NewSaveDNSHandler(svc contracts.DNSServer, log *slog.Logger) gin.HandlerFunc {
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

		if err := svc.SaveDNS(dns); err != nil {
			if errors.Is(err, service.ErrDNSAlreadyExists) {
				ctx.AbortWithStatusJSON(http.StatusConflict, presentation.DNSAlreadyExists)
				log.Warn("dns already exists", slog.String("dns", dns))
				return
			}
			ctx.AbortWithStatusJSON(http.StatusInternalServerError, presentation.InternalError)
			log.Error("failed to save dns", slog.String("dns", dns), slog.Any("error", err))
			return
		}

		log.Info("dns added", slog.String("dns", dns))
		ctx.Status(http.StatusCreated)
	}
}
