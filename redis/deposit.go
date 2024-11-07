package rdb

import (
	"context"
	"job-test/types"
	"math/big"
)

func (r *Rdb) InsertDeposit(ctx context.Context, id string, denom string, amount big.Int) error {

	return nil
}

func (r *Rdb) DeleteDeposit(ctx context.Context, id string, denom string) error {

	return nil
}

func (r *Rdb) SetDepositAmount(ctx context.Context, id string, denom string, amount big.Int) error {

	return nil
}

func (r *Rdb) Deposit(ctx context.Context, id string, denom string, deltaAmt big.Int) error {
	return nil
}

func (r *Rdb) Withdraw(ctx context.Context, id string, denom string, deltaAmt big.Int) error {
	return nil
}

func (r *Rdb) GetDepositByCustomer(ctx context.Context, id string) ([]types.DepositItem, error) {

	return nil, nil
}
