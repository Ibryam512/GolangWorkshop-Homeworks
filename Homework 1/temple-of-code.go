package main

import (
	"fmt"
	"math"
)

func templeOfCode() {
	var number int
	var sumOfPositiveNumbers int
	var sumOfNegativeNumbers int

	fmt.Scan(&number)
	for number != 0 {
		if number > 0 {
			sumOfPositiveNumbers += number
		} else {
			sumOfNegativeNumbers += number
		}
		fmt.Scan(&number)
	}

	sumOfNegativeNumbers = -sumOfNegativeNumbers

	if sumOfPositiveNumbers == sumOfNegativeNumbers {
		fmt.Println("Yes")
	} else {
		fmt.Println(math.Abs(float64(sumOfPositiveNumbers - sumOfNegativeNumbers)))
	}
}
