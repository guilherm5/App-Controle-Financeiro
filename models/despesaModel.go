package models

import (
	"time"

	"github.com/shopspring/decimal"
)

type Despesas struct {
	IDDespesa        int             `json:"id_despesa"`
	NomeDespesa      string          `json:"nome_despesa"`
	DescricaoDespesa string          `json:"descricao_despesa"`
	MesDespesa       time.Time       `json:"mes_despesa"`
	Valor            decimal.Decimal `json:"valor"`
	NameUser         string          `json:"name_user"`
	IDUser           int             `json:"id_user"`
}
