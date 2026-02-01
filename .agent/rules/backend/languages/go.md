---
trigger: always_on
---

# Go Backend Rules

> Contextual Language Rules for Go Projects

These rules apply **ONLY** when working on Go-based backend code.
They are loaded dynamically by the Backend Specialist Agent.

---

## 1. Core Go Principles (MANDATORY)

You MUST follow these principles at all times:

- Prefer simplicity over cleverness
- Favor explicit behavior over abstraction
- Avoid unnecessary layers or indirection
- Write code that is easy to read, test, and reason about

If a pattern increases cognitive load, it is likely wrong.

---

## 2. Package & Layering Rules

### Domain vs Infrastructure Separation

You MUST:

- Keep **domain logic** free from:
  - HTTP frameworks
  - Database drivers
  - Message brokers
  - Serialization formats
- Isolate infrastructure concerns in adapter layers
- Ensure domain packages can be tested without external dependencies

Domain code MUST compile without importing infrastructure packages.

---

## 3. Standard Project Layout (MANDATORY)

You MUST follow the standard Go project layout:

- `/cmd` — Entry points (main.go)
- `/internal` — Private application code (not importable by others)
  - `/domain` — Pure business logic, entities, interfaces
  - `/usecase` — Application services, orchestrating domains
  - `/infra` — Concrete implementations (DB, HTTP, Brokers)
- `/pkg` — Public library code (safe for external import)
- `/api` — API definitions (OpenAPI, gRPC proto)

Avoid flat structures for non-trivial services.

---

## 3. Interface Design Rules

- Define interfaces at the **consumer side**, not the producer
- Keep interfaces small (1–3 methods)
- Avoid “god interfaces”
- Use interfaces to enable testing and substitution, not abstraction for its own sake

Do NOT create interfaces unless at least one alternative implementation is expected or testing requires it.

---

## 4. Error Handling

You MUST:

- Handle errors explicitly
- Return errors, do not panic (except at process boundaries)
- Wrap errors with context using `%w`
- Avoid sentinel errors unless strictly necessary
- Never ignore returned errors

Errors are part of the API contract.

---

## 5. Constructors & Dependency Injection

- Prefer explicit constructors
- Inject dependencies via parameters, not globals
- Avoid package-level mutable state
- Ensure constructors are deterministic and side-effect free

If an object cannot be constructed without side effects, reconsider the design.

---

## 6. Testing Rules

- Design for testability first
- Avoid tightly coupling logic to concrete implementations
- Use interfaces to mock external dependencies
- Prefer table-driven tests
- Keep tests deterministic and isolated

Test code MUST NOT influence production behavior.

---

## 7. Concurrency & Goroutines

You MUST:

- Keep concurrency explicit and controlled
- Avoid launching goroutines without ownership
- Use context for cancellation
- Ensure goroutines can terminate cleanly

Leaking goroutines is a production bug.

---

## 8. Context Hygiene (MANDATORY)

- **Propagation**: Always pass `context.Context` as the first argument to functions that perform I/O.
- **Boundaries**: Start contexts at the entry point (HTTP handler, CLI command) and propagate downwards.
- **Values**: Use context values **ONLY** for request-scoped data (Trace IDs, User identity). NEVER use context for optional parameters or dependencies.
- **Cleanup**: Always call the `cancel` function returned by `WithTimeout` or `WithCancel` (usually via `defer`).

---

## 8. Anti-Patterns (Forbidden)

You MUST NOT:

- Hide logic behind reflection or magic
- Overuse generics
- Introduce frameworks into domain code
- Use global mutable state
- Mix HTTP/database logic with business rules

---

## Final Rule

Idiomatic Go is **boring, explicit, and predictable**.

If the code feels “clever”, it is probably wrong.
