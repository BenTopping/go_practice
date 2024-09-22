package thefarm

import (
	"errors"
	"fmt"
)

func DivideFood(fc FodderCalculator, cows int) (float64, error) {
	amount, err := fc.FodderAmount(cows)
	if (err != nil) {
		return 0, err
	}
	factor, err := fc.FatteningFactor()
	if (err != nil) {
		return 0, err
	}
	return (amount * factor) / float64(cows), nil
}

func ValidateInputAndDivideFood(fc FodderCalculator, cows int) (float64, error) {
	if (cows > 0) {
		return DivideFood(fc, cows)
	}
	return 0, errors.New("invalid number of cows")
}
type InvalidCowsError struct {
	cows int
	message string
}

func (e *InvalidCowsError) Error() string {
	return fmt.Sprintf("%v cows are invalid: %s", e.cows, e.message)
}

func ValidateNumberOfCows(cows int) error {
	if (cows < 0) {
		return &InvalidCowsError{cows: cows, message: "there are no negative cows"}
	} else if (cows == 0) {
		return &InvalidCowsError{cows: cows, message: "no cows don't need food"}
	}
	return nil
}

// Your first steps could be to read through the tasks, and create
// these functions with their correct parameter lists and return types.
// The function body only needs to contain `panic("")`.
//
// This will make the tests compile, but they will fail.
// You can then implement the function logic one by one and see
// an increasing number of tests passing as you implement more
// functionality.
