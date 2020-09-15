package main

import "fmt"

const prefixoOlaPortugues = "Olá, "
const prefixoOlaIngles = "Hello, "
const prefixoOlaEspanhol = "Hola, "

func Ola(nome string, idioma string) string {
	if nome == "" {
		nome = "Mundo"
	}

	prefixo := prefixoOlaPortugues

	switch idioma {
	case "ingles":
		prefixo = prefixoOlaIngles
	case "espanhol":
		prefixo = prefixoOlaEspanhol
	}

	return prefixo + nome
}

func main() {
	fmt.Println(Ola("main", ""))
}
