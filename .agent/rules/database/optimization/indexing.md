# Indexing Optimization Rules
> Query Performance Rules

---

## Core Principles

Indexes exist to serve queries, not aesthetics.

---

## Rules

You MUST:

- Index columns used in WHERE, JOIN, ORDER BY
- Remove unused indexes
- Measure index effectiveness

Indexes have a write cost.

---

## Composite Index Rules

You SHOULD:

- Match index order to query order
- Avoid unnecessary wide composite indexes

---

## Anti-Patterns

You MUST NOT:

- Index everything
- Guess index needs without measurement
