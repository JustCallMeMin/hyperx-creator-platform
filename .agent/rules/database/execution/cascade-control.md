# Cascade Control Rules
> Execution Chain & Procedure Optimization

---

## Core Principle

Implicit cascading execution is a hidden performance cost.

---

## Rules

You MUST:

- Avoid trigger → function → trigger chains
- Prefer explicit execution entry points
- Document execution chains when unavoidable

---

## Function Rules

You MUST:

- Avoid loops with internal queries
- Prefer set-based operations
- Assign correct volatility (IMMUTABLE / STABLE / VOLATILE)

---

## Measurement

You MUST:

- Measure call count per transaction
- Use EXPLAIN ANALYZE on function calls

---

## Anti-Patterns

You MUST NOT:

- Implement business workflows in triggers
- Create recursive execution chains
