package ecobee

import (
	"log"

	"github.com/rspier/go-ecobee/ecobee"
)

func status(apiKey string) {
	thermostat := ecobee.NewClient(apiKey, ".go-ecobee.yaml")
	log.Printf("Thermostat: %v", thermostat)
}
