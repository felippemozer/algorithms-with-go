package module01

import "fmt"

// DecToBase will return a string representing
// the provided decimal number in the provided base.
// This is limited to bases 2-16 for simplicity.
//
// Eg:
//
//	DecToBase(14, 16) => "E"
//	DecToBase(14, 2) => "1110"
func DecToBase(dec, base int) string {
	var result string

	for dec > 0 {
		remainder := dec % base

		if remainder >= 10 {
			remainder = remainder + 55
			remainder_letter := rune(remainder)
			result = string(remainder_letter) + result
		} else {
			result = fmt.Sprint(remainder) + result
		}

		dec = dec / base
	}

	return result
}
