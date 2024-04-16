package main

import (
	"fmt"
	"log"
	"net/http"
	"text/template"
)

type Film struct {
	Title    string
	Director string
}

func main() {
	fmt.Println("Hello world")

	h1 := func(w http.ResponseWriter, r *http.Request) {
		templ := template.Must(template.ParseFiles("index.html"))
		films := map[string][]Film{
			"Films": {
				{Title: "The Godfather", Director: "Francis Ford Coppola"},
				{Title: "Blade Runner", Director: "Ridley Scott"},
				{Title: "The Thing", Director: "John Carpenter"},
			},
		}
		err := templ.Execute(w, films)
		if err != nil {
			return
		}
	}
	h2 := func(w http.ResponseWriter, r *http.Request) {
		title := r.PostFormValue("title")
		director := r.PostFormValue("director")
		htmlStr := fmt.Sprintf("<li class='list-group-item bg-primary text-white'>%s - %s</li>", title, director)
		tmpl, _ := template.New("t").Parse(htmlStr)
		err := tmpl.Execute(w, nil)
		if err != nil {
			return
		}
	}
	http.HandleFunc("/", h1)
	http.HandleFunc("/add-film/", h2)

	log.Fatal(http.ListenAndServe(":8080", nil))
}
