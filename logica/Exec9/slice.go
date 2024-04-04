package main

import "fmt"

func main() {
	//cria uma constante com o valor "Hello"

	const hello = "Hello"

	s := make([]string, 0, 4)

	s = append(s, hello, "Amanda", "Elisa", "Brayan")

	v := s[2]

	fmt.Println("esse é meu slice:, s")
	//imprimindo a frase "Hell, Fulano"
	fmt.Println(s[0], v)
}
