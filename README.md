# Go REST API Template

This project provides a simple REST API for managing users using Go and PostgreSQL. It follows a minimal layered architecture and can be run in Docker.

## Running with Docker Compose

```bash
docker-compose up --build
```

The application will be available on `http://localhost:8080` and connects to a PostgreSQL database exposed on port `5432`.
When the database container starts for the first time it executes `db/init.sql` to
create the `users` table automatically.

## Endpoints

- `POST /register` – register a user
- `GET /users` – list users
- `GET /users/show?id=1` – get user by ID
- `POST /login` – obtain a JWT token using a username

All `/users` endpoints require a valid `Authorization: Bearer <token>` header.

