package models 

import (
	"time"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Recipe struct {
	ID            primitive.ObjectID `bson:"_id"`
    Name string `bson:"Name,omitempty"`
    Type string `bson:"Type,omitempty"`
    Yield int `bson:"Yield,omitempty"`
    Ingredients []string `bson:"Ingredients,omitempty"`
    CaloriesPerServing int `bson:"CaloriesPerServing,omitempty"`
    ProteinPerServing int `bson:"ProteinPerServing,omitempty"`
}