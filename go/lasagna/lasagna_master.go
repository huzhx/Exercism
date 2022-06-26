package lasagna

// PreparationTime2 accepts a slice of layers as a []string and the average preparation time per layer in minutes as an int, and returns the estimate for the total preparation time
func PreparationTime2(layers []string, avgTime int) int {
	defaultVal := 2
	numLayers := len(layers)
	if avgTime == 0 {
		return numLayers * defaultVal
	}
	return numLayers * avgTime
}

// Quantities that takes a slice of layers as parameter as a []string and returns the quantity of noodles and sauce needed
func Quantities(layers []string) (int, float64) {
	noodleNeededPerLayer := 50
	sauceNeededPerLayer := 0.2
	noodleNeededTotal := 0
	sauceNeededTotal := 0.0
	for _, layer := range layers {
		switch layer {
		case "noodles":
			noodleNeededTotal += noodleNeededPerLayer
		case "sauce":
			sauceNeededTotal += sauceNeededPerLayer
		}
	}
	return noodleNeededTotal, sauceNeededTotal
}

// AddSecretIngredient accetps the list your friend sent you and the ingredient list of your own recipe replace "?" in your list with the last item from your friends list
func AddSecretIngredient(friendsList []string, myList []string) {
	myList[len(myList)-1] = friendsList[len(friendsList)-1]
}

// ScaleRecipe takes a slice of float64 amounts needed for 2 portions and the number of portions you want to cook and returns a slice of float64 of the amounts needed for the desired number of portions
func ScaleRecipe(quantities []float64, portions int) []float64 {
	output := make([]float64, len(quantities))
	for i, quantityFor2 := range quantities {
		output[i] = quantityFor2 / 2 * float64(portions)
	}
	return output
}
