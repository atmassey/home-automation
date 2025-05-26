package main

import "os"

var (
	office_fan   string
	office_light string
)

func setup_env() {
	office_fan = os.Getenv("OFFICE_FAN")
	office_light = os.Getenv("OFFICE_LIGHT")
	if office_fan == "" {
		panic("OFFICE_FAN environment variable is not set")
	}
	if office_light == "" {
		panic("OFFICE_LIGHT environment variable is not set")
	}
}
