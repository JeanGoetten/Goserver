package main

import (
	"fmt"
	"html/template"
	"net/http"
)

type NewPage struct {
	Title string
	News  string
}

func Home(w http.ResponseWriter, r *http.Request) {
	p := NewPage{Title: "Title Home", News: "News Home"}
	t, _ := template.ParseFiles("index.html")
	err := t.Execute(w, p)
	fmt.Println(err)
}

func Page1(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<h1>Page 1<h1")
}

func rotes() {
	http.HandleFunc("/", Home)
	http.HandleFunc("/page1", Page1)
}

func server_head() {
	rotes()

	fmt.Println("Server running on http://localhost:1337")
	http.ListenAndServe(":1337", nil) //defaultservermux
}

func main() {
	server_head()
}
