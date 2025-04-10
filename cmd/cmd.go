package cmd

import (
	"gogofly/conf"
	"gogofly/global"
	"gogofly/router"
	"gogofly/utils"
)

func Start() {
	var initErr error
	// = init system config； viper config
	conf.InitConfig()
	// = init Logger component
	global.Logger = conf.InitLogger()
	// = init DB conn
	db, err := conf.InitDB()
	global.DB = db
	if err != nil {
		initErr = utils.AppendError(initErr, err)
	}
	if initErr != nil {
		if global.Logger != nil {
			global.Logger.Error(initErr.Error())
		}
		panic(initErr.Error())
	}

	// = init router
	router.InitRouter()
}

func Clean() {

}
