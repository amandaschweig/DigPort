package main

import (
	"fmt"
)

func main() {
	var numero int
	fmt.Printf("Escolha um numero: ")
	fmt.Scanf("%d", &numero)

	if numero > 0 {
		fmt.Println("POSITIVO!")
	} else if numero < 0 {
		fmt.Println("NEGATIVO!")
	} else {
		fmt.Println("Zero!")
	}
}
