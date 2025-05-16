# Chirpy 🐦 — A Lightweight Twitter-Like API Server in Go

Chirpy is a simple, Twitter-inspired API server written in Go. It allows users to register, log in, post short messages ("chirps"), and manage sessions. It includes support for JWT authentication, user upgrades via a `polka` webhook, and basic admin metrics.

## 📦 Features

- User registration, login, and profile updates
- JWT-based authentication & session refresh/revoke
- Chirp creation, listing, deletion, and filtering by author
- Admin metrics and server health checks
- Webhook handler to upgrade user roles
- Uses PostgreSQL for persistent storage

---

## 🚀 Getting Started

### Prerequisites

- Go 1.20+
- PostgreSQL
- Make sure you have a `.env` file or set environment variables as below

---

### 📁 Environment Variables

Create a `.env` file in the root directory:

```env
DB_URL=postgres://<username>:<password>@localhost/<dbname>?sslmode=disable
PLATFORM=production
JWT_SECRET=your-secret-key
POLKA_KEY=your-polka-key
```
