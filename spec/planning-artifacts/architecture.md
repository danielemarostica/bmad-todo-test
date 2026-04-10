---
stepsCompleted: ['step-01-init', 'step-02-context', 'step-03-starter', 'step-04-decisions', 'step-05-patterns', 'step-06-structure', 'step-07-validation', 'step-08-complete']
status: 'complete'
completedAt: '2026-04-10'
inputDocuments: ['spec/planning-artifacts/prd.md']
workflowType: 'architecture'
project_name: 'bmad-todo-test'
user_name: 'Maro'
date: '2026-04-10'
---

# Architecture Decision Document

_This document builds collaboratively through step-by-step discovery. Sections are appended as we work through each architectural decision together._

## Project Context Analysis

### Requirements Overview

**Functional Requirements:**
28 FRs across 6 areas. The core is straightforward CRUD — no complex state machines, no async workflows, no inter-service communication. The highest-complexity FRs are around deployment (Docker Compose orchestration, health checks, volume persistence).

**Non-Functional Requirements:**
16 NFRs set specific, measurable targets. Performance requirements (< 200ms UI, < 100ms API) are easily achievable with any modern stack. Security is basic hygiene (XSS/injection prevention). Accessibility (WCAG AA) requires intentional frontend design but no architectural complexity.

**Scale & Complexity:**

- Primary domain: Full-stack web application
- Complexity level: Low
- Estimated architectural components: 3 (SPA frontend, REST API, database)

### Technical Constraints & Dependencies

- SPA frontend communicates with backend exclusively via JSON REST API
- Frontend and backend must be independently containerizable
- Database must support persistent volume mounts in Docker
- No server-side rendering, no WebSockets, no session management
- Modern evergreen browsers only — no legacy compatibility

### Cross-Cutting Concerns Identified

- **Input validation** — must occur on both client (UX feedback) and server (data integrity)
- **Error handling** — consistent error states in UI, appropriate HTTP error codes from API
- **Health monitoring** — health check endpoints for Docker orchestration
- **Data persistence** — single source of truth in database, no client-side caching complexity

## Starter Template Evaluation

### Primary Technology Domain

Full-stack web application: React SPA (TypeScript) + Go API (Gin) + MongoDB

### Technology Stack

| Layer | Technology | Version | Rationale |
|-------|-----------|---------|-----------|
| Frontend | React + TypeScript | via Vite (react-ts template) | Fast dev server, native TS support, SWC transpilation |
| Backend | Go + Gin | Gin v1.12.x | High-performance REST framework, minimal boilerplate |
| Database | MongoDB | 8.0 (Docker image) | Document store, schema-flexible, simple for CRUD |
| Build/Dev | Vite | Latest | Esbuild-powered, fast HMR |
| Frontend Tests | Vitest + React Testing Library | Latest | Native Vite integration, Jest-compatible API |
| Backend Tests | Go testing + testify | stdlib + testify | Go's built-in testing with assertion helpers |
| E2E Tests | Playwright | Latest | Cross-browser, reliable, MCP integration |
| Containerization | Docker + Docker Compose | Latest | Required by FR24-FR28 |

### Initialization Approach

No single starter template covers this stack. Each layer scaffolded independently:

**Frontend:**
```bash
npm create vite@latest frontend -- --template react-ts
```

**Backend:**
```bash
mkdir backend && cd backend
go mod init bmad-todo-test
go get github.com/gin-gonic/gin
go get go.mongodb.org/mongo-driver/mongo
```

**Project Root Structure:**
```
bmad-todo-test/
├── frontend/          # React SPA (Vite + TypeScript)
├── backend/           # Go API (Gin + MongoDB driver)
├── docker-compose.yml # Orchestration
├── Dockerfile.frontend
├── Dockerfile.backend
└── spec/              # BMAD artifacts
```

### Architectural Decisions Provided by Stack

- **Languages:** TypeScript (frontend), Go (backend) — two-language stack with clear separation
- **Build Tooling:** Vite for frontend (Esbuild + Rollup), Go compiler for backend
- **Testing:** Vitest + RTL (frontend), Go testing + testify (backend), Playwright (E2E)
- **Code Organization:** Monorepo with `frontend/` and `backend/` directories, Docker Compose at root
- **Dev Experience:** Vite HMR for frontend, `air` for Go hot reload

## Core Architectural Decisions

### Decision Priority Analysis

**Critical Decisions (Block Implementation):**
- Data model, API contract, frontend state approach, Docker networking

**Deferred Decisions (Post-MVP):**
- Authentication mechanism, caching strategy, CI/CD pipeline, monitoring/logging infrastructure

### Data Architecture

- **Database:** MongoDB 8.0 via Docker image
- **Collection:** Single `todos` collection
- **Document schema:** `{ _id: ObjectID, text: string, completed: boolean, createdAt: Date }`
- **Validation:** Dual-layer — Gin struct binding/validation on server, form validation in React
- **Migration:** Not needed — schema-flexible document store, no migrations for V1
- **Caching:** None — direct DB reads for every request, sufficient for single-user scale

### Authentication & Security

- **Authentication:** None for V1 (per PRD scope)
- **Input sanitization:** Gin middleware sanitizes all request bodies; React escapes output by default (XSS protection)
- **CORS:** Backend allows configurable frontend origin via `CORS_ORIGIN` environment variable
- **API security:** Request body size limits, content-type validation, proper HTTP status codes

### API & Communication Patterns

- **Style:** RESTful JSON API
- **Base path:** `/api/v1/todos`
- **Endpoints:**
  - `GET /api/v1/todos` — list all todos
  - `POST /api/v1/todos` — create todo `{ text: string }`
  - `PATCH /api/v1/todos/:id` — update todo `{ completed: boolean }`
  - `DELETE /api/v1/todos/:id` — delete todo
  - `GET /api/health` — health check
- **Error format:** `{ "error": "message" }` with appropriate HTTP status codes (400, 404, 500)
- **Success format:** Direct resource representation (single object or array)

### Frontend Architecture

- **State management:** React `useState`/`useEffect` — no external state library
- **HTTP client:** Native `fetch` API — zero dependencies
- **Styling:** Plain CSS — minimal, no build tooling overhead
- **Component structure:** Flat, simple components: `App`, `TodoList`, `TodoItem`, `TodoInput`
- **Routing:** None — single-page, single-view application
- **Error handling:** Try/catch around fetch calls, error state in component

### Infrastructure & Deployment

- **Container architecture:** Three services in Docker Compose: `frontend`, `backend`, `mongo`
- **Frontend container:** Multi-stage build — Vite build → nginx serve static assets
- **Backend container:** Multi-stage build — Go compile → scratch/alpine runtime
- **MongoDB:** Official `mongo:8.0` image with named volume for persistence
- **Networking:** Docker Compose internal network, frontend proxies API via nginx config
- **Environment config:** `.env` file for Docker Compose, environment variables for all configurable values (`MONGO_URI`, `CORS_ORIGIN`, `PORT`)
- **Health checks:** Backend `/api/health` endpoint checks MongoDB connectivity; Docker Compose healthcheck on all services

### Decision Impact Analysis

**Implementation Sequence:**
1. Backend API + MongoDB connection (foundation)
2. Frontend React app with API integration
3. Docker Compose orchestration
4. Test suites (unit, integration, E2E)

**Cross-Component Dependencies:**
- Frontend depends on API contract (endpoint paths, request/response shapes)
- Backend depends on MongoDB connection string (environment variable)
- Docker Compose depends on both Dockerfiles and health check endpoints
- E2E tests depend on full stack running via Docker Compose

## Implementation Patterns & Consistency Rules

### Naming Patterns

**Database (MongoDB):**
- Collection names: lowercase plural (`todos`)
- Field names: camelCase (`createdAt`, `completed`) — MongoDB convention
- IDs: MongoDB `ObjectID`, serialized as `id` (not `_id`) in API responses

**API:**
- Endpoints: lowercase plural nouns (`/api/v1/todos`)
- Route parameters: `:id` (Gin convention)
- JSON fields in request/response: camelCase (`{ "text": "...", "completed": false, "createdAt": "..." }`)
- Dates: ISO 8601 strings (`2026-04-10T12:00:00Z`)

**Go Backend Code:**
- Packages: lowercase single-word (`handlers`, `models`, `config`)
- Structs/types: PascalCase (`Todo`, `CreateTodoRequest`)
- Functions: PascalCase for exported, camelCase for unexported
- Variables: camelCase
- Files: lowercase with underscores (`todo_handler.go`, `health_handler.go`)

**TypeScript Frontend Code:**
- Components: PascalCase files and names (`TodoItem.tsx`, `TodoList.tsx`)
- Non-component files: camelCase (`api.ts`, `types.ts`)
- CSS files: match component name (`TodoItem.css`)
- Functions/variables: camelCase
- Types/interfaces: PascalCase (`Todo`, `ApiError`)

### Structure Patterns

**Backend (`backend/`):**
```
backend/
├── main.go
├── go.mod
├── handlers/
│   ├── todo_handler.go
│   ├── todo_handler_test.go
│   └── health_handler.go
├── models/
│   └── todo.go
├── config/
│   └── config.go
└── middleware/
    └── middleware.go
```

**Frontend (`frontend/src/`):**
```
frontend/src/
├── App.tsx
├── App.css
├── main.tsx
├── components/
│   ├── TodoInput.tsx
│   ├── TodoInput.css
│   ├── TodoList.tsx
│   ├── TodoList.css
│   ├── TodoItem.tsx
│   └── TodoItem.css
├── api/
│   └── todos.ts
├── types/
│   └── todo.ts
└── __tests__/
    ├── TodoInput.test.tsx
    ├── TodoList.test.tsx
    └── TodoItem.test.tsx
```

**E2E Tests:**
```
e2e/
├── playwright.config.ts
└── tests/
    ├── todo-crud.spec.ts
    └── error-states.spec.ts
```

### Format Patterns

**API Response — Success:**
```json
// GET /api/v1/todos → array
[{ "id": "...", "text": "...", "completed": false, "createdAt": "2026-04-10T12:00:00Z" }]

// POST /api/v1/todos → single object
{ "id": "...", "text": "...", "completed": false, "createdAt": "2026-04-10T12:00:00Z" }
```

**API Response — Error:**
```json
{ "error": "text is required" }        // 400
{ "error": "todo not found" }          // 404
{ "error": "internal server error" }   // 500
```

No response wrapper — success returns the resource directly, errors return `{ "error": "..." }`.

### Process Patterns

**Error Handling — Backend:**
- Gin recovery middleware catches panics → 500
- Validation errors → 400 with descriptive message
- Not found → 404
- All errors logged to stdout with timestamp

**Error Handling — Frontend:**
- Try/catch around every fetch call
- Error state stored in component: `{ error: string | null }`
- Error UI displayed inline, not as alerts/modals
- Network errors: generic "Failed to save. Please try again."

**Loading States — Frontend:**
- Boolean `loading` state per component
- Initial load: show loading indicator
- Mutations (create/toggle/delete): simple await + refresh for V1

### Enforcement Guidelines

**All AI Agents MUST:**
- Follow the naming conventions above — no deviations
- Place files in the specified directory structure
- Use the exact API response formats defined
- Co-locate tests with source in Go, use `__tests__/` directory in React
- Use environment variables for all configuration, never hardcode values

## Project Structure & Boundaries

### Complete Project Directory Structure

```
bmad-todo-test/
├── README.md
├── .gitignore
├── .env.example
├── docker-compose.yml
├── Dockerfile.frontend
├── Dockerfile.backend
├── nginx.conf
│
├── backend/
│   ├── main.go
│   ├── go.mod
│   ├── go.sum
│   ├── config/
│   │   └── config.go
│   ├── models/
│   │   └── todo.go
│   ├── handlers/
│   │   ├── todo_handler.go
│   │   ├── todo_handler_test.go
│   │   ├── health_handler.go
│   │   └── health_handler_test.go
│   └── middleware/
│       └── middleware.go
│
├── frontend/
│   ├── package.json
│   ├── tsconfig.json
│   ├── vite.config.ts
│   ├── vitest.config.ts
│   ├── index.html
│   ├── public/
│   └── src/
│       ├── main.tsx
│       ├── App.tsx
│       ├── App.css
│       ├── types/
│       │   └── todo.ts
│       ├── api/
│       │   └── todos.ts
│       ├── components/
│       │   ├── TodoInput.tsx
│       │   ├── TodoInput.css
│       │   ├── TodoList.tsx
│       │   ├── TodoList.css
│       │   ├── TodoItem.tsx
│       │   └── TodoItem.css
│       └── __tests__/
│           ├── TodoInput.test.tsx
│           ├── TodoList.test.tsx
│           └── TodoItem.test.tsx
│
├── e2e/
│   ├── package.json
│   ├── playwright.config.ts
│   └── tests/
│       ├── todo-crud.spec.ts
│       └── error-states.spec.ts
│
└── spec/
    └── planning-artifacts/
        ├── prd.md
        └── architecture.md
```

### Architectural Boundaries

**API Boundary:**
- Frontend → Backend: HTTP only via `/api/v1/*` endpoints
- Frontend never accesses MongoDB directly
- nginx reverse-proxies `/api/*` to backend container

**Component Boundaries (Frontend):**
- `App` owns all state (`todos[]`, `loading`, `error`)
- Props down, callbacks up — standard React prop drilling
- `api/todos.ts` is the only module that makes HTTP calls

**Data Boundary (Backend):**
- `handlers/` receive HTTP requests, validate, call MongoDB, return responses
- `models/` define data structures and validation rules
- MongoDB driver used directly in handlers — no ORM, no repository layer
- All MongoDB operations through the official Go driver

### Requirements to Structure Mapping

| FR Category | Backend Files | Frontend Files |
|-------------|--------------|----------------|
| Task Management (FR1-FR7) | `handlers/todo_handler.go`, `models/todo.go` | `App.tsx`, `components/*`, `api/todos.ts` |
| Input Handling (FR8-FR10) | `models/todo.go` (validation tags) | `TodoInput.tsx` |
| UI States (FR11-FR14) | — | `App.tsx`, `TodoList.tsx`, `TodoItem.tsx` |
| Responsive (FR15-FR17) | — | `*.css` files |
| API (FR18-FR23) | `handlers/todo_handler.go`, `health_handler.go` | `api/todos.ts` |
| Deployment (FR24-FR28) | `Dockerfile.backend` | `Dockerfile.frontend`, `docker-compose.yml`, `nginx.conf` |

### Data Flow

```
User Action → React Component → api/todos.ts → fetch()
  → nginx proxy → Gin handler → MongoDB → response
  → Gin handler → JSON → fetch response → setState → re-render
```

## Architecture Validation Results

### Coherence Validation ✅

**Decision Compatibility:** All technology choices are proven together — React/Vite/TypeScript is a standard frontend stack; Go/Gin/MongoDB is a well-established backend combination. No version conflicts.

**Pattern Consistency:** Naming conventions follow idiomatic standards per language. API contract (camelCase JSON) bridges both sides cleanly. Structure patterns align with Go and React conventions.

**Structure Alignment:** Monorepo with clear `frontend/` and `backend/` separation maps directly to independent Docker containers. nginx proxy bridges the two at deployment.

### Requirements Coverage ✅

All 28 FRs and 16 NFRs have explicit architectural support. Every FR category maps to specific files in the project structure. No orphaned requirements.

### Implementation Readiness ✅

**Decision Completeness:** All critical decisions documented with specific versions. API contract fully specified with endpoints, request/response shapes, and error formats.

**Structure Completeness:** Every file in the project tree has a defined purpose. Requirements-to-structure mapping is explicit.

**Pattern Completeness:** Naming, structure, format, and process patterns cover all identified conflict points between AI agents.

### Architecture Completeness Checklist

- [x] Project context thoroughly analyzed
- [x] Scale and complexity assessed
- [x] Technology stack fully specified with versions
- [x] API contract defined (endpoints, formats, errors)
- [x] Data model defined (MongoDB schema)
- [x] Naming conventions established (Go, TypeScript, API, DB)
- [x] Directory structure complete with file-level detail
- [x] Component boundaries and data flow documented
- [x] Requirements mapped to specific files
- [x] Testing strategy defined per layer
- [x] Docker deployment architecture specified
- [x] Environment configuration approach defined

### Architecture Readiness Assessment

**Overall Status:** READY FOR IMPLEMENTATION

**Confidence Level:** High

**Key Strengths:**
- Radically simple architecture matching radically simple product
- Clear separation of concerns (frontend/backend/database)
- Explicit API contract prevents integration mismatches
- Every requirement traceable to specific architectural components

**First Implementation Priority:** Backend API + MongoDB connection, then frontend, then Docker Compose orchestration.
