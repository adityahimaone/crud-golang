package main

import (
	"net/http"
)

func OpenServer() {
	directory := http.Dir("./resources")
	fileServer := http.FileServer(directory)

	mux := http.NewServeMux()
	mux.HandleFunc("/", TemplateIndex)
	mux.HandleFunc("/form", TemplateForm)
	mux.Handle("/static/",http.StripPrefix("/static", fileServer))
	server := http.Server{
		Addr:    "localhost:8080",
		Handler: mux,
	}
	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}
