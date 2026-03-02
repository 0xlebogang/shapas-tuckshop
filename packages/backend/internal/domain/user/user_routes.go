package user

import "github.com/gin-gonic/gin"

func RegisterUserRoutes(r gin.RouterGroup) {
	users := r.Group("/user")

	users.POST("")
}
