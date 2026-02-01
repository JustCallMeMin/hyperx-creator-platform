---
description: Add or update features in existing application. Used for iterative development.
---

---
description: Iterative enhancement workflow for adding or modifying features in an existing application.
---

## Invocation Rules

This workflow MAY be invoked only when:
- Adding new features to an existing application.
- Modifying or refactoring existing logic.
- Extending functionality of an already initialized project.

This workflow MUST NOT be invoked when:
- Creating a new application from scratch (use `/create`).
- Performing investigation or root cause analysis only (use `/debug`).
- A deployment-only action is required (use `/deploy`).

---

# /enhance — Update Existing Application

$ARGUMENTS

---

## Purpose

Apply controlled and incremental changes to an existing application.
This workflow ensures that enhancements are planned, reviewed,
and applied without breaking existing behavior.

---

## Behavior

When `/enhance` is triggered:

1. **Current State Assessment**
   - Load and review the current project state.
   - Identify existing features, architecture, and constraints.
   - Confirm the scope of change.

2. **Change Planning**
   - Define what will be added, modified, or removed.
   - Identify affected files, modules, and dependencies.
   - Assess risk and potential side effects.

3. **Approval Gate**
   - For significant changes, present a clear change plan.
   - Wait for explicit approval before proceeding.

4. **Implementation**
   - Invoke appropriate agents based on scope.
   - Apply changes incrementally.
   - Run relevant tests and validations.

5. **Verification and Preview**
   - Verify that existing functionality still works.
   - Update preview or runtime environment if applicable.

---

## Output Expectations

An enhancement result MUST include:
- Clear description of changes made.
- Confirmation that existing behavior is preserved.
- Evidence of verification (tests or checks).
- Identification of any follow-up work if required.

---

## Usage Examples

```
/enhance add dark mode
/enhance build admin panel
/enhance integrate payment system
/enhance add search feature
/enhance update profile page
/enhance improve responsiveness
```

---

## Caution and Constraints

- Major changes require explicit approval.
- Conflicting architectural decisions must be flagged.
- Each enhancement should be applied in logical, reviewable units.
- Changes should be traceable and reversible when possible.

---

## Key Principles

- Prefer small, incremental changes.
- Preserve existing behavior unless explicitly changed.
- Validate before and after applying enhancements.
- Avoid introducing hidden dependencies.
