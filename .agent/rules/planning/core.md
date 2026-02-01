# Planning Core Principles

## Purpose

Define the foundational mindset and workflow for structured task planning.

---

## The 4-Phase Workflow

All non-trivial development MUST follow this sequence:

1. **ANALYSIS**: Research, brainstorm, explore codebase (Survey mode).
2. **PLANNING**: Decompose into tasks and create `{task-slug}.md`.
3. **SOLUTIONING**: Architecture, detailed design, and technical decisions.
4. **IMPLEMENTATION**: Code writing and unit testing.
5. **VERIFICATION (Phase X)**: Final integration checks and performance audits.

---

## Planning Discipline

- **Planning precedes building**: No code is written during Analysis or Planning phases.
- **English-only**: All plans and documentation MUST be written in English for cross-agent compatibility.
- **Dynamic Naming**: Plans are named `{task-slug}.md` (e.g., `auth-fix.md`) and stored in the project root.
- **Traceability**: Every task MUST be traceable to a specific requirement or business outcome.

---

## STOP Gate

You MUST STOP if:

- The high-level objective is ambiguous.
- OS-specific constraints are not confirmed (Windows vs macOS/Linux).
- Existing plan files in the root have not been reviewed.
