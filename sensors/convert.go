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
