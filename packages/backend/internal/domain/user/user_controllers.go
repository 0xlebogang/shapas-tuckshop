package user

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Controller struct {
	service *Service
}

func NewController(service *Service) *Controller {
	return &Controller{
		service: service,
	}
}

func (c *Controller) CreateUser() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var json User
		_ = ctx.BindJSON(json)

		user, err := c.service.CreateUser(&json)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"error": "Failed to create user",
			})
		}

		ctx.JSON(http.StatusCreated, gin.H{
			"user": user,
		})
	}
}
