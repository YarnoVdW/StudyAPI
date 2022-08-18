package main

import (
	"log"
	"os"

	"github.com/YarnoVdW/StudyAPI/go-service/cmd/service/config"
	"github.com/YarnoVdW/StudyAPI/go-service/cmd/service/migration"
	"github.com/YarnoVdW/StudyAPI/go-service/cmd/service/route"
	"github.com/gin-gonic/gin"
)

func init() {
	db := config.Init()
	migration.Migrate(db)
}

func main() {
	gin.SetMode(gin.ReleaseMode)
	router := route.SetupRoutes()

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	if err := router.Run(":" + port); err != nil {
		log.Panicf("error: %s", err)
	}
}
