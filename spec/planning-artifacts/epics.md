---
stepsCompleted: ['step-01-validate-prerequisites', 'step-02-design-epics', 'step-03-create-stories', 'step-04-final-validation']
status: 'complete'
completedAt: '2026-04-10'
inputDocuments: ['spec/planning-artifacts/prd.md', 'spec/planning-artifacts/architecture.md']
---

# bmad-todo-test - Epic Breakdown

## Overview

This document provides the complete epic and story breakdown for bmad-todo-test, decomposing the requirements from the PRD and Architecture into implementable stories.

## Requirements Inventory

### Functional Requirements

- FR1: User can create a new todo by entering a text description
- FR2: User can view a list of all todos (active and completed)
- FR3: User can mark a todo as completed
- FR4: User can mark a completed todo as active again
- FR5: User can delete a todo
- FR6: System persists all todos across browser sessions
- FR7: System persists all todos across server restarts
- FR8: System prevents creation of empty or whitespace-only todos
- FR9: System handles long text descriptions without breaking layout
- FR10: System provides clear feedback when input validation fails
- FR11: System displays an empty state when no todos exist
- FR12: System displays a loading state while fetching todos
- FR13: System displays an error state when API requests fail
- FR14: Completed todos are visually distinguishable from active todos
- FR15: User can perform all actions on desktop viewports
- FR16: User can perform all actions on mobile viewports (320px minimum)
- FR17: UI adapts layout appropriately across viewport sizes
- FR18: API supports creating a todo via POST request
- FR19: API supports retrieving all todos via GET request
- FR20: API supports updating a todo's completion status via PATCH/PUT request
- FR21: API supports deleting a todo via DELETE request
- FR22: API returns appropriate error responses for invalid requests
- FR23: API exposes a health check endpoint
- FR24: Application runs via Docker Compose as containerized services
- FR25: Frontend and backend run as separate containers
- FR26: Containers report health status via health checks
- FR27: Application supports environment-based configuration
- FR28: Data volume persists across container restarts

### NonFunctional Requirements

- NFR1: UI interactions (create, toggle, delete) complete with visual feedback in < 200ms
- NFR2: Initial page load renders the full todo list in < 1 second
- NFR3: List remains scrollable and responsive with 100+ todos
- NFR4: API responses return in < 100ms under normal conditions
- NFR5: All user input is sanitized to prevent XSS attacks
- NFR6: API validates and sanitizes all incoming request data to prevent injection
- NFR7: API rejects malformed requests with appropriate error codes
- NFR8: No sensitive data (credentials, secrets) exposed in client-side code or logs
- NFR9: Application meets WCAG 2.1 AA compliance
- NFR10: All interactive elements are keyboard navigable
- NFR11: Screen readers can interpret all UI states (empty, loading, error, todo list)
- NFR12: Color contrast ratios meet AA minimum (4.5:1 for text, 3:1 for large text)
- NFR13: Lighthouse accessibility audit scores >= 90
- NFR14: No data loss on server restart or container restart
- NFR15: Failed API requests do not corrupt existing data
- NFR16: Application recovers gracefully from transient network failures

### Additional Requirements

- Starter template: Vite react-ts for frontend, Go mod init for backend
- API contract: `/api/v1/todos` — GET, POST, PATCH /:id, DELETE /:id, GET /api/health
- Data model: `{ _id: ObjectID, text: string, completed: boolean, createdAt: Date }`
- Docker: 3-service compose (frontend/nginx, backend/Go, mongo:8.0)
- nginx reverse proxy for `/api/*` routing to backend
- Environment config: `MONGO_URI`, `CORS_ORIGIN`, `PORT`
- Frontend: React useState/useEffect, plain CSS, native fetch
- Backend: Go Gin, MongoDB driver, dual-layer validation
- Testing: Vitest + RTL (frontend), Go testing + testify (backend), Playwright (E2E)

### UX Design Requirements

No UX design document — requirements derived from PRD user journeys and architecture decisions.

### FR Coverage Map

| FR | Epic | Description |
|----|------|-------------|
| FR1 | Epic 2 | Create todo |
| FR2 | Epic 2 | View all todos |
| FR3 | Epic 2 | Mark completed |
| FR4 | Epic 2 | Mark active again |
| FR5 | Epic 2 | Delete todo |
| FR6 | Epic 2 | Persist across browser sessions |
| FR7 | Epic 1 | Persist across server restarts |
| FR8 | Epic 2 | Prevent empty todos |
| FR9 | Epic 2 | Handle long text |
| FR10 | Epic 2 | Validation feedback |
| FR11 | Epic 2 | Empty state |
| FR12 | Epic 2 | Loading state |
| FR13 | Epic 2 | Error state |
| FR14 | Epic 2 | Visual completion distinction |
| FR15 | Epic 2 | Desktop actions |
| FR16 | Epic 2 | Mobile actions |
| FR17 | Epic 2 | Adaptive layout |
| FR18 | Epic 1 | POST endpoint |
| FR19 | Epic 1 | GET endpoint |
| FR20 | Epic 1 | PATCH endpoint |
| FR21 | Epic 1 | DELETE endpoint |
| FR22 | Epic 1 | Error responses |
| FR23 | Epic 1 | Health check |
| FR24 | Epic 1 | Docker Compose |
| FR25 | Epic 1 | Separate containers |
| FR26 | Epic 1 | Health status |
| FR27 | Epic 1 | Environment config |
| FR28 | Epic 1 | Volume persistence |

## Epic List

### Epic 1: Project Foundation & Core API
Users (developers) can clone the repo, run `docker-compose up`, and have a running backend API connected to MongoDB with health checks — the foundation everything else builds on.
**FRs covered:** FR7, FR18, FR19, FR20, FR21, FR22, FR23, FR24, FR25, FR26, FR27, FR28

### Epic 2: Todo Management UI
Users can create, view, complete, uncomplete, and delete todos through a responsive, accessible interface with instant feedback.
**FRs covered:** FR1, FR2, FR3, FR4, FR5, FR6, FR8, FR9, FR10, FR11, FR12, FR13, FR14, FR15, FR16, FR17

### Epic 3: Quality Assurance & Test Coverage
The application has comprehensive test coverage across all layers — unit, integration, and E2E — ensuring reliability and confidence for future changes.
**FRs covered:** Cross-cutting (validates all FRs)

## Epic 1: Project Foundation & Core API

Users (developers) can clone the repo, run `docker-compose up`, and have a running backend API connected to MongoDB with health checks.

### Story 1.1: Project Scaffolding & Docker Compose Setup

As a developer,
I want to scaffold the project with frontend and backend directories and a working Docker Compose configuration,
So that I have a reproducible development environment from day one.

**Acceptance Criteria:**

**Given** a fresh clone of the repository
**When** I run `docker-compose up`
**Then** three containers start: frontend (nginx), backend (Go), and mongo
**And** all containers reach healthy status
**And** the frontend is accessible at `http://localhost:3000`
**And** the backend is accessible at `http://localhost:8080`

**Given** Docker Compose is running
**When** I check environment configuration
**Then** `MONGO_URI`, `CORS_ORIGIN`, and `PORT` are configurable via `.env`
**And** an `.env.example` file documents all required variables

### Story 1.2: Health Check Endpoint & MongoDB Connection

As a developer,
I want a health check endpoint that verifies MongoDB connectivity,
So that I can confirm the backend and database are operational.

**Acceptance Criteria:**

**Given** the backend is running and MongoDB is connected
**When** I send `GET /api/health`
**Then** I receive a 200 response with status OK

**Given** MongoDB is unreachable
**When** I send `GET /api/health`
**Then** I receive a 503 response indicating database connectivity failure

**Given** Docker Compose healthcheck is configured
**When** the backend container starts
**Then** Docker reports the container as healthy only after `/api/health` returns 200

### Story 1.3: Create Todo API Endpoint

As a user (via API),
I want to create a new todo by sending a POST request,
So that my task is persisted in the database.

**Acceptance Criteria:**

**Given** the API is running
**When** I send `POST /api/v1/todos` with `{ "text": "Buy groceries" }`
**Then** I receive a 201 response with `{ "id": "...", "text": "Buy groceries", "completed": false, "createdAt": "..." }`
**And** the todo is persisted in MongoDB

**Given** I send `POST /api/v1/todos` with `{ "text": "" }`
**When** the API validates the request
**Then** I receive a 400 response with `{ "error": "text is required" }`

**Given** I send `POST /api/v1/todos` with `{ "text": "   " }`
**When** the API validates the request
**Then** I receive a 400 response rejecting whitespace-only text

### Story 1.4: List Todos API Endpoint

As a user (via API),
I want to retrieve all todos,
So that I can see my complete task list.

**Acceptance Criteria:**

**Given** todos exist in the database
**When** I send `GET /api/v1/todos`
**Then** I receive a 200 response with an array of all todos
**And** each todo has `id`, `text`, `completed`, and `createdAt` fields

**Given** no todos exist
**When** I send `GET /api/v1/todos`
**Then** I receive a 200 response with an empty array `[]`

### Story 1.5: Update Todo Completion API Endpoint

As a user (via API),
I want to update a todo's completion status,
So that I can mark tasks as done or reactivate them.

**Acceptance Criteria:**

**Given** a todo exists with `completed: false`
**When** I send `PATCH /api/v1/todos/:id` with `{ "completed": true }`
**Then** I receive a 200 response with the updated todo
**And** the todo's `completed` field is `true` in MongoDB

**Given** I send `PATCH /api/v1/todos/:invalidId`
**When** the todo does not exist
**Then** I receive a 404 response with `{ "error": "todo not found" }`

### Story 1.6: Delete Todo API Endpoint

As a user (via API),
I want to delete a todo,
So that I can remove tasks I no longer need.

**Acceptance Criteria:**

**Given** a todo exists in the database
**When** I send `DELETE /api/v1/todos/:id`
**Then** I receive a 204 response (no content)
**And** the todo is removed from MongoDB

**Given** I send `DELETE /api/v1/todos/:invalidId`
**When** the todo does not exist
**Then** I receive a 404 response with `{ "error": "todo not found" }`

### Story 1.7: Data Persistence Across Restarts

As a user,
I want my todos to survive server and container restarts,
So that I never lose my data.

**Acceptance Criteria:**

**Given** todos exist in the database
**When** I restart the backend container via `docker-compose restart backend`
**Then** all todos are still retrievable via `GET /api/v1/todos`

**Given** todos exist in the database
**When** I run `docker-compose down` and `docker-compose up`
**Then** all todos are still retrievable because MongoDB data volume persists

## Epic 2: Todo Management UI

Users can create, view, complete, uncomplete, and delete todos through a responsive, accessible interface with instant feedback.

### Story 2.1: App Shell & Todo List Display

As a user,
I want to see all my todos when I open the app,
So that I immediately know what I need to do.

**Acceptance Criteria:**

**Given** the app is loaded and todos exist in the database
**When** the page renders
**Then** all todos are displayed in a list with their text and completion status
**And** the page loads in under 1 second

**Given** no todos exist
**When** the page renders
**Then** an empty state message is displayed: "No todos yet. Add one above."

**Given** the API is being called
**When** todos are loading
**Then** a loading indicator is displayed until data arrives

### Story 2.2: Create Todo from UI

As a user,
I want to type a task and press Enter to add it,
So that I can quickly capture tasks without friction.

**Acceptance Criteria:**

**Given** the app is loaded
**When** I type "Buy groceries" in the input field and press Enter
**Then** the todo appears in the list immediately
**And** the input field is cleared
**And** the todo is persisted via the API

**Given** I submit an empty input or whitespace only
**When** I press Enter
**Then** no todo is created
**And** the input shows a validation hint

**Given** I type a very long task description (200+ characters)
**When** I submit it
**Then** the todo is created and the UI handles the long text without breaking layout

### Story 2.3: Toggle Todo Completion

As a user,
I want to click a checkbox to mark a todo as done or undone,
So that I can track my progress.

**Acceptance Criteria:**

**Given** an active todo is displayed
**When** I click its checkbox
**Then** the todo visually changes to completed (strikethrough, muted color)
**And** the completion status is persisted via the API

**Given** a completed todo is displayed
**When** I click its checkbox
**Then** the todo returns to active visual state
**And** the status change is persisted via the API

**Given** the API call fails
**When** I toggle a todo
**Then** an error message is displayed inline
**And** the todo reverts to its previous visual state

### Story 2.4: Delete Todo

As a user,
I want to delete a todo I no longer need,
So that my list stays clean and relevant.

**Acceptance Criteria:**

**Given** a todo is displayed
**When** I click the delete button
**Then** the todo is removed from the list
**And** the deletion is persisted via the API

**Given** the API call fails
**When** I delete a todo
**Then** an error message is displayed
**And** the todo remains in the list

### Story 2.5: Error State Handling

As a user,
I want to see clear error messages when something goes wrong,
So that I understand what happened and can try again.

**Acceptance Criteria:**

**Given** the backend is unreachable
**When** the app tries to load todos
**Then** an error state is displayed instead of the todo list
**And** the message indicates the failure (e.g., "Failed to load todos. Please try again.")

**Given** a create/toggle/delete operation fails
**When** the API returns an error
**Then** an inline error message is shown near the failed action
**And** the error does not break the rest of the UI

### Story 2.6: Responsive Layout & Accessibility

As a user,
I want the app to work on my phone and be usable with keyboard and screen readers,
So that I can manage todos from any device and with any ability.

**Acceptance Criteria:**

**Given** a viewport of 320px width
**When** the app renders
**Then** all elements are visible and functional without horizontal scrolling
**And** touch targets are at least 44x44px

**Given** a desktop viewport
**When** the app renders
**Then** the layout uses available space appropriately

**Given** I navigate with keyboard only
**When** I tab through the interface
**Then** I can reach the input, every checkbox, and every delete button
**And** focus indicators are clearly visible

**Given** a screen reader is active
**When** it reads the page
**Then** all UI states (empty, loading, error, todo list) are announced correctly
**And** interactive elements have appropriate ARIA labels

**Given** I run a Lighthouse accessibility audit
**When** the audit completes
**Then** the score is >= 90

## Epic 3: Quality Assurance & Test Coverage

The application has comprehensive test coverage across all layers ensuring reliability and confidence for future changes.

### Story 3.1: Backend Integration Tests

As a developer,
I want integration tests for all API endpoints,
So that I can verify the backend works correctly against a real MongoDB instance.

**Acceptance Criteria:**

**Given** a test MongoDB instance is available
**When** I run `go test ./...`
**Then** all CRUD endpoints are tested: create, list, update, delete
**And** validation error cases are tested (empty text, invalid ID, not found)
**And** health check endpoint is tested (connected and disconnected states)
**And** all tests pass with 100% endpoint coverage

### Story 3.2: Frontend Component Tests

As a developer,
I want unit tests for all React components,
So that I can verify UI behavior in isolation.

**Acceptance Criteria:**

**Given** the test suite is configured with Vitest and React Testing Library
**When** I run `npm test`
**Then** TodoInput is tested: renders, submits valid text, rejects empty input, clears after submit
**And** TodoList is tested: renders todos, shows empty state, shows loading state, shows error state
**And** TodoItem is tested: renders text, toggles completion, triggers delete
**And** test coverage is >= 70%

### Story 3.3: End-to-End Tests

As a developer,
I want E2E tests covering all user journeys,
So that I can verify the full stack works together from the user's perspective.

**Acceptance Criteria:**

**Given** the full stack is running via Docker Compose
**When** I run the Playwright test suite
**Then** the happy path is tested: create todo, complete todo, uncomplete todo, delete todo
**And** the empty state is verified on first load
**And** input validation is tested (empty submit)
**And** error handling is tested (backend unavailable scenario)
**And** all tests pass across Chromium
