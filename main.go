package main

import (
	"html/template"
	"net/http"
	"strconv"
	"sync"
)

var (
	tmpl      = template.Must(template.ParseGlob("templates/*.html"))
	todoItems = []string{}
	mu        sync.Mutex
)

func main() {
	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/add", addHandler)
	http.HandleFunc("/delete", deleteHandler)
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

	http.ListenAndServe(":8080", nil)
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	mu.Lock()
	defer mu.Unlock()
	tmpl.ExecuteTemplate(w, "index.html", todoItems)
}

func addHandler(w http.ResponseWriter, r *http.Request) {
	mu.Lock()
	defer mu.Unlock()
	if r.Method == http.MethodPost {
		todo := r.FormValue("todo")
		todoItems = append(todoItems, todo)
	}

	tmpl.ExecuteTemplate(w, "todo.html", todoItems)
}

func deleteHandler(w http.ResponseWriter, r *http.Request) {
	mu.Lock()
	defer mu.Unlock()
	if r.Method == http.MethodPost {
		index := r.FormValue("index")
		i, _ := strconv.Atoi(index)
		todoItems = append(todoItems[:i], todoItems[i+1:]...)
	}
	tmpl.ExecuteTemplate(w, "todo.html", todoItems)
}
