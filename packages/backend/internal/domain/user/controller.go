package user

import (
	"errors"
	"log"
	"net/http"

	"github.com/0xlebogang/shapas/internal/domain/models"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
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
		var json models.User
		_ = ctx.BindJSON(&json)

		user, err := c.service.CreateUser(&json)
		if err != nil {
			if errors.Is(err, gorm.ErrDuplicatedKey) {
				ctx.JSON(http.StatusConflict, gin.H{
					"message": "User with provided email already exists",
				})
				return
			}

			log.Printf("User creation failed: %v\n", err)
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"message": "Failed to create user",
			})
			return
		}

		ctx.JSON(http.StatusOK, gin.H{
			"user": user.ToResponse(),
		})
	}
}
