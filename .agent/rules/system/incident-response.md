# Incident Response & Postmortem Policy

## Purpose

Define the mandatory protocol for handling critical failures, regressions, or security breaches detected during Phase X verification or production monitoring.

---

## 1. Incident Detection & Triggering

An **Incident Mode** MUST be triggered by the `orchestrator` when:
- **Phase X Verification fails** on a CRITICAL or HIGH risk task.
- **Production regression** is reported (via telemetry or user report).
- **Security vulnerability** is identified in published code.

### Incident Command
The `security-auditor` and `orchestrator` form the **Incident Command**. No other agent may propose changes without their explicit coordination.

---

## 2. Response Flow (MANDATORY)

1. **Containment**: Immediate stop of deployment or activation of rollback procedures (Reference: `deployment-procedures.md`).
2. **Isolation**: The `debugger` must isolate the root cause without modifying production state.
3. **Emergency Planning**: Use `.agent/templates/incident-plan.md`. This plan overrides all regular task plans.
4. **Resolution**: Apply the minimum viable fix to restore stability.

---

## 3. Postmortem Protocol

Every incident MUST conclude with a **Postmortem Document** covering:
- **Timeline**: When it happened, when it was detected, when it was fixed.
- **Root Cause**: Why did the existing rules/gates fail?
- **Impact**: Domain, data, and business impact assessment.
- **Action Items**: Explicit tasks to prevent recurrence.

---

## 4. AOS Hardening (The Feedback Loop)

The most critical step of an incident is the **Hardening Phase**:
- Identify the specific **Rule File** (`.agent/rules/*`) that allowed the failure.
- Update the rule to cover the edge case.
- Update the **Orchestrator Quality Gate** to enforce this new check in future tasks.

**An incident is not "Closed" until the rules are updated.**

---

## STOP Rule

You MUST NOT exit Incident Mode until:
1. Stability is verified by `qa-automation-engineer`.
2. Postmortem is approved by `orchestrator`.
3. Rule updates are merged and pushed to origin.
