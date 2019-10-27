package db

import (
	"database/sql"
	"log"
	"os"
	"strconv"
	"strings"

	_ "github.com/lib/pq"
)

var database *sql.DB = nil

//Init this function opens a postgresql database
func Init() {
	var err error
	database, err = sql.Open("postgres", os.Getenv("DATABASE_URL"))
	if err != nil {
		log.Fatal(err)
	}
}

//GetDB get *sql.DB. if the variable databse is nil, then it calls Init()
func GetDB() *sql.DB {
	if database == nil {
		Init()
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
