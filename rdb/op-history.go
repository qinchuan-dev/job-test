package rdb

import (
	"context"
	"encoding/json"
	"errors"
	"job-test/types"
	"log"
	"math/big"
	"time"
)

const (
	PrefixDepositHistory = "depositHistory:"
	PrefixSendHistory    = "sendHistory:"
)

func (r *Rdb) InsertDepositHistory(ctx context.Context, id, denom string, amount big.Int, date time.Time, op types.OpType, memo string) error {
	amtStr, ok := new(big.Int).SetString(amount.String(), 0)
	if !ok {
		return errors.New("invalid amount")
	}

	item := &types.DepositHistory{
		Id:     id,
		Denom:  denom,
		Amount: amtStr.String(),
		OpType: string(rune(op)),
		Date:   date.String(),
		Memo:   memo,
	}
	bz, err := json.Marshal(item)
	if err != nil {
		return err
	}
	_, err = r.db.RPush(ctx, PrefixDepositHistory+id, string(bz)).Result()
	if err != nil {
		return err
	}

	return nil
}

func (r *Rdb) GetDepositHistoryByCustomer(ctx context.Context, id string) ([]types.DepositHistory, error) {
	var historyItems []types.DepositHistory

	r.db.LRange(ctx, PrefixDepositHistory+id, 0, -1).Result()

	listLength, err := r.db.LLen(ctx, PrefixDepositHistory+id).Result()
	if err != nil {
		log.Fatalf("Failed to get list length: %v", err)
	}

	for i := int64(0); i < listLength; i++ {
		bz, err := r.db.LIndex(ctx, PrefixDepositHistory+id, i).Result()
		if err != nil {
			log.Fatalf("Failed to get list element: %v", err)
		}

		var item types.DepositHistory
		err = json.Unmarshal([]byte(bz), &item)
		if err != nil {
			log.Fatalf("Failed to unmarshal JSON to person: %v", err)
		}

		historyItems = append(historyItems, item)
	}

	return historyItems, nil
}

func (r *Rdb) InsertSendHistory(ctx context.Context, sender, receiver, denom string, amount big.Int, date time.Time, memo string) error {
	amtStr, ok := new(big.Int).SetString(amount.String(), 0)
	if !ok {
		return errors.New("invalid amount")
	}

	item := &types.SendHistory{
		Sender:   sender,
		Receiver: receiver,
		Denom:    denom,
		Amount:   amtStr.String(),
		Date:     date.String(),
		Memo:     memo,
	}
	bz, err := json.Marshal(item)
	if err != nil {
		return err
	}

	_, err = r.db.RPush(ctx, PrefixSendHistory+sender, string(bz)).Result()
	if err != nil {
		return err
	}
	return nil
}

func (r *Rdb) GetSendHistoryByCustomer(ctx context.Context, sender string) ([]types.SendHistory, error) {
	var historyItems []types.SendHistory

	r.db.LRange(ctx, PrefixSendHistory+sender, 0, -1).Result()

	listLength, err := r.db.LLen(ctx, PrefixSendHistory+sender).Result()
	if err != nil {
		log.Fatalf("Failed to get list length: %v", err)
	}

	for i := int64(0); i < listLength; i++ {
		bz, err := r.db.LIndex(ctx, PrefixSendHistory+sender, i).Result()
		if err != nil {
			log.Fatalf("Failed to get list element: %v", err)
		}

		var item types.SendHistory
		err = json.Unmarshal([]byte(bz), &item)
		if err != nil {
			log.Fatalf("Failed to unmarshal JSON to person: %v", err)
		}

		historyItems = append(historyItems, item)
	}

	return historyItems, nil
}
