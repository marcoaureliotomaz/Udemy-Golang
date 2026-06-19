package main

import "fmt"

func main() {
	fmt.Println("Estrutura de controle")

	numero := 10

	if numero > 15 {
		fmt.Println("Maior que 15")
	} else {
		fmt.Println("Menor ou igual a 15")
	}

	if outroNumero := numero - 10; outroNumero > 15 {
		fmt.Println("Numero é maior que 0")
	} else {
		fmt.Println("Numero é menor que 0")
	}

}
