package main

import (
	"fmt"
	"math"
)

func main() {

	fmt.Println("Please enter number.")
	var num1 float64
	fmt.Scanln(&num1)

	//fmt.Println(math.Sqrt(num1))

	var squar1 float64
	squar1 = math.Sqrt(num1)

	fmt.Println("squar root is", squar1)
}
