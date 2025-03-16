package main

import (
	"fmt"
)

type transformationFunction func(int) int

func main() {
	numbers := []int{1, 2, 3, 4, 5}
	fmt.Println(transformNumbers(numbers, addOne))
	fmt.Println(transformNumbers(numbers, createMultiplyByANumberFunction(3)))
	variadicFunction(1, 2, 3, 4, 5)
	variadicFunction(numbers...)
}

// This function takes a slice of numbers and a function that transforms a number
// It returns a slice of numbers that have been transformed by the function
// The function passed in is a higher order function
func transformNumbers(numbers []int, f transformationFunction) []int {
	var result []int
	for _, number := range numbers {
		result = append(result, f(number))
	}
	return result
}

func addOne(number int) int {
	return number + 1
}

// This function returns a function that multiplies a number by the multiplier
// The function returned is a closure that captures the multiplier variable
func createMultiplyByANumberFunction(multiplier int) transformationFunction {
	return func(number int) int {
		return number * multiplier
	}
}

func variadicFunction(numbers ...int) {
	for _, number := range numbers {
		fmt.Println(number)
	}
}
