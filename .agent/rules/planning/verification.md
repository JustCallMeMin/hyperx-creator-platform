# Verification & Phase X Rules

## Purpose

Define the "Definition of Done" and enforce mandatory verification scripts.

---

## Phase X: The Final Gate

A project is NOT complete until the "Phase X" checklist is finalized in the plan file.

### Required Checks

1. **LINT**: `npm run lint` or equivalent must pass without errors.
2. **BUILD**: `npm run build` must succeed (for compiled projects).
3. **SECURITY**: `python .agent/skills/vulnerability-scanner/scripts/security_scan.py .` MUST pass.
4. **UX/UI**: `python .agent/skills/frontend-design/scripts/ux_audit.py .` (if applicable).
5. **PERFORMANCE**: `python .agent/skills/performance-profiling/scripts/lighthouse_audit.py <URL>`.

---

## Verification Discipline

- **Execute, don't assume**: Never mark a task as complete `[x]` without actually running the verification steps listed in the plan.
- **Evidence-based**: If a script fails, the task is incomplete. Fix the issue before moving to the next task.
- **Outcome Verification**: Manual verification is required for UX/UI changes that scripts cannot fully capture.

---

## STOP Gate

You MUST STOP if:

- A P0 verification script fails.
- The dev server (`npm run dev`) cannot start in a working state.
