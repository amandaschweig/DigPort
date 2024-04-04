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
	fmt.Printf("Seu IMC é: %.2f\n", imc)
}
