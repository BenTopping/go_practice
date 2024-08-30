package purchase

// NeedsLicense determines whether a license is needed to drive a type of vehicle. Only "car" and "truck" require a license.
func NeedsLicense(kind string) bool {
	return (kind == "car") || (kind == "truck")
}

// ChooseVehicle recommends a vehicle for selection. It always recommends the vehicle that comes first in lexicographical order.
func ChooseVehicle(option1, option2 string) string {
	betterChoice := option1
	if (option2 < option1) {
		betterChoice = option2
	}
	return betterChoice + " is clearly the better choice."
}

// CalculateResellPrice calculates how much a vehicle can resell for at a certain age.
func CalculateResellPrice(originalPrice, age float64) float64 {
	if (int(age) >= 10) {
		return float64(originalPrice * 0.5)
	} else if (int(age) >= 3) {
		return float64(originalPrice * 0.7)
	} else {
		return float64(originalPrice * 0.8)
	}
}
