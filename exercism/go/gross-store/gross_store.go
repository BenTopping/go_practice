package gross

// Units stores the Gross Store unit measurements.
func Units() map[string]int {
	return map[string]int{
		"quarter_of_a_dozen": 3,
		"half_of_a_dozen": 6,
		"dozen": 12,
		"small_gross": 120,
		"gross": 144,
		"great_gross": 1728}
}

// NewBill creates a new bill.
func NewBill() map[string]int {
	return map[string]int{}
}

// AddItem adds an item to customer bill.
func AddItem(bill, units map[string]int, item, unit string) bool {
	unitsValue, unitsExists := units[unit]
	if (unitsExists == true) {
		billValue, billExists := bill[item]
		if (billExists == true) {
			bill[item] = billValue + unitsValue
		} else {
			bill[item] = unitsValue
		}
	}
	return unitsExists
}

// RemoveItem removes an item from customer bill.
func RemoveItem(bill, units map[string]int, item, unit string) bool {
	billValue, billExists := bill[item]
	unitsValue, unitsExists := units[unit]
	if billExists == true && unitsExists == true {
		newQuantity := billValue - unitsValue
		if newQuantity > 0 {
			bill[item] = newQuantity
			return true
		} else if newQuantity == 0 {
			delete(bill, item)
			return true
		} else {
			return false
		}
	} else {
		return false
	}
}

// GetItem returns the quantity of an item that the customer has in his/her bill.
func GetItem(bill map[string]int, item string) (int, bool) {
	value, exists := bill[item]
	return value, exists
}
