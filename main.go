package main
import (
	"net/http"
	"fmt"
	"text/template"
	"github.com/thewhitetulip/picsort/database"
)

type Photo struct {
    Name string
    Remaining int
}

func baseHandler(w http.ResponseWriter, r *http.Request) {
	method := r.Method
	if method == "POST" {
		r.ParseForm()
		tags := r.Form.Get("tags")
		photoname := r.Form.Get("photoname")
		database.Update(tags, photoname)
	}
	t, _ := template.ParseFiles("home.gtpl")
	photo_name, remaining := database.GetPhoto()
	p := &Photo{
        Name: photo_name,
        Remaining: remaining,
	}
	
	t.Execute(w, p)
}


func imageHandler(w http.ResponseWriter, r *http.Request){
	database.SortImages()
	w.Write([]byte("Sorting done"))
}

func main() {
	
	
	PORT := ":8080"
	http.HandleFunc("/", baseHandler)
	database.Initialize()
	
	http.HandleFunc("/sort/", imageHandler)
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("./public/"))))
	fmt.Println("running on port 8080\n")
	http.ListenAndServe(PORT, nil)
}
