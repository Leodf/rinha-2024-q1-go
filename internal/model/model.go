package model

import (
	"fmt"
	"time"

	"github.com/Leodf/leodf-go/internal/db"
	"github.com/Leodf/leodf-go/internal/dto"
)

var (
	sqlFunction = map[string]string{"d": "debito", "c": "credito"}
)

func SaveTransaction(id int, body *dto.TransactionRequest) (*dto.TransactionResponse, error) {
	var balance int
	var limit int

	callsqlfunction := fmt.Sprintf("SELECT saldo_atual, limite FROM %s($1, $2, $3)", sqlFunction[body.Type])

	err := db.PG.QueryRow(callsqlfunction, id, body.Value, body.Description).Scan(&balance, &limit)
	if err != nil {
		if err.Error() == "pq: saldo insuficiente" {
			return nil, err
		}
		fmt.Println(err.Error())
	}

	output := &dto.TransactionResponse{
		Limit:   limit,
		Balance: balance,
	}

	return output, nil
}

func GetClientBalance(id int) (*dto.StatmentResponse, error) {

	var balance int
	var limit int
	date := time.Now()

	err := db.PG.QueryRow("SELECT saldo, limite FROM clientes WHERE id = $1 FOR UPDATE", id).Scan(&balance, &limit)
	if err != nil {
		return nil, err
	}
	statmentHead := &dto.StatmentHead{
		Total: balance,
		Limit: limit,
		Date:  date,
	}
	rows, err := db.PG.Query(`
		SELECT valor, tipo, descricao, realizada_em 
		FROM transacoes 
		WHERE cliente_id = $1 
		ORDER BY realizada_em DESC 
		LIMIT 10
	`, id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var lastTransactions []*dto.LastTransactions
	for rows.Next() {
		var i dto.LastTransactions
		if err := rows.Scan(
			&i.Value,
			&i.Type,
			&i.Description,
			&i.CreatedAt,
		); err != nil {
			return nil, err
		}
		lastTransactions = append(lastTransactions, &i)
	}
	output := &dto.StatmentResponse{
		StatmentHead:     statmentHead,
		LastTransactions: lastTransactions,
	}
	return output, nil

}
