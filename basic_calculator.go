package main

import "fmt"

func main() {

	fmt.Println("Which opperator you want to use. +=1, -=2, *=3, /=4")
	var num3 int
	fmt.Scanln(&num3)

	fmt.Println("Please enter number 1.")
	var num1 float64
	fmt.Scanln(&num1)

	fmt.Println("Please enter number 2.")
	var num2 float64
	fmt.Scanln(&num2)

	switch num3 {
	case 1:
		var sum float64
		sum = num1 + num2
		fmt.Println("sum is", sum)

	case 2:
		var sub float64
		sub = num1 - num2
		fmt.Println("subtraction is", sub)

	case 3:
		var mul float64
		mul = num1 * num2
		fmt.Println("multiplication is", mul)

	case 4:
		var div float64
		div = num1 / num2
		fmt.Println("division is", div)
	}

}
