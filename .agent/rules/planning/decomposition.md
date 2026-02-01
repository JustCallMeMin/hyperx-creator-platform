# Task Decomposition Rules

## Purpose

Define how to break down complex requests into manageable, verifiable units of work.

---

## Project Type Detection (MANDATORY)

Before assigning tasks, you MUST classify the project type to ensure correct agent selection:

| Target Context | Project Type | Primary Agent |
| :--- | :--- | :--- |
| Mobile app, iOS, Android, RN, Flutter | **MOBILE** | `mobile-developer` |
| Website, Web App, Next.js, React (web) | **WEB** | `frontend-specialist` |
| API, Backend, Standalone Database | **BACKEND** | `backend-specialist` |

---

## Task Sizing & Scope

- **Duration**: Each task should represent 2-10 minutes of work.
- **Outcome**: Each task MUST have one clear, singular outcome.
- **Atomicity**: Tasks should be small enough to allow for easy rollback if verification fails.

---

## Task Format (MANDATORY)

Every task in a plan MUST include:

- `task_id`: Numeric or alphanumeric ID.
- `name`: Concise description.
- `agent`: Assigned specialist agent.
- `skills`: Required skills for implementation.
- `priority`: P0 (Critical) to P3 (Optional).
- `dependencies`: List of blocking Task IDs.
- **INPUT → OUTPUT → VERIFY**:
  - **INPUT**: What the agent needs to start.
  - **OUTPUT**: What the agent produces.
  - **VERIFY**: Concrete steps to prove the output is correct.

---

## STOP Gate

You MUST STOP if:

- A task lacks explicit verification criteria.
- The project type is "Mixed" but primary domain ownership is not assigned.
