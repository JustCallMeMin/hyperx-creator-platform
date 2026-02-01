# Agent Selection Rules

## Purpose

Standardize how tasks are assigned to specialist agents based on domain boundaries.

---

## Domain Ownership Matrix

Assign tasks based on the primary domain:

| Domain | Agent | Exclusions |
| :--- | :--- | :--- |
| **Mobile** | `mobile-developer` | Is Full-Stack for mobile (UI + Logic + Local DB). |
| **Web UI** | `frontend-specialist` | Focus on UI, CSS, and Client-side logic. |
| **Backend** | `backend-specialist` | Focus on APIs, Server logic, and Frameworks. |
| **Database** | `database-architect` | Mandatory for schema changes and migrations. |
| **Security** | `security-auditor` | Mandatory for auth design and audit. |
| **Infra** | `devops-engineer` | Mandatory for CI/CD and deployments. |

---

## Selection Heuristics

- **Web vs. Mobile**: NEVER use `frontend-specialist` for mobile apps.
- **Logic vs. UI**: If a task is 80% business logic, prefer `backend-specialist` even if it affects the UI layer.
- **Verification**: Always assign `test-engineer` for building dedicated test suites or E2E harnesses.

---

## STOP Gate

You MUST STOP if:

- A task boundary is "fuzzy" and touches multiple domains without a lead agent.
- An agent is assigned a task outside its CORE responsibility (e.g. Designer writing SQL).
