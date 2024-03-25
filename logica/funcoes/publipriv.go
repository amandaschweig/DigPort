package main

import "fmt"

// função pública
func PublicFunction() {
	fmt.Println("Esta é uma função pública") // variavel acessível em todas as funções desse arquivo - Função que inicia em mausculo é global e pode ser chamada fora do pacote
}

// função privada
func privateFunction() {
	fmt.Println("Esta é uma função privada") // Função que inicia em minusculo não pode ser chamada fora do pacote
}

func main() {
	PublicFunction()  // chamada para função pública
	privateFunction() // chamada para função privada (dentro do mesmo pacote)
}
