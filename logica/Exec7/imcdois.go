// Esse exercicio calcula o IMC e gera números aleatórios. Este programa contem uma CLASSE que calcule as duas operações:
// O IMC De uma pessoa e a Fórmula: IMC = Peso ÷ (Altura × Altura)

package main

import (
	"fmt"
)

func main() {
	var peso, altura float64

	fmt.Printf("Qual é a sua altura em metros? ")
	fmt.Scan(&altura)

	// verifica se altura é válida
	if altura <= 0 {
		fmt.Printf("A altura não pode ser zero ou menor que zero")
		return // encerra o programa se a altura não for válida
	}

	// Solicita peso em quilogramas
	fmt.Printf("Quantos quilos você pesa em quilogramas? ")
	fmt.Scan(&peso)

	// calcula IMC
	imc := peso / (altura * altura)

	// analisa o IMC e imprime o resultado correspondente
	if imc < 18.5 {
		fmt.Println("Você está abaixo do peso, seu imc é:", imc)
	} else if imc >= 18.5 && imc <= 24.9 {
		fmt.Println("Você está no peso ideal. Parabéns! Seu imc é:", imc)
	} else if imc >= 25.0 && imc <= 29.9 {
		fmt.Println("Você está levemente acima do peso. Seu imc é:", imc)
	} else if imc >= 30.0 && imc <= 34.9 {
		fmt.Println("Você está com obesidade grau I. Seu imc é:", imc)
	} else if imc >= 35.0 && imc <= 39.9 {
		fmt.Println("Você está com obesidade grau II (severa). Seu imc é:", imc)
	} else {
		fmt.Println("Você está com obesidade grau III (Mórbida). Seu imc é:", imc)
	}
}
