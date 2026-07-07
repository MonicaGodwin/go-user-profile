package main

import (
	"log"
	"net/http"
	"text/template"
)

var tmpl = template.Must(template.ParseFiles("templates/index.html"))

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.Error(w, "path not found", http.StatusNotFound)
		return
	}
	if r.Method != http.MethodGet {
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}

	data, err := UserProfile()
	if err != nil {
		http.Error(w, "server error", http.StatusInternalServerError)
		return
	}
	err = tmpl.Execute(w, data)
	log.Println(err)
	if err != nil {
		http.Error(w, "template error", http.StatusInternalServerError)
		return
	}
}
