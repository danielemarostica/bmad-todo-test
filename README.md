# bmad-todo-test

A full-stack todo application with a Go backend, React frontend, and MongoDB database, fully containerized with Docker Compose.

## Tech Stack

- **Backend:** Go (Gin), MongoDB driver v2
- **Frontend:** React 19, TypeScript, Vite
- **Database:** MongoDB 8.0
- **Proxy:** Nginx (production)
- **CI:** GitHub Actions (Go tests, ESLint, frontend build & tests)

## API Endpoints

| Method | Path | Description |
|--------|------|-------------|
| GET | `/api/health` | Health check |
| GET | `/api/v1/todos` | List all todos |
| POST | `/api/v1/todos` | Create a todo |
| PATCH | `/api/v1/todos/:id` | Update a todo |

## Getting Started

### Prerequisites

- Docker and Docker Compose

### Environment Variables

Copy the example env file:

```sh
cp .env.example .env
```

| Variable | Default | Description |
|----------|---------|-------------|
| `PORT` | `8080` | Backend server port |
| `MONGO_URI` | `mongodb://mongo:27017/todos` | MongoDB connection string |
| `CORS_ORIGIN` | `http://localhost:3000` | Allowed CORS origin |

### Run (Production)

```sh
docker compose up --build
```

- Frontend: http://localhost:3000
- Backend: http://localhost:8080

Nginx serves the frontend and proxies `/api/` requests to the backend.

### Run (Development)

```sh
docker compose -f docker-compose.dev.yml up --build --watch
```

- Frontend (Vite dev server): http://localhost:5173
- Backend (with Air hot-reload): http://localhost:8080

File changes sync automatically via Docker Compose watch.

## Testing

### Backend

```sh
cd backend && go test ./... -v
```

### Frontend

```sh
cd frontend && npm ci && npm test
```

## Project Structure

```
├── backend/
│   ├── config/         # Environment config
│   ├── handlers/       # HTTP handlers + tests
│   ├── middleware/      # CORS setup
│   ├── models/         # MongoDB models & store
│   └── main.go
├── frontend/
│   └── src/            # React + TypeScript app
├── docker-compose.yml          # Production
├── docker-compose.dev.yml      # Development (watch mode)
├── Dockerfile.backend
├── Dockerfile.backend.dev
├── Dockerfile.frontend
├── Dockerfile.frontend.dev
└── nginx.conf
```
