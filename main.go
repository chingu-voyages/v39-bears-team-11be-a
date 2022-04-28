package main

import (
    "os"
	"github.com/UfiairENE/litetalk/v39-bears-team-11be-a/routes"
	"github.com/UfiairENE/litetalk/v39-bears-team-11be-a/middleware"
	
    "github.com/gin-gonic/gin"
    _ "github.com/heroku/x/hmetrics/onload"
)

func main() {
    port := os.Getenv("PORT")

    if port == "" {
        port = "8080"
    }

    router := gin.New()
    router.Use(gin.Logger())
    routes.UserRoutes(router)

    router.Use(middleware.Authentication())

    
    router.Run(":" + port)
}