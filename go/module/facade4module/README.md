# facade4module

Business logic that reads and writes this module's data through
[`dal-go/dalgo`](https://github.com/dal-go/dalgo).

`facade.go` shows the current idiom end to end: `SaveGreeting` runs a single
read-write transaction that reads the existing record first and only then
writes (collect-then-apply - required under Firestore-profile semantics,
which reject a read after a write in the same transaction); `GetGreeting`
does a plain read outside a transaction. See `facade_test.go` for a test
built on the `mocks4module` in-memory DB.
