# orbihub-registry

> REST API and registry service for OrbiHub written in Go — manages app listings, metadata, and versioning for the rocketry software marketplace.

## Overview

`orbihub-registry` is the backend registry service powering [OrbiHub](https://github.com/FIU-SEDS/OrbiHub), a desktop marketplace for rocketry software. It exposes a RESTful API backed by a PostgreSQL database (Supabase) and is designed to replace hardcoded app listings with a live, queryable registry.

## Endpoints

| Method   | Endpoint     | Auth     | Description                     |
| -------- | ------------ | -------- | ------------------------------- |
| `GET`    | `/apps`      | None     | List all apps in the registry   |
| `GET`    | `/apps/{id}` | None     | Get a specific app by ID        |
| `POST`   | `/apps`      | Required | Publish a new app               |
| `PUT`    | `/apps/{id}` | Required | Update an existing app          |
| `DELETE` | `/apps/{id}` | Required | Remove an app from the registry |

Write endpoints (`POST`, `PUT`, `DELETE`) require an `Authorization: Bearer <API_KEY>` header.

## Stack

- **Language** — Go 1.25
- **Database** — PostgreSQL via Supabase
- **Driver** — `pgx/v5`
- **Deployment** — Render

## Project Structure

```
orbihub-registry/
├── cmd/
│   └── server/
│       └── main.go          ← entrypoint, starts the server
├── internal/
│   ├── handler/
│   │   └── apps.go          ← HTTP handlers
│   ├── middleware/
│   │   └── auth.go          ← API key auth middleware
│   ├── store/
│   │   └── apps.go          ← database queries
│   └── model/
│       └── app.go           ← App struct definition
├── db/
│   └── seed.sql             ← initial registry data
├── .env                     ← environment variables (gitignored)
├── go.mod
└── go.sum
```

- **`cmd/server`** — entrypoint only, kept thin. Each binary gets its own folder under `cmd/`
- **`internal/`** — private to this module, not importable by external packages
  - **`handler`** — knows about HTTP, calls store, knows nothing about the database
  - **`middleware`** — auth logic applied to write endpoints
  - **`store`** — knows about the database, knows nothing about HTTP
  - **`model`** — plain structs, no logic, imported by both handler and store
- **`db/`** — SQL files kept separate from Go code

## Getting Started

### Prerequisites

- Go 1.22+
- A [Supabase](https://supabase.com) project with the `apps` table (see `db/seed.sql`)

### Setup

1. Clone the repository

```bash
git clone https://github.com/erielC/orbihub-registry
cd orbihub-registry
```

2. Install dependencies

```bash
go mod download
```

3. Create a `.env` file in the project root

```env
DATABASE_URL=user=... password=... host=... port=6543 dbname=postgres sslmode=require statement_cache_mode=describe
API_KEY=your-secret-key
```

4. Run the server

```bash
go run cmd/server/main.go
```

The server starts at `http://localhost:8000`.

## Usage

### List all apps

```bash
curl http://localhost:8000/apps
```

### Get app by ID

```bash
curl http://localhost:8000/apps/telemetry-viewer
```

### Publish a new app

```bash
curl -X POST http://localhost:8000/apps \
  -H "Content-Type: application/json" \
  -H "Authorization: Bearer your-api-key" \
  -d '{
    "id": "my-app",
    "name": "My App",
    "description": "A rocketry tool",
    "version": "1.0.0",
    "repo": "https://github.com/username/my-app",
    "author": "Your Name",
    "image": "my_app_logo.png"
  }'
```

### Update an app

```bash
curl -X PUT http://localhost:8000/apps/my-app \
  -H "Content-Type: application/json" \
  -H "Authorization: Bearer your-api-key" \
  -d '{"version": "1.1.0"}'
```

### Delete an app

```bash
curl -X DELETE http://localhost:8000/apps/my-app \
  -H "Authorization: Bearer your-api-key"
```

## Database

Run `db/seed.sql` in the Supabase SQL editor to create the `apps` table and populate it with initial data.

## License

MIT
