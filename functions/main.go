package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	apiMux := http.NewServeMux()
	apiMux.HandleFunc("/", HomeHandler)
	apiMux.Handle("/statics/", http.StripPrefix("/statics/", http.FileServer(http.Dir("statics"))))
	fmt.Println("serving...")
	err := http.ListenAndServe(":8080", apiMux)
	if err != nil {
		log.Fatal(err)
	}
}
