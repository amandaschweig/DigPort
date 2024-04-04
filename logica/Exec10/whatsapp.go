package main

import "fmt"

func main() {
	// criando um mapa de contatos
	listaDeContatos := map[string]string{
		"Amanda":  "Leitura",
		"Giovane": "Futebol",
		"Isabel":  "Escrita",
		"Brayan":  "Música",
		"Sara":    "Maquiagem",
	}

	for nome, hobby := range listaDeContatos {
		fmt.Printf("O hobby de %s é %s.\n", nome, hobby)
	}

}
