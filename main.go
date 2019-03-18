package main

import (
	. "github.com/LG_Tokens/Logger"
	ctrl "github.com/LG_Tokens/controller"
	db "github.com/LG_Tokens/db"
	"github.com/LG_Tokens/routers"
	. "github.com/LG_Tokens/util"
	"github.com/robfig/cron"
	"github.com/spf13/viper"
)

func main() {
	InitViper()
	db.InitDB()
	// ctrl.GetRateByOutApi()
	c := cron.New()
	c.AddFunc("@hourly", ctrl.GetRateByOutApi)
	c.Start()

	router := routers.InitRouter()
	Log.Info("start server at:" + viper.GetString("server.port"))
	router.Run(":" + viper.GetString("server.port"))
}
