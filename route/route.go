package route

import (
	"log"
	"net/http"

	"github.com/YarnoVdW/StudyAPI/go-service/cmd/service/auth"
	"github.com/YarnoVdW/StudyAPI/go-service/cmd/service/controller"
	"github.com/gin-gonic/gin"
)

func SetupRoutes() *gin.Engine {
	router := gin.Default()
	authMiddleware, err := auth.SetupAuth()

	if err != nil {
		log.Fatal("JWT error: " + err.Error())
	}
	router.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "Welcome to my Study app")
	})

	v1 := router.Group("/v1")
	{
		v1.POST("/login", authMiddleware.LoginHandler)
		v1.POST("/register", controller.RegisterEndPoint)

		studyItem := v1.Group("studyItem")
		{
			studyItem.POST("/create", authMiddleware.MiddlewareFunc(), controller.CreateStudyItem)
			studyItem.GET("/all", authMiddleware.MiddlewareFunc(), controller.FetchAllStudyitems)
			studyItem.GET("/get/:id", authMiddleware.MiddlewareFunc(), controller.FetchSingleTask)

		}
	}

	authorization := router.Group("/auth")
	authorization.GET("/refresh_token", authMiddleware.RefreshHandler)
	return router
}
