package src

import (
	Config "mLibAPI/src/config"
	Router "mLibAPI/src/router"
)

func App() {

	App := Router.CreateServer()

	App.Run(Config.Port)
}

func Init() {
	Config.Init()

	Logs()
	App()

}
