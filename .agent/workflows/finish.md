---
description: Finish development work, merge, or create PR using finishing-a-development-branch skill.
---

---
description: Complete a development cycle. Verify tests, choose integration strategy (Merge/PR), and cleanup workspace.
---

## Invocation Rules

This workflow MAY be invoked only when:
- Development work is complete.
- All tests are believed to be passing.
- The user is ready to merge or submit the code.

This workflow MUST NOT be invoked when:
- Work is still in progress (use `/status` or just continue coding).
- Tests are known to be failing (fix them first).

---

# /finish — Finish Development

$ARGUMENTS

---

## Purpose

Safely wrap up a feature branch, ensuring quality gates are passed before merging or publishing.

---

## Behavior

When `/finish` is triggered:

1. **Verify Quality**
   - Run the full test suite.
   - If tests fail, STOP and report.

2. **Present Options**
   - **REQUIRED SKILL**: `finishing-a-development-branch`
   - Offer standard choices:
     1. Merge locally
     2. Push & Create PR
     3. Keep as-is
     4. Discard

3. **Execute & Cleanup**
   - Perform the selected action.
   - Clean up the worktree if appropriate.

---

## Output Expectations

- Test results verified.
- User choice executed (Merge/PR/etc).
- Workspace cleaned up (if merged/discarded).

---

## Usage Examples

```
/finish
```
