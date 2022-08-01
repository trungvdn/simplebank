package db

import (
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
	"github.com/trungvdn/simplebank/util"
)

func createNewTransfer(t *testing.T, fromAccountID, toAccountID int64) Transfer {
	args := CreateTransferParams{
		FromAccountID: fromAccountID,
		ToAccountID:   toAccountID,
		Amount:        util.RandomMoney(),
	}
	transfer, err := testQueries.CreateTransfer(context.Background(), args)
	require.NoError(t, err)
	require.NotEmpty(t, transfer)

	return transfer
}

func TestCreateTransfer(t *testing.T) {
	fromAccount := createNewAccount(t)
	toAccount := createNewAccount(t)
	createNewTransfer(t, fromAccount.ID, toAccount.ID)
}

func TestGetTransfer(t *testing.T) {
	fromAccount := createNewAccount(t)
	toAccount := createNewAccount(t)
	transfer := createNewTransfer(t, fromAccount.ID, toAccount.ID)
	gotTransfer, err := testQueries.GetTransfer(context.Background(), transfer.ID)
	require.NoError(t, err)
	require.NotEmpty(t, gotTransfer)

	require.Equal(t, transfer.FromAccountID, gotTransfer.FromAccountID)
	require.Equal(t, transfer.ToAccountID, gotTransfer.ToAccountID)
	require.Equal(t, transfer.Amount, gotTransfer.Amount)
	require.WithinDuration(t, transfer.CreatedAt, gotTransfer.CreatedAt, time.Second)
}

func TestListTransfers(t *testing.T) {
	fromAccount := createNewAccount(t)
	toAccount := createNewAccount(t)
	for i := 0; i < 10; i++ {
		createNewTransfer(t, fromAccount.ID, toAccount.ID)
	}
	args := ListTransfersParams{
		FromAccountID: fromAccount.ID,
		ToAccountID:   toAccount.ID,
		Limit:         5,
		Offset:        5,
	}
	gotTransfers, err := testQueries.ListTransfers(context.Background(), args)
	require.NoError(t, err)
	require.Len(t, gotTransfers, 5)

	for _, transfer := range gotTransfers {
		require.NotEmpty(t, transfer)
		require.Equal(t, fromAccount.ID, transfer.FromAccountID)
		require.Equal(t, toAccount.ID, transfer.ToAccountID)
	}
}
