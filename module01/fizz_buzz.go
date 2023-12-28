package module01

import (
	"fmt"
)

// FizzBuzz will print out all of the numbers
// from 1 to N replacing any divisible by 3
// with "Fizz", and divisible by 5 with "Buzz",
// and any divisible by both with "Fizz Buzz".
//
// Note: The test for this is a little
// complicated so that you can just use the
// `fmt` package and print to standard out.
// I wouldn't normally recommend this, but did
// it here to make life easier for beginners.
func FizzBuzz(n int) {
	for i := 1; i <= n; i++ {
		if i%15 == 0 {
			fmt.Printf("Fizz Buzz")
		} else if i%5 == 0 {
			fmt.Printf("Buzz")
		} else if i%3 == 0 {
			fmt.Printf("Fizz")
		} else {
			fmt.Printf("%d", i)
		}

		if i != n {
			fmt.Printf(", ")
		}
	}
	fmt.Println()
}
