package db

import (
	"context"
	"database/sql"
	"testing"
	"time"

	"github.com/anshoryihsan/golearn/util"
	"github.com/stretchr/testify/require"
)

func createRandomAccount(t *testing.T) Account {
	arg := CreateAccountParams{
		Owner:    util.RandomOwner(),    //create random func
		Balance:  util.RandomMoney(),    //create random func
		Currency: util.RandomCurrency(), //create random func
	}

	account, err := testingQueries.CreateAccount(context.Background(), arg)
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
	accont1 := createRandomAccount(t)
	accont2, err := testingQueries.GetAccounts(context.Background(), accont1.ID)
	require.NoError(t, err)
	require.NotEmpty(t, accont2)

	require.Equal(t, accont1.ID, accont2.ID)
	require.Equal(t, accont1.Owner, accont2.Owner)
	require.Equal(t, accont1.Balance, accont2.Balance)
	require.Equal(t, accont1.Currency, accont2.Currency)
	require.WithinDuration(t, accont1.CreatedAt, accont2.CreatedAt, time.Second)
}

func TestUpdateAccount(t *testing.T) {
	accont1 := createRandomAccount(t)

	arg := UpdateAccountsParams{
		ID:      accont1.ID,
		Balance: util.RandomMoney(),
	}
	accont2, err := testingQueries.UpdateAccounts(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, accont2)

	require.Equal(t, accont1.ID, accont2.ID)
	require.Equal(t, accont1.Owner, accont2.Owner)
	require.Equal(t, arg.Balance, accont2.Balance)
	require.Equal(t, accont1.Currency, accont2.Currency)
	require.WithinDuration(t, accont1.CreatedAt, accont2.CreatedAt, time.Second)
}

func TestDeleteAccount(t *testing.T) {
	account1 := createRandomAccount(t)
	err := testingQueries.DeleteAccounts(context.Background(), account1.ID)
	require.NoError(t, err)

	account2, err := testingQueries.GetAccounts(context.Background(), account1.ID)
	require.Error(t, err)
	require.EqualError(t, err, sql.ErrNoRows.Error())
	require.Empty(t, account2)
}

func TestListAccounts(t *testing.T) {
	for i := 0; i < 10; i++ {
		createRandomAccount(t)
	}

	arg := ListAccountsParams{
		Limit:  5,
		Offset: 5,
	}

	accounts, err := testingQueries.ListAccounts(context.Background(), arg)
	require.NoError(t, err)
	require.Len(t, accounts, 5)

	for _, accaccount := range accounts {
		require.NotEmpty(t, accaccount)
	}
}
