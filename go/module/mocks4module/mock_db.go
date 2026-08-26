// Package mocks4module provides test doubles for this module's own tests.
package mocks4module

import (
	"context"
	"testing"

	"github.com/dal-go/dalgo/dal"
	"github.com/sneat-co/sneat-go-core/sneatcoretesting"
)

// NewTestContext returns a context with a fresh in-memory, Firestore-profile
// dal.DB installed as the facade DB (see facade.GetSneatDB), plus the DB
// itself for direct assertions in tests.
//
// It is a thin wrapper over sneat-go-core's sneatcoretesting helper - the
// fleet's single canonical choice of in-memory backend profile for a Sneat
// module - so this module's tests never have to pick a dalgo2memory profile
// on their own (and never call dalgo2memory directly).
func NewTestContext(t *testing.T) (context.Context, dal.DB) {
	t.Helper()
	return sneatcoretesting.SetupMemoryDB(t)
}
