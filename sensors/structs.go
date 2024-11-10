package sensors

// Aqara door sensor struct
type Door struct {
	Battery            int  `json:"battery"`
	Contact            bool `json:"contact"`
	Link_Quality       int  `json:"linkquality"`
	Power_Outage_Count int  `json:"power_outage_count"`
	Voltage            int  `json:"voltage"`
	Device_Temperature int  `json:"device_temperature"`
}
