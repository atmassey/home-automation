package main

import (
	"fmt"
	"html"
	"log"
	"net/http"
	"time"
)

func HttpServer() {
	s := &http.Server{
		Addr:              ":8081",
		ReadHeaderTimeout: time.Second * 10,
	}
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))
	})
	err := s.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
}
