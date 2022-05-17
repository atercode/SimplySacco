package db

import (
	"context"
	"database/sql"
	"testing"
	"time"

	"github.com/atercode/SimplySacco/utils"
	"github.com/brianvoe/gofakeit/v6"
	"github.com/stretchr/testify/require"
)

func createRandomMember(t *testing.T) Member {
	hashedPassword, err := utils.HashPassword(gofakeit.DigitN(8))
	require.NoError(t, err)

	args := CreateMemberParams{
		FullName:       gofakeit.Name(),
		Email:          gofakeit.Email(),
		StatusCode:     "TEST",
		HashedPassword: hashedPassword,
	}

	member, err := testQueries.CreateMember(context.Background(), args)
	require.NoError(t, err)
	require.NotEmpty(t, member)

	require.Equal(t, args.FullName, member.FullName)
	require.Equal(t, args.Email, member.Email)
	require.Equal(t, args.StatusCode, member.StatusCode)

	require.NotZero(t, member.ID)
	require.NotZero(t, member.CreatedAt)
	return member
}

func TestCreateMember(t *testing.T) {
	createRandomMember(t)
}

func TestGetMember(t *testing.T) {
	created_member := createRandomMember(t)
	retrieved_member, err := testQueries.GetMember(context.Background(), created_member.ID)
	require.NoError(t, err)
	require.NotEmpty(t, retrieved_member)

	require.Equal(t, created_member.ID, retrieved_member.ID)
	require.Equal(t, created_member.FullName, retrieved_member.FullName)
	require.Equal(t, created_member.Email, retrieved_member.Email)
	require.Equal(t, created_member.StatusCode, retrieved_member.StatusCode)
	require.WithinDuration(t, created_member.CreatedAt.Time, retrieved_member.CreatedAt.Time, time.Minute)
}

func TestGetMemberByEmail(t *testing.T) {
	created_member := createRandomMember(t)
	retrieved_member, err := testQueries.GetMemberByEmail(context.Background(), created_member.Email)
	require.NoError(t, err)
	require.NotEmpty(t, retrieved_member)

	require.Equal(t, created_member.ID, retrieved_member.ID)
	require.Equal(t, created_member.FullName, retrieved_member.FullName)
	require.Equal(t, created_member.Email, retrieved_member.Email)
	require.Equal(t, created_member.StatusCode, retrieved_member.StatusCode)
	require.WithinDuration(t, created_member.CreatedAt.Time, retrieved_member.CreatedAt.Time, time.Minute)
}

func TestUpdateMember(t *testing.T) {
	created_member := createRandomMember(t)
	args := UpdateMemberParams{
		ID:         created_member.ID,
		FullName:   gofakeit.Name(),
		Email:      gofakeit.Email(),
		StatusCode: "SWITCHED",
	}

	updated_member, err := testQueries.UpdateMember(context.Background(), args)

	require.NoError(t, err)
	require.NotEmpty(t, updated_member)

	require.Equal(t, created_member.ID, updated_member.ID)
	require.NotEqualValues(t, created_member.FullName, updated_member.FullName)
	require.NotEqualValues(t, created_member.Email, updated_member.Email)
	require.NotEqualValues(t, created_member.StatusCode, updated_member.StatusCode)
}

func TestDeleteMember(t *testing.T) {
	created_member := createRandomMember(t)

	err := testQueries.DeleteMember(context.Background(), created_member.ID)

	require.NoError(t, err)
	//make sure that created account is deleted
	retrieved_member, err := testQueries.GetMember(context.Background(), created_member.ID)
	require.Error(t, err)
	require.EqualError(t, err, sql.ErrNoRows.Error())
	require.Empty(t, retrieved_member)
}

func TestListMembers(t *testing.T) {
	const max_created = 5
	for a := 0; a < max_created; a++ {
		createRandomMember(t)
	}
	members, err := testQueries.ListMembers(context.Background())

	require.NoError(t, err)
	require.NotEmpty(t, members)
	require.GreaterOrEqual(t, len(members), max_created)

	for i := 0; i < max_created; i++ {
		require.NotEmpty(t, members[i])
	}
}

func TestListMembersPaginated(t *testing.T) {
	const max_created = 5
	for a := 0; a < max_created; a++ {
		createRandomMember(t)
	}
	args := ListMembersPaginatedParams{
		Limit:  3,
		Offset: 2,
	}
	members, err := testQueries.ListMembersPaginated(context.Background(), args)

	require.NoError(t, err)
	require.NotEmpty(t, members)
	require.Len(t, members, int(args.Limit))

	for _, member := range members {
		require.NotEmpty(t, member)
	}
}
