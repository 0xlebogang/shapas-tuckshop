package constants

import "time"

const (
	AuthorizationHeaderKey  = "Authorization"
	AuthorizationTypeBearer = "bearer"
	AuthorizationPayloadKey = "payload"

	AuthAccessTokenLifeTime  = 60 * time.Minute
	AuthRefreshTokenLifeTime = 7 * 24 * time.Hour
)
