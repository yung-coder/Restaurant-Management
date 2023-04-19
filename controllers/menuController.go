package controllers

import (
	"backend/database"
	"context"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

var menuColletion *mongo.Collection = database.OpenCollection(database.Client, "menu")

func GetMenus() gin.HandlerFunc {
	return func(c *gin.Context) {
		var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
		result, err := menuColletion.Find(context.TODO(), bson.M{})
		defer cancel()
		if err != nil {
			c.JSONP(http.StatusInternalServerError, gin.H{"error": "error occcured while menu items"})
		}

		var allMenus []bson.M

		if err = result.All(ctx, &allMenus); err != nil {
			log.Fatal(err)
		}
		c.JSON(http.StatusOK, allMenus)
	}
}

func GetMenu() gin.HandlerFunc {
	return func(ctx *gin.Context) {

	}
}

func CreateMenu() gin.HandlerFunc {
	return func(ctx *gin.Context) {

	}
}

func UpdateMenu() gin.HandlerFunc {
	return func(ctx *gin.Context) {

	}
}
