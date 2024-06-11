package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
)

type Page struct {
	Path  string
	Title string
	Body  []byte
}

func loadFile(filename string) (*Page, error) {
	path := "./public/" + filename + ".html"
	text, err := os.ReadFile(path)

	if err != nil {
		return nil, err
	}

	return &Page{Path: path, Title: filename, Body: text}, nil
}

func handler(w http.ResponseWriter, r *http.Request) {
	path := r.URL.Path

	if path == "/" {
		path = "/index"
	}

	txt, err := loadFile(path[1:])
	if err != nil {
		fmt.Fprintf(w, "Path '%s' not found", path)
		return
	}
	fmt.Fprint(w, string(txt.Body))
}

type HelloPage struct {
	Name string
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("./public/hello.html")
	name := r.URL.Path[len("/hello/"):]
	t.Execute(w, HelloPage{Name: name})
}

func main() {
	http.HandleFunc("/", handler)
	http.HandleFunc("/hello/", helloHandler)
	log.Fatal(http.ListenAndServe(":8000", nil))
}
