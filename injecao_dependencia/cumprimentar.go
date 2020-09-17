package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

// io.Writer é uma interface
// Programar voltado para interface e não para a implementação
func Cumprimenta(escritor io.Writer, nome string) {
	fmt.Fprintf(escritor, "Olá, %s", nome)
}

func HandlerMeuCumprimento(w http.ResponseWriter, r *http.Request) {
	Cumprimenta(w, "mundo")
}

func main() {
	Cumprimenta(os.Stdout, "Elodie")
	err := http.ListenAndServe(":5000", http.HandlerFunc(HandlerMeuCumprimento))

	if err != nil {
		fmt.Println(err)
	}
}
