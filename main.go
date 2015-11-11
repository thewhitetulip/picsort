package main

import (
	"fmt"
	"github.com/thewhitetulip/picsort/database"
	"net/http"
	"text/template"
)

type Photo struct {
	Name      string
	Remaining int
}

var t *template.Template
var tmpl *template.Template
var err error

func baseHandler(w http.ResponseWriter, r *http.Request) {
	method := r.Method
	if method == "POST" {
		r.ParseForm()
		tags := r.Form.Get("tags")
		photoname := r.Form.Get("photoname")
		database.Update(tags, photoname)
	}

	photo_name, remaining := database.GetPhoto()
	p := &Photo{
		Name:      photo_name,
		Remaining: remaining,
	}

	t.Execute(w, p)
}

func imageHandler(w http.ResponseWriter, r *http.Request) {
	database.SortImages()
	tmpl.Execute(w, nil)
}

func deleteHandler(w http.ResponseWriter, r *http.Request) {
	truth := database.DeleteImages()
	if truth != false {
		w.Write([]byte("Delete done. stop the server now, output is copied in public/result folder! please cut the folders after you are done"))
	}
}

func main() {
	t, err = template.ParseFiles("home.gtpl")
	if err != nil {
		fmt.Println("Error parsing file")
	}
	message := "<html>sorting done, <a href='/delete'> delete </a> db entries and old files?</html>"
	tmpl, err = template.New("delete").Parse(message)

	if err != nil {
		fmt.Println("Error parsing file")
	}

	PORT := ":8080"
	database.Initialize()
	defer database.Close()
	http.HandleFunc("/", baseHandler)
	http.HandleFunc("/sort/", imageHandler)
	http.HandleFunc("/delete/", deleteHandler)
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("./public/"))))
	fmt.Println("running on port 8080\n")
	fmt.Println(http.ListenAndServe(PORT, nil))
}
