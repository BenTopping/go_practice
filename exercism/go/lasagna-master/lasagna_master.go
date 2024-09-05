package lasagna

func PreparationTime(layers []string, avgLayerTime int) int {
	if (avgLayerTime == 0) {
		avgLayerTime = 2
	}
	return len(layers) * avgLayerTime
}

func Quantities(layers []string) (noodles int, sauce float64) {
	for _, layer := range layers {
		if (layer == "noodles") {
			noodles += 50
		} else if (layer == "sauce") {
			sauce += 0.2
		}
	}
	return noodles, sauce
}

func AddSecretIngredient(friendList []string, ownList[]string) {
	ownList[len(ownList)-1] = friendList[len(friendList)-1]
}

func ScaleRecipe(portionAmount []float64, portionNumber int) (scaledPortions[]float64) {
	for _, portion := range portionAmount {
		singlePortion := portion / 2
		scaledPortions = append(scaledPortions, singlePortion * float64(portionNumber))
	}
	return scaledPortions
}
