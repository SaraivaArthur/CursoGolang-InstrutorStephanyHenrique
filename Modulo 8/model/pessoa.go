package model

import (
	"fmt"
	"time"
)

type Pessoa struct {
	Nome             string
	Endereco         Endereco
	DataDeNascimento time.Time
	Idade            int
}

func (p *Pessoa) CalculaIdade() int {
	anoDeNascimento := p.DataDeNascimento.Year()
	anoAtual := time.Now().Year()
	p.Idade = anoAtual - anoDeNascimento
	fmt.Println("Idade: ", p.Idade)
	return p.Idade
}
