package rdb

import (
	"context"
	"job-test/types"
	"math/big"
)

func (r *Rdb) InsertDepositHistory(ctx context.Context, id, denom string, amount big.Int, opType types.OpType, memo string) error {

	return nil
}

func (r *Rdb) GetDepositHistoryByCustomer(ctx context.Context, id string) ([]types.DepositHistoryItem, error) {
	var historyItems []types.DepositHistoryItem

	return historyItems, nil
}

func (r *Rdb) InsertSendHistory(ctx context.Context, sender, receiver, denom string, amount big.Int, memo string) error {

	return nil
}

func (r *Rdb) GetSendHistoryByCustomer(ctx context.Context, snStart uint64, customerId string) ([]types.SendHistoryItem, error) {
	var historyItems []types.SendHistoryItem

	return historyItems, nil
}
