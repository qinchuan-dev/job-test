package rdb

import (
	"context"
	"github.com/stretchr/testify/require"
	"job-test/types"
	"testing"
	"time"
)

func TestRdb_GetDepositHistoryByCustomer(t *testing.T) {
	rdbInstance, err := NewRdb(connString)
	require.NoError(t, err)

	ctx := context.Background()

	err = rdbInstance.db.FlushDB(ctx).Err()
	require.NoError(t, err)

	date := time.Now()
	in := types.DepositHistory{
		Id:     "id",
		Denom:  "denom",
		Amount: "1000",
		OpType: "1",
		Date:   date.String(),
		Memo:   "memo",
	}

	err = rdbInstance.InsertDepositHistory(ctx, in)
	require.NoError(t, err)
	err = rdbInstance.InsertDepositHistory(ctx, in)
	require.NoError(t, err)

	out, err := rdbInstance.GetDepositHistoryByCustomer(ctx, "id")
	require.NoError(t, err)
	require.Equal(t, len(out), 2)
}

func TestRdb_GetSendHistoryByCustomer(t *testing.T) {
	rdbInstance, err := NewRdb(connString)
	require.NoError(t, err)

	ctx := context.Background()

	err = rdbInstance.db.FlushDB(ctx).Err()
	require.NoError(t, err)

	date := time.Now()
	in := types.SendHistory{
		Sender:   "sender",
		Receiver: "receiver",
		Denom:    "denom",
		Amount:   "1000",
		Date:     date.String(),
		Memo:     "memo",
	}

	err = rdbInstance.InsertSendHistory(ctx, in)
	require.NoError(t, err)
	err = rdbInstance.InsertSendHistory(ctx, in)
	require.NoError(t, err)

	out, err := rdbInstance.GetSendHistoryByCustomer(ctx, "sender")
	require.NoError(t, err)
	require.Equal(t, len(out), 2)
}

func TestRdb_InsertDepositHistory(t *testing.T) {

}

func TestRdb_InsertSendHistory(t *testing.T) {

}
