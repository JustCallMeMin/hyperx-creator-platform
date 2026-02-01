---
description: Execute an implementation plan using executing-plans skill.
---

---
description: Execute an approved implementation plan task-by-task with checkpoints. Uses executing-plans and using-git-worktrees skills.
---

## Invocation Rules

This workflow MAY be invoked only when:
- A valid implementation plan exists (e.g. `docs/plans/YYYY-MM-DD-feature.md`).
- The user has approved the plan for execution.

This workflow MUST NOT be invoked when:
- No written plan exists (use `/plan` first).
- The plan is vague or lacks verification steps.

---

# /execute — Plan Execution

$ARGUMENTS

---

## Purpose

Systematically implement a feature based on a written plan, ensuring isolation and verification at every step.

---

## Behavior

When `/execute` is triggered:

1. **Setup Workspace (MANDATORY)**
   - **REQUIRED SKILL**: `using-git-worktrees`
   - Create an isolated worktree for this feature.
   - Do NOT work directly on the main branch.

2. **Execute Plan (MANDATORY)**
   - **REQUIRED SKILL**: `executing-plans`
   - Load the specified plan file.
   - Execute tasks in batches (default: 3 tasks).
   - Stop and verify after each batch.

3. **Report**
   - Show progress against the plan.
   - Confirm verification results.

---

## Output Expectations

- An isolated worktree created.
- Code changes implemented according to plan.
- Tests passing for implemented features.
- Plan file updated with progress.

---

## Usage Examples

```
/execute docs/plans/2025-02-01-auth-feature.md
/execute
```
