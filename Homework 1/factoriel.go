package main

import (
	"fmt"
	"strconv"
)

func factorielResult(number int) int {
	result := 1
	for number > 1 {
		result *= number
		number--
	}
	return result
}

func countNumbers(number int) int {
	result := 1
	for number > 10 {
		number /= 10
		result++
	}
	return result
}

func factoriel() {
	result := []int{}

	for i := 0; i < 5; i++ {
		var input string
		fmt.Scan(&input)
		number, _ := strconv.Atoi(input)
		count := countNumbers(factorielResult(number))
		result = append(result, count)
	}

	for i := 0; i < len(result); i++ {
		fmt.Printf("%d ", result[i])
	}

}
