package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func ServeIndex(c *gin.Context) {
	c.String(http.StatusOK, "hello world")
}
