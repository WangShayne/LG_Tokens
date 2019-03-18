package util

import (
	"fmt"
	"os"
	"strings"

	"github.com/spf13/viper"
)

const CONFIG = "config"

func InitViper() {
	viper.New()
	viper.SetEnvPrefix(CONFIG)
	viper.AutomaticEnv()
	replacer := strings.NewReplacer(".", "_")
	viper.SetEnvKeyReplacer(replacer)
	viper.SetConfigName(CONFIG)
	viper.AddConfigPath("./")

	err := viper.ReadInConfig()
	if err != nil {
		fmt.Println(fmt.Errorf("Fatal error when reading %s config file:%s", CONFIG, err))
		os.Exit(1)
	}
}
