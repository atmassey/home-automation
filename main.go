package main

import (
	"log/slog"
	"os"
)

var logger = slog.New(slog.NewJSONHandler(os.Stdout, nil))

func main() {
	logger.Info("Starting up")
	// go StartMqtt()
	go HttpServer()
	select {}
}
