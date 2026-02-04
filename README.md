# Snippetbox

A web application for creating and sharing text snippets with expiration dates, built with Go and PostgreSQL.

## Features

- Create and share text snippets with configurable expiration (1 day, 7 days, or 1 year)
- User authentication system (signup, login, logout)
- Session management with PostgreSQL-backed sessions
- CSRF protection on all state-changing operations
- Secure password hashing with bcrypt
- Server-side form validation
- Flash messages for user feedback
- Clean, responsive UI

## Tech Stack

- **Go 1.25**: Backend language
- **PostgreSQL**: Database with pgx/v5 driver
- **Chi**: Lightweight HTTP router
- **sqlx**: Extensions for database/sql
- **SCS**: Session management with PostgreSQL store
- **Alice**: Middleware chaining
- **bcrypt**: Password hashing
- **HTML templates**: Server-side rendering
- **nosurf**: CSRF protection

## Project Structure

```
.
├── cmd/web/              # Application entry point and HTTP handlers
│   ├── main.go          # Main application setup
│   ├── handlers.go      # HTTP request handlers
│   ├── routes.go        # Route definitions
│   ├── middleware.go    # Custom middleware
│   ├── helpers.go       # Helper functions
│   └── templates.go     # Template rendering
├── internal/
│   ├── models/          # Data models and database logic
│   │   ├── snippets.go  # Snippet CRUD operations
│   │   ├── users.go     # User authentication
│   │   └── errors.go    # Custom error types
│   └── validator/       # Form validation logic
└── ui/
    ├── html/            # HTML templates
    │   ├── base.tmpl
    │   ├── pages/
    │   └── partials/
    └── static/          # Static assets (CSS, JS, images)
```

## Prerequisites

- Go 1.25 or higher
- PostgreSQL database
- Environment variables configured in `.env` file

## Setup

1. Clone the repository:
```bash
git clone https://github.com/thedeepak12/snippetbox.git
cd snippetbox
```

2. Install dependencies:
```bash
go mod download
```

3. Create a `.env` file with the following variables:
```
PORT=4000
DSN=postgres://username:password@localhost/snippetbox?sslmode=disable
```

4. Set up the PostgreSQL database with the required tables:
   - `snippets` table: id, title, content, created, expires
   - `users` table: id, name, email, hashed_password, created
   - `sessions` table: token, data, expiry (for SCS session storage)

5. Run the application:
```bash
go run ./cmd/web
```

The server will start on the port specified in your `.env` file (default: 4000).

## Usage

1. Visit `http://localhost:4000` to view the home page with the latest snippets
2. Sign up for an account to start creating snippets
3. Log in and create your first snippet with a title, content, and expiration period
4. Share snippet URLs with others
5. Snippets automatically expire and become inaccessible after their expiration date

## Acknowledgments

- **Alex Edwards** for the comprehensive [Let's Go](https://lets-go.alexedwards.net/) book
