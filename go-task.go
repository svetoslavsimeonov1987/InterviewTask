// Question: What will happen when you run this GO code?

package main

import (
	"fmt"
)

const number = 5

func main() {
	var number int64
	number = 10 + number
	addFive(addFive(number))
	fmt.Print(number, "\n")
}

func addFive(num int64) int64 {
	newNumber := num + 5
	return newNumber
}
