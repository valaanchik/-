package main

import (
	"fmt"
)

func multiplyMat(number int) {
	for i := 1; i <= number; i++ {
		for j := 1; j <= number; j++ {
			
			fmt.Print(i*j, " ")
		}
		fmt.Println("")
	}
}

func main() {
	var number int

	fmt.Print("Введите число: ")
	fmt.Scan(&number)

	multiplyMat(number)
}
