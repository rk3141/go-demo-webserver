package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func loadFile(filename string) ([]byte, error) {
	text, err := os.ReadFile(filename)

	if err != nil {
		return nil, err
	}

	return text, nil
}

func handler(w http.ResponseWriter, r *http.Request) {
	path := r.URL.Path

	if path == "/" {
		path = "/index.html"
	}

	txt, err := loadFile("./public" + path)
	if err != nil {
		fmt.Fprintf(w, "Path '%s' not found", path)
		return
	}
	fmt.Fprint(w, string(txt))
}

func main() {
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe(":8000", nil))
}
