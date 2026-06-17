package main

import "fmt"

func main() {

	//OPERACORES ARITIMETICOS
	soma := 1 + 1
	subtracao := 1 - 1
	divisao := 10 / 4
	multiplicacao := 10 * 4
	restoDaDivisao := 10 % 3

	fmt.Println(soma, subtracao, divisao, multiplicacao, restoDaDivisao)

	var numero1 int16 = 10
	var numero2 int16 = 20

	soma2 := numero1 + numero2
	fmt.Println(soma2)

	//ATRIBUICAO

	var variavel1 string = "String"
	variavel2 := "String2"

	fmt.Println(variavel1, variavel2)

	//OPERADORES RELACIONAIS
	fmt.Println(1 > 2)
	fmt.Println(1 > 2)
	fmt.Println(1 >= 2)
	fmt.Println(1 == 2)
	fmt.Println(1 < 2)
	fmt.Println(1 <= 2)
	fmt.Println(1 != 2)

	//OPERADORES LOGICOS
	fmt.Println("-------------------")
	verdadeiro, falso := true, false
	fmt.Println(verdadeiro && falso)
	fmt.Println(verdadeiro || falso)
	fmt.Println(!verdadeiro)

	//OPERADORES UNARIOS

	numero := 30
	numero++
	numero += 15
	fmt.Println(numero)

	numero--
	numero -= 20
	fmt.Println(numero)

	numero *= 3
	fmt.Println(numero)

	numero /= 5
	fmt.Println(numero)

	numero %= 2
	fmt.Println(numero)

}
