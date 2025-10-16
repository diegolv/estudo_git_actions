package main

import "fmt"

func main() {
	fmt.Printf("O resultado é: %d", Soma(10, 10))
}

func Soma(a, b int) int {
	return a + b
}
