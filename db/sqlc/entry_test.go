package db

import (
	"context"
	"testing"
	"time"

	"github.com/VL-037/go-bank/util"
	"github.com/stretchr/testify/require"
)

func createRandomEntry(t *testing.T, account Account) Entry {
	arg := CreateEntryParams{
		AccountID: account.ID,
		Amount:    util.RandomMoney(),
	}

	entry, err := testQueries.CreateEntry(context.Background(), arg)

	require.NoError(t, err)
	require.NotEmpty(t, entry)

	require.Equal(t, arg.AccountID, entry.AccountID)
	require.Equal(t, arg.Amount, entry.Amount)

	require.NotZero(t, entry.ID)
	require.NotZero(t, entry.CreatedAt)
	require.NotZero(t, entry.UpdatedAt)
	require.False(t, entry.MarkForDelete)

	return entry
}

func TestCreateEntry(t *testing.T) {
	account := createRandomAccount(t)
	createRandomEntry(t, account)
}

func TestGetEntry(t *testing.T) {
	account := createRandomAccount(t)
	savedEntry := createRandomEntry(t, account)
	entry, err := testQueries.GetEntry(context.Background(), savedEntry.ID)
	require.NoError(t, err)
	require.NotEmpty(t, entry)

	require.Equal(t, savedEntry.ID, entry.ID)
	require.Equal(t, savedEntry.AccountID, entry.AccountID)
	require.Equal(t, savedEntry.Amount, entry.Amount)
	require.WithinDuration(t, savedEntry.CreatedAt, entry.CreatedAt, time.Second)
}

func TestListEntries(t *testing.T) {
	account := createRandomAccount(t)
	for i := 0; i < 10; i++ {
		createRandomEntry(t, account)
	}

	arg := ListEntriesParams{
		AccountID: account.ID,
		Limit:     5,
		Offset:    5,
	}

	entries, err := testQueries.ListEntries(context.Background(), arg)
	require.NoError(t, err)
	require.Len(t, entries, 5)

	for _, entry := range entries {
		require.NotEmpty(t, entry)
		require.Equal(t, arg.AccountID, entry.AccountID)
	}
}
