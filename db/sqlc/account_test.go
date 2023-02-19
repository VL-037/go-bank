package db

import (
	"context"
	"database/sql"
	"testing"
	"time"

	"github.com/VL-037/go-bank/util"
	"github.com/stretchr/testify/require"
)

func createRandomAccount(t *testing.T) Account {
	user := createRandomUser(t)

	arg := CreateAccountParams{
		Owner:    user.Username,
		Balance:  util.RandomMoney(),
		Currency: util.RandomCurrency(),
	}

	account, err := testQueries.CreateAccount(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, account)

	require.Equal(t, arg.Owner, account.Owner)
	require.Equal(t, arg.Balance, account.Balance)
	require.Equal(t, arg.Currency, account.Currency)

	require.NotZero(t, account.ID)
	require.NotZero(t, account.CreatedAt)
	require.NotZero(t, account.UpdatedAt)
	require.False(t, account.MarkForDelete)

	return account
}

func TestCreateAccount(t *testing.T) {
	createRandomAccount(t)
}

func TestGetAccount(t *testing.T) {
	savedAccount := createRandomAccount(t)
	account, err := testQueries.GetAccount(context.Background(), savedAccount.ID)

	require.NoError(t, err)
	require.NotEmpty(t, account)

	require.Equal(t, savedAccount.ID, account.ID)
	require.Equal(t, savedAccount.Owner, account.Owner)
	require.Equal(t, savedAccount.Balance, account.Balance)
	require.Equal(t, savedAccount.Currency, account.Currency)
	require.WithinDuration(t, savedAccount.CreatedAt, account.CreatedAt, time.Second)
	require.WithinDuration(t, savedAccount.UpdatedAt, account.UpdatedAt, time.Second)
	require.Equal(t, savedAccount.MarkForDelete, account.MarkForDelete)
}

func TestUpdateAccount(t *testing.T) {
	savedAccount := createRandomAccount(t)

	arg := UpdateAccountParams{
		ID:      savedAccount.ID,
		Balance: util.RandomMoney(),
	}

	account, err := testQueries.UpdateAccount(context.Background(), arg)

	require.NoError(t, err)
	require.NotEmpty(t, account)

	require.Equal(t, savedAccount.ID, account.ID)
	require.Equal(t, savedAccount.Owner, account.Owner)
	require.Equal(t, arg.Balance, account.Balance)
	require.Equal(t, savedAccount.Currency, account.Currency)
	require.WithinDuration(t, savedAccount.CreatedAt, account.CreatedAt, time.Second)
	require.Equal(t, savedAccount.MarkForDelete, account.MarkForDelete)
}

func TestDeleteAccount(t *testing.T) {
	savedAccount := createRandomAccount(t)

	err := testQueries.DeleteAccount(context.Background(), savedAccount.ID)

	require.NoError(t, err)

	account, err := testQueries.GetAccount(context.Background(), savedAccount.ID)
	require.Error(t, err)
	require.Equal(t, sql.ErrNoRows, err)
	require.Empty(t, account)
}

func TestListAccounts(t *testing.T) {
	var lastAccount Account
	for i := 0; i < 10; i++ {
		lastAccount = createRandomAccount(t)
	}

	arg := ListAccountsParams{
		Owner: lastAccount.Owner,
		Limit:  5,
		Offset: 0,
	}

	accounts, err := testQueries.ListAccounts(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, accounts)

	for _, account := range accounts {
		require.NotEmpty(t, account)
		require.Equal(t, lastAccount.Owner, account.Owner)
	}
}
