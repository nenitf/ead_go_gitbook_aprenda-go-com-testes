package main

import (
	"bytes"
	"reflect"
	"testing"
	"time"
)

/* Observações pertinentes sobre o trabalho de mockar:

Se seu código de mock estiver ficando complicado ou você tem que mockar muita
coisa para testar algo, você deve prestar mais atenção a essa sensação ruim e
pensar sobre o seu código. Geralmente isso é sinal de que:

- A coisa que você está testando está tendo que fazer coisas demais
> Modularize a função para que faça menos coisas

- Suas dependências estão muito desacopladas
> Pense e uma forma de consolidar algumas das dependências em um módulo útil

- Você está se preocupando demais com detalhes de implementação
> Dê prioridade em testar o comportamento esperado ao invés da implementação

Normalmente, muitos pontos de mock são sinais de abstração ruim no seu código.
*/

type TempoSpy struct {
	duracaoPausa time.Duration
}

func (t *TempoSpy) Pausa(duracao time.Duration) {
	t.duracaoPausa = duracao
}

// Spies (espiões) são um tipo de mock em que podemos gravar como uma
// dependência é usada
type SleeperSpy struct {
	Chamadas int
}

// implementa a interface Sleeper da main
// para criar a própria versão de sleep
func (s *SleeperSpy) Pausa() {
	s.Chamadas++
}

type SpyContagemOperacoes struct {
	Chamadas []string
}

func (s *SpyContagemOperacoes) Pausa() {
	s.Chamadas = append(s.Chamadas, pausa)
}

func (s *SpyContagemOperacoes) Write(p []byte) (n int, err error) {
	s.Chamadas = append(s.Chamadas, escrita)
	return
}

const escrita = "escrita"
const pausa = "pausa"

func TestContagem(t *testing.T) {

	t.Run("imprime 3 até Vai!", func(t *testing.T) {
		buffer := &bytes.Buffer{}
		Contagem(buffer, &SpyContagemOperacoes{})

		resultado := buffer.String()
		esperado := `3
2
1
Vai!`

		if resultado != esperado {
			t.Errorf("resultado '%s', esperado '%s'", resultado, esperado)
		}
	})

	t.Run("pausa antes de cada impressão", func(t *testing.T) {
		spyImpressoraSleep := &SpyContagemOperacoes{}
		Contagem(spyImpressoraSleep, spyImpressoraSleep)

		esperado := []string{
			pausa,
			escrita,
			pausa,
			escrita,
			pausa,
			escrita,
			pausa,
			escrita,
		}

		if !reflect.DeepEqual(esperado, spyImpressoraSleep.Chamadas) {
			t.Errorf("esperado %v chamadas, resultado %v", esperado, spyImpressoraSleep.Chamadas)
		}
	})
}

func TestSleeperConfiguravel(t *testing.T) {
	tempoPausa := 5 * time.Second

	tempoSpy := &TempoSpy{}
	sleeper := SleeperConfiguravel{tempoPausa, tempoSpy.Pausa}
	sleeper.Pausa()

	if tempoSpy.duracaoPausa != tempoPausa {
		t.Errorf("deveria ter pausado por %v, mas pausou por %v", tempoPausa, tempoSpy.duracaoPausa)
	}
}
