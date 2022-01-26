package controllers

import (
	context "context"
	http "net/http"
	time "time"

	gin "github.com/gin-gonic/gin"
	validator "github.com/go-playground/validator/v10"
	bson "go.mongodb.org/mongo-driver/bson"
	mongo "go.mongodb.org/mongo-driver/mongo"

	Database "mLibAPI/src/database"
	// Models "mLibAPI/src/models"

	Helpers "mLibAPI/src/helpers"
)

var BookCollection *mongo.Collection = Database.GetCollection(Database.DB, "books")
var Validate = validator.New()

func ServeIndex(c *gin.Context) {
	c.HTML(http.StatusOK, "index.html", gin.H{})
}

func GetBooks(c *gin.Context) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	var books []bson.M
	defer cancel()

	cursor, err := BookCollection.Find(ctx, bson.M{})
	Helpers.PrintError(err)

	err = cursor.All(ctx, &books)
	Helpers.PrintError(err)

	c.JSON(http.StatusOK, books)

}
