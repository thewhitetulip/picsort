package database

import (
	"fmt"
	_"github.com/mattn/go-sqlite3" 
	"database/sql"
	"strings"
	"os"
	"io/ioutil"
)

var db *sql.DB
var err error

func SortImages() {
	distinctTags := "select distinct tags from picture;"
	var tag string
	result := "./public/result"
	tagRows, errRows := db.Query(distinctTags)
	checkError(errRows, "sortImages")
	os.Chdir(result)
	for tagRows.Next(){
		errRows = tagRows.Scan(&tag)
		itag := strings.Split(tag, ",")
		
		for _, value := range itag {
			tag = strings.Trim(value, " ")
			fileerr := os.Mkdir(tag, 0777)
			checkError(fileerr, "sortImages")
		}
	}
	
	sqlQuery := "select name, tags from picture where length(tags)>0;"
	os.Chdir("../")
	var name string
	var tags string
	rows, err := db.Query(sqlQuery)
	checkError(err, "sortImages")
	for rows.Next(){
		err = rows.Scan(&name, &tags)
		all_tags := strings.Split(tags, ",")
		for _, value := range all_tags {
			//fmt.Println("copying ", name, " to ",value)
			data, readErr := ioutil.ReadFile(name)
			checkError(readErr, "file read")
			value = strings.Trim(value, " ")
			err = ioutil.WriteFile("result/"+value+"/"+name, data, 0644)
			checkError(err, "file copy")
		}
	}
}

func checkError(err error, name string) {
	if err != nil {
		fmt.Println(err , " ",  name)
	}
}
func Initialize() {
	db, err = sql.Open("sqlite3", "./pics.db")
	if err != nil {
		fmt.Println(err)
	}
}

func GetPhoto()  (string, int) {	
	var name string
	var remaining int
	
	rows, err := db.Query("select name from picture where length(tags)<1 limit 1;")

	checkError(err, "GetPhoto")
	
	defer rows.Close()
	if rows.Next() {
		err = rows.Scan(&name)
		checkError(err, "GetPhoto")
	}
	
	rows, err = db.Query("select count(*) from picture where length(tags)<1 limit 1;")
	checkError(err, "GetPhoto")
	
	defer rows.Close()
	if rows.Next(){
		err = rows.Scan(&remaining)
		checkError(err, "GetPhoto")		
	}
	
	return name, remaining
}


func Update(tags string, photoname string) bool {
	var updateQuery string = "update picture set tags = ? where name = ?"
	_, err := db.Exec(updateQuery, tags, photoname)
	if err!=nil{
		fmt.Println(err)
		return false
	}
	return true
}
