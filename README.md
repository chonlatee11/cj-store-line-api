# CJ Store LINE API

Backend API service for webhook → LINE Official Account integration.

This project handles webhook events from external systems and communicates with **LINE Official Account** using [`line-bot-sdk-go`](https://github.com/line/line-bot-sdk-go). It sends different message types (text, image, template, flex, etc.) to LINE users.

## Architecture

This project follows Clean Architecture principles:

- **Domain**: Pure models + interfaces + errors (no framework imports)
- **Use case/Service**: Business logic orchestrations, calling domain interfaces
- **Adapters**: HTTP/Gin handlers, DB repositories, external services (LINE OA, etc.)
- **Infrastructure**: DB clients (GORM + PostgreSQL), LINE bot SDK integration, config, DI

## Tech Stack

- **Golang**: Gin framework, GORM ORM
- **Database**: PostgreSQL
- **External Integration**: LINE Official Account API
- **Infrastructure**: Docker/Compose, Kubernetes, GitLab CI/CD

## Getting Started

### Prerequisites

- Go 1.21+
- PostgreSQL
- LINE Official Account (Channel Secret & Access Token)

### Installation

1. Clone the repository:
```bash
git clone https://github.com/chonlatee11/cj-store-line-api.git
cd cj-store-line-api
```

2. Copy environment file:
```bash
cp .env.example .env
```

3. Update `.env` with your configuration

4. Install dependencies:
```bash
go mod download
```

5. Run migrations:
```bash
make migrate-up
```

6. Start the server:
```bash
make run
```

The API will be available at `http://localhost:8080`

## API Endpoints

### Health Checks
- `GET /healthz` - Liveness probe
- `GET /readyz` - Readiness probe

### Webhook
- `POST /webhook/line` - Receive LINE webhook events

## Development

### Available Make Commands

```bash
make run          # Run the application
make test         # Run tests with race detection and coverage
make lint         # Run golangci-lint
make migrate-up   # Run database migrations
make migrate-down # Rollback database migrations
make seed         # Seed database with test data
```

### Testing

```bash
# Run all tests
make test

# Run tests with verbose output
go test -v -race -cover ./...
```

## Deployment

This project uses GitLab CI/CD with the following stages:
- `build`: Build the application
- `test`: Run tests and linting
- `migrate`: Run database migrations
- `deploy`: Deploy to target environment

## Contributing

1. Follow Conventional Commits for commit messages
2. Use feature branches: `feat/`, `fix/`, `chore/`
3. Ensure all tests pass before submitting PR
4. Add tests for new functionality

## License

[Add your license here]