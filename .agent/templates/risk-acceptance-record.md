# Risk Acceptance Record (RAR)

Record ID: [RAR-YYYYMMDD-XXX]
Task Link: [task-slug]
Status: [ACTIVE | EXPIRED | RESOLVED]

---

## 1. Risk Description

- **Nature of Risk**: [Security / Technical Debt / QA / Performance]
- **Specialist Agent Report**: Link to the specific logic or test that failed.
- **Severity**: [HIGH | MEDIUM | LOW]
- **Impact Analysis**: What happens if this risk manifests?

---

## 2. Business Justification

- **Why are we accepting this?**: [e.g., Critical market deadline, low probability of exploit, temporary fix planned]
- **Trade-off Analysis**: Benefit of shipping vs. Risk of failure.

---

## 3. Governance & Ownership

- **Risk Owner**: @[agent-name or human]
- **Explicit Acceptance Statement**: "I accept the responsibility for this risk under the conditions defined in this record."
- **Date of Acceptance**: [YYYY-MM-DD]

---

## 4. Mitigation & Resolution (MANDATORY)

- **Interim Mitigations**: What stops this from being a total failure?
- **TTL (Expiry Date)**: [YYYY-MM-DD] (Max 14 days unless justified)
- **Follow-up Task ID**: [backlog-id] - Assigned to resolve this debt.

---

## 5. Closure Criteria

- How will we know this risk is neutralized?
- Verification agent required for closure: @[specialist-agent]

---

Approved by: @orchestrator
Logged to: project-governance.log
