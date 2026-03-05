package user

import (
	"errors"
	"log"
	"net/http"

	"github.com/0xlebogang/shapas/internal/constants"
	"github.com/0xlebogang/shapas/internal/domain/models"
	"github.com/0xlebogang/shapas/internal/token"
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
				"message": "An unexpected error occured",
			})
			return
		}

		ctx.JSON(http.StatusOK, gin.H{
			"user": user.ToResponse(),
		})
	}
}

func (c *Controller) GetUser() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		payload := ctx.MustGet(constants.AuthorizationPayloadKey).(*token.UserClaims)
		user, err := c.service.GetUser(payload.UserID)
		if err != nil {
			switch {
			case errors.Is(err, gorm.ErrRecordNotFound):
				ctx.JSON(http.StatusNotFound, gin.H{
					"message": "User not found",
				})
				return

			default:
				log.Printf("Failed to get user: %v\n", err)
				ctx.JSON(http.StatusInternalServerError, gin.H{
					"message": "An unexpected error occured",
				})
				return
			}
		}

		ctx.JSON(http.StatusOK, gin.H{
			"user": user.ToResponse(),
		})
	}
}
