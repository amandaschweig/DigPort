// Esse exercicio solicita que o usuário escolha um numero e some com o numero da variavel que foi declarada por mim (6).

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
