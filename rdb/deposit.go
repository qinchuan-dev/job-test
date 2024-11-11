package rdb

import (
	"context"
	"errors"
	"github.com/redis/go-redis/v9"
	"job-test/types"
	"math/big"
)

const (
	PrefixDeposit = "deposit:"
)

func (r *Rdb) SetDeposit(ctx context.Context, id string, denom string, amt big.Int) error {
	_, err := r.db.Set(ctx, PrefixDeposit+id+":"+denom+":", amt.String(), 0).Result()
	return err
}

func (r *Rdb) DeleteDeposit(ctx context.Context, id string, denom string) error {
	_, err := r.db.Del(ctx, PrefixDeposit+id+":"+denom+":").Result()
	return err
}

func (r *Rdb) Deposit(ctx context.Context, id string, denom string, deltaAmt big.Int) error {
	k := PrefixDeposit + id + ":" + denom + ":"

	oldAmtStr, err := r.db.Get(ctx, k).Result()
	if err != nil {
		if errors.Is(err, redis.Nil) {
			_, err := r.db.Set(ctx, k, deltaAmt.String(), 0).Result()
			if err != nil {
				return err
			}
		} else {
			return err
		}
	}
	oldAmt, ok := new(big.Int).SetString(oldAmtStr, 0)
	if !ok {
		return errors.New("invalid old amount")
	}
	newAmt := oldAmt.Add(oldAmt, &deltaAmt)

	_, err = r.db.Set(ctx, k, newAmt.String(), 0).Result()
	if err != nil {
		return err
	}

	return nil
}

func (r *Rdb) Withdraw(ctx context.Context, id string, denom string, deltaAmt big.Int) error {
	oldAmtStr, err := r.db.Get(ctx, PrefixDeposit+id+":"+denom+":").Result()
	if err != nil {
		return err
	}
	oldAmt, ok := new(big.Int).SetString(oldAmtStr, 0)
	if !ok {
		return errors.New("invalid old amount")
	}

	newAmt := oldAmt.Sub(oldAmt, &deltaAmt)

	r.db.Set(ctx, PrefixDeposit+id+":"+denom+":", newAmt.String(), 0)

	return nil
}

func (r *Rdb) GetDepositByCustomerDenom(ctx context.Context, id, denom string) (types.DepositItem, error) {
	amtStr, err := r.db.Get(ctx, PrefixDeposit+id+":"+denom+":").Result()
	if err != nil {
		return types.DepositItem{}, err
	}

	var item types.DepositItem
	item.Denom = denom
	item.Amount = amtStr

	return item, nil
}

func (r *Rdb) GetDepositByCustomer(ctx context.Context, id string) ([]types.DepositItem, error) {
	var cursor uint64 = 0
	var keys []string

	for {
		var scanKeys []string
		var err error
		pattern := PrefixDeposit + id + ":*"

		scanKeys, cursor, err = r.db.Scan(ctx, cursor, pattern, 0).Result()
		if err != nil {
			return nil, err
		}
		keys = append(keys, scanKeys...)
		if cursor == 0 {
			break
		}
	}

	var items []types.DepositItem
	for _, key := range keys {
		item := types.DepositItem{}
		amtStr, err := r.db.Get(ctx, key).Result()
		if err != nil {
			return nil, err
		}
		item.Denom = key
		item.Amount = amtStr
		items = append(items, item)
	}

	return items, nil
}
