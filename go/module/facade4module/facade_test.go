package facade4module

import (
	"testing"

	"github.com/dal-go/record"
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
