package utils

import (
	"math/big"
)

func IsNumeric(word string) bool {
	_, ok := new(big.Float).SetString(word)
	return ok
}

func ValidName(word string) bool {
	_, ok := new(big.Float).SetString(word)
	return !ok
}
