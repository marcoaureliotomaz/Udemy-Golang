package main

import (
	"errors"
	"fmt"
)

func main() {
	//NUMEROS INTEIROS
	var numero int16 = 10000
	fmt.Println(numero)

	var numero2 uint = 1
	fmt.Println(numero2)

	//Alias INT32 = rune
	var numero3 rune = 123456
	fmt.Println(numero3)

	//Alias INT8 = byte
	var numero4 byte = 123
	fmt.Println(numero4)

	//NUMEROS REAIS
	var numeroReal1 float32 = 1234567890.99
	fmt.Println(numeroReal1)

	var numeroReal2 float64 = 1234567891234567890.45
	fmt.Println(numeroReal2)

	//STRING
	var str string = "Texto"
	fmt.Println(str)

	str2 := "Texto2"
	fmt.Println(str2)

	//Tabela ASCII
	char := 'B'
	fmt.Println(char)

	var texto string
	fmt.Println(texto)

	var inteiro int16
	fmt.Println(inteiro)

	var boolean1 bool = true
	fmt.Println(boolean1)

	var boolean2 bool
	fmt.Println(boolean2)

	var erro error
	fmt.Println(erro)

	var erro1 error = errors.New("Erro interno")
	fmt.Println(erro1)

}
