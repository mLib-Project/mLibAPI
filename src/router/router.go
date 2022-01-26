package router

import (
	Controllers "mLibAPI/src/controllers"

	// static "github.com/gin-contrib/static"
	cors "github.com/gin-contrib/cors"
	gin "github.com/gin-gonic/gin"
)

func CreateServer() *gin.Engine {

	R := gin.Default()
	R.Use(cors.Default())

	// MIDDLEWARES
	R.SetTrustedProxies([]string{"192.168.1.2"})

	R.GET("/api", Controllers.GetBooks)

	return R
}
