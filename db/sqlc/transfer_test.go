package sqlc

import (
	"context"
	"database/sql"
	"github.com/stretchr/testify/require"
	"testing"
	"time"
)

func TestQueries_CreateTransfer(t *testing.T) {
	createRandomTransfer(t)
}

func createRandomTransfer(t *testing.T) Transfer {
	account1 := createRandomAccount(t)
	account2 := createRandomAccount(t)
	createTransferParams := CreateTransferParams{
		FromAccountID: account1.ID,
		ToAccountID:   account2.ID,
		Amount:        account1.Balance,
	}
	transfer, err := testQueries.CreateTransfer(context.Background(), createTransferParams)
	require.NoError(t, err)
	require.NotEmpty(t, transfer)
	require.Equal(t, account1.ID, transfer.FromAccountID)
	require.Equal(t, account2.ID, transfer.ToAccountID)
	require.Equal(t, account1.Balance, transfer.Amount)

	return transfer
}

func TestQueries_GetTransfer(t *testing.T) {
	transfer := createRandomTransfer(t)
	get, err := testQueries.GetTransfer(context.Background(), transfer.ID)
	require.NoError(t, err)
	require.NotEmpty(t, get)
	require.Equal(t, transfer.ID, get.ID)
	require.Equal(t, transfer.FromAccountID, get.FromAccountID)
	require.Equal(t, transfer.ToAccountID, get.ToAccountID)
	require.Equal(t, transfer.Amount, get.Amount)
	require.WithinDuration(t, transfer.CreatedAt, get.CreatedAt, time.Second)
}

func TestQueries_ListTransfers(t *testing.T) {
	for i := 0; i < 10; i++ {
		createRandomTransfer(t)
	}
	transfers, err := testQueries.ListTransfers(context.Background(), ListTransfersParams{
		Limit:  5,
		Offset: 5,
	})
	require.NoError(t, err)
	require.NotEmpty(t, transfers)
	require.Equal(t, 5, len(transfers))
	for _, transfer := range transfers {
		require.NotEmpty(t, transfer)
	}
}

func TestQueries_DeleteTransfer(t *testing.T) {
	transfer := createRandomTransfer(t)
	err := testQueries.DeleteTransfer(context.Background(), transfer.ID)
	require.NoError(t, err)
	get, err := testQueries.GetTransfer(context.Background(), transfer.ID)
	require.Error(t, err, sql.ErrNoRows.Error())
	require.Empty(t, get)
}
