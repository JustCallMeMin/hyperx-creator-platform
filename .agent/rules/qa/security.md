# Negative Testing & Chaos Rules

## Purpose

Enforce testing of failure modes, edge cases, and unexpected system behavior.

---

## The "Unhappy Path"

Developers test the "Happy Path". QA MUST test the chaos:

| Scenario | Automation Strategy |
| :--- | :--- |
| **Slow Network** | Throttle network (3G/Poor connection simulation). |
| **Server Failures** | Mock 4xx/5xx responses mid-flow for resilience validation. |
| **Race Conditions** | Extreme parallelism or high-frequency "double clicks". |
| **Auth Transitions** | Expire tokens or change permissions during active sessions. |
| **Data Corruption** | Inject invalid payloads into API bodies and input fields. |

---

## Negative Validation

- **Error Catching**: Verify that the system fails gracefully with friendly errors, not stack traces.
- **Security Scans**: Integrate automated tools (e.g., OWASP ZAP) into the E2E flow where possible.
- **Resource Limits**: Test system behavior under memory or CPU pressure.

---

## STOP Gate

You MUST STOP if:

- Only "Happy Path" scenarios are covered.
- System error messages leak sensitive implementation details.
