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

// Aqara door sensor conversion struct
type DoorSensorConversion struct {
	Battery           float32
	Contact           bool
	LinkQuality       int
	PowerOutageCount  int
	Voltage           float32
	DeviceTemperature float32
}

// Aqara Temp sensor struct
type Temp struct {
	Battery            int     `json:"battery"`
	Link_Quality       int     `json:"linkquality"`
	Power_Outage_Count int     `json:"power_outage_count"`
	Voltage            int     `json:"voltage"`
	Device_Temperature int     `json:"device_temperature"`
	Humidity           float32 `json:"humidity"`
	Pressure           float32 `json:"pressure"`
}

// Aqara Temp sensor conversion struct
type TempSensorConversion struct {
	Battery           float32
	LinkQuality       int
	PowerOutageCount  int
	Voltage           float32
	DeviceTemperature float32
	Humidity          float32
	Pressure          float32
}
