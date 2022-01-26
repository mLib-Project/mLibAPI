package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func ServeIndex(c *gin.Context) {
	c.HTML(http.StatusOK, "index.html", gin.H{})
}

func GetBooks(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "pong",
	})
}
