package user

import "github.com/gin-gonic/gin"

func RegisterUserRoute(r *gin.Engine, ctrl *Controller) {
	route := r.Group("/user")

	route.POST("", ctrl.CreateUser())
}
