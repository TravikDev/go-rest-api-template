# Go REST API Template

This project provides a simple REST API for managing users using Go and PostgreSQL. It follows a minimal layered architecture and can be run in Docker.

## Running with Docker Compose

```bash
docker-compose up --build
```

The application will be available on `http://localhost:8080` and connects to a PostgreSQL database exposed on port `5432`.

## Endpoints

- `POST /users` – create a user
- `GET /users` – list users
- `GET /users/show?id=1` – get user by ID

