package models

import (
	"time"

	"github.com/shopspring/decimal"
)

type Salario struct {
	IDSalario  int             `json:"id_salario"`
	MesSalario time.Time       `json:"mes_salario`
	Salario    decimal.Decimal `json:"salario"`
	RendaExtra decimal.Decimal `json:"renda_extra"`
	NameUser   string          `json:"name_user"`
	IDUser     int             `json:"id_user"`
}
