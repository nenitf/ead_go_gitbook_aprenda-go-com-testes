package structs

import "testing"

func TestPerimetro(t *testing.T) {
	retangulo := Retangulo{10.0, 10.0}
	resultado := Perimetro(retangulo)
	esperado := 40.0

	if resultado != esperado {
		t.Errorf("resultado %.2f, esperado %.2f", resultado, esperado)
	}
}

// Table driven tests são úteis quando você quer construir uma lista de casos
// de testes que podem ser testados da mesma forma.
// https://github.com/golang/go/wiki/TableDrivenTests
func TestArea(t *testing.T) {
	testesArea := []struct {
		nome    string
		forma   Forma
		temArea float64
	}{
		// O teste é lido de forma mais clara, como se fosse uma afirmação da
		// verdade, não uma sequência de operações:
		// {Retangulo{12, 6}, 72.0},
		// {Circulo{10}, 314.1592653589793},
		// {Triangulo{12, 6}, 36.0},
		{nome: "Retângulo", forma: Retangulo{Largura: 12, Altura: 6}, temArea: 72.0},
		{nome: "Círculo", forma: Circulo{Raio: 10}, temArea: 314.1592653589793},
		{nome: "Triângulo", forma: Triangulo{Base: 12, Altura: 6}, temArea: 36.0},
	}

	for _, tt := range testesArea {
		t.Run(tt.nome, func(t *testing.T) {
			resultado := tt.forma.Area()
			if resultado != tt.temArea {
				t.Errorf("%#v resultado %.2f, esperado %.2f", tt.forma, resultado, tt.temArea)
			}
		})
	}
}
