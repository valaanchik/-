package main

import (
	"fmt"
)

func simple(n int) bool {
	for i := 2; i*i <= n; i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}

func simpleNum(num1 int, num2 int) []int {
	listNum := []int{}
	for i := num1; i <= num2; i++ {
		if simple(i) {
			listNum = append(listNum, i)
		}

	}
	return listNum
}

func main() {
	var number1 int
	var number2 int

	fmt.Print("Введите число 1: ")
	fmt.Scan(&number1)
	fmt.Print("Введите число 2: ")
	fmt.Scan(&number2)

	result := simpleNum(number1, number2)
	fmt.Println(result)
}
