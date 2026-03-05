package auth

import "github.com/gin-gonic/gin"

func RegisterAuthRoute(r *gin.Engine, ctrl *Controller) {
	auth := r.Group("/auth")

	auth.POST("/token", ctrl.Authenticate())
	auth.POST("/refresh", ctrl.RefreshToken())
}
