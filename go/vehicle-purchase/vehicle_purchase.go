package purchase

import "fmt"

// NeedsLicense determines whether a license is needed to drive a type of vehicle. Only "car" and "truck" require a license.
func NeedsLicense(kind string) bool {
	if kind == "car" || kind == "truck" {
		return true
	}
	return false
}

// ChooseVehicle recommends a vehicle for selection. It always recommends the vehicle that comes first in lexicographical order.
func ChooseVehicle(option1, option2 string) string {
	output := option1
	if option1 > option2 {
		output = option2
	}
	return fmt.Sprintf("%s is clearly the better choice.", output)
}

// CalculateResellPrice calculates how much a vehicle can resell for at a certain age.
func CalculateResellPrice(originalPrice, age float64) float64 {
	ratio := 1.0
	if age < 3 {
		ratio = 0.8
	} else if age >= 3 && age < 10 {
		ratio = 0.7
	} else if age >= 10 {
		ratio = 0.5
	}
	return originalPrice * ratio
}
