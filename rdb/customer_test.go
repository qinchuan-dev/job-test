package rdb

import (
	"context"
	"github.com/stretchr/testify/require"
	"testing"
)

const (
	connString = "redis://localhost:55002/0"
)

func TestRdb_DeleteCustomerById(t *testing.T) {
	rdbInstance, err := NewRdb(connString)
	require.NoError(t, err)

	ctx := context.Background()

	err = rdbInstance.db.FlushDB(ctx).Err()
	require.NoError(t, err)

	err = rdbInstance.InsertCustomer(ctx, "id", "name")
	require.NoError(t, err)

	err = rdbInstance.DeleteCustomerById(ctx, "id")
	require.NoError(t, err)

	_, err = rdbInstance.GetCustomer(ctx, "id")
	require.Error(t, err)
}

func TestRdb_GetCustomer(t *testing.T) {
	rdbInstance, err := NewRdb(connString)
	require.NoError(t, err)

	ctx := context.Background()

	err = rdbInstance.db.FlushDB(ctx).Err()
	require.NoError(t, err)

	err = rdbInstance.InsertCustomer(ctx, "id", "name")
	require.NoError(t, err)

	name, err := rdbInstance.GetCustomer(ctx, "id")
	require.NoError(t, err)
	require.Equal(t, "name", name)
}
