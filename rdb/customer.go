package rdb

import (
	"context"
)

const (
	PrefixCustomer = "customer_"
)

type Customer struct {
	Name string `json:"name"`
}

func (r *Rdb) InsertCustomer(ctx context.Context, Id string, name string) error {
	r.db.HSet(ctx, PrefixCustomer, Id, name)
	return nil
}

func (r *Rdb) DeleteCustomerById(ctx context.Context, id string) {
	r.db.HDel(ctx, PrefixCustomer, id)
	return
}

func (r *Rdb) GetCustomer(ctx context.Context, id string) (string, error) {
	cmdReturn := r.db.HGetAll(ctx, PrefixCustomer+id)
	var c Customer
	if err := cmdReturn.Scan(&c); err != nil {
		panic(err)
	}
	return c.Name, nil
}
