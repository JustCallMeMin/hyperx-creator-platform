---
description: Preview server start, stop, and status check. Local development server management.
---

---
description: Local preview server management workflow for starting, stopping, restarting, and checking development server status.
---

## Invocation Rules

This workflow MAY be invoked only when:
- Starting, stopping, restarting, or checking the status of a local preview server.
- Verifying UI or logic changes in a local or preview environment.

This workflow MUST NOT be invoked when:
- Running production deployments (use `/deploy`).
- Creating project plans or modifying planning artifacts (use `/plan`).

---

# /preview — Preview Server Management

$ARGUMENTS

---

## Purpose

Manage the lifecycle of a local preview server in a controlled and observable manner.
This workflow supports rapid verification without affecting production systems.

---

## Supported Commands

```
/preview
/preview start
/preview stop
/preview restart
/preview check
```

---

## Behavior

When `/preview` is triggered:

1. **Status Inspection**
   - Detect whether a preview server is currently running.
   - Report server URL, type, and health status if available.

2. **Server Start**
   - Start the local preview server using the configured framework.
   - Select an available port or report conflicts.

3. **Server Stop**
   - Gracefully stop the running preview server.

4. **Restart**
   - Stop and immediately restart the preview server.

5. **Health Check**
   - Perform a basic health check to confirm responsiveness.

---

## Output Expectations

The workflow MUST report:
- Server status (running or stopped)
- Active port and URL (if running)
- Framework or project type (if detectable)
- Health status

---

## Output Format

### Status Report

```markdown
## Preview Status

- URL: http://localhost:3000
- Project Path: /path/to/project
- Type: nextjs
- Health: OK
```

---

### Start Result

```markdown
## Preview Started

- Port: 3000
- Type: Next.js
- URL: http://localhost:3000
```

---

### Error: Port Conflict

```markdown
## Preview Start Failed

Reason:
Port 3000 is already in use.

Available Options:
1. Start on the next available port
2. Stop the process currently using the port
3. Specify a different port explicitly
```

---

## Technical Notes

Preview server management is implemented via the automation script:

```bash
python .agent/scripts/auto_preview.py start [port]
python .agent/scripts/auto_preview.py stop
python .agent/scripts/auto_preview.py restart
python .agent/scripts/auto_preview.py status
```

---

## Usage Examples

```
/preview
/preview start
/preview restart
/preview stop
/preview check
```

---

## Key Principles

- Preview environments are local and disposable.
- No production data or credentials should be used.
- Preview does not imply deployment.
- Failures must be reported clearly.
