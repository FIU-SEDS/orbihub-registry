# orbihub-registry

REST API and registry service for OrbiHub written in Go — manages app listings, metadata, and versioning for the rocketry software marketplace

### File Structure

```bash
orbihub-registry/
├── cmd/
│   └── server/
│       └── main.go          ← entrypoint, starts the server
├── internal/
│   ├── handler/
│   │   └── apps.go          ← HTTP handlers (GET /apps, GET /apps/:id, etc.)
│   ├── store/
│   │   └── apps.go          ← database queries (all SQL lives here)
│   └── model/
│       └── app.go           ← App struct definition
├── db/
│   └── seed.sql             ← your seed file goes here
├── .env                     ← DATABASE_URL (gitignored)
├── .gitignore
├── go.mod
└── go.sum
```

- `cmd/server` — entrypoint only, kept thin. Real projects often have multiple binaries (a CLI tool, a migration runner, etc.) so each gets its own folder under cmd/
- `internal/` — Go's way of saying "this code is private to this module, not importable by others"
  - `handler` — knows about HTTP, calls store
  - `store` — knows about the database, knows nothing about HTTP
  - `model` — plain structs, no logic, imported by both handler and store
- `db/` — keeps your SQL files organized, separate from Go code

### How to add new app

A `POST` request has already been setup so you can update the registry for a new application

```bash
curl -X POST http://localhost:8000/apps \
  -H "Content-Type: application/json" \
  -d '{"id":"test-app","name":"Test App","description":"A test application","version":"1.0.0","repo":"https://github.com/test/test-app","author":"Test","image":"test.png"}'
```

