package core

import (
	// . "GenesisDAT/src/config"
	controllers "mLib/src/controllers"

	"github.com/gin-gonic/gin"
)

func Init() {
	App := gin.Default()

	App.GET("/", controllers.ServeIndex)
	App.Run(":7223")
}
