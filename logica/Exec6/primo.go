package main

import (
	"fmt"
)

func main() {
	var numero int

	fmt.Print("Escolha um numero: ")
	fmt.Scan(&numero)

	if isPrime(numero) {
		fmt.Print("Primo!")
	} else {
		fmt.Print("Não é primo!")
	}
}

func isPrime(n int) bool {
	if n <= 1 {
		return false
	}
	for i := 2; i*i <= n; i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}
