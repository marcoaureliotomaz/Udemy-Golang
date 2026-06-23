package main

import "fmt"

func soma(numeros ...int) int {
	total := 0

	for _, numero := range numeros {
		total += numero
	}
	return total
}

func escrever(texto string, numeros ...int) {
	for _, numero := range numeros {
		fmt.Println(texto, numero)
	}
}

func main() {

	fmt.Println(soma(1, 2, 3, 5, 7, 9, 11))
	escrever("Marco", 9, 8, 7)
}
