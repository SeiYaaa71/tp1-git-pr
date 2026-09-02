package main

import "fmt"

func Add(a, b int) int {
	return a + b
}

func Multiply(a, b int) int {
	return a * b
}

func main() {
	fmt.Println("Add(2, 3) =", Add(2, 3))
	fmt.Println("Multiply(4, 5) =", Multiply(4, 5))
}