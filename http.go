package main

import (
	"bytes"
	"encoding/json"
	"io"
	"log"
	"log/slog"
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
	http.HandleFunc("/api/v1/toggle/", toggle_shelly)
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

func toggle_shelly(w http.ResponseWriter, r *http.Request) {
	var requestBody map[string]interface{}
	body, err := io.ReadAll(r.Body)
	if err != nil {
		slog.Error("Error reading body", "error", err)
		http.Error(w, "Error reading body", http.StatusInternalServerError)
		return
	}
	err = json.Unmarshal(body, &requestBody)
	if err != nil {
		slog.Error("Error unmarshaling body", "error", err)
		http.Error(w, "Error unmarshaling body", http.StatusInternalServerError)
		return
	}
	shellyLocation, ok := requestBody["location"].(string)
	if !ok {
		slog.Error("Invalid location", "location", requestBody["location"])
		http.Error(w, "Invalid location", http.StatusBadRequest)
		return
	}
	slog.Info("Toggling shelly", "location", shellyLocation)
	// TODO: Add code to toggle the shelly device
}
