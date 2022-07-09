package main

import (
	"errors"
	"fmt"
)

var ErrNoRowFound = errors.New("no row found")

type errStruct struct {
	Sql string
	Err error
}

func queryDAO() (string, error) {
	sql := "SELECT * FROM user WHERE id = 1"

	return sql, ErrNoRowFound
}

func main() {
	tmpErr := queryService()
	fmt.Println(tmpErr)
}

func queryService() errStruct {
	var errMsg errStruct

	if tmpSql, tmpErr := queryDAO(); errors.Is(tmpErr, ErrNoRowFound) {
		errors.Unwrap(tmpErr)
		errMsg.Err = fmt.Errorf("%w", tmpErr)
		errMsg.Sql = tmpSql
	}
	return errMsg
}
