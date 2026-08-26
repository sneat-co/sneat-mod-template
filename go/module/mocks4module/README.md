# mocks4module

Test doubles for this module's own tests.

`mock_db.go` builds an in-memory `dal.DB` via
`sneat-go-core/sneatcoretesting.SetupMemoryDB` - the fleet's canonical
Firestore-profile in-memory test DB for a Sneat module - and returns it
installed in a context, ready for `facade4module` code to consume.
