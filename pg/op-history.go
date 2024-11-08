package pg

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v5"
	"job-test/types"
	"math/big"
	"time"
)

func (pg *Postgres) InsertDepositHistory(ctx context.Context, id, denom string, amount big.Int, date time.Time, opType types.OpType, memo string) error {
	query := `INSERT INTO deposit_history (id, denom, amount, date, type, memo) VALUES (@id, @denom, @amount, @date, @type, @memo)`
	args := pgx.NamedArgs{
		"id":     id,
		"denom":  denom,
		"amount": amount.String(),
		"date":   date,
		"type":   opType,
		"memo":   memo,
	}
	_, err := pg.db.Exec(ctx, query, args)
	if err != nil {
		return fmt.Errorf("unable to insert row: %w", err)
	}

	return nil
}

func (pg *Postgres) GetDepositHistoryByCustomer(ctx context.Context, id string) ([]types.DepositHistory, error) {
	query := `SELECT id, denom, amount, type, date, memo from deposit_history where id = @id`
	args := pgx.NamedArgs{
		"id": id,
	}
	rows, err := pg.db.Query(ctx, query, args)
	if err != nil {
		return nil, fmt.Errorf("unable to insert row: %w", err)
	}

	var historyItems []types.DepositHistory
	for rows.Next() {
		item := types.DepositHistory{}
		err := rows.Scan(&item.Id, &item.Denom, &item.Amount, &item.Date, &item.OpType, &item.Memo)
		if err != nil {
			return nil, fmt.Errorf("unable to scan row: %w", err)
		}
		historyItems = append(historyItems, item)
	}
	return historyItems, nil
}

func (pg *Postgres) InsertSendHistory(ctx context.Context, sender, receiver, denom string, amount big.Int, date time.Time, memo string) error {
	query := `INSERT INTO send_history (sender, receiver, denom, amount, date, memo) VALUES (@sender, @receiver, @denom, @amount, @date, @memo)`
	args := pgx.NamedArgs{
		"sender":   sender,
		"receiver": receiver,
		"denom":    denom,
		"amount":   amount.String(),
		"date":     date,
		"memo":     memo,
	}
	_, err := pg.db.Exec(ctx, query, args)
	if err != nil {
		return fmt.Errorf("unable to insert row: %w", err)
	}

	return nil
}

func (pg *Postgres) GetSendHistoryByCustomer(ctx context.Context, sender string) ([]types.SendHistory, error) {
	query := `SELECT sender, receiver, denom, amount, date, memo FROM send_history where sender = @sender`

	args := pgx.NamedArgs{
		"sender": sender,
	}
	rows, err := pg.db.Query(ctx, query, args)
	if err != nil {
		return nil, fmt.Errorf("unable to query users: %w", err)
	}
	defer rows.Close()

	var historyItems []types.SendHistory
	for rows.Next() {
		item := types.SendHistory{}
		err := rows.Scan(&item.Sender, &item.Receiver, &item.Denom, &item.Amount, &item.Date, &item.Memo)
		if err != nil {
			return nil, fmt.Errorf("unable to scan row: %w", err)
		}
		historyItems = append(historyItems, item)
	}

	return historyItems, nil
}
