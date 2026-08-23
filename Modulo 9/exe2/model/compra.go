package model

import (
	"errors"
	"time"
)

type Compra struct {
	Mercado string
	Data    time.Time
	Itens   []ItemCompra
}

type ItemCompra struct {
	Nome string
}

func NewCompra(mercado string, data time.Time, nomeDosItens []string) (*Compra, error) {
	var itens []ItemCompra

	if mercado == "" {
		return nil, errors.New("Mercado é obrigatório")
	}

	if len(nomeDosItens) == 0 {
		return nil, errors.New("Itens são obrigatórios")
	}

	for _, nome := range nomeDosItens {
		itens = append(itens, ItemCompra{Nome: nome})
	}

	return &Compra{
		Mercado: mercado,
		Data:    time.Now(),
		Itens:   itens,
	}, nil
}
