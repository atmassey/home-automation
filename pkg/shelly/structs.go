package shelly

// Status of Shelly 1 Relay
type RelayStatus struct {
	IsOn           bool   `json:"ison"`
	HasTimer       bool   `json:"has_timer"`
	TimerStarted   bool   `json:"timer_started"`
	TimerDuration  int    `json:"timer_duration"`
	TimerRemaining int    `json:"timer_remaining"`
	Overpower      bool   `json:"overpower"`
	Source         string `json:"source"`
}

type Shelly struct {
	RelayIP string
}

type ShellyOptions struct {
	Timer string
	Turn  string
}

type ShellyHAndT struct {
	IsValid        bool           `json:"is_valid"`
	Temp           ShellyTemp     `json:"tmp"`
	Humidity       ShellyHumidity `json:"hum"`
	Battery        ShellyBattery  `json:"bat"`
	Action         []string       `json:"act_reasons"`
	ConnectRetries int            `json:"connect_retries"`
}

type ShellyTemp struct {
	Value   float64 `json:"value"`
	Units   string  `json:"units"`
	TC      float64 `json:"tC"`
	TF      float64 `json:"tF"`
	IsValid bool    `json:"is_valid"`
}

type ShellyHumidity struct {
	Value   float64 `json:"value"`
	IsValid bool    `json:"is_valid"`
}

type ShellyBattery struct {
	Value   float64 `json:"value"`
	Voltage float64 `json:"voltage"`
}

//TODO: Add Shelly plug structs
