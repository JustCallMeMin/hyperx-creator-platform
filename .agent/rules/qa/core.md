# QA Core Principles

## Purpose

Define the foundational mindset and standards for high-value quality assurance.

---

## Risk-Based Testing

- **Value over Count**: Success is measured by the severity of rủi ro (risk) mitigated, not the number of test cases.
- **Critical Path First**: Prioritize testing the core business value (e.g., checkout, login, data integrity).
- **Proactive Discovery**: QA exists to find problems before they reach production, not just to verify tasks.

---

## The "Safety System" Mindset

- **Automation is a requirement**: If a critical path isn't automated, it is considered unsafe.
- **Trust is earned**: A passing test suite must be a reliable indicator of system health.
- **Verification over Opinion**: Every quality claim must be backed by reproducible evidence (logs, traces, screenshots).

---

## STOP Gate

You MUST STOP if:

- The primary risk being tested is not identified.
- The environment is too unstable to provide deterministic results.
- Test data management strategy is missing.
