---
description: Coordinate multiple specialized agents to solve complex, multi-domain tasks in a controlled and verifiable manner.
---

## Invocation Rules

This workflow MAY be invoked only when:
- A task spans multiple domains and cannot be solved reliably by a single agent.
- At least three specialized agents are required.
- The user explicitly requests orchestration or uses `/orchestrate`.

This workflow MUST NOT be invoked when:
- The task is simple or confined to a single domain.
- No plan exists and the task has not entered Planning Phase (unless explicitly allowed).

---

# /orchestrate — Multi-Agent Coordination

$ARGUMENTS

---

## Purpose

Coordinate multiple specialized agents to address complex tasks. 
This workflow enforces planning discipline, agent boundaries, and verification requirements across domains.

---

## Minimum Agent Requirement

Orchestration requires a minimum of **three different agents**.

If fewer than three agents are used:
- The task is considered delegation, not orchestration.
- The workflow MUST NOT be marked complete.

---

## Agent Selection Guidelines

| Task Type | Required Agents (minimum) |
| :--- | :--- |
| Web Application | frontend-specialist, backend-specialist, test-engineer |
| API / Backend | backend-specialist, security-auditor, test-engineer |
| UI / Design | frontend-specialist, seo-specialist, performance-optimizer |
| Database | database-architect, backend-specialist, security-auditor |
| Full Stack | project-planner, frontend-specialist, backend-specialist, devops-engineer |
| Debugging | debugger, explorer-agent, test-engineer |
| Security | security-auditor, penetration-tester, devops-engineer |

---

## Execution Model

Orchestration is executed in two strict phases.

### Phase 1: Planning
- Only the following agents MAY be invoked:
  - `project-planner`
  - `explorer-agent` (optional, for discovery)
- Output:
  - A formally written plan file (`{task-slug}.md`)
- **Strictly No implementation** is allowed during this phase.

### Approval Gate
Execution MUST pause until the plan is explicitly approved.

---

### Phase 2: Implementation
- After approval, multiple agents MAY be invoked in parallel.
- Typical execution groups:
  - **Foundation**: database-architect, security-auditor
  - **Core**: backend-specialist, frontend-specialist
  - **Validation**: test-engineer, devops-engineer

---

## Context Passing Requirements

Every agent invocation MUST include:
1. Original user request
2. All decisions made so far
3. Summary of previous agent outputs
4. Current plan state (if any)

Invoking an agent without full context is a protocol violation.

---

## Verification Requirements

Before orchestration is considered complete:
- Verification scripts MUST be executed as appropriate.
- Security and quality checks are mandatory.
- Results MUST be recorded in the final report.

---

## Output Expectations

The orchestration result MUST include:
- List of agents invoked (minimum three)
- Summary of each agent’s contribution
- Verification results
- Final synthesized outcome

---

## Output Format

```markdown
## Orchestration Report

### Task
[Summary of the original request]

### Mode
[Current AOS mode: planning / execution]

### Agents Invoked
| Agent | Domain | Status |
| :--- | :--- | :--- |
| project-planner | Planning | Done |
| backend-specialist | Backend | Done |
| test-engineer | Verification | Done |

### Verification
- Security checks: Pass / Fail
- Quality checks: Pass / Fail

### Key Findings
- [Agent]: Key result
- [Agent]: Key result

### Deliverables
- Plan file
- Implemented changes
- Verification evidence

### Summary
[Unified synthesis of all agent work]
```

---

## Exit Gate

Before marking orchestration complete, confirm:
- At least three agents were invoked.
- Required verification steps were executed.
- A final orchestration report was produced.

If any condition is not met, orchestration MUST continue.

---

## Usage Examples

- **Multi-domain**: `/orchestrate refactor the API and update the React dashboard`
- **Security Audit**: `/orchestrate review the auth system for vulnerabilities and fix them`
- **Performance**: `/orchestrate optimize database queries and implement server-side caching`

---

## Pre-Flight Questions

- Does this task require expertise from more than one domain?
- Can this task be completed reliably by fewer than three specialized agents?
- Is there already an approved plan (`{task-slug}.md`) for this objective?
