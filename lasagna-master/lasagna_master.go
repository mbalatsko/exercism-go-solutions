package lasagna

// TODO: define the 'PreparationTime()' function
func PreparationTime(layers []string, time int) int {
	if time == 0 {
		time = 2
	}
	return len(layers) * time
}

func CountByLayerName(layers []string, layerName string) (cnt int) {
	for _, l := range layers {
		if l == layerName {
			cnt += 1
		}
	}
	return
}

// TODO: define the 'Quantities()' function
func Quantities(layers []string) (int, float64) {
	return CountByLayerName(layers, "noodles") * 50, float64(CountByLayerName(layers, "sauce")) * .2
}

// TODO: define the 'AddSecretIngredient()' function
func AddSecretIngredient(friendsList []string, myList []string) {
	myList[len(myList)-1] = friendsList[len(friendsList)-1]
}

// TODO: define the 'ScaleRecipe()' function
func ScaleRecipe(quantities []float64, scale int) []float64 {
	scaledQuantities := make([]float64, 0, len(quantities))
	for _, q := range quantities {
		scaledQuantities = append(scaledQuantities, (q/2)*float64(scale))
	}

	return scaledQuantities
}

// Your first steps could be to read through the tasks, and create
// these functions with their correct parameter lists and return types.
// The function body only needs to contain `panic("")`.
//
// This will make the tests compile, but they will fail.
// You can then implement the function logic one by one and see
// an increasing number of tests passing as you implement more
// functionality.
