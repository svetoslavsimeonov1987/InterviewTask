package main

import (
	"fmt"
	"math"
)

func main() {
	NUMBER := 2
	ARR := []int{1, 2, 3, 4, 5}
	
	addFive(addFive(NUMBER))
	
	NUMBER += int(math.Pow(2, 3)) + ARR[len(ARR)]
	fmt.Println(NUMBER)
}

func addFive(num int) int {
	newNumber := num + 5
	return newNumber
}
