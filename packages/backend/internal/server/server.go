package server

import (
	"fmt"
	"net/http"
	"time"

	"github.com/0xlebogang/shapas/internal/config"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type Server struct {
	config    *config.Config
	db        *gorm.DB
	router    *gin.Engine
	startTime time.Time
}

func New(config *config.Config, db *gorm.DB) *Server {
	return &Server{
		config:    config,
		db:        db,
		router:    gin.Default(),
		startTime: time.Now(),
	}
}

func (s *Server) Start() error {
	s.router.GET("/healthz", s.healthCheck())

	routes := NewRoutes(s.router, s.db)
	routes.SetupUserRoutes()

	svr := s.createHttpServer()
	return svr.ListenAndServe()
}

func (s *Server) createHttpServer() *http.Server {
	return &http.Server{
		Addr:    fmt.Sprintf(":%s", s.config.Port),
		Handler: s.router.Handler(),
	}
}
