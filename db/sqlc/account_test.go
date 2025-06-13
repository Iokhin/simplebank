package sqlc

import (
	"context"
	"database/sql"
	"github.com/stretchr/testify/require"
	"simplebank/util"
	"testing"
	"time"
)

func TestQueries_CreateAccount(t *testing.T) {
	createRandomAccount(t)
}

func createRandomAccount(t *testing.T) Account {
	user := createRandomUser(t)
	accountParams := CreateAccountParams{
		Owner:    user.Username,
		Balance:  util.RandomMoney(),
		Currency: util.RandomCurrency(),
	}
	account, err := testQueries.CreateAccount(context.Background(), accountParams)
	require.NoError(t, err)
	require.NotEmpty(t, account)
	require.Equal(t, accountParams.Owner, account.Owner)
	require.Equal(t, accountParams.Balance, account.Balance)
	require.Equal(t, accountParams.Currency, account.Currency)
	require.NotZero(t, account.ID)
	require.NotZero(t, account.CreatedAt)

	return account
}

func TestQueries_GetAccount(t *testing.T) {
	account := createRandomAccount(t)
	get, err := testQueries.GetAccount(context.Background(), account.ID)

	require.NoError(t, err)
	require.NotEmpty(t, get)

	require.Equal(t, account.ID, get.ID)
	require.Equal(t, account.Owner, get.Owner)
	require.Equal(t, account.Balance, get.Balance)
	require.Equal(t, account.Currency, get.Currency)
	require.WithinDuration(t, account.CreatedAt, get.CreatedAt, time.Second)
}

func TestQueries_UpdateAccountAccount(t *testing.T) {
	account := createRandomAccount(t)
	params := UpdateAccountParams{
		ID:      account.ID,
		Balance: util.RandomMoney(),
	}
	update, err := testQueries.UpdateAccount(context.Background(), params)

	require.NoError(t, err)
	require.NotEmpty(t, update)

	require.Equal(t, account.ID, update.ID)
	require.Equal(t, account.Owner, update.Owner)
	require.Equal(t, params.Balance, update.Balance)
	require.Equal(t, account.Currency, update.Currency)
	require.WithinDuration(t, account.CreatedAt, update.CreatedAt, time.Second)
}

func TestQueries_DeleteAccount(t *testing.T) {
	account := createRandomAccount(t)
	err := testQueries.DeleteAccount(context.Background(), account.ID)
	require.NoError(t, err)

	get, err := testQueries.GetAccount(context.Background(), account.ID)
	require.Error(t, err)
	require.EqualError(t, err, sql.ErrNoRows.Error())
	require.Empty(t, get)
}

func TestQueries_ListAccounts(t *testing.T) {
	for i := 0; i < 10; i++ {
		createRandomAccount(t)
	}
	accounts, err := testQueries.ListAccounts(context.Background(), ListAccountsParams{
		Limit:  5,
		Offset: 5,
	})

	require.NoError(t, err)
	require.Equal(t, 5, len(accounts))

	for _, acc := range accounts {
		require.NotEmpty(t, acc)
	}
}
