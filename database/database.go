package database

import (
	"database/sql"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
	"io/ioutil"
	"os"
	"strings"
)

var db *sql.DB
var err error
var tag string
var name string
var tags string

func SortImages() {
	distinctTags := "select distinct tags from picture;"
	result := "./public/result"
	tagRows, errRows := db.Query(distinctTags)
	checkError(errRows, "sortImages")
	os.Chdir(result)
	for tagRows.Next() {
		errRows = tagRows.Scan(&tag)
		itag := strings.Split(tag, ",")

		for _, value := range itag {
			tag = strings.Trim(value, " ")
			fileerr := os.Mkdir(tag, 0777)
			checkError(fileerr, "sortImages")
		}
	}

	tagRows.Close()

	sqlQuery := "select name, tags from picture where length(tags)>0;"
	os.Chdir("../")
	rows, err := db.Query(sqlQuery)
	checkError(err, "sortImages")
	for rows.Next() {
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
	rows.Close()
}

func checkError(err error, name string) {
	if err != nil {
		fmt.Println(err, " ", name)
	}
}
func Initialize() {
	db, err = sql.Open("sqlite3", "./pics.db")
	if err != nil {
		fmt.Println(err)
	}

	rows, err := db.Query("select count(*) from picture")
	checkError(err, "select count")

	var count int
	insertStmt := "insert into picture(name) values("
	if rows.Next() {
		rows.Scan(&count)
	}
	fmt.Println(count)
	rows.Close()
	if count == 0 {
		files, err := ioutil.ReadDir("./public")
		checkError(err, "reading files")
		for _, value := range files {
			if !value.IsDir() {
				_, err := db.Exec(insertStmt + "\"" + strings.Trim(value.Name(), " ") + "\");")
				checkError(err, "inserting values")
			}
		}
	}
}

func GetPhoto() (string, int) {
	var name string
	var remaining int

	rows, err := db.Query("select name from picture where tags is null limit 1;")
	checkError(err, "GetPhoto")

	if rows.Next() {
		err = rows.Scan(&name)
		checkError(err, "GetPhoto")
	}

	rows.Close()

	rows, err = db.Query("select count(*) from picture where tags is null limit 1;")
	checkError(err, "GetPhoto")

	if rows.Next() {
		err = rows.Scan(&remaining)
		checkError(err, "GetPhoto")
	}

	rows.Close()

	return name, remaining
}

func Update(tags string, photoname string) bool {
	var updateQuery string = "update picture set tags = ? where name = ?"
	_, err = db.Exec(updateQuery, tags, photoname)
	if err != nil {
		fmt.Println(err)
		return false
	}
	return true
}

func Close() {
	db.Close()
}

func DeleteImages() bool {
	var deleteQuery string = "delete from picture"
	_, err = db.Exec(deleteQuery)
	if err != nil {
		fmt.Println(err)
		return false
	}
	return true
}
