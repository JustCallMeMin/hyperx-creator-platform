---
description: Create technical implementation plan using writing-plans skill.
---

---
description: Create a technical implementation plan for a feature, refactor, or task. Uses writing-plans skill to generate a standardized, verifiable plan.
---

## Invocation Rules

This workflow MAY be invoked only when:
- The user is ready to move from concept/design to technical planning.
- A clear goal or design exists (e.g. from `/brainstorm`).
- Implementation details need to be defined before coding.

This workflow MUST NOT be invoked when:
- The requirements are still vague (use `/brainstorm` first).
- The task is a trivial one-line change.

---

# /plan — Implementation Planning

$ARGUMENTS

---

## Purpose

Create a detailed, actionable, and verifiable implementation plan ("The Recipe") that any engineer could execute.

---

## Behavior

When `/plan` is triggered:

1. **Check Context**
   - Is the goal clear? If not, ask.
   - Is there a design? If not, briefly clarify approach.

2. **Execute Planning**
   - **MANDATORY**: Use the `writing-plans` skill.
   - Do NOT write the plan manually or invent a format.
   - Let the skill handle the file creation and structure.

3. **Plan Location**
   - The skill will save the plan to `docs/plans/YYYY-MM-DD-<task-slug>.md`.

4. **Handoff**
   - Present the created plan file path.
   - Offer execution options (Subagent-Driven or Parallel Session) as defined in the skill.

---

## Output Expectations

- A new plan file created by `writing-plans`.
- No code implementation (planning only).
- Clear next steps for execution.

---

## Usage Examples

```
/plan add user authentication
/plan refactor database schema
/plan implement dark mode
```
