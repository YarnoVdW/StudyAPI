package auth

import (
	"time"

	"github.com/YarnoVdW/StudyAPI/config"
	jwtapple2 "github.com/appleboy/gin-jwt/v2"
)

func SetupAuth() (*jwtapple2.GinJWTMiddelware, error) {
	authMiddleware, err := jwtapple2.New(&jwtapple2.GinJWTMiddelware{
		Realm:           "apistudy",
		Key:             []byte(config.Key),
		Timeout:         time.Hour * 24,
		MaxRefresh:      time.Hour,
		IdentityKey:     config.IdentityKey,
		PayloadFunc:     payload,
		IdentityHandler: identityHandler,
		Authenticator:   authenticator,
		Authorizator:    authorizator,
		Unauthorized:    unauthorized,
		LoginResponse:   loginResponse,
		TokenLookup:     "header: Authorization, query: token, cookie: jwtapple2",
		TokenHeadName:   "Bearer",
		TimeFunc:        time.Now,
	})
	return authMiddleware, err
}
