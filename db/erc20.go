package db

import (
	"fmt"

	. "github.com/LG_Tokens/Logger"
	. "github.com/LG_Tokens/common"
)

func QueryCommonToken() (tokenSlice []Tokens, errCode int) {
	sqlStr := `SELECT * FROM tokens where isCommon = 1`
	stmt, err := DB.Prepare(sqlStr)
	if err != nil {
		Log.Warn(err)
		return tokenSlice, 10001
	}

	defer func() {
		if stmt != nil {
			stmt.Close()
		}
	}()

	result, err := stmt.Query()
	if err != nil {
		Log.Warn(err)
		return tokenSlice, 10001
	}
	for result.Next() {
		var id int64
		var token Tokens
		err := result.Scan(&id, &token.Code, &token.IsCommon, &token.Name, &token.EnName, &token.CnName, &token.Address, &token.Logo)
		if err != nil {
			Log.Warn(err)
		}
		fmt.Println(token)
		tokenSlice = append(tokenSlice, token)
	}
	return tokenSlice, 00000
}

func QueryKeywordToken(keyword string) (tokenSlice []Tokens, errCode int) {
	sqlStr := `SELECT * FROM tokens where name like ? or enName like ? or cnName like ? or address like ?`
	stmt, err := DB.Prepare(sqlStr)
	if err != nil {
		Log.Error(err)
		return tokenSlice, 10001
	}

	defer func() {
		if stmt != nil {
			stmt.Close()
		}
	}()
	keyword = "%" + keyword + "%"
	result, err := stmt.Query(keyword, keyword, keyword, keyword)
	if err != nil {
		Log.Warn(err)
	}
	for result.Next() {
		var id int64
		var token Tokens
		err := result.Scan(&id, &token.Code, &token.IsCommon, &token.Name, &token.EnName, &token.CnName, &token.Address, &token.Logo)
		if err != nil {
			Log.Warn(err)
		}
		tokenSlice = append(tokenSlice, token)
	}
	return tokenSlice, 00000
}
