package src

import (
	Config "mLibAPI/src/config"
	Database "mLibAPI/src/database"
	Router "mLibAPI/src/router"

	gin "github.com/gin-gonic/gin"
)

func App() {
	Database.Init()

	if Config.Mode == "PROD" {
		gin.SetMode(gin.ReleaseMode)
	}

	App := Router.CreateServer()

	App.Run(Config.Port)

	defer Database.Client.Disconnect(Database.Ctx)
}

func Init() {
	Logs()

	App()

}
