package ecobee

import (
	"log/slog"
	"os"

	"github.com/rspier/go-ecobee/ecobee"
)

var apiKey string

func init() {
	apiKey = os.Getenv("ECOBEE_API_KEY")
	if apiKey == "" {
		slog.Error("No ecobee api key provided")
	}
}

func Status(apiKey string) {
	thermostat := ecobee.NewClient(apiKey, ".go-ecobee.yaml")
	selection := ecobee.Selection{
		IncludeSensors: true,
		IncludeAlerts:  true,
	}
	summary, err := thermostat.GetThermostatSummary(selection)
	if err != nil {
		slog.Error("error while gathering summary", "err", err)
	}
	slog.Info("Ecobee response", "summary", summary)
}
