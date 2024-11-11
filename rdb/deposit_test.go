package rdb

import (
	"context"
	"github.com/stretchr/testify/require"
	"math/big"
	"testing"
)

func TestRdb_SetDeposit(t *testing.T) {
	rdbInstance, err := NewRdb(connString)
	require.NoError(t, err)

	ctx := context.Background()

	err = rdbInstance.db.FlushDB(ctx).Err()
	require.NoError(t, err)

	err = rdbInstance.SetDeposit(ctx, "id", "denom", *big.NewInt(1000))
	require.NoError(t, err)

	item, err := rdbInstance.GetDepositByCustomerDenom(ctx, "id", "denom")
	require.NoError(t, err)
	require.Equal(t, "1000", item.Amount)
}

func TestRdb_DeleteDeposit(t *testing.T) {
	rdbInstance, err := NewRdb(connString)
	require.NoError(t, err)

	ctx := context.Background()

	err = rdbInstance.db.FlushDB(ctx).Err()
	require.NoError(t, err)

	err = rdbInstance.SetDeposit(ctx, "id", "denom", *big.NewInt(1000))
	require.NoError(t, err)

	err = rdbInstance.DeleteDeposit(ctx, "id", "denom")
	require.NoError(t, err)

	_, err = rdbInstance.GetDepositByCustomerDenom(ctx, "id", "denom")
	require.Error(t, err)
}

func TestRdb_Deposit(t *testing.T) {
	rdbInstance, err := NewRdb(connString)
	require.NoError(t, err)

	ctx := context.Background()

	err = rdbInstance.db.FlushDB(ctx).Err()
	require.NoError(t, err)

	err = rdbInstance.SetDeposit(ctx, "id", "denom", *big.NewInt(1000))
	require.NoError(t, err)

	item1, err := rdbInstance.GetDepositByCustomerDenom(ctx, "id", "denom")
	require.NoError(t, err)
	require.Equal(t, "1000", item1.Amount)

	err = rdbInstance.Deposit(ctx, "id", "denom", *big.NewInt(1000))
	require.NoError(t, err)

	item2, err := rdbInstance.GetDepositByCustomerDenom(ctx, "id", "denom")
	require.NoError(t, err)
	require.Equal(t, "2000", item2.Amount)
}

func TestRdb_Withdraw(t *testing.T) {
	rdbInstance, err := NewRdb(connString)
	require.NoError(t, err)

	ctx := context.Background()

	err = rdbInstance.db.FlushDB(ctx).Err()
	require.NoError(t, err)

	err = rdbInstance.SetDeposit(ctx, "id", "denom", *big.NewInt(1000))
	require.NoError(t, err)

	item1, err := rdbInstance.GetDepositByCustomerDenom(ctx, "id", "denom")
	require.NoError(t, err)
	require.Equal(t, "1000", item1.Amount)

	err = rdbInstance.Withdraw(ctx, "id", "denom", *big.NewInt(500))
	require.NoError(t, err)

	item2, err := rdbInstance.GetDepositByCustomerDenom(ctx, "id", "denom")
	require.NoError(t, err)
	require.Equal(t, "500", item2.Amount)
}

func TestRdb_GetDepositByCustomer(t *testing.T) {
	rdbInstance, err := NewRdb(connString)
	require.NoError(t, err)

	ctx := context.Background()

	err = rdbInstance.db.FlushDB(ctx).Err()
	require.NoError(t, err)

	err = rdbInstance.SetDeposit(ctx, "id", "usd", *big.NewInt(1000))
	require.NoError(t, err)
	err = rdbInstance.SetDeposit(ctx, "id", "rmb", *big.NewInt(2000))
	require.NoError(t, err)

	items, err := rdbInstance.GetDepositByCustomer(ctx, "id")
	require.NoError(t, err)
	require.Equal(t, len(items), 2)
}
