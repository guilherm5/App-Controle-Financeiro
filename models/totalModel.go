package models

type TotalDespesasResponse struct {
	NameUser      string  `json:"name_user"`
	IDUser        int     `json:"id_user"`
	TotalDespesas float64 `json:"total_despesas"`
}

type TotalSalarioResponse struct {
	NameUserSaldo string  `json:"name_user"`
	IDUserSaldo   int     `json:"id_user"`
	TotalSaldo    float64 `json:"total_salario"`
}
