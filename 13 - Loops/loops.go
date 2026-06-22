package main

import (
	"fmt"
	"time"
)

func main() {

	usuario := map[string]string{
		"nome":      "Leonardo",
		"sobrenome": "Silva",
	}

	for chave, valor := range usuario {
		fmt.Println(chave, valor)
	}

	nomes := [3]string{"Marco", "Aurélio", "Tomaz"}

	for indice, nome := range nomes {

		fmt.Println(indice, nome)

	}

	for _, nome := range nomes {

		fmt.Println(nome)

	}

	for {
		fmt.Println("Executando infinitamente")
		time.Sleep(time.Second)

	}

	fmt.Println("Loops")

	for j := 0; j < 10; j += 2 {
		fmt.Println("Incrementando J", j)
		time.Sleep(time.Second)
	}

	i := 0

	for i < 10 {
		i++
		fmt.Println("Incrementando I", i)
		time.Sleep(time.Second)
	}
	fmt.Println(i)

}
