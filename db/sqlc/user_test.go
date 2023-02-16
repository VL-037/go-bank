package db

import (
	"context"
	"github.com/VL-037/go-bank/util"
	"github.com/stretchr/testify/require"
	"testing"
	"time"
)

func createRandomUser(t *testing.T) User {
	hashedPassword, err := util.HashPassword(util.RandomString(6))
	require.NoError(t, err)

	arg := CreateUserParams{
		Username:       util.RandomOwner(),
		HashedPassword: hashedPassword,
		FullName:       util.RandomOwner(),
		Email:          util.RandomEmail(),
	}

	user, err := testQueries.CreateUser(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, user)

	require.Equal(t, arg.Username, user.Username)
	require.Equal(t, arg.HashedPassword, user.HashedPassword)
	require.Equal(t, arg.FullName, user.FullName)
	require.Equal(t, arg.Email, user.Email)

	require.NotZero(t, user.CreatedAt)
	require.NotZero(t, user.UpdatedAt)
	require.False(t, user.MarkForDelete)

	return user
}

func TestCreateUser(t *testing.T) {
	createRandomUser(t)
}

func TestGetUser(t *testing.T) {
	savedUser := createRandomUser(t)
	user, err := testQueries.GetUser(context.Background(), savedUser.Username)

	require.NoError(t, err)
	require.NotEmpty(t, user)

	require.Equal(t, savedUser.Username, user.Username)
	require.Equal(t, savedUser.HashedPassword, user.HashedPassword)
	require.Equal(t, savedUser.FullName, user.FullName)
	require.Equal(t, savedUser.Email, user.Email)
	require.WithinDuration(t, savedUser.PasswordUpdatedAt, user.PasswordUpdatedAt, time.Second)
	require.WithinDuration(t, savedUser.CreatedAt, user.CreatedAt, time.Second)
	require.WithinDuration(t, savedUser.UpdatedAt, user.UpdatedAt, time.Second)
	require.Equal(t, savedUser.MarkForDelete, user.MarkForDelete)
}
