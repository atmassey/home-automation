package weather

import (
	"errors"

	owm "github.com/briandowns/openweathermap"
)

func CurrentWeather(apiKey string, city string) (*owm.CurrentWeatherData, error) {
	if city == "" {
		return nil, errors.New("city is required")
	}
	w, err := owm.NewCurrent("F", "en", apiKey)
	if err != nil {
		return nil, err
	}
	w.CurrentByName(city)
	return w, nil
}

func Forecast(apiKey string, city string) (*owm.ForecastWeatherData, error) {
	if city == "" {
		return nil, errors.New("city is required")
	}
	f, err := owm.NewForecast("5", "F", "en", apiKey)
	if err != nil {
		return nil, err
	}
	f.DailyByName(city, 5)
	return f, nil
}
