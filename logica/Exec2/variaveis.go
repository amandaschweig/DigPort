package main

import (
	"fmt"
)

func main() {
	var b int
	a := 6
	fmt.Printf("Escolha um numero: ")
	fmt.Scanf("%d", &b)
	soma := a + b // ler o numero

	fmt.Println("O resultado é: ", soma) // escrever o resultado do numero
}
