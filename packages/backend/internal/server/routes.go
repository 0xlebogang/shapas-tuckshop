package server

import (
	"github.com/0xlebogang/shapas/internal/config"
	"github.com/0xlebogang/shapas/internal/domain/auth"
	"github.com/0xlebogang/shapas/internal/domain/user"
	"github.com/0xlebogang/shapas/internal/middlewares"
	"github.com/0xlebogang/shapas/internal/token"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type Routes struct {
	config     *config.Config
	db         *gorm.DB
	router     *gin.Engine
	middleware *middlewares.Middleware
}

func NewRoutes(config *config.Config, db *gorm.DB, router *gin.Engine, middleware *middlewares.Middleware) *Routes {
	return &Routes{config: config, db: db, router: router, middleware: middleware}
}

func (r *Routes) SetupUserRoutes() {
	userRepo := user.NewRepository(r.db)
	userService := user.NewService(userRepo)
	userCtrl := user.NewController(userService)

	user.Register(r.router, userCtrl, r.middleware)
}

func (r *Routes) SetupAuthRoutes() {
	userRepo := user.NewRepository(r.db)
	authRepo := auth.NewRepository(r.db)

	tokenFactory := token.NewTokenFactory(r.config.JWTSecret)
	authService := auth.NewService(tokenFactory, userRepo, authRepo)
	authCtrl := auth.NewController(authService)

	auth.RegisterAuthRoute(r.router, authCtrl)
}
