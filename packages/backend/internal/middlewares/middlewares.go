package middlewares

import (
	"github.com/0xlebogang/shapas/internal/token"
)

type Middleware struct {
	tokenFactory *token.TokenFactory
}

func New(tokenFactory *token.TokenFactory) *Middleware {
	return &Middleware{tokenFactory: tokenFactory}
}
