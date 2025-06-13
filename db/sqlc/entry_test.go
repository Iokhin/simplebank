package sqlc

import (
	"context"
	"database/sql"
	"github.com/stretchr/testify/require"
	"simplebank/util"
	"testing"
)

func TestQueries_CreateEntry(t *testing.T) {
	createRandomEntry(t)
}

func createRandomEntry(t *testing.T) Entry {
	account := createRandomAccount(t)
	params := CreateEntryParams{
		AccountID: account.ID,
		Amount:    account.Balance,
	}
	entry, err := testQueries.CreateEntry(context.Background(), params)
	require.NoError(t, err)
	require.NotEmpty(t, entry)
	require.Equal(t, account.ID, entry.AccountID)
	require.Equal(t, account.Balance, entry.Amount)

	return entry
}

func TestQueries_GetEntry(t *testing.T) {
	entry := createRandomEntry(t)
	get, err := testQueries.GetEntry(context.Background(), entry.ID)
	require.NoError(t, err)
	require.NotEmpty(t, get)
	require.Equal(t, entry.ID, get.ID)
	require.Equal(t, entry.AccountID, get.AccountID)
	require.Equal(t, entry.Amount, get.Amount)
	require.Equal(t, entry.CreatedAt, get.CreatedAt)
}

func TestQueries_UpdateEntry(t *testing.T) {
	entry := createRandomEntry(t)
	params := UpdateEntryParams{
		ID:     entry.ID,
		Amount: util.RandomMoney(),
	}
	update, err := testQueries.UpdateEntry(context.Background(), params)
	require.NoError(t, err)
	require.NotEmpty(t, update)
	require.Equal(t, params.ID, update.ID)
	require.Equal(t, params.Amount, update.Amount)
}

func TestQueries_DeleteEntry(t *testing.T) {
	entry := createRandomEntry(t)
	err := testQueries.DeleteEntry(context.Background(), entry.ID)
	require.NoError(t, err)
	get, err := testQueries.GetEntry(context.Background(), entry.ID)
	require.Error(t, err, sql.ErrNoRows.Error())
	require.Empty(t, get)
}

func TestQueries_ListEntries(t *testing.T) {
	for i := 0; i < 10; i++ {
		createRandomEntry(t)
	}

	entries, err := testQueries.ListEntries(context.Background(), ListEntriesParams{
		Limit:  5,
		Offset: 5,
	})

	require.NoError(t, err)
	require.NotEmpty(t, entries)
	require.Equal(t, 5, len(entries))

	for _, entry := range entries {
		require.NotEmpty(t, entry)
	}
}
