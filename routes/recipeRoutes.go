package routes

import (
    controller "github.com/junaama/hackpserver/controllers"

    "github.com/gin-gonic/gin"
)

//UserRoutes function
func RecipeRoutes(incomingRoutes *gin.Engine) {
	incomingRoutes.GET("/recipes", controller.GetMyRecipe())
	incomingRoutes.DELETE("/recipes/:id", controller.DeleteMyRecipe())
	incomingRoutes.GET("/recipes/:id", controller.GetOneRecipe())
}