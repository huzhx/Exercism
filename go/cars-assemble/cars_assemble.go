package cars

const productionAtSpeedOne float64 = 221

// CalculateProductionRatePerHour for the assembly line, taking into account
// its success rate
func CalculateProductionRatePerHour(speed int) float64 {
	return productionAtSpeedOne * float64(speed) * successRate(speed)
}

// CalculateProductionRatePerMinute describes how many working items are
// produced by the assembly line every minute
func CalculateProductionRatePerMinute(speed int) int {
	return int(productionAtSpeedOne * float64(speed) * successRate(speed) / 60)
}

// successRate is used to calculate the ratio of an item being created without
// error for a given speed
func successRate(speed int) float64 {
	if speed == 0 {
		return 0
	} else if speed >= 1 && speed <= 4 {
		return 1
	} else if speed >= 5 && speed <= 8 {
		return 0.9
	} else {
		return 0.77
	}
}
