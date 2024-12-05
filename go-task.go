package main

import (
	"fmt"
	"math"
)

func main() {
	NUMBER := 2
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("An error occurred.")
		}
	}()
	ARR := []int{1, 2, 3, 4, 5}
	NUMBER = 2 + int(math.Pow(2, 3)) + ARR[len(ARR)-1]
	addFive(addFive(NUMBER))
	fmt.Println(NUMBER)
}

func addFive(num int) int {
	newNumber := num + 5
	return newNumber
}
