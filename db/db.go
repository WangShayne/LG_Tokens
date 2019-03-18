package db

import (
	"database/sql"
	"fmt"

	. "github.com/LG_Tokens/Logger"
	. "github.com/LG_Tokens/common"
	_ "github.com/go-sql-driver/mysql"
	"github.com/spf13/viper"
)

var DB *sql.DB

func InitDB() error {
	var DBconfig DBconfig

	DBconfig.Host = viper.GetString("db.Host")
	DBconfig.Port = viper.GetString("db.Port")
	DBconfig.User = viper.GetString("db.User")
	DBconfig.Password = viper.GetString("db.Password")
	DBconfig.DBname = viper.GetString("db.DBname")

	var err error
	DB, err = sql.Open("mysql", fmt.Sprintf("%v:%v@tcp(%v:%v)/%v?parseTime=true", DBconfig.User, DBconfig.Password, DBconfig.Host, DBconfig.Port, DBconfig.DBname))
	if err != nil {
		Log.Fatal(err)
	}

	DB.SetMaxIdleConns(20)
	DB.SetMaxOpenConns(100)

	if err := DB.Ping(); err != nil {
		Log.Fatal(err)
	}
	return nil
}
