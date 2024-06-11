package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
	"strings"
)

var contentType map[string]string = make(map[string]string)

type Page struct {
	Path  string
	Title string
	Body  []byte
}

func loadFile(filename string) (*Page, error) {
	path := "./public" + filename
	text, err := os.ReadFile(path)

	if err != nil {
		return nil, err
	}

	return &Page{Path: path, Title: filename, Body: text}, nil
}

func handler(w http.ResponseWriter, r *http.Request) {
	path := r.URL.Path

	fmt.Printf("req: '%s'\n", path)

	if path == "/" {
		path = "/index.html"
	}

	txt, err := loadFile(path)
	if err != nil {
		fmt.Fprintf(w, "Path '%s' not found", path)
		return
	}

	s := strings.Split(path, ".")
	extension := s[len(s)-1]

	w.Header().Add("Content-Type", contentType[extension])
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

	contentType["html"] = "text/html"
	contentType["css"] = "text/css"

	http.HandleFunc("/", handler)
	http.HandleFunc("/hello/", helloHandler)
	log.Fatal(http.ListenAndServe(":8000", nil))
}
