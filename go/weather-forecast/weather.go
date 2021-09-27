// Package weather provides ability to check weather information
package weather

// CurrentCondition stores current weather condition
var CurrentCondition string

// CurrentLocation stores current location information
var CurrentLocation string

// Forecast returns a string that shows the current weather condition at the current location
func Forecast(city, condition string) string {
	CurrentLocation, CurrentCondition = city, condition
	return CurrentLocation + " - current weather condition: " + CurrentCondition
}
