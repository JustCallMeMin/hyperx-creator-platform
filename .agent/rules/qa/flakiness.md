# Test Determinism & Flakiness Discipline

## Purpose

Maintain the highest standards of stability for automated test suites.

---

## The Determinism Rule

- **No Arbitrary Waits**: `sleep()` or fixed timeouts are strictly FORBIDDEN.
- **State-Based Assertions**: Always wait for a specific system state (e.g., `toBeVisible()`, `toReceiveMessage()`).
- **Environmental Factors**: Tests must handle slow network or high CPU load without failing (e.g., using extended timeouts for CI).

---

## Flakiness Hunting

- **Identification**: Any test that passes and then fails without a code change is considered "Flaky".
- **Zero Tolerance**: Flaky tests MUST be isolated (skipped/quarantined) immediately to preserve pipeline trust.
- **Root Cause Analysis (RCA)**: Always identify the timing or race condition causing the flake before re-enabling.

---

## STOP Gate

You MUST STOP if:

- The "Flaky test ratio" exceeds 5% of the total suite.
- Re-running a failed test suite "to see if it passes" is used as a standard practice.
