package main

import (
	"fmt"
	"strings"
)

func raffle() {
	var number string
	fmt.Scan(&number)
	var lastRepeatedNumber string
	var previousNumbers string

	for i := 0; i < len(number); i++ {
		numberToCheck := string(number[i])
		if strings.Contains(previousNumbers, numberToCheck) {
			lastRepeatedNumber = numberToCheck
		}
		previousNumbers += numberToCheck
	}

	if lastRepeatedNumber == "" {
		fmt.Println("No")
	} else {
		fmt.Println(lastRepeatedNumber)
	}
}
