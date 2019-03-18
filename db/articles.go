package db

import (
	"fmt"
	"github.com/LG_Tokens/Logger"
	"github.com/LG_Tokens/common"
)

func QueryArticles(key string) (article common.Article) {
	sqlStr := `SELECT * FROM articles WHERE keyName = ?`
	stmt, err := DB.Prepare(sqlStr)
	if err != nil {
		Logger.Log.Warn(err)
		fmt.Println(err)
		return
	}

	defer func() {
		if stmt != nil {
			stmt.Close()
		}
	}()

	rows, err := stmt.Query(key)
	if err != nil {
		Logger.Log.Warn(err)
		fmt.Println(err)
		return
	}
	for rows.Next() {
		var id int32
		var keyname string
		err := rows.Scan(&id, &article.Name, &keyname, &article.Codes)
		if err != nil {
			Logger.Log.Warn(err)
			fmt.Println(err)
		}
	}
	return
}
