# Testing Strategy Rules

## Purpose

Define the structure and triggers for different levels of automated testing.

---

## The Testing Tiers

### 1. Smoke Suite (P0)

- **Goal**: Rapid verification of core functionality (< 2 mins).
- **Triggers**: Every commit, initial dev environment setup.
- **Coverage**: Login, Critical Path, Basic connectivity.

### 2. Regression Suite (P1)

- **Goal**: Deep functional coverage and boundary testing.
- **Triggers**: Pre-merge to `main`, Nightly builds.
- **Coverage**: All user stories, complex edge cases, cross-browser/platform checks.

### 3. Visual Regression

- **Goal**: Detect unintended UI changes.
- **Tools**: Pixelmatch, Percy, or native Playwright snapshots.
- **Rule**: Baseline snapshots must be updated intentionally; any drift is a failure.

---

## Strategy Gates

- **Necessity Gate**: Do not write an E2E test if the same risk can be covered by a Unit or Integration test.
- **Coverage Gate**: Every PR affecting core logic MUST have an associated test update or verification in the suite.

---

## STOP Gate

You MUST STOP if:

- The test suite execution time exceeds 5 minutes without a parallelism plan.
- There is no distinction between "Smoke" and "Full" regression.
