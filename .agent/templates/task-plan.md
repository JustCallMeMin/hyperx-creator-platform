# Task: [Clear and Concise Task Name]
**Slug**: [task-slug]
**Lead Agent**: @[agent-name]
**Status**: [DRAFT | REVIEW | APPROVED]

---

## 1. Business Context & Objective
- **Purpose**: Why are we doing this? What problem are we solving for the user or the business?
- **Business Rules**: List specific business invariants that must be upheld (Reference `business-rules.md`).
- **Terminology**: Ensure alignment with definitions in `glossary.md`.

---

## Decision Log (Optional but Recommended)
- **Decision**: [The primary decision made]
- **Decision Owner**: [The agent or human who made it]
- **Alternatives considered**: [Brief list]
- **Trade-offs accepted**: [What we gave up for this decision]
- **Date**: [YYYY-MM-DD]

---

# Task: [Short business-facing title]

Slug: [task-slug]
Primary Agent: @[agent-name]
Related Agents: @[agent-name], @[agent-name]
Status: DRAFT | REVIEW | APPROVED
Created: YYYY-MM-DD
Last updated: YYYY-MM-DD
Priority: P0 | P1 | P2
Change Risk: Low | Medium | High | Emergency

---

## 0. Executive summary (1–3 lines)
- What this does
- Business outcome / KPI impacted
- Required approval(s)

---

## 1. Business Context (MANDATORY)
- **Goal:** Why this task exists; business problem and metric(s) to move.
- **Users / personas affected.**
- **Business rules & invariants:** links to `business-rules.md` (explicit citations).  
  *Rule:* If a business invariant is affected, include the exact clause and who owns it. :contentReference[oaicite:4]{index=4}

---

## 2. Scope & Out-of-scope (MANDATORY)
- **In-scope:** exact boundaries of the change.
- **Out-of-scope:** explicitly list what will NOT be changed.

---

## 3. System / Technical Flow (Current vs Proposed) (MANDATORY)
### 3.1 Current flow (step-by-step)
- Diagram reference (link / ASCII)
- Modules / agents / services involved

### 3.2 Proposed flow (step-by-step)
- Exactly which steps change (numbered)
- Data flow & control flow changes
- Sequence diagrams or request traces if available

*Best practice:* Keep the “delta” explicit — reviewers should see exactly what differs. :contentReference[oaicite:5]{index=5}

---

## 4. Stakeholders & Roles (MANDATORY)
- Business owner: @
- Primary implementer (agent): @
- Reviewers / approvers (Orchestrator, DB Architect, Security, QA): @

---

## 5. Detailed Task Breakdown (Actionable checklist)
- [ ] Task A — description — Owner: @ — Est: Xh — Dep: task-slug-X
- [ ] Task B — description — Owner: @ — Est: Xh
- Include sub-tasks for docs, tests, monitoring, rollbacks.

*Best practice:* keep tasks small and reviewable (PR-able). :contentReference[oaicite:6]{index=6}

---

## 6. Risk Assessment & Mitigations (MANDATORY)
- **Risk matrix:** (Likelihood × Impact)
  - Risk 1: description → mitigation / rollback
  - Risk 2: description → mitigation / rollback
- **Blast radius:** services / regions / customers affected
- **Canary / phased rollout plan** (if applicable)

*Use change classification (Low/Medium/High/Emergency) to route approvals.* :contentReference[oaicite:7]{index=7}

---

## 7. Rollout & Rollback Plan (MANDATORY if Risk ≥ Medium)
- Pre-deploy checks (metrics, backups)
- Deployment steps (commands / pipeline stages)
- Verification steps (healthchecks, smoke tests)
- Rollback conditions & exact rollback commands
- Post-rollback validation

*Rollback is a feature — design it and test it.* :contentReference[oaicite:8]{index=8}

---

## 8. Testing & Acceptance Criteria (MANDATORY)
- Unit tests required
- Integration/E2E tests required
- Performance test criteria (if applicable)
- Acceptance Criteria (explicit pass/fail conditions)

*No production push without test coverage matching the risk level.* :contentReference[oaicite:9]{index=9}

---

## 9. Observability & SLOs to Monitor (MANDATORY)
- Key metrics to watch (latency, error rate, throughput)
- Dashboards / alerts to temporarily enable
- Rollout watch window & escalation path

*If monitoring isn't available → do not deploy.* :contentReference[oaicite:10]{index=10}

---

## 10. Security, Compliance & Data Migration Notes
- Data model changes and migration scripts
- PII / compliance impact and approvals
- Secrets/credentials handling

*If DB schema or migration is involved → route to `database-architect` and Orchestrator gate.* :contentReference[oaicite:11]{index=11}

---

## 11. Documentation & ADR / RFC references
- Linked RFC/ADR (if decision-level)
- Docs to update (README, API reference, runbook)
- Post-deploy docs / runbook update required? (Yes/No) — who owns it

*Large decisions must have ADRs recorded.* :contentReference[oaicite:12]{index=12}

---

## 12. Timeline & Rollout Phases
- Planned start / end
- Canary windows
- Maintenance windows (if needed)

---

## 13. Approvals (MANDATORY)
- [ ] Business Analyst
- [ ] Orchestrator
- [ ] Database Architect (if DB touched)
- [ ] Security (if applicable)
- [ ] DevOps (for production deploy)

---

## 14. Postmortem / Follow-up
- Owner for post-deploy verification
- Date to review and close task
- Postmortem required if incident or if RCA determines systemic causes. :contentReference[oaicite:13]{index=13}

## 2. Technical Analysis & Flow
- **Current State**: Describe the existing implementation or flow.
- **Proposed Changes**: Describe the new flow or structural changes step-by-step.
- **Architectural Impact**: How does this affect other modules, performance, or system boundaries?

## 3. Detailed Task Breakdown
Complete the following steps in sequence:
- [ ] Task 1: [Technical description] -> Assigned to: @[agent]
- [ ] Task 2: [Technical description] -> Assigned to: @[agent]

## 4. Risk Mitigation & Verification
- **Identified Risks**: What could go wrong? What are the edge cases?
- **Rollback Strategy**: How do we revert changes if something fails in production?
- **Acceptance Criteria (AC)**: How do we prove the task is completed and correct?
- **Testing Strategy**: List required unit, integration, or characterization tests.

---

## 5. Documentation Impacts
- [ ] Update `README.md`
- [ ] Update API Documentation
- [ ] Record Architectural Decisions (ADR)
