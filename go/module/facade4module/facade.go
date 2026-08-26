// Package facade4module holds this module's business logic - the code that
// reads and writes the module's own data through dal-go/dalgo.
package facade4module

import (
	"context"
	"fmt"

	"github.com/dal-go/dalgo/dal"
	"github.com/dal-go/record"
	"github.com/sneat-co/sneat-go-core/facade"
	"github.com/sneat-co/sneat-mod-module/go/module/models4module"
)

// thingCollection is the example collection name used by SaveGreeting and
// GetGreeting below.
const thingCollection = "things"

// SaveGreeting inserts or updates a models4module.Thing record identified by
// id with the given greeting text.
//
// The transaction callback follows the current dal-go/dalgo best practice of
// collect-then-apply: it reads the existing record first and performs its one
// write only after that read has completed. A Firestore-profile backend
// (production and sneatcoretesting's in-memory test DB alike) rejects a read
// that happens after a write in the same transaction with
// dal.ErrReadAfterWriteInTransaction, so this ordering is not optional.
func SaveGreeting(ctx context.Context, id, greeting string) error {
	key := record.NewKeyWithID(thingCollection, id)
	return facade.RunReadwriteTransaction(ctx, func(ctx context.Context, tx dal.ReadwriteTransaction) error {
		thing := new(models4module.Thing)
		existing, err := dal.GetRecordWithIDIntoData(ctx, tx, key, id, thing)
		if err != nil && !record.IsNotFound(err) {
			return fmt.Errorf("failed to get existing %s record: %w", thingCollection, err)
		}

		// All reads for this transaction are done - only writes from here on.
		thing.Greeting = greeting

		if existing.Record.Exists() {
			return tx.Set(ctx, existing.Record)
		}
		_, err = dal.InsertRecordWithDataAndID(ctx, tx, key, id, thing)
		return err
	})
}

// GetGreeting returns the greeting stored for id.
//
// If no record exists yet it returns "" and a non-nil error for which
// record.IsNotFound(err) is true.
func GetGreeting(ctx context.Context, id string) (string, error) {
	db, err := facade.GetSneatDB(ctx)
	if err != nil {
		return "", fmt.Errorf("failed to get facade DB: %w", err)
	}
	thing := new(models4module.Thing)
	key := record.NewKeyWithID(thingCollection, id)
	if _, err = dal.GetRecordWithIDIntoData(ctx, db, key, id, thing); err != nil {
		return "", err
	}
	return thing.Greeting, nil
}
