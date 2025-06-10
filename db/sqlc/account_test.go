package sqlc

import (
	"context"
	"database/sql"
	"github.com/stretchr/testify/assert"
	"simplebank/util"
	"testing"
	"time"
)

func TestQueries_CreateAccount(t *testing.T) {
	createRandomAccount(t)
}

func createRandomAccount(t *testing.T) Account {
	accountParams := CreateAccountParams{
		Owner:    util.RandomOwner(),
		Balance:  util.RandomMoney(),
		Currency: util.RandomCurrency(),
	}
	account, err := testQueries.CreateAccount(context.Background(), accountParams)
	assert.NoError(t, err)
	assert.NotEmpty(t, account)
	assert.Equal(t, accountParams.Owner, account.Owner)
	assert.Equal(t, accountParams.Balance, account.Balance)
	assert.Equal(t, accountParams.Currency, account.Currency)
	assert.NotZero(t, account.ID)
	assert.NotZero(t, account.CreatedAt)

	return account
}

func TestQueries_GetAccount(t *testing.T) {
	account := createRandomAccount(t)
	get, err := testQueries.GetAccount(context.Background(), account.ID)

	assert.NoError(t, err)
	assert.NotEmpty(t, get)

	assert.Equal(t, account.ID, get.ID)
	assert.Equal(t, account.Owner, get.Owner)
	assert.Equal(t, account.Balance, get.Balance)
	assert.Equal(t, account.Currency, get.Currency)
	assert.WithinDuration(t, account.CreatedAt, get.CreatedAt, time.Second)
}

func TestQueries_UpdateAccountAccount(t *testing.T) {
	account := createRandomAccount(t)
	params := UpdateAccountParams{
		ID:      account.ID,
		Balance: util.RandomMoney(),
	}
	update, err := testQueries.UpdateAccount(context.Background(), params)

	assert.NoError(t, err)
	assert.NotEmpty(t, update)

	assert.Equal(t, account.ID, update.ID)
	assert.Equal(t, account.Owner, update.Owner)
	assert.Equal(t, params.Balance, update.Balance)
	assert.Equal(t, account.Currency, update.Currency)
	assert.WithinDuration(t, account.CreatedAt, update.CreatedAt, time.Second)
}

func TestQueries_DeleteAccount(t *testing.T) {
	account := createRandomAccount(t)
	err := testQueries.DeleteAccount(context.Background(), account.ID)
	assert.NoError(t, err)

	get, err := testQueries.GetAccount(context.Background(), account.ID)
	assert.Error(t, err)
	assert.EqualError(t, err, sql.ErrNoRows.Error())
	assert.Empty(t, get)
}

func TestQueries_ListAccounts(t *testing.T) {
	for i := 0; i < 10; i++ {
		createRandomAccount(t)
	}
	accounts, err := testQueries.ListAccounts(context.Background(), ListAccountsParams{
		Limit:  5,
		Offset: 5,
	})

	assert.NoError(t, err)
	assert.Equal(t, 5, len(accounts))

	for _, acc := range accounts {
		assert.NotEmpty(t, acc)
	}
}
