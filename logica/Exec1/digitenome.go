//Esse exercício solicita o nome do usuário, e escaneia o mesmo para dizer "Bem vindo, 'fulano'"

package main

import (
	"fmt"
)

func main() {
	var name string

	fmt.Println("Digite seu nome: ")
	fmt.Scan(&name)
	fmt.Println("Bem vindo, " + name)

}
