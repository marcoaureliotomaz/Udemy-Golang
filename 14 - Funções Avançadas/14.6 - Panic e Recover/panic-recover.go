package main

import "fmt"

func alunoEstaAprovado(n1, n2 float64) bool {

	defer recuperarExecucao()
	media := (n1 + n2) / 2

	if media > 6 {
		return true
	} else if media < 6 {
		return false
	}

	panic("A mpedia é exatamente 6")

}

func recuperarExecucao() {

	if r := recover(); r != nil {
		fmt.Println("Tentantiva de recuperar")
	}
}

func main() {

	fmt.Println("Panic Recover")
	alunoEstaAprovado(6, 6)
	fmt.Println("pós execução")
}
