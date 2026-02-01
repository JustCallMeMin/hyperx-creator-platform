# CI/CD Testing Integration

## Purpose

Standardize how automated tests are executed within the delivery pipeline.

---

## Pipeline Integration

- **Atomic Pipelines**: Tests should run in isolated environments (e.g., Docker) to ensure consistency.
- **Fail-Fast**: Smoke tests must run at the earliest stage possible to catch regressions quickly.
- **Parallelism**: Large suites MUST be parallelized to keep developer feedback loops short (< 10 mins).

---

## Observability & Artifacts

- **Trace Viewer**: Screenshots, videos, and network traces MUST be captured and stored on failure.
- **Reporting**: Test results should be aggregated and visible in the CI UI (e.g., GitHub Actions Summary).
- **History**: Maintain a history of test runs to identify trends in failure or performance degradation.

---

## STOP Gate

You MUST STOP if:

- Tests require manual steps to run in CI.
- CI pipeline results do not provide actionable failure logs.
