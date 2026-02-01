---
description: Create new application command. Triggers App Builder skill and starts interactive dialogue with user.
---

---
description: Create a new application or major system from scratch. Initializes planning and orchestration flow.
---

## Invocation Rules

This workflow MAY be invoked only when:
- The user wants to start a completely new application or project.
- A new major system, module, or microservice is being initialized.
- No existing architecture or plan has been approved yet.

This workflow MUST NOT be invoked when:
- Adding features to an existing application (use `/enhance`).
- Debugging or fixing issues in an existing system (use `/debug`).
- A task-plan (`{task-slug}.md`) has already been approved.

---

# /create — Create New Application

$ARGUMENTS

---

## Purpose

Start a structured application creation process.
This workflow ensures that new systems are properly analyzed, planned, and approved
before any implementation begins.

---

## Behavior

When `/create` is triggered:

1. **Request Analysis**
   - Understand the user’s intent and high-level goals.
   - Identify missing or ambiguous information.
   - Ask clarifying questions if required.

2. **Project Planning**
   - Invoke `project-planner` to:
     - Break down scope into tasks.
     - Propose architecture and tech stack.
     - Plan file and directory structure.
     - Generate a `{task-slug}.md` plan file.

3. **Approval Gate**
   - Pause execution until the plan is reviewed and approved.
   - No implementation may start before approval.

4. **Application Building (Post-Approval)**
   - Orchestrate execution using appropriate agents:
     - `database-architect` for data modeling.
     - `backend-specialist` for APIs and services.
     - `frontend-specialist` or `mobile-developer` for UI.

5. **Preview and Handoff**
   - Generate a runnable preview if applicable.
   - Present access details or next steps to the user.

---

## Output Expectations

- A clear understanding of application scope and intent.
- A formally written `{task-slug}.md` plan file.
- Explicit approval checkpoint before implementation.
- No production code written prior to approval.

---

## Usage Examples

```
/create blog site
/create e-commerce application with product listing and cart
/create todo application
/create social media platform
/create CRM system with customer management
```

---

## Pre-Flight Questions

If the request is unclear, ask:
- What type of application is being built?
- Who are the intended users?
- What are the core features or goals?

Defaults MAY be proposed but MUST be confirmed.
