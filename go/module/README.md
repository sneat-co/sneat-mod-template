# Module: {MODULE_ID}

This module ...

## Packages

- [api4module](api4module) - HTTP endpoints
- [const4module](const4module) - Constants
- [dto4module](dto4module) - Data transfer objects
- [facade4module](facade4module) - Facade (e.g. business logic); shows the
  current `dal-go/dalgo` read-write-transaction idiom (reads before writes)
- [mocks4module](mocks4module) - Mocks for tests; an in-memory `dal.DB` via
  `sneat-go-core/sneatcoretesting`
- [models4module](models4module) - Models (e.g. database models); a minimal
  example record used by facade4module
