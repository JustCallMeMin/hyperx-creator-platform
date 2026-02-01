---
description: Display agent and project status. Progress tracking and status board.
---

---
description: Display project, agent, file, and preview server status for situational awareness.
---

## Invocation Rules

This workflow MAY be invoked only when:
- An overview of current project progress or agent activity is required.
- Inspecting preview server status or health.
- Tracking file creation or modification counts.

This workflow MUST NOT be invoked when:
- Implementing or modifying code.
- Running debugging, deployment, or planning workflows.

---

# /status — System Status Overview

$ARGUMENTS

---

## Purpose

Provide a consolidated and read-only overview of the current project state.
This workflow is intended for visibility and coordination, not execution.

---

## Behavior

When `/status` is triggered:

1. **Project Overview**
   - Display project name, path, and classification.
   - Show detected technology stack.
   - List implemented and pending features.

2. **Agent Status Board**
   - Show active, completed, and waiting agents.
   - Indicate current focus or progress percentage if available.

3. **File Statistics**
   - Count files created and modified in the current session.

4. **Preview Status**
   - Report preview server status.
   - Display URL and health check result if running.

---

## Output Expectations

The status report MUST include:
- Project summary
- Agent activity overview
- File statistics
- Preview server status

---

## Output Format

```markdown
## Project Status

### Project
- Name: my-ecommerce
- Path: /projects/my-ecommerce
- Type: nextjs-ecommerce
- State: active

### Tech Stack
- Framework: Next.js
- Database: PostgreSQL
- Authentication: Clerk
- Payments: Stripe

### Features
Completed:
- product-listing
- cart
- checkout
- user-auth
- order-history

Pending:
- admin-panel
- email-notifications

### Files
- Created: 73
- Modified: 12

### Agents
| Agent | Status | Notes |
|------|--------|-------|
| database-architect | Completed | |
| backend-specialist | Completed | |
| frontend-specialist | In Progress | Dashboard (60%) |
| test-engineer | Waiting | |

### Preview
- URL: http://localhost:3000
- Health: OK
```

---

## Key Principles

- Read-only visibility.
- No side effects.
- Status reflects current session state only.
