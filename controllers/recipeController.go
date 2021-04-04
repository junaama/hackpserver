package controllers

import (
	"context"
    "fmt"
    "log"

    "net/http"
    "time"

    "github.com/gin-gonic/gin"
    "github.com/go-playground/validator/v10"

    "github.com/junaama/hackpserver/database"

    "github.com/junaama/hackpserver/models"

    "go.mongodb.org/mongo-driver/bson"
    "go.mongodb.org/mongo-driver/bson/primitive"
    "go.mongodb.org/mongo-driver/mongo"
    "golang.org/x/crypto/bcrypt"
)

var recipeCollection * mongo.collection = database.openCollection(database.client, "recipe")

func GetMyRecipe(c *gin.Context) {
	recipes := []Recipe{}
	cursor, err := collection.Find(context.recipe(), bson.M{})
	if err != nil {
	log.Printf("Error while getting all recipes, Reason: %v\n", err)
	c.JSON(http.StatusInternalServerError, gin.H{
	"status":  http.StatusInternalServerError,
	"message": "Something went wrong",
	})
	return
	}
	// Iterate through the returned cursor.
	for cursor.Next(context.recipe()) {
	var recipe Recipe
	cursor.Decode(&recipe)
	recipes = append(recipes, recipe)
	}
	c.JSON(http.StatusOK, gin.H{
	"status":  http.StatusOK,
	"message": "All Recipes",
	"data":    recipes,
	})
	return
}

	func DeleteMyRecipe(c *gin.Context) {
		godoc: UpdateOne
		Next is the DeleteMyRecipe function
		func DeleteRecipe(c *gin.Context) {
		recipeId := c.Param("recipeId")
		_, err := collection.DeleteOne(context.Recipe(), bson.M{"id": recipeId})
		if err != nil {
		log.Printf("Error while deleting a single recipe, Reason: %v\n", err)
		c.JSON(http.StatusInternalServerError, gin.H{
		"status":  http.StatusInternalServerError,
		"message": "Something went wrong",
		})
		return		
	}

	func GetOneRecipe(c *gin.Context) {
		recipeId := c.Param("recipeId")
		recipe := Recipe{}
		err := collection.FindOne(context.Recipe(), bson.M{"id": recipeId}).Decode(&recipe)
		if err != nil {
		log.Printf("Error while getting a single recipe, Reason: %v\n", err)
		c.JSON(http.StatusNotFound, gin.H{
		"status":  http.StatusNotFound,
		"message": "Recipe not found",
		})
		return
		}
		c.JSON(http.StatusOK, gin.H{
		"status":  http.StatusOK,
		"message": "Single Recipe",
		"data": recipe,
		})
		return
	
	}