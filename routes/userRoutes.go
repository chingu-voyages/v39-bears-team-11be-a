package routes

import (

  controller "github.com/UfiairENE/litetalk/v39-bears-team-11be-a/controllers"

    "github.com/gin-gonic/gin"
)

//UserRoutes function
func UserRoutes(incomingRoutes *gin.Engine) {
    incomingRoutes.POST("/users/signup", controller.SignUp())
    incomingRoutes.POST("/users/login", controller.Login())
}
