package pg

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v5"
	"job-test/types"
	"math/big"
	"time"
)

func (pg *Postgres) InsertDepositHistory(ctx context.Context, id, denom string, amount big.Int, opType types.OpType, memo string) error {
	query := `INSERT INTO deposit_history (id, denom, amount, date, type, memo) VALUES (@id, @denom, @amount, @date, @type, @memo)`
	args := pgx.NamedArgs{
		"id":     id,
		"denom":  denom,
		"amount": amount.String(),
		"date":   time.Now(),
		"type":   opType,
		"memo":   memo,
	}
	_, err := pg.db.Exec(ctx, query, args)
	if err != nil {
		return fmt.Errorf("unable to insert row: %w", err)
	}

	return nil
}

func (pg *Postgres) GetDepositHistoryByCustomer(ctx context.Context, id string) ([]types.DepositHistoryItem, error) {
	query := `SELECT * from deposit_history where id = @id`
	args := pgx.NamedArgs{
		"id": id,
	}
	rows, err := pg.db.Query(ctx, query, args)
	if err != nil {
		return nil, fmt.Errorf("unable to insert row: %w", err)
	}

	var historyItems []types.DepositHistoryItem
	for rows.Next() {
		item := types.DepositHistoryItem{}
		err := rows.Scan(&item.CustomerId, &item.Denom, &item.Amount, &item.Date, &item.Type)
		if err != nil {
			return nil, fmt.Errorf("unable to scan row: %w", err)
		}
		historyItems = append(historyItems, item)
	}
	return historyItems, nil
}

func (pg *Postgres) InsertSendHistory(ctx context.Context, sender, receiver, denom string, amount big.Int, memo string) error {
	query := `INSERT INTO send_history (sender, receiver, denom, amount, date, memo) VALUES (@sender, @receiver, @denom, @amount, @date, @memo)`
	args := pgx.NamedArgs{
		"sender":   sender,
		"receiver": receiver,
		"denom":    denom,
		"amount":   amount.String(),
		"date":     time.Now().String(),
		"memo":     memo,
	}
	_, err := pg.db.Exec(ctx, query, args)
	if err != nil {
		return fmt.Errorf("unable to insert row: %w", err)
	}

	return nil
}

func (pg *Postgres) GetSendHistoryByCustomer(ctx context.Context, snStart uint64, customerId string) ([]types.SendHistoryItem, error) {
	query := `SELECT * FROM send_history where @sn >= sn LIMIT 10 `

	args := pgx.NamedArgs{
		"customer_id": customerId,
	}
	rows, err := pg.db.Query(ctx, query, args)
	if err != nil {
		return nil, fmt.Errorf("unable to query users: %w", err)
	}
	defer rows.Close()

	var historyItems []types.SendHistoryItem
	for rows.Next() {
		item := types.SendHistoryItem{}
		err := rows.Scan(&item.Serial, &item.Sender, &item.Receiver, &item.Denom, &item.Amount)
		if err != nil {
			return nil, fmt.Errorf("unable to scan row: %w", err)
		}
		historyItems = append(historyItems, item)
	}

	return historyItems, nil
}
