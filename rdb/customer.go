package rdb

import (
	"context"
)

const (
	PrefixCustomer = "customer:"
)

func (r *Rdb) InsertCustomer(ctx context.Context, id string, name string) error {
	_, err := r.db.Set(ctx, PrefixCustomer+id, name, 0).Result()
	return err
}

func (r *Rdb) DeleteCustomerById(ctx context.Context, id string) error {
	_, err := r.db.Del(ctx, PrefixCustomer+id).Result()
	return err
}

func (r *Rdb) GetCustomer(ctx context.Context, id string) (string, error) {
	return r.db.Get(ctx, PrefixCustomer+id).Result()
}
