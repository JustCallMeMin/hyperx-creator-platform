# Postmortem Report Template

Incident ID: [INC-YYYYMMDD-XXX]
Date: [YYYY-MM-DD]
Severity: [CRITICAL | HIGH | MEDIUM]
Author: @[agent-name] (Orchestrator/Security-Auditor)

---

## 1. Executive Summary

Briefly describe the incident, duration, and final outcome.

- **Total Downtime / Unavailability**: [duration]
- **Business Impact**: [e.g., loss of revenue, user trust, data exposure]

---

## 2. Timeline

Chronological list of key events from detection to closure.

- `HH:MM UTC`: Incident occurred.
- `HH:MM UTC`: Detected via [Alerting/Phase X/User Report].
- `HH:MM UTC`: Stabilization actions taken (Containment).
- `HH:MM UTC`: Root cause identified.
- `HH:MM UTC`: Fix deployed and verified.
- `HH:MM UTC`: Resolution confirmed and AOS Hardening started.

---

## 3. Root Cause Analysis (RCA)

- **Technical Trigger**: What exactly failed in the code or infra?
- **Process Failure**: Why was this not caught by the `orchestrator` or `qa-automation-engineer` earlier?
- **AOS Gap**: Which rule file (`.agent/rules/*`) was insufficient or ignored?

---

## 4. AOS Hardening (The Feedback Loop)

**MANDATORY: Link to specific PRs or file changes that harden the system.**

- [ ] **Rule Update**: [e.g., Updated `security/auth.md` to enforce X check]
- [ ] **Orchestrator Gate Update**: [e.g., Added Y to the Security Review Checklist]
- [ ] **QA Suite Update**: [e.g., Added regression test Z to Playwright suite]

---

## 5. Lessons Learned

- What went well? (e.g., Containment speed)
- What could be improved? (e.g., Detection lag)
- Unexpected behaviors during the incident.

---

## 6. Closure Approval

- [ ] Incident Response satisfied.
- [ ] AOS Hardening merged.
- [ ] Evidence preserved.

Approved by: @orchestrator
Date: [YYYY-MM-DD]
