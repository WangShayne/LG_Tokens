package db

import (
	"fmt"
	"github.com/LG_Tokens/Logger"
	"github.com/LG_Tokens/common"
)

func QueryArticles(key string) (article common.Article, errCode int) {
	sqlStr := `SELECT name,type,codes FROM articles WHERE keyName = ?`
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

	result := stmt.QueryRow(key)
	result.Scan(&article.Name, &article.Type, &article.Codes)

	if article.Name == "" {
		errCode = 60001
	} else {
		errCode = 0
	}
	return
}
