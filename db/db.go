package db

import (
	"database/sql"
	"log"
	"strconv"
	"strings"

	_ "github.com/lib/pq"
)

var database *sql.DB = nil

//Init this function opens a postgresql database
func Init(dataSource string) {
	var err error
	database, err = sql.Open("postgres", dataSource)
	if err != nil {
		log.Fatal(err)
	}
}

//GetDB get *sql.DB.
func GetDB() *sql.DB {
	if database == nil {
		return nil
	}
	return database
}

//Close close a postgresql database
func Close() {
	if database != nil {
		err := database.Close()
		if err != nil {
			log.Fatal(err)
		}
		database = nil
	}
}

//util
//////////////////////////////////////

//Model the type of model which has below functions.
//ToSlice() this function returns a slice in which has member variables as elements.
type Model interface {
	To_Slice() []interface{}
}

//GenerateSQLForMultiRows this function generates a sql query and variables for query
//input: an array of data
//output: variables for query and a sql query
func GenerateSQLForMultiRows(datas []Model) ([]interface{}, string) {
	vals := []interface{}{}
	sqlStr := ""
	count := 1
	for _, data := range datas {
		slicedData := data.To_Slice()
		temp := "("
		length := len(slicedData)
		for i := 0; i < length-1; i++ {
			temp += "$" + strconv.Itoa(count) + ","
			vals = append(vals, slicedData[i])
			count++
		}
		temp += "$" + strconv.Itoa(count) + "),"
		vals = append(vals, slicedData[length-1])
		count++
		sqlStr += temp
	}
	sqlStr = strings.TrimSuffix(sqlStr, ",")
	return vals, sqlStr
}

//SafeQuery
//コネクションリークを起こさない関数
/*
引数
query string 実行するSQL文
args []interface{} queryに展開する引数
返り値
[][]interface{} クエリーの実行結果
error
*/
func SafeQuery(query string, args []interface{}) ([][]interface{}, error) {
}
