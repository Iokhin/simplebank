package sqlc

import (
	"context"
	"database/sql"
	"github.com/stretchr/testify/assert"
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
	assert.NoError(t, err)
	assert.NotEmpty(t, entry)
	assert.Equal(t, account.ID, entry.AccountID)
	assert.Equal(t, account.Balance, entry.Amount)

	return entry
}

func TestQueries_GetEntry(t *testing.T) {
	entry := createRandomEntry(t)
	get, err := testQueries.GetEntry(context.Background(), entry.ID)
	assert.NoError(t, err)
	assert.NotEmpty(t, get)
	assert.Equal(t, entry.ID, get.ID)
	assert.Equal(t, entry.AccountID, get.AccountID)
	assert.Equal(t, entry.Amount, get.Amount)
	assert.Equal(t, entry.CreatedAt, get.CreatedAt)
}

func TestQueries_UpdateEntry(t *testing.T) {
	entry := createRandomEntry(t)
	params := UpdateEntryParams{
		ID:     entry.ID,
		Amount: util.RandomMoney(),
	}
	update, err := testQueries.UpdateEntry(context.Background(), params)
	assert.NoError(t, err)
	assert.NotEmpty(t, update)
	assert.Equal(t, params.ID, update.ID)
	assert.Equal(t, params.Amount, update.Amount)
}

func TestQueries_DeleteEntry(t *testing.T) {
	entry := createRandomEntry(t)
	err := testQueries.DeleteEntry(context.Background(), entry.ID)
	assert.NoError(t, err)
	get, err := testQueries.GetEntry(context.Background(), entry.ID)
	assert.Error(t, err, sql.ErrNoRows.Error())
	assert.Empty(t, get)
}

func TestQueries_ListEntries(t *testing.T) {
	for i := 0; i < 10; i++ {
		createRandomEntry(t)
	}

	entries, err := testQueries.ListEntries(context.Background(), ListEntriesParams{
		Limit:  5,
		Offset: 5,
	})

	assert.NoError(t, err)
	assert.NotEmpty(t, entries)
	assert.Equal(t, 5, len(entries))

	for _, entry := range entries {
		assert.NotEmpty(t, entry)
	}
}
