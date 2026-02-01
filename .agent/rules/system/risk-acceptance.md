# Explicit Risk Acceptance Policy

## Purpose

Define the mandatory protocol for identifying, evaluating, and explicitly accepting risks when project constraints (e.g., deadlines) conflict with safety or quality standards (Security, QA, Performance).

---

## 1. Risk vs. Ignorance

- **Risk Acceptance**: A conscious, documented business decision to proceed with a known vulnerability or technical debt under defined conditions.
- **Risk Ignorance**: Proceeding without identifying or documenting risks. This is a **Hard Protocol Failure**.

---

## 2. Risk Ownership (Logical Roles)

Agents MUST NOT accept risks on behalf of the business. Risk acceptance requires a **Risk Owner**:

| Context | Authorized Risk Owner |
| :--- | :--- |
| **Product / Feature** | `product-owner` |
| **System / Security** | `human (USER)` |
| **Deployment / Infra** | `devops-engineer` (for non-security risks) |

*Note: `security-auditor` and `qa-automation-engineer` can only recommend mitigation; they do not own the risk.*

---

## 3. Mandatory Acceptance Criteria

A risk can ONLY be accepted if the following exist:

1. **Explicit Risk Record**: Documented using `.agent/templates/risk-acceptance-record.md`.
2. **Fixed TTL (Time To Live)**: The risk acceptance expires after a defined period (e.g., 7 days, 1 sprint).
3. **Mandatory Follow-up Task**: A verified task entry in the backlog to resolve the debt/risk.
4. **Mitigation Plan**: Steps taken to minimize the impact of the accepted risk.

---

## 4. Hard Rules (Enforcement)

- **CRITICAL Security Risks**: CANNOT be accepted for production without immediate mitigation that lowers the severity below CRITICAL.
- **Silent Acceptance**: Orchestrators found "silently" accepting risk by ignoring agent reports will be moved to **Incident Mode**.
- **Human Escalation**: Any risk with **HIGH** impact on PII or financial data MUST be escalated to the `human (USER)`.

---

## 5. The Risk Acceptance Gate

During the **Verification Phase**, if an agent reports a `FAIL`:

1. Orchestrator triggers the **Risk Evaluation**.
2. Specialist Agent provides `Severity` and `Impact`.
3. Risk Owner provides `Reasoning` and `Acceptance`.
4. Orchestrator logs the `Risk Acceptance Record`.
5. Orchestrator creates the `Follow-up Task`.

Only then can the task be marked as `COMPLETE` or `SHIPPED`.
