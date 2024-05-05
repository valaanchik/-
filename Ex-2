
package main

import (
	"fmt"
)

func findResult(numbers []int) []int {

	minNumber := numbers[0]
	for _, number := range numbers {
		if number < minNumber {
			minNumber = number
		}
	}

	div := make([]int, 0)
	for i := 1; i <= minNumber; i++ {
		if minNumber%i == 0 {
			div = append(div, i)
		}
	}

	commonDiv := make([]int, 0)
	for _, divisor := range div {
		boolCommonDiv := true
		for _, number := range numbers {
			if number%divisor != 0 {
				boolCommonDiv = false
				break
			}
		}
		if boolCommonDiv {
			commonDiv = append(commonDiv, divisor)
		}
	}

	return commonDiv
}

func main() {
	numbers := []int{42, 12, 18}
	commonDiv := findResult(numbers)
	fmt.Println("Результат:", commonDiv)
}

