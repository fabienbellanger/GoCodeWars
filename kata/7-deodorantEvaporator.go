package kata

// Evaporator reports the nth day (as an integer) on which the evaporator will be out of use
func Evaporator(content float64, evapPerDay int, threshold int) int {
	nbDays := 0
	limit := content * float64(threshold) / 100
	decrement := 1 - float64(evapPerDay)/100

	for remaing := content; remaing >= limit; remaing *= decrement {
		nbDays++
	}

	return nbDays
}
