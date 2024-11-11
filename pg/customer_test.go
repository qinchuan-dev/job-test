package pg

import (
	"context"
	"github.com/stretchr/testify/require"
	"testing"
)

const (
	pgConnStr = "postgresql://localhost:55000/postgres"
)

func TestPostgres_DeleteCustomerById(t *testing.T) {
	ctx := context.Background()
	db, err := NewPG(ctx, pgConnStr)
	require.NoError(t, err)

	err = db.InsertCustomer(ctx, "id", "denom")
	require.NoError(t, err)

	_, err = db.GetCustomer(ctx, "id")

}

func TestPostgres_GetCustomer(t *testing.T) {

}

func TestPostgres_InsertCustomer(t *testing.T) {

}
