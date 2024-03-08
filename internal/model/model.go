package model

import (
	"context"
	"time"

	"github.com/Leodf/leodf-go/internal/db"
	"github.com/Leodf/leodf-go/internal/dto"
)

const (
	transaction string = "SELECT saldo_atual, limite FROM transacao($1, $2, $3, $4)"
	client      string = "SELECT saldo, limite FROM clientes WHERE id = $1"
	statment    string = `
		SELECT valor, tipo, descricao, realizada_em
		FROM transacoes t
		WHERE t.cliente_id = $1
		ORDER BY realizada_em DESC
		LIMIT 10
	`
)

func SaveTransaction(ctx context.Context, id int, body *dto.TransactionRequest) (*dto.TransactionResponse, error) {
	var balance int
	var limit int
	err := db.PG.QueryRow(ctx, transaction, id, body.Value, body.Type, body.Description).Scan(&balance, &limit)
	if err != nil {
		return nil, err
	}
	output := &dto.TransactionResponse{
		Limit:   limit,
		Balance: balance,
	}
	return output, nil
}

func GetClientBalance(ctx context.Context, id int) (*dto.StatmentResponse, error) {
	var limit, balance int
	err := db.PG.QueryRow(ctx, client, id).Scan(&balance, &limit)
	if err != nil {
		return nil, err
	}
	rows, err := db.PG.Query(ctx, statment, id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	lastTransactions := make([]*dto.LastTransactions, 0)
	for rows.Next() {
		var lt dto.LastTransactions
		if err := rows.Scan(
			&lt.Value,
			&lt.Type,
			&lt.Description,
			&lt.CreatedAt,
		); err != nil {
			return nil, err
		}
		lastTransactions = append(lastTransactions, &lt)
	}

	statHead := &dto.StatmentHead{
		Total: balance,
		Limit: limit,
		Date:  time.Now(),
	}
	output := &dto.StatmentResponse{
		StatmentHead:     statHead,
		LastTransactions: lastTransactions,
	}
	return output, nil
}
