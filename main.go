package main

import (
	"fmt"
)

func main() {
	fmt.Println("Hello World!")
}

func FizzBuzz(number int) string {

	if number%15 == 0 {
		return "FizzBuzz"
	}

	if number%5 == 0 {
		return "Buzz"
	}

	if number%3 == 0 {
		return "Fizz"
	}

	return fmt.Sprint(number)
}
