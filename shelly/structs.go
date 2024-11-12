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
