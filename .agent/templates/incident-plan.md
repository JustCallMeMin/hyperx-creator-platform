# Incident Plan Template

Incident ID: [INC-YYYYMMDD-XXX]
Severity: LOW | MEDIUM | HIGH | CRITICAL
Status: OPEN | MITIGATED | RESOLVED
Primary Agent: @[agent-name]
Timestamp: [UTC]

---

## 1. Situation Overview (What is happening?)

- Observable symptoms (errors, downtime, data inconsistency)
- Affected users / systems
- When did it start?

⚠️ Facts only. No assumptions.

---

## 2. Immediate Risk Assessment

- Is data integrity at risk?
- Is security impacted? (PII exposure, unauthorized access)
- Is system performance degrading gracefully or crashing?
- Is the blast radius growing?

If YES → escalation required.

---

## 3. Stabilization Actions (MANDATORY FIRST)

Goal: **Stop the bleeding**

- [ ] Disable feature / flag
- [ ] Rollback deployment
- [ ] Scale / isolate components
- [ ] Apply temporary guard (rate limit, block path)
- [ ] **Security**: Invalidate sessions / Rotate keys (if compromised)
- [ ] **Performance**: Activate "Limp Mode" (disable non-essential services)

⚠️ No refactor. No optimization.

---

## 4. Evidence Preservation (MANDATORY)

- [ ] Capture logs and metrics
- [ ] Capture heap dumps / traces (for performance)
- [ ] Secure audit logs (for security)
- [ ] Isolate forensic data

---

## 5. Root Cause Hypothesis

- Initial hypothesis (clearly marked as hypothesis)
- What evidence supports this?
- What evidence is missing?

---

## 6. Ownership & Next Steps

- Incident Owner (Decision Authority): @[agent / human]
- Who investigates root cause?
- Is a full `{task-slug}.md` required after resolution?

---

## 7. Resolution Criteria

- What does “resolved” mean?
- Metrics / signals confirming stability
- Zero-trust validation (re-verify security after fix)

---

## 8. Post-Incident Follow-up (AFTER RESOLUTION)

- Root Cause Analysis document required?
- Preventive actions needed? (New lint rules, automated tests)
- Should this become a planned task?
