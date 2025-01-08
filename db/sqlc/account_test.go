package db

import (
	"bank/util"
	"context"
	"testing"

	"github.com/stretchr/testify/require"
)

func createMultipleAccountsForOneOwner(t *testing.T) Account {
	arg := CreateAccountParams{
		Owner:    "Owner",
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

	return account
}

func createRandomAccount(t *testing.T) Account {
	arg := CreateAccountParams{
		Owner:    util.RandomOwner(),
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

	return account
}

func TestCreateAccount(t *testing.T) {
	createRandomAccount(t)
}

func TestGetAccount(t *testing.T) {
	createAccount := createRandomAccount(t)
	account1, err := testQueries.GetAccount(context.Background(), createAccount.ID)
	require.NoError(t, err)
	require.NotEmpty(t, account1)

	require.Equal(t, createAccount.ID, account1.ID)
	require.Equal(t, createAccount.Owner, account1.Owner)
	require.Equal(t, createAccount.Balance, account1.Balance)
	require.Equal(t, createAccount.Currency, account1.Currency)
	require.WithinDuration(t, createAccount.CreatedAt, account1.CreatedAt, 0)

}

func TestUpdateAccount(t *testing.T) {
	createAccount := createRandomAccount(t)
	arg := UpdateAccountParams{
		ID:      createAccount.ID,
		Balance: util.RandomMoney(),
	}
	account, err := testQueries.UpdateAccount(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, account)

	require.Equal(t, createAccount.ID, account.ID)
	require.Equal(t, createAccount.Owner, account.Owner)
	require.Equal(t, arg.Balance, account.Balance)
	require.Equal(t, createAccount.Currency, account.Currency)
	require.WithinDuration(t, createAccount.CreatedAt, account.CreatedAt, 0)
}

func TestDeleteAccount(t *testing.T) {
	createAccount := createRandomAccount(t)
	err := testQueries.DeleteAccount(context.Background(), createAccount.ID)
	require.NoError(t, err)

	account, err := testQueries.GetAccount(context.Background(), createAccount.ID)
	require.Error(t, err)
	require.Empty(t, account)
}

func TestListAccounts(t *testing.T) {
	var lastAccount Account
	for i := 0; i < 5; i++ {
		lastAccount = createMultipleAccountsForOneOwner(t)
	}

	arg := ListAccountsParams{
		Owner:  lastAccount.Owner,
		Limit:  5,
		Offset: 0,
	}

	accounts, err := testQueries.ListAccounts(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, accounts)
	require.Len(t, accounts, 5)

	for _, account := range accounts {
		require.NotEmpty(t, account)
		require.Equal(t, lastAccount.Owner, account.Owner)
	}
}
