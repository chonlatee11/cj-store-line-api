# Copilot Instructions for This Repository

> Use these rules whenever generating code, tests, docs, or scripts for this project. If something is unclear, prefer reading existing code and config files over guessing.

---

## 1) Project Context (Read First)

* **Domain**: Backend API service for webhook → LINE OA integration.
* **Description**: This project handles webhook events from external systems and communicates with **LINE Official Account** using [`line-bot-sdk-go`](https://github.com/line/line-bot-sdk-go). It sends different message types (text, image, template, flex, etc.) to LINE users.
* **Primary stack**:

  * **Golang**: Gin framework, GORM ORM, PostgreSQL as the main datastore.
  * **Infra**: Docker/Compose, Kubernetes, GitLab CI/CD.
* **Architectural style**: Clean Code / Clean Architecture (domain → use case/service → adapters → infra).
* **Datastores**: PostgreSQL (primary), Redis/MongoDB optional for caching or session if required.

### Always infer from the repo

* Read `README.md`, `Makefile`, `docker-compose.yml`, `/scripts`, `/deploy`, `/k8s`, `/migrations`, and `/docs` before generating.
* Mirror folder conventions and naming already present in the repo.
* If a type/schema/DTO exists, **reuse** it—do not make new ones with overlapping meaning.

---

## 2) Code Style & Conventions

### Go (Gin + GORM)

* **Version**: Match `go.mod`.
* **Style**: `golangci-lint` compliant. Keep functions small, clean, and readable.
* **Errors**: Wrap with context via `fmt.Errorf("…: %w", err)`. Return sentinel errors from domain packages.
* **HTTP**: DTOs in `/internal/transport/http`. Always use `context.Context` with timeouts for requests.
* **DB**: Use transactions for multi-step writes. Define repository interfaces in domain layer, implement with GORM in infra layer.
* **GORM**:

  * Always use context-aware methods: `db.WithContext(ctx)`.
  * Avoid raw SQL unless necessary.
  * Handle `ErrRecordNotFound` explicitly.
* **Testing**: `go test -race -cover`. Provide table-driven tests & small fakes. Do not hit prod resources. Use test containers or mocks.

---

## 3) Architecture Rules

* Follow **Clean Architecture** boundaries:

  * **domain**: pure models + interfaces + errors (no framework imports).
  * **usecase/service**: business logic orchestrations, calling domain interfaces.
  * **adapters**: http/gin handlers, db repositories, external services (LINE OA, etc.).
  * **infra**: db clients (GORM + PostgreSQL), LINE bot SDK integration, config, DI.
* No circular dependencies. One direction: outer depends on inner.
* Do not leak DB entities to API responses; always map to DTOs.

---

## 4) Security & Compliance

* **Secrets**: Never hardcode. Load from env or secret manager.
* **Auth**: Integrate JWT or external IAM when specified. Middleware should enforce auth/role checks.
* **Input validation**: All external inputs validated and sanitized.
* **Logging**: No PII in logs. Use structured logs (Zap/Logrus). Keep correlation IDs.
* **Dependency safety**: Prefer pinned versions. Avoid abandoned libs.

---

## 5) Observability

* Use request IDs (trace IDs) and propagate via context.
* Emit metrics for latency, error counts, and critical paths.
* Health checks: `/healthz` (liveness) and `/readyz` (readiness).

---

## 6) Data & Migrations

* Migrations live in `/migrations` (golang-migrate or equivalent).
* Schema changes must be backward compatible for rolling deploys.
* Provide rollback scripts.

---

## 7) CI/CD & Dev Tooling

* **Git**: Conventional Commits. Branching model per repo (`feat/`, `fix/`, `chore/`).
* **CI**: GitLab CI with stages `build`, `test`, `migrate`, `deploy`. Fail fast.
* **Docker**: Multi-stage builds; small runtime images; healthcheck. No root user where possible.
* **Makefile**: Add targets for `run`, `test`, `lint`, `migrate-up`, `migrate-down`, `seed`.

---

## 8) API Design Checklist

* [ ] Endpoint name and method follow REST semantics.
* [ ] Request/response schemas defined and validated.
* [ ] Error model documented (code, message, details, correlationId).
* [ ] Pagination, filtering, sorting (if applicable).
* [ ] Auth/authorization + scopes.
* [ ] Idempotency for unsafe methods.
* [ ] Tests: unit + handler + integration (happy & failure).
* [ ] Observability: logs + metrics + traces.
* [ ] Documentation snippet added to `/docs/api`.

---

## 9) Testing Policy

* Unit tests do not call real external services. Use fakes/mocks.
* DB tests run against ephemeral/local containers, not prod.
* Include race detection and coverage gates.
* For flaky timing code, use virtual clocks or deterministic helpers.
* When testing LINE OA integration, mock `line-bot-sdk-go` client.

---

## 10) What Copilot **Must Not** Do

* ❌ Invent endpoints, fields, or env vars that are not present in code/docs.
* ❌ Commit secrets or long random tokens in examples.
* ❌ Introduce breaking schema changes without a migration plan.
* ❌ Bypass validation, error handling, or logging policies.

---

## 11) Preferred Patterns & Snippets

### Go (Gin handler with service)

```go
func (h *WebhookHandler) HandleWebhook(c *gin.Context) {
    var req WebhookRequest
    if err := c.ShouldBindJSON(&req); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request"})
        return
    }

    ctx, cancel := context.WithTimeout(c.Request.Context(), 5*time.Second)
    defer cancel()

    if err := h.service.ProcessWebhook(ctx, req); err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    c.JSON(http.StatusOK, gin.H{"status": "ok"})
}
```

### Go (Service using LINE Bot SDK)

```go
func (s *WebhookService) ProcessWebhook(ctx context.Context, req WebhookRequest) error {
    // Example: reply with a text message
    message := linebot.NewTextMessage("Hello, this is an automated reply!")
    _, err := s.lineClient.PushMessage(req.UserID, message).Do()
    return err
}
```

---

## 12) Example Prompts for Copilot

* “Generate a Gin handler `POST /webhook/line` that receives LINE webhook events and passes them to `WebhookService.ProcessWebhook`.”
* “Add a GORM repository method `SaveWebhookEvent(ctx, event)` with unit tests (table-driven).”
* “Implement a service method using `line-bot-sdk-go` to send Flex messages to a LINE user.”
* “Create GitLab CI job `migrate` that runs golang-migrate migrations before `deploy`, using env `DATABASE_URL`.”

---

## 13) Directory Conventions (suggested)

```
/internal
  /domain
  /usecase
  /adapters
  /infra
/migrations
/scripts
/k8s | /deploy
/docs
```

---

## 14) Review Checklist (PRs)

* [ ] Code follows style & boundaries
* [ ] Tests cover success and failure paths
* [ ] No secrets, no PII in logs
* [ ] Migrations + rollbacks included
* [ ] Docs/README updated

---

**Keep changes small, typed, and well-tested. Prefer clarity over cleverness.**
