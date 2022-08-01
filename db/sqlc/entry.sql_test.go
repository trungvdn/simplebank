package db

import (
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
	"github.com/trungvdn/simplebank/util"
)

func createNewEntry(t *testing.T, account Account) Entry {
	args := CreateEntryParams{
		AccountID: account.ID,
		Amount:    util.RandomMoney(),
	}
	entry, err := testQueries.CreateEntry(context.Background(), args)
	require.NoError(t, err)
	require.NotEmpty(t, entry)

	require.Equal(t, args.AccountID, entry.AccountID)
	require.Equal(t, args.Amount, entry.Amount)

	require.NotZero(t, entry.ID)
	require.NotZero(t, entry.CreatedAt)
	return entry
}

func TestCreateEntry(t *testing.T) {
	account := createNewAccount(t)
	createNewEntry(t, account)
}

func TestGetEntry(t *testing.T) {
	account := createNewAccount(t)
	entry := createNewEntry(t, account)
	gotEntry, err := testQueries.GetEntry(context.Background(), entry.ID)
	require.NoError(t, err)
	require.NotEmpty(t, gotEntry)
	require.Equal(t, entry.AccountID, gotEntry.AccountID)
	require.Equal(t, entry.Amount, gotEntry.Amount)
	require.Equal(t, entry.ID, gotEntry.ID)
	require.WithinDuration(t, entry.CreatedAt, gotEntry.CreatedAt, time.Second)
}

func TestListEntries(t *testing.T) {
	account := createNewAccount(t)
	for i := 0; i < 10; i++ {
		createNewEntry(t, account)
	}
	args := ListEntriesParams{
		AccountID: account.ID,
		Limit:     5,
		Offset:    5,
	}
	entries, err := testQueries.ListEntries(context.Background(), args)
	require.NoError(t, err)
	require.Len(t, entries, 5)

	for i := 0; i < 5; i++ {
		require.NotEmpty(t, entries[i])
		require.Equal(t, args.AccountID, entries[i].AccountID)
	}
}
