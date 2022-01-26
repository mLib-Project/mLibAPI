package src

import (
	Config "mLibAPI/src/config"
	Database "mLibAPI/src/database"
	Router "mLibAPI/src/router"
)

func App() {

	Database.ConnectDB()
	App := Router.CreateServer()

	App.Run(Config.Port)
	// defer Database.Client.Disconnect(Database.Ctx)
}

func Init() {
	Config.Init()

	Logs()
	App()

}
