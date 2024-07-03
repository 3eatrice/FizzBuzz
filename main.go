package main

import (
	"fmt"
)

func main() {
	fmt.Println("Hello World!")
}

func FizzBuzz(number int) string {

	myString := ""

	if number%3 == 0 {
		myString += "Fizz"
	}

	if number%5 == 0 {
		myString += "Buzz"
	}

	if myString == "" {
		myString = fmt.Sprint(number)
	}

	return myString
}
