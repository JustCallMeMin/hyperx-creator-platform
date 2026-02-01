# PostgreSQL Platform Rules
> Database Platform Rules for PostgreSQL

These rules apply ONLY when PostgreSQL is the selected database platform.

---

## Core Capabilities

PostgreSQL SHOULD be used when:

- Strong relational integrity is required
- Advanced constraints and indexing are needed
- JSONB or semi-structured data is involved
- Extensions (pgvector, PostGIS, full-text search) are required

---

## Schema & Constraint Rules

You MUST:

- Prefer `BIGINT` or `UUID` for primary keys
- Use `CHECK` constraints for domain validation
- Enforce relationships using `FOREIGN KEY`
- Avoid application-only enforcement of invariants

Business correctness belongs in the database.

---

## Indexing Rules

You SHOULD:

- Use `btree` for equality and range queries
- Use `GIN` for JSONB and array fields
- Avoid over-indexing write-heavy tables

Indexes MUST be driven by real query patterns.

---

## Performance & Planning

You MUST:

- Use `EXPLAIN ANALYZE`
- Monitor sequential scans on large tables
- Tune based on actual execution plans, not assumptions

---

## Anti-Patterns

You MUST NOT:

- Disable constraints for convenience
- Treat PostgreSQL like MySQL
- Rely on ORM defaults blindly
