package auth

import (
	"time"

	"github.com/YarnoVdW/StudyAPI/go-service/cmd/service/config"
	"github.com/YarnoVdW/StudyAPI/go-service/cmd/service/model"
	jwtapple2 "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
)

func SetupAuth() (*jwtapple2.GinJWTMiddelware, error) {
	authMiddleware, err := jwtapple2.New(&jwtapple2.GinJWTMiddelware{
		Realm:           "apistudy",
		Key:             []byte(config.Key),
		Timeout:         time.Hour * 24,     //life time of generated token
		MaxRefresh:      time.Hour,          // allows clients to refresh their token until its value has passed
		IdentityKey:     config.IdentityKey, // string value to identify elements from claims
		PayloadFunc:     payload,            //will be called during login
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
func payload(data interface{}) jwtapple2.MapClaims {
	if v, ok := data.(*model.User); ok {
		return jwtapple2.MapClaims{
			config.IdentityKey: v.ID,
		}
	}
	return jwtapple2.MapClaims{}
}

// callback function used to identify data using the claims information
func identityHandler(c *gin.Context) interface{} {
	claims := jwtapple2.ExtractClaims(c)
	var user model.User
	config.GetDb().Where("id = ?", claims[config.IdentityKey]).First(&user)

	return user
}

// callback function that performs the authentication based on the login info
func authenticator(c *gin.Context) (interface{}, error) {
	var loginVals model.User

	if err := c.ShouldBind(&loginVals); err != nil {
		return "", jwtapple2.ErrMissingLoginValues
	}
	var result model.User
	config.GetDb().Where("username = ? AND password = ?",
		loginVals.Username, loginVals.Password).First(&result)

	if result.ID == 0 {
		return nil, jwtapple2.ErrFailedAuthentication
	}

	return &result, nil
}
