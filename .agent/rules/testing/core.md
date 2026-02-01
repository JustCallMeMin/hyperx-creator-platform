# Testing Core Principles & Strategy

## 1. The Testing Mindset
- **Testing is Risk Management**: We don't test everything; we test where failure has the highest impact or likelihood.
- **Tests as Documentation**: A good test explains the system's behavior better than a comment.
- **Verification over Assumption**: Never assume a feature works; prove it with an automated, reproducible artifact.

---

## 2. The Testing Pyramid
Maintain a healthy balance of tests to ensure speed and reliability:

- **Unit Tests (Many)**: Focus on small, isolated logic. Fast, reliable, and precise.
- **Integration Tests (Some)**: Focus on the "glue" between components (API, DB, Services).
- **Component Tests (Few)**: Focus on UI behavior in isolation.
- **E2E Tests (Very Few)**: Focus on critical user journeys. High maintenance, but high confidence.

---

## 3. Heuristics for Test Selection
- If it's a **pure function** → Unit Test.
- If it's a **database query / repository** → Integration Test.
- If it's an **API endpoint** → Integration Test + Contract Test.
- If it's a **UI component** → Unit/Component Test with RTL.
- If it's a **critical business flow** (e.g., Checkout) → E2E Test.

---

## 4. Stability & Determinism
- **Flaky tests are worse than no tests**: They erode trust.
- **No arbitrary sleeps**: Use polling or event-based waiting.
- **Environmental Isolation**: Tests must run anywhere (local, CI) with the same result.

---

## 5. Definition of Done
A feature is done ONLY when:
1. Observable behavior is verified.
2. Edge cases are handled.
3. Tests are integrated into the CI pipeline.
4. No regressions are introduced.
