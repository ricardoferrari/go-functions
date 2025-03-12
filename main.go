package main

import (
	"fmt"
)

type transformationFunction func(int) int

func main() {
	numbers := []int{1, 2, 3, 4, 5}
	fmt.Println(transformNumbers(numbers, addOne))
}

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
