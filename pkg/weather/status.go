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
	err = w.CurrentByName(city)
	if err != nil {
		return nil, err
	}
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
	err = f.DailyByName(city, 5)
	if err != nil {
		return nil, err
	}
	return f, nil
}
