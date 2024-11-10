package sensors

// Convert Door sensor temp to non-commy units
func (d *Door) Convert() {
	d.Device_Temperature = d.Device_Temperature / 100
}
