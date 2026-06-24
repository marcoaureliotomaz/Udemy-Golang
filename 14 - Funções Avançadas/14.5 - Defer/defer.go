package main

import "fmt"

func funcao1() {
	fmt.Println("Função 1")
}

func funcao2() {
	fmt.Println("Funcao 2")
}

func alunoEstaAprovdo(n1, n2 float32) bool {
	defer fmt.Println("Média Calculada")
	fmt.Println("Entrando na função de aprovação")
	media := (n1 + n2) / 2

	if media >= 6 {

		return true
	}

	return false

}

func main() {

	fmt.Println("Defer")

	//defer funcao1()

	//funcao2()

	fmt.Println(alunoEstaAprovdo(7, 8))
}
