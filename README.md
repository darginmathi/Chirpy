# Chirpy — A Twitter-Like API Server in Go

Chirpy is a simple, Twitter-inspired API server written in Go. It allows users to register, log in, post short messages ("chirps"), and manage sessions. It includes support for JWT authentication and basic admin metrics.

## Features

- User registration, login, and profile updates
- JWT-based authentication & session refresh/revoke
- Chirp creation, listing, deletion, and filtering by author
- Admin metrics and server health checks
- Webhook handler to upgrade user roles
- PostgreSQL for persistent storage

---

## Getting Started

### Prerequisites

- Go 1.20+
- PostgreSQL
- Make sure you have a `.env` file or set environment variables as below

---

### Environment Variables

Create a `.env` file in the root directory:

```env
DB_URL="postgres://<username>:<password>@localhost/<dbname>?sslmode=disable"
PLATFORM=production
JWT_SECRET=your-secret-key
POLKA_KEY=your-polka-key
```

## Build & Run

```bash
go build -o chirpy
./chirpy
```

## API Endpoints
### Health
`GET /api/healthz` – Check server readiness

### User Management
`POST /api/users` – Register a user

`PUT /api/users` – Update user email and password

`POST /api/login` – Log in (returns access + refresh token)

`POST /api/refresh` – Refresh access token

`POST /api/revoke` – Revoke refresh token

### Chirps
`POST /api/chirps` – Create a new chirp (JWT required)

`GET /api/chirps` – Get all chirps
Optional query parameters:
`author_id=uuid`
`sort=asc|desc`

`GET /api/chirps/{chirpID}` – Get a single chirp

`DELETE /api/chirps/{chirpID}` – Delete a chirp (JWT required)

### Admin
`GET /admin/metrics` – View file server hit count

`POST /admin/reset` – Reset metrics (if platform is dev)

### Webhooks
`POST /api/polka/webhooks` – Handle user upgrade


