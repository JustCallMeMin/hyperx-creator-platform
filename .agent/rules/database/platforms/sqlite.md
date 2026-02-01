# SQLite Platform Rules
> Database Platform Rules for SQLite

These rules apply ONLY when SQLite is used.

---

## When SQLite Is Appropriate

SQLite SHOULD be used for:

- Embedded systems
- Local-first or edge deployments
- Read-heavy workloads
- Small to medium datasets

---

## Schema Rules

You MUST:

- Enable foreign keys explicitly
- Avoid complex concurrency assumptions
- Design schemas defensively

SQLite enforces less by default. Be explicit.

---

## Concurrency Considerations

You MUST:

- Assume single-writer concurrency
- Avoid long-running write transactions
- Batch writes where possible

---

## Anti-Patterns

You MUST NOT:

- Assume horizontal scalability
- Use SQLite as a drop-in replacement for PostgreSQL
