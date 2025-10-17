package main

import (
	"fmt"
	"html/template"
	"net/http"
	"os"
)

func main() {
	listTemplates, errtemplates := template.ParseGlob("./templates/*.html")
	if errtemplates != nil {
		fmt.Println(errtemplates)
		os.Exit(1)
	}
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		listTemplates.ExecuteTemplate(w, "index", nil)
	})

	http.ListenAndServe("localhost:8000", nil)
}
