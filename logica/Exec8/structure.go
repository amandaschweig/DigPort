// Esse código nomeia 3 funcinários, seus nomes, funções, idades e salários,
// e printa na tela as informações do segundo funcionário.
package main

import "fmt"

// Definindo a estrutura Funcionario
type Funcionario struct {
	Nome    string
	Funcao  string
	Idade   int
	Salario float64
}

func main() {

	// Criação de funcionário
	funcionario1 := Funcionario{
		Nome:    "Amanda Schweig",
		Funcao:  "Estagiária",
		Idade:   23,
		Salario: 500.000,
	}
	funcionario2 := Funcionario{
		Nome:    "Elisa Vilande",
		Funcao:  "Estagiária",
		Idade:   22,
		Salario: 501.000,
	}

	funcionario3 := Funcionario{
		Nome:    "Brayan Coimbra",
		Funcao:  "Estagiário",
		Idade:   23,
		Salario: 502.000,
	}

	// Imprimindo os detalhes do segundo funcionário

	fmt.Println("Nome:", funcionario2.Nome)
	fmt.Println("Função:", funcionario2.Funcao)
	fmt.Println("Idade:", funcionario2.Idade)
	fmt.Println("Salário:", funcionario2.Salario)
}
