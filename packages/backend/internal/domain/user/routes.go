package user

import (
	"github.com/0xlebogang/shapas/internal/middlewares"
	"github.com/gin-gonic/gin"
)

func Register(r *gin.Engine, ctrl *Controller, middleware *middlewares.Middleware) {
	route := r.Group("/user")

	route.POST("", ctrl.CreateUser())
	route.GET("", middleware.AuthMiddleware(), ctrl.GetUser())
}
