---
description: Structured brainstorming for projects and features. Explores multiple options before implementation.
---

## Invocation Rules

This workflow MAY be invoked when:
- Requirements involve new features, architectural decisions, or complex trade-offs.
- Intent exploration is explicitly requested or required by AOS decision gates.
- User needs to explore design alternatives before committing to implementation.

This workflow MUST NOT be invoked when:
- A task-plan (`{task-slug}.md`) has already been approved.
- Implementation has already started.
- The task is a simple bug fix or trivial change.

---

# /brainstorm — Structured Idea Exploration

$ARGUMENTS

---

## Purpose

Turn vague ideas into fully formed designs through Socratic discovery and structured exploration.
This workflow reduces uncertainty, compares approaches, and clarifies intent before any code is written.

---

## Behavior

When `/brainstorm` is triggered:

1. **Execute Brainstorming Skill (MANDATORY)**
   - **REQUIRED SKILL**: `brainstorming`
   - The skill will guide the entire discovery process.
   - Do NOT manually implement brainstorming logic.

2. **Socratic Discovery**
   - Ask questions **one at a time** to refine the idea.
   - Prefer multiple choice questions when possible.
   - Focus on: purpose, constraints, success criteria.

3. **Explore Alternatives**
   - Propose 2-3 different approaches with trade-offs.
   - Present options conversationally with recommendation.
   - Lead with the recommended option and explain why.

4. **Present Design Incrementally**
   - Break design into sections (200-300 words each).
   - Ask after each section: "Does this look right so far?"
   - Cover: architecture, components, data flow, error handling, testing.

5. **Document Design**
   - Write validated design to `docs/plans/YYYY-MM-DD-<topic>-design.md`.
   - Commit the design document.

6. **Handoff to Implementation (Optional)**
   - Ask: "Ready to set up for implementation?"
   - If yes:
     - Use `using-git-worktrees` to create isolated workspace.
     - Use `writing-plans` to create detailed implementation plan.

---

## Output Expectations

- A validated design document saved to `docs/plans/`.
- Clear understanding of the problem and chosen approach.
- User confirmation before proceeding to implementation.

---

## Key Principles

- **One question at a time** - Don't overwhelm with multiple questions.
- **Multiple choice preferred** - Easier to answer than open-ended.
- **YAGNI ruthlessly** - Remove unnecessary features from all designs.
- **Explore alternatives** - Always propose 2-3 approaches before settling.
- **Incremental validation** - Present design in sections, validate each.
- **Be flexible** - Go back and clarify when something doesn't make sense.

---

## Usage Examples

```
/brainstorm authentication system
/brainstorm state management for complex form
/brainstorm database schema for social app
/brainstorm caching strategy
```

---

## Integration

This workflow is the **entry point** for the Feature Development Lifecycle:

1. **Brainstorm** (this workflow) → Design validated
2. **Plan** (`/plan`) → Implementation plan created
3. **Execute** (`/execute`) → Code implemented
4. **Finish** (`/finish`) → Merge/PR

---
