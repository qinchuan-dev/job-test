package pg

import (
	"context"
	"errors"
	"fmt"
	"github.com/jackc/pgx/v5"
)

func (pg *Postgres) InsertCustomer(ctx context.Context, Id string, name string) error {
	query := `INSERT INTO customer (id, name) VALUES (@id, @name)`
	args := pgx.NamedArgs{
		"id":   Id,
		"name": name,
	}
	_, err := pg.db.Exec(ctx, query, args)
	if err != nil {
		return fmt.Errorf("unable to insert row: %w", err)
	}

	return nil
}

func (pg *Postgres) DeleteCustomerById(ctx context.Context, id string) error {
	query := `DELETE from customer where id=@id`
	args := pgx.NamedArgs{
		"id": id,
	}
	_, err := pg.db.Exec(ctx, query, args)
	if err != nil {
		return fmt.Errorf("unable to delete row: %w", err)
	}

	return nil
}

func (pg *Postgres) GetCustomer(ctx context.Context, id string) error {
	query := `SELECT * FROM customer where id=@id`
	args := pgx.NamedArgs{
		"id": id,
	}
	row := pg.db.QueryRow(ctx, query, args)
	err := row.Scan()
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil
		} else {
			return fmt.Errorf("unable to get row: %w", err)
		}
	}

	return nil
}
