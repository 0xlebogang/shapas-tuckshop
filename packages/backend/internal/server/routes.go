package server

import (
	"github.com/0xlebogang/shapas/internal/domain/user"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type Routes struct {
	router *gin.Engine
	db     *gorm.DB
}

func NewRoutes(router *gin.Engine, db *gorm.DB) *Routes {
	return &Routes{router: router, db: db}
}

func (r *Routes) SetupUserRoutes() {
	userRepo := user.NewRepository(r.db)
	userService := user.NewService(userRepo)
	userCtrl := user.NewController(userService)

	user.RegisterUserRoute(r.router, userCtrl)
}
