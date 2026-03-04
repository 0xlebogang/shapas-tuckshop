package server

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func (s *Server) healthCheck() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		db, err := s.db.DB()
		if err != nil || db.Ping() != nil {
			ctx.JSON(http.StatusServiceUnavailable, gin.H{
				"status": "Unhealthy",
				"db":     "disconnected",
			})
			return
		}

		ctx.JSON(http.StatusOK, gin.H{
			"status":    "UP",
			"version":   "0.0.0",
			"uptime":    time.Since(s.startTime).String(),
			"timestamp": time.Now().Format(time.RFC1123),
			"dependecies": gin.H{
				"DB": "OK",
			},
		})
	}
}
