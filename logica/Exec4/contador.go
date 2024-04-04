//Esse exercicio é um contador, e faz contagem regressiva e progressiva.

package main

import (
	"fmt"
)

func main() {
	// diminuindo
	for contador := 10; contador > 0; contador-- {
		fmt.Println("Contagem Regressiva", contador)
	}
	{
		// aumentando
		for contador := 0; contador < 10; contador++ {
			fmt.Println("Lá vamos nós", contador)
		}
	}
	fmt.Printf("Arrasou!")

}
