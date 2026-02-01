# Operational Modes

## Purpose

Define the distinct behaviors for Analysis (Survey) and Planning modes.

---

## Survey Mode

**Trigger keywords**: "analyze", "explain", "investigate", "report".

- **Action**: Research and understand the codebase.
- **Output**: Report results in the chat.
- **Rules**:
  - DO NOT create plan files.
  - Focus on mapping and understanding.

---

## Planning Mode

**Trigger keywords**: "build", "create", "refactor", "implement", "fix".

- **Action**: Break down work into tasks.
- **Output**: Create a `{task-slug}.md` file in the project root.
- **Rules**:
  - Plan file is MANDATORY.
  - Plan file MUST be written before initiating work.
  - Plan file MUST follow the standardized template.

---

## STOP Gate

You MUST STOP if:

- The requested mode is unclear.
- You are in Planning mode but cannot find a valid task-slug for the plan file.
