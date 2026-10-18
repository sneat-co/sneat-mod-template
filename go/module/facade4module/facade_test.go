package facade4module

import (
	"context"
	"errors"
	"testing"

	"github.com/dal-go/dalgo/dal"
	"github.com/dal-go/record"
	"github.com/sneat-co/sneat-go-core/facade"
	"github.com/sneat-co/sneat-mod-module/go/module/mocks4module"
	"github.com/stretchr/testify/require"
)

func TestSaveGreeting_insertsThenUpdates(t *testing.T) {
	ctx, _ := mocks4module.NewTestContext(t)

	require.NoError(t, SaveGreeting(ctx, "user1", "hello"))

	greeting, err := GetGreeting(ctx, "user1")
	require.NoError(t, err)
	require.Equal(t, "hello", greeting)

	// Second call exercises the update path: SaveGreeting must read the
	// existing record before it writes, proving the reads-before-writes
	// ordering required inside a single transaction actually works.
	require.NoError(t, SaveGreeting(ctx, "user1", "hi again"))

	greeting, err = GetGreeting(ctx, "user1")
	require.NoError(t, err)
	require.Equal(t, "hi again", greeting)
}

func TestGetGreeting_notFound(t *testing.T) {
	ctx, _ := mocks4module.NewTestContext(t)

	_, err := GetGreeting(ctx, "missing")
	require.Error(t, err)
	require.True(t, record.IsNotFound(err))
}

func TestGetGreeting_dbError(t *testing.T) {
	ctx := facade.WithSneatDBProvider(context.Background(), func(ctx context.Context) (dal.DB, error) {
		return nil, errors.New("simulated db provider error")
	})
	_, err := GetGreeting(ctx, "user1")
	require.Error(t, err)
	require.Contains(t, err.Error(), "failed to get facade DB")
}

type mockGetErrorTx struct {
	dal.ReadwriteTransaction
}

func (m mockGetErrorTx) Get(ctx context.Context, record record.Record) error {
	return errors.New("simulated get error")
}

type mockGetErrorDB struct {
	dal.DB
}

func (m mockGetErrorDB) RunReadwriteTransaction(ctx context.Context, f func(ctx context.Context, tx dal.ReadwriteTransaction) error, opts ...dal.TransactionOption) error {
	return m.DB.RunReadwriteTransaction(ctx, func(ctx context.Context, tx dal.ReadwriteTransaction) error {
		return f(ctx, mockGetErrorTx{ReadwriteTransaction: tx})
	}, opts...)
}

func TestSaveGreeting_getError(t *testing.T) {
	_, memDB := mocks4module.NewTestContext(t)
	db := mockGetErrorDB{DB: memDB}
	ctx := facade.WithSneatDB(context.Background(), db)

	err := SaveGreeting(ctx, "user1", "hello")
	require.Error(t, err)
	require.Contains(t, err.Error(), "failed to get existing")
}
