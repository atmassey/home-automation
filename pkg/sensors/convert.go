package sensors

// Convert Door sensor temp to freedom units
func (d *Door) Convert() *DoorSensorConversion {
	ConvertedTemperature := (d.Device_Temperature * 9 / 5) + 32
	ConvertedVoltage := d.Voltage / 1000
	return &DoorSensorConversion{
		Battery:           float32(d.Battery),
		Contact:           d.Contact,
		LinkQuality:       d.Link_Quality,
		PowerOutageCount:  d.Power_Outage_Count,
		Voltage:           float32(ConvertedVoltage),
		DeviceTemperature: float32(ConvertedTemperature),
	}
}

// Convert Temp sensor temp to freedom units
func (t *Temp) Convert() *TempSensorConversion {
	ConvertedTemperature := (t.Device_Temperature * 9 / 5) + 32
	ConvertedVoltage := t.Voltage / 1000
	return &TempSensorConversion{
		Battery:           float32(t.Battery),
		LinkQuality:       t.Link_Quality,
		PowerOutageCount:  t.Power_Outage_Count,
		Voltage:           float32(ConvertedVoltage),
		DeviceTemperature: float32(ConvertedTemperature),
		Humidity:          t.Humidity,
		Pressure:          t.Pressure,
	}
}
