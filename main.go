package main

func main() {
	runSystray()
}

func minimizeDisplay(intValue uint64) (float64, string) {
	units := []string{"B", "KB", "MB", "GB", "TB", "PB"}
	unit := 0
	value := float64(intValue)
	for value > 1024 {
		value = value / 1024
		unit++
	}
	return value, units[unit]
}
