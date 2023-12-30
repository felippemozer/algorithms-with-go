package module01

import "math"

// BaseToDec takes in a number and the base it is currently
// in and returns the decimal equivalent as an integer.
//
// Eg:
//
//	BaseToDec("E", 16) => 14
//	BaseToDec("1110", 2) => 14
func BaseToDec(value string, base int) int {
	var result int
	reversed_value := Reverse(value)
	for idx, r := range reversed_value {
		num := int(r)
		if num >= 65 && num <= 90 { // if value is a uppercase letter (A, B, C, etc) in ASCII
			num = num - 55
		} else if num >= 48 && num <= 57 { // if value is a number in ASCII
			num = num - 48
		}
		result = result + (num * (int(math.Pow(float64(base), float64(idx)))))
	}
	return result
}
