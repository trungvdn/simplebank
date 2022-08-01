package db

import (
	"context"
	"testing"
	"time"

	_ "github.com/lib/pq"
	"github.com/stretchr/testify/require"
	"github.com/trungvdn/simplebank/util"
)

func createNewAccount(t *testing.T) Account {
	args := CreateAccountParams{
		Owner:    util.RandomOwner(),
		Balance:  util.RandomMoney(),
		Currency: util.RandomCurrency(),
	}
	account, err := testQueries.CreateAccount(context.Background(), args)
	require.NoError(t, err)
	require.NotEmpty(t, account)

	require.Equal(t, args.Balance, account.Balance)
	require.Equal(t, args.Currency, account.Currency)
	require.Equal(t, args.Owner, account.Owner)

	require.NotZero(t, account.ID)
	require.NotZero(t, account.CreatedAt)
	return account
}

func TestCreateAccount(t *testing.T) {
	createNewAccount(t)
}

func TestGetAccount(t *testing.T) {
	account := createNewAccount(t)

	gotAccount, err := testQueries.GetAccount(context.Background(), account.ID)
	require.NoError(t, err)
	require.NotEmpty(t, account)

	require.Equal(t, account.Balance, gotAccount.Balance)
	require.Equal(t, account.ID, gotAccount.ID)
	require.Equal(t, account.Owner, gotAccount.Owner)
	require.Equal(t, account.Currency, gotAccount.Currency)
	require.WithinDuration(t, account.CreatedAt, gotAccount.CreatedAt, time.Second)
}

func TestUpdateAccount(t *testing.T) {
	account := createNewAccount(t)
	args := UpdateAccountParams{
		ID:      account.ID,
		Balance: util.RandomMoney(),
	}
	err := testQueries.UpdateAccount(context.Background(), args)
	require.NoError(t, err)
}

func TestDeleteAccount(t *testing.T) {
	account := createNewAccount(t)
	err := testQueries.DeleteAccount(context.Background(), account.ID)
	require.NoError(t, err)
}

func TestListAccounts(t *testing.T) {
	for i := 0; i < 10; i++ {
		createNewAccount(t)
	}
	args := ListAccountsParams{
		Limit:  5,
		Offset: 5,
	}
	accounts, err := testQueries.ListAccounts(context.Background(), args)
	require.NoError(t, err)
	require.Len(t, accounts, 5)
	for i := 0; i < len(accounts); i++ {
		require.NotEmpty(t, accounts[i])
	}
}
