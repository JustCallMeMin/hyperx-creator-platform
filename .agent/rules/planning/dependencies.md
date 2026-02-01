# Dependency & Ordering Rules

## Purpose

Ensure a stable and efficient execution sequence while preventing merge conflicts.

---

## Blocking Analysis

- **Hard Blockers**: Only include dependencies that are strictly necessary for the next task to start.
- **Circuit Breakers**: Identify "Critical Path" tasks (P0) that must pass verification before any downstream work begins.

---

## Parallel vs. Serial Execution

- **Parallelism**: Tasks can be parallel if they involve different agents AND different files.
- **Serialization**: Tasks MUST be serial if they:
  - Modify the same file.
  - Follow a "Provider → Consumer" pattern (e.g., Schema first, then API).
  - Involve breaking changes that require immediate sync.

---

## STOP Gate

You MUST STOP if:

- Circular dependencies are detected.
- The execution order requires multiple agents to edit the same file simultaneously.
