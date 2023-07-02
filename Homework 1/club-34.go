package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func club34() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	input := strings.ToUpper(scanner.Text())
	result := ""

	for i := len(input) - 1; i >= 0; i-- {
		number := int(input[i])
		//using the ASCII table numbers - 65 stands for A, 90 stands for Z
		if number >= 65 && number <= 90 {
			result += strconv.Itoa(number - 64)
		}
	}

	fmt.Print(result)
}
