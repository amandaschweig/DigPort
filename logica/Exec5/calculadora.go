//Esse exercicio é uma calculadora, com operações de soma, divisão, subtração e multiplicação.

package main

import (
	"fmt"
)

func main() {
	var numero1, numero2, operacao, resultado int

	fmt.Println("Escolha um numero: ")
	fmt.Scan(&numero1)

	fmt.Println("Escolha outro numero: ")
	fmt.Scan(&numero2)

	fmt.Println(`Qual operação você quer realizar:
	1. Soma
	2. Subtração
	3. Multiplicação
	4. Divisão`)
	fmt.Scan(&operacao)

	// Verificação se a operação está dentro do intervalo válido
	if operacao < 1 || operacao > 4 {
		fmt.Println("Operação inválida!")
		return
	}

	if operacao == 1 {
		resultado = Somar(numero1, numero2)
	} else if operacao == 2 {
		resultado = numero1 - numero2
	} else if operacao == 3 {
		resultado = numero1 * numero2
	} else if operacao == 4 {
		if numero2 != 0 {
			resultado = numero1 / numero2
		} else {
			fmt.Println("Não é possível dividir por zero.")
			return
		}
	}

	fmt.Println("O resultado é:", resultado)
}

func Somar(num1, num2 int) int {
	return num1 + num2
}
