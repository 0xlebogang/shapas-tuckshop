package server

import (
	"fmt"
	"net/http"

	"github.com/0xlebogang/shapas/backend/internal/config"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type Server struct {
	config *config.Env
	db     *gorm.DB
	router *gin.Engine
}

func New(config *config.Env, db *gorm.DB) *Server {
	return &Server{
		config: config,
		db:     db,
		router: gin.Default(),
	}
}

func (s *Server) createHttpServer() *http.Server {
	return &http.Server{
		Addr:    fmt.Sprintf(":%s", s.config.Port),
		Handler: s.router.Handler(),
	}
}

func (s *Server) Start() error {
	// Attatch healthcheck endpoint
	s.router.GET("/healthz", s.healthCheckEndpoint())

	server := s.createHttpServer()
	return server.ListenAndServe()
}

func (s *Server) healthCheckEndpoint() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"status": "healthy",
		})
	}
}
