package main

import (
	"bytes"
	"log"
	"net/http"
	"os"
	"text/template"
	"time"
)

var templates *template.Template

func load_templates() {
	// Load templates
	var err error
	templates, err = template.ParseFS(os.DirFS("./web/templates"), "*.html")
	if err != nil {
		log.Fatal(err)
	}
}

func HttpServer() {
	logger.Info("Starting http server")
	s := &http.Server{
		Addr:              ":8081",
		ReadHeaderTimeout: time.Second * 10,
	}
	load_templates()
	http.HandleFunc("/", html_main)
	http.HandleFunc("/sensors", html_sensors)
	err := s.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
}

func html_main(w http.ResponseWriter, r *http.Request) {
	b := new(bytes.Buffer)
	err := templates.ExecuteTemplate(w, "main.html", b)
	if err != nil {
		log.Fatal(err)
	}
	w.Header().Set("Content-Type", "text/html")
	w.WriteHeader(http.StatusOK)
	_, err = w.Write(b.Bytes())
	if err != nil {
		log.Fatal(err)
	}
}

func html_sensors(w http.ResponseWriter, r *http.Request) {
	b := new(bytes.Buffer)
	err := templates.ExecuteTemplate(w, "sensors.html", b)
	if err != nil {
		log.Fatal(err)
	}
	w.Header().Set("Content-Type", "text/html")
	w.WriteHeader(http.StatusOK)
	_, err = w.Write(b.Bytes())
	if err != nil {
		log.Fatal(err)
	}
}
