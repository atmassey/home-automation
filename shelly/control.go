package shelly

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func NewShelly(RelayID string) *Shelly {
	return &Shelly{RelayID: RelayID}
}

// GetRelayStatus returns the status of the relay
func (s *Shelly) GetRelayStatus() (*RelayStatus, error) {
	var status RelayStatus
	req, err := http.NewRequest("GET", fmt.Sprintf("http://%s%s%s", s.RelayIP, relay, s.RelayID), nil)
	if err != nil {
		return nil, err
	}
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}
	if err := json.NewDecoder(resp.Body).Decode(&status); err != nil {
		return nil, err
	}
	return &status, nil
}

// ToggleRelay toggles the relay on
func (s *Shelly) TurnRelayOn(opts *ShellyOptions) error {
	if opts.Timer == "" {
		opts.Timer = "0"
	}
	req, err := http.NewRequest("GET", fmt.Sprintf("http://%s%s%s", s.RelayIP, relay, s.RelayID), nil)
	if err != nil {
		return err
	}
	q := req.URL.Query()
	q.Add("turn", "on")
	q.Add("timer", opts.Timer)
	req.URL.RawQuery = q.Encode()
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}
	return nil
}

// TurnRelayOff turns off the relay
func (s *Shelly) TurnRelayOff(opts *ShellyOptions) error {
	if opts.Timer == "" {
		opts.Timer = "0"
	}
	req, err := http.NewRequest("GET", fmt.Sprintf("http://%s%s%s", s.RelayIP, relay, s.RelayID), nil)
	if err != nil {
		return err
	}
	q := req.URL.Query()
	q.Add("turn", "off")
	q.Add("timer", opts.Timer)
	req.URL.RawQuery = q.Encode()
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}
	return nil
}

// ToggleRelay toggles the relay state
func (s *Shelly) ToggleRelay(opts *ShellyOptions) error {
	if opts.Timer == "" {
		opts.Timer = "0"
	}
	req, err := http.NewRequest("GET", fmt.Sprintf("http://%s%s%s", s.RelayIP, relay, s.RelayID), nil)
	if err != nil {
		return err
	}
	q := req.URL.Query()
	q.Add("turn", "toggle")
	q.Add("timer", opts.Timer)
	req.URL.RawQuery = q.Encode()
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}
	return nil
}
