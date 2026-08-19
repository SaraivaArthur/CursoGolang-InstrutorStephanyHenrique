package model

import "time"

type Pessoa struct {
	Nome             string
	Endereco         Endereco
	DataDeNascimento time.Time
	Idade            int
}

func (p Pessoa) IdadeAtual() int {
	anoDeNascimento := p.DataDeNascimento.Year()
	anoAtual := time.Now().Year()
	return anoAtual - anoDeNascimento
}

func (p *Pessoa) CalculaIdade() int {
	anoDeNascimento := p.DataDeNascimento.Year()
	anoAtual := time.Now().Year()
	return anoAtual - anoDeNascimento
}
