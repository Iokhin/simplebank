package sqlc

import (
	"context"
	"github.com/stretchr/testify/require"
	"simplebank/util"
	"testing"
	"time"
)

func TestQueries_CreateUser(t *testing.T) {
	createRandomUser(t)
}

func createRandomUser(t *testing.T) User {
	createUserParams := CreateUserParams{
		Username:       util.RandomOwner(),
		HashedPassword: "secret",
		FullName:       util.RandomOwner(),
		Email:          util.RandomEmail(),
	}
	user, err := testQueries.CreateUser(context.Background(), createUserParams)
	require.NoError(t, err)
	require.NotEmpty(t, user)
	require.Equal(t, createUserParams.Username, user.Username)
	require.Equal(t, createUserParams.HashedPassword, user.HashedPassword)
	require.Equal(t, createUserParams.FullName, user.FullName)
	require.Equal(t, createUserParams.Email, user.Email)
	require.True(t, user.PasswordChangedAt.IsZero())
	require.NotZero(t, user.CreatedAt)

	return user
}

func TestQueries_GetUser(t *testing.T) {
	user := createRandomUser(t)
	get, err := testQueries.GetUser(context.Background(), user.Username)

	require.NoError(t, err)
	require.NotEmpty(t, get)

	require.Equal(t, user.Username, get.Username)
	require.Equal(t, user.HashedPassword, get.HashedPassword)
	require.Equal(t, user.FullName, get.FullName)
	require.Equal(t, user.Email, get.Email)
	require.WithinDuration(t, user.PasswordChangedAt, get.PasswordChangedAt, time.Second)
	require.WithinDuration(t, user.CreatedAt, get.CreatedAt, time.Second)
}
