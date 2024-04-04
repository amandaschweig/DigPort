package main

import (
	"fmt"
	"unicode"
)

// Variável acessível em todas as funções desse arquivo
var FraseDoUsuario string

// Função que inicia em maiúsculo é global e pode ser chamada fora do pacote
func Publica(nomeDoUsuario string, idade int, numeroDaCor int) {
	corPreferida := encontrarCorPreferida(numeroDaCor) // Aqui está a correção
	nomeDoUsuarioCorrigido := colocarPrimeiraLetraEmMaiusculo(nomeDoUsuario)

	fraseFormatada := fmt.Sprintf("Olá! Meu nome é %s, minha idade é %d e minha cor preferida é %s",
		nomeDoUsuarioCorrigido,
		idade,
		corPreferida,
	)

	FraseDoUsuario = fraseFormatada
}

// Função que inicia em minúsculo não pode ser chamada fora do pacote
func encontrarCorPreferida(numeroDaCor int) string {
	if numeroDaCor == 1 {
		return "laranja"
	} else if numeroDaCor == 2 {
		return "roxo"
	} else if numeroDaCor == 3 {
		return "rosa"
	} else if numeroDaCor == 4 {
		return "branco"
	} else if numeroDaCor == 5 {
		return "preto"
	} else {
		return "nenhuma"
	}
}

func colocarPrimeiraLetraEmMaiusculo(nome string) string {
	if len(nome) == 0 {
		return ""
	}

	runas := []rune(nome)
	runas[0] = unicode.ToUpper(runas[0])

	for i := 1; i < len(runas); i++ {
		runas[i] = unicode.ToLower(runas[i])
	}

	return string(runas)
}

func second() {
	Publica("carlos", 30, 2)
	fmt.Println(FraseDoUsuario)
}
