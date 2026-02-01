---
description: Create clean, conventional commits. Review changes, stage logically, and check for secrets before committing.
---

## Invocation Rules

This workflow MAY be invoked when:
- Work has been completed and needs to be saved.
- The user asks to "commit", "save work", or "checkpoint".
- A logical unit of work (feature, fix, refactor) is done.

This workflow MUST NOT be invoked when:
- Tests are currently failing (unless committing a "wip: failing test" explicitly).
- The worktree is empty.

---

# /commit — Commit Work

$ARGUMENTS

---

## Purpose

To ensure all commits in the repo are high-quality, reviewable, and safe.
This workflow forces a "Review -> Stage -> Commit" discipline instead of blind `git commit -am`.

---

## Behavior

When `/commit` is triggered:

1. **Inspect Status**
   - Run `git status` and `git diff --stat`.
   - Identify what has changed.

2. **Execute Commit Skill (MANDATORY)**
   - **REQUIRED SKILL**: `commit-work`
   - Ask: "Single commit or split?" if unclear.
   - For each logical change:
     - Review diff (`git diff`).
     - Stage (`git add`).
     - Write Conventional Commit message.
     - Verify (quick lint/test).
     - Commit.

3. **Report**
   - Show the commit hash and message.
   - Confirm working tree status.

---

## Output Expectations

- Clean, conventional commit message(s).
- No accidental files (secrets, logs) committed.
- Clear summary of what was saved.

---

## Usage Examples

```
/commit
/commit "feat: add user login"
```
