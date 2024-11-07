package pg

import (
	"context"
	"errors"
	"fmt"
	"github.com/jackc/pgx/v5"
	"job-test/types"
	"math/big"
)

func (pg *Postgres) InsertDeposit(ctx context.Context, id string, denom string, amount big.Int) error {
	query := `INSERT INTO deposit (id, denom, amount) VALUES (@id, @denom, @amount)`
	args := pgx.NamedArgs{
		"id":     id,
		"denom":  denom,
		"amount": amount.String(),
	}
	_, err := pg.db.Exec(ctx, query, args)
	if err != nil {
		return fmt.Errorf("unable to insert row: %w", err)
	}

	return nil
}

func (pg *Postgres) DeleteDeposit(ctx context.Context, id string, denom string) error {
	query := `DELETE FROM deposit where id=@id and denom=@denom`
	args := pgx.NamedArgs{
		"id":    id,
		"denom": denom,
	}
	_, err := pg.db.Exec(ctx, query, args)
	if err != nil {
		return fmt.Errorf("unable to delete row: %w", err)
	}

	return nil
}

func (pg *Postgres) SetDepositAmount(ctx context.Context, id string, denom string, amount big.Int) error {
	query := `UPDATE deposit SET amount = @amount WHERE id = @id and denom = @denom`
	args := pgx.NamedArgs{
		"amount": amount.String(),
		"id":     id,
		"denom":  denom,
	}
	_, err := pg.db.Exec(ctx, query, args)
	if err != nil {
		return fmt.Errorf("unable to update row: %w", err)
	}

	return nil
}

func (pg *Postgres) Deposit(ctx context.Context, id string, denom string, deltaAmt big.Int) error {
	query := `select amount from deposit where id = @id and denom = @denom`
	args := pgx.NamedArgs{
		"id":    id,
		"denom": denom,
	}
	row := pg.db.QueryRow(ctx, query, args)

	oldAmtStr := ""
	err := row.Scan(&oldAmtStr)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return pg.InsertDeposit(ctx, id, denom, deltaAmt)
		} else {
			return err
		}
	} else {
		oldAmt, ok := new(big.Int).SetString(oldAmtStr, 0)
		if !ok {
			return fmt.Errorf("invalid old amount %s", oldAmtStr)
		}
		newAmt := oldAmt.Add(oldAmt, &deltaAmt)

		return pg.SetDepositAmount(ctx, id, denom, *newAmt)
	}
}

func (pg *Postgres) Withdraw(ctx context.Context, id string, denom string, deltaAmt big.Int) error {
	query := `select amount from deposit where id = @id and denom = @denom`
	args := pgx.NamedArgs{
		"id":    id,
		"denom": denom,
	}
	row := pg.db.QueryRow(ctx, query, args)

	oldAmtStr := ""
	err := row.Scan(&oldAmtStr)
	if err != nil {
		return err
	} else {
		oldAmt, ok := new(big.Int).SetString(oldAmtStr, 0)
		if !ok {
			return fmt.Errorf("invalid old amount %s", oldAmtStr)
		}

		flag := oldAmt.Cmp(&deltaAmt)
		if -1 == flag {
			return fmt.Errorf("not enough deposit amount to withdraw:old=%s, delta=%s", oldAmtStr, deltaAmt.String())
		} else if 0 == flag {
			return pg.DeleteDeposit(ctx, id, denom)
		} else {
			newAmt := oldAmt.Sub(oldAmt, &deltaAmt)
			return pg.SetDepositAmount(ctx, id, denom, *newAmt)
		}
	}
}

func (pg *Postgres) GetDepositByCustomer(ctx context.Context, id string) ([]types.DepositItem, error) {
	query := `SELECT * from deposit where id=@id`

	args := pgx.NamedArgs{
		"id": id,
	}

	rows, err := pg.db.Query(ctx, query, args)
	if err != nil {
		return nil, fmt.Errorf("unable to query deposit: %w", err)
	}
	defer rows.Close()

	var items []types.DepositItem
	for rows.Next() {
		item := types.DepositItem{}

		err := rows.Scan(&item.Denom, &item.Amount)
		if err != nil {
			return nil, fmt.Errorf("unable to scan row: %w", err)
		}

		items = append(items, item)
	}

	return items, nil
}
