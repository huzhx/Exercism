package lasagna

// TODO: define the 'OvenTime' constant
const OvenTime = 40

// TODO: define the 'RemainingOvenTime()' function
func RemainingOvenTime(timeSpentInMins int) int {
	if timeSpentInMins <= OvenTime {
		return OvenTime - timeSpentInMins
	}
	return 0
}

// TODO: define the 'PreparationTime()' function
func PreparationTime(numLayers int) int {
	return 2 * numLayers
}

// TODO: define the 'ElapsedTime()' function
func ElapsedTime(numLayers, timeSpentInMins int) int {
	return PreparationTime(numLayers) + timeSpentInMins
}
