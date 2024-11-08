package rdb

import (
	"context"
	"errors"
	"github.com/redis/go-redis/v9"
	"job-test/types"
	"log"
	"math/big"
)

const (
	PrefixDeposit = "deposit_"
)

func (r *Rdb) SetDeposit(ctx context.Context, id string, denom string, amt big.Int) error {
	amtStr, ok := new(big.Int).SetString(amt.String(), 0)
	if !ok {
		return errors.New("invalid old amount")
	}

	r.db.HSet(ctx, PrefixDeposit+id, denom, amtStr)

	return nil
}

func (r *Rdb) DeleteDeposit(ctx context.Context, id string, denom string) error {
	_, err := r.db.HDel(ctx, PrefixDeposit+id, denom).Result()
	if err != nil {
		return err
	}
	return nil
}

func (r *Rdb) Deposit(ctx context.Context, id string, denom string, deltaAmt big.Int) error {
	oldAmtStr, err := r.db.HGet(ctx, PrefixDeposit+id, denom).Result()
	if err != nil {
		if errors.Is(err, redis.Nil) {
			_, err := r.db.HSet(ctx, PrefixDeposit+id, denom, deltaAmt.String()).Result()
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

	r.db.HSet(ctx, PrefixDeposit+id, denom, newAmt.String())

	return nil
}

func (r *Rdb) Withdraw(ctx context.Context, id string, denom string, deltaAmt big.Int) error {
	oldAmtStr, err := r.db.HGet(ctx, PrefixDeposit+id, denom).Result()
	if err != nil {
		return err
	}
	oldAmt, ok := new(big.Int).SetString(oldAmtStr, 0)
	if !ok {
		return errors.New("invalid old amount")
	}

	newAmt := oldAmt.Sub(oldAmt, &deltaAmt)

	r.db.HSet(ctx, PrefixDeposit+id, denom, newAmt.String())

	return nil
}

func (r *Rdb) GetDepositByCustomer(ctx context.Context, id string) ([]types.DepositItem, error) {
	hashFields, err := r.db.HGetAll(ctx, PrefixDeposit+id).Result()
	if err != nil {
		log.Fatalf("Failed to get hash fields and values: %v", err)
	}

	var items []types.DepositItem
	for denom, amtStr := range hashFields {
		item := types.DepositItem{}
		item.Denom = denom
		item.Amount = amtStr

		items = append(items, item)
	}
	return items, nil
}
