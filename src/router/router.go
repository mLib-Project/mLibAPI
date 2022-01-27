package router

import (
	cors "github.com/gin-contrib/cors"
	static "github.com/gin-contrib/static"
	gin "github.com/gin-gonic/gin"

	Controllers "mLibAPI/src/controllers"
)

func CreateServer() *gin.Engine {

	R := gin.Default()
	R.Use(cors.Default())

	R.LoadHTMLGlob("public/*.html")
	R.Use(static.Serve("/", static.LocalFile("public/", false)))
	R.NoRoute(Controllers.ServeIndex)

	R.SetTrustedProxies([]string{"192.168.1.2"})
	R.GET("/api", Controllers.GetBooks)

	return R
}
