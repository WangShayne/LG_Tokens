package db

import (
	"fmt"
	. "github.com/LG_Tokens/Logger"
	. "github.com/LG_Tokens/common"
)

func InsertRate(USDToCNYRate *USDToCNYRate) {
	sqlStr := "INSERT INTO rate(name,fBuyPri,update_time) VALUE(?,?,?) ON DUPLICATE KEY UPDATE `name` = ?,`fBuyPri` = ?,`update_time` = ?"
	stmt, err := DB.Prepare(sqlStr)
	if err != nil {
		fmt.Println(err)
	}
	defer func() {
		if stmt != nil {
			stmt.Close()
		}
	}()
	result, err := stmt.Exec(
		USDToCNYRate.Name,
		USDToCNYRate.FBuyPri,
		USDToCNYRate.UpdateTime,
		USDToCNYRate.Name,
		USDToCNYRate.FBuyPri,
		USDToCNYRate.UpdateTime,
	)
	_ = result
	if err != nil {
		Log.Warn(err)
	}

}

func QueryRate() USDToCNYRate {
	sqlStr := `SELECT * FROM rate LIMIT 0,1`
	stmt, err := DB.Prepare(sqlStr)
	if err != nil {
		Log.Warn(err)
	}
	defer func() {
		if stmt != nil {
			stmt.Close()
		}
	}()

	result, err := stmt.Query()
	if err != nil {
		Log.Warn(err)
	}

	var usdToCny USDToCNYRate
	for result.Next() {
		err := result.Scan(&usdToCny.Name, &usdToCny.FBuyPri, &usdToCny.UpdateTime)
		if err != nil {
			Log.Warn(err)
		}
	}

	return usdToCny
}
