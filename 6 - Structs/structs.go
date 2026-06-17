package main

import "fmt"

type usuario struct {
	nome     string
	idade    uint8
	endereco endereco
}

type endereco struct {
	logradouro string
	numero     uint32
}

func main() {

	fmt.Println("Arquivo Structs")

	var u usuario

	fmt.Println(u)

	u.nome = "Marco Aurelio Tomaz"
	u.idade = 49

	enderecoExemplo := endereco{"R M F", 404}

	fmt.Println(u)

	usuario2 := usuario{"Guilherme", 22, enderecoExemplo}
	fmt.Println(usuario2)

	usuario3 := usuario{nome: "Nícolas"}
	fmt.Println(usuario3)

}
