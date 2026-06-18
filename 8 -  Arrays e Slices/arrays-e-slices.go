package main

import (
	"fmt"
	"reflect"
)

func main() {
	fmt.Println("Arrays e Slices")

	var array1 [5]string

	array1[0] = "Teste"

	fmt.Println(array1)

	array2 := [5]string{"98as$=", "@", "--", "ad", "654"}
	fmt.Println(array2)

	array3 := [...]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	fmt.Println(array3)

	slice := []int{1, 2, 32, 4, 5, 66, 7, 88, 99}

	fmt.Println(slice)

	fmt.Println(reflect.TypeOf(slice))
	fmt.Println(reflect.TypeOf(array3))

	slice = append(slice, 100)
	fmt.Println(slice)

	slice2 := array2[1:3]
	fmt.Println(slice2)

	array2[1] = "Pocição alterada"
	fmt.Println(slice2)
}
