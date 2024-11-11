package rdb

import (
	"context"
	"encoding/json"
	"job-test/types"
)

const (
	PrefixDepositHistory = "depositHistory:"
	PrefixSendHistory    = "sendHistory:"
)

func (r *Rdb) InsertDepositHistory(ctx context.Context, item types.DepositHistory) error {
	bz, err := json.Marshal(item)
	if err != nil {
		return err
	}
	_, err = r.db.RPush(ctx, PrefixDepositHistory+item.Id, string(bz)).Result()
	if err != nil {
		return err
	}

	return nil
}

func (r *Rdb) GetDepositHistoryByCustomer(ctx context.Context, id string) ([]types.DepositHistory, error) {
	var historyItems []types.DepositHistory

	_, err := r.db.LRange(ctx, PrefixDepositHistory+id, 0, -1).Result()
	if err != nil {
		return nil, err
	}

	listLength, err := r.db.LLen(ctx, PrefixDepositHistory+id).Result()
	if err != nil {
		return nil, err
	}

	for i := int64(0); i < listLength; i++ {
		bz, err := r.db.LIndex(ctx, PrefixDepositHistory+id, i).Result()
		if err != nil {
			return nil, err
		}

		var item types.DepositHistory
		err = json.Unmarshal([]byte(bz), &item)
		if err != nil {
			return nil, err
		}

		historyItems = append(historyItems, item)
	}

	return historyItems, nil
}

func (r *Rdb) InsertSendHistory(ctx context.Context, item types.SendHistory) error {
	bz, err := json.Marshal(item)
	if err != nil {
		return err
	}

	_, err = r.db.RPush(ctx, PrefixSendHistory+item.Sender, string(bz)).Result()
	if err != nil {
		return err
	}
	return nil
}

func (r *Rdb) GetSendHistoryByCustomer(ctx context.Context, sender string) ([]types.SendHistory, error) {
	var historyItems []types.SendHistory

	_, err := r.db.LRange(ctx, PrefixSendHistory+sender, 0, -1).Result()
	if err != nil {
		return nil, err
	}

	listLength, err := r.db.LLen(ctx, PrefixSendHistory+sender).Result()
	if err != nil {
		return nil, err
	}

	for i := int64(0); i < listLength; i++ {
		bz, err := r.db.LIndex(ctx, PrefixSendHistory+sender, i).Result()
		if err != nil {
			return nil, err
		}

		var item types.SendHistory
		err = json.Unmarshal([]byte(bz), &item)
		if err != nil {
			return nil, err
		}

		historyItems = append(historyItems, item)
	}

	return historyItems, nil
}
