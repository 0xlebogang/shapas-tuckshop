package auth

import (
	"errors"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type Controller struct {
	service *Service
}

func NewController(service *Service) *Controller {
	return &Controller{service: service}
}

func (c *Controller) Authenticate() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var json AuthRequest
		_ = ctx.BindJSON(&json)

		tokens, err := c.service.Authenticate(&json)
		if err != nil {
			switch {
			case errors.Is(err, gorm.ErrRecordNotFound) || errors.Is(err, bcrypt.ErrMismatchedHashAndPassword):
				ctx.JSON(http.StatusBadRequest, gin.H{
					"message": "Invalid email or password",
				})
				return

			default:
				log.Printf("Failed to created tokens: %v", err)
				ctx.JSON(http.StatusInternalServerError, gin.H{
					"message": "An unexpected error occured",
				})
				return
			}
		}

		ctx.JSON(http.StatusOK, *tokens)
	}
}
