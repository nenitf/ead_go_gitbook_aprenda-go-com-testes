package ponteiros

import (
	"errors"
	"fmt"
)

// https://golang.org/pkg/fmt/#Stringer
type Stringer interface {
	String() string
}

type Bitcoin int

type Carteira struct {
	saldo Bitcoin
}

// Carteira atualiza o próprio saldo
// func (c Carteira) Depositar(quantidade int) {
// func (c *Carteira) Depositar(quantidade int) {
func (c *Carteira) Depositar(quantidade Bitcoin) {
	c.saldo += quantidade
}

var ErroSaldoInsuficiente = errors.New("não é possível retirar: saldo insuficiente")

func (c *Carteira) Retirar(quantidade Bitcoin) error {
	if quantidade > c.saldo {
		return ErroSaldoInsuficiente
	}

	c.saldo -= quantidade
	return nil
}

// func (c Carteira) Saldo() int {
// func (c *Carteira) Saldo() int {
func (c *Carteira) Saldo() Bitcoin {
	return c.saldo
}

func (b Bitcoin) String() string {
	return fmt.Sprintf("%d BTC", b)
}
