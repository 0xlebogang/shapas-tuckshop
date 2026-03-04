package middlewares

import "github.com/0xlebogang/shapas/internal/token"

type Middleware struct {
	tokenFactory *token.Factory
}

func New(tokenFactory *token.Factory) *Middleware {
	return &Middleware{tokenFactory: tokenFactory}
}
