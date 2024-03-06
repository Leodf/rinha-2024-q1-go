package dto

import "time"

type TransactionRequest struct {
	Value       int    `json:"valor" validate:"required,gt=0"`
	Type        string `json:"tipo" validate:"required,oneof=c d"`
	Description string `json:"descricao" validate:"required,min=1,max=10"`
}

type TransactionResponse struct {
	Limit   int `json:"limite"`
	Balance int `json:"saldo"`
}

type StatmentHead struct {
	Total int       `json:"total"`
	Limit int       `json:"limite"`
	Date  time.Time `json:"data_extrato"`
}

type LastTransactions struct {
	Value       int       `json:"valor"`
	Type        string    `json:"tipo"`
	Description string    `json:"descricao"`
	CreatedAt   time.Time `json:"realizada_em"`
}

type StatmentResponse struct {
	StatmentHead     *StatmentHead       `json:"saldo"`
	LastTransactions []*LastTransactions `json:"ultimas_transacoes"`
}
