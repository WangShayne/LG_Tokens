package db

import (
	"fmt"
	"github.com/LG_Tokens/Logger"
	"github.com/LG_Tokens/common"
)

func InsertENS(name, key string) (ens common.Ens, errCode int) {
	sqlStr := "INSERT INTO ens(domainName,pubKey) VALUE(?,?)"
	stmt, err := DB.Prepare(sqlStr)
	if err != nil {
		Logger.Log.Warn(err)
		errCode = 40001
		return
	}

	defer func() {
		if stmt != nil {
			stmt.Close()
		}
	}()

	result, err := stmt.Exec(
		name, key,
	)
	if err != nil {
		errCode = 40002
		return
	}

	id, err := result.LastInsertId()

	if err != nil {
		Logger.Log.Warn(err)
		errCode = 40001
		return
	}

	ens = common.Ens{
		id,
		name,
		key,
	}
	errCode = 0

	return
}

func QueryENS(name string) (ens common.Ens, errCode int) {
	sqlStr := "SELECT * FROM ens WHERE domainName = ?"
	stmt, err := DB.Prepare(sqlStr)
	if err != nil {
		fmt.Println(err)
		errCode = 10001
		return
	}
	defer func() {
		if stmt != nil {
			stmt.Close()
		}
	}()

	result, err := stmt.Query(name)
	if err != nil {
		fmt.Println(err)
		errCode = 10001
		return
	}
	for result.Next() {

		err := result.Scan(&ens.ID, &ens.DomainName, &ens.PubKey)
		if err != nil {
			Logger.Log.Warn(err)
			fmt.Println(err)
			errCode = 10001
		}
	}
	errCode = 0
	return
}
