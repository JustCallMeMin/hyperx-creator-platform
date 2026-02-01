# E2E Automation Standards

## Purpose

Standardize browser and user-flow automation patterns for reliability and maintainability.

---

## Tooling Standards

- **Playwright (Preferred)**: Use for multi-tab, parallel execution, and trace/video capture.
- **Cypress**: Use for high-quality component and frontend integration tests.
- **Selectors**: Use user-facing attributes (role, text, label) first. NEVER use unstable CSS classes or XPath.

---

## Page Object Model (POM)

- **Abstraction**: Test files MUST NOT contain raw selectors. All interactions are abstracted into Page Classes.
- **Methods**: Page methods should represent user actions (e.g., `loginPage.submit()`), not direct clicks.
- **Readability**: A test script should read like a user story, not a sequence of DOM manipulations.

---

## Data Management

- **Isolation**: Each test MUST be responsible for its own data lifecycle (Create → Use → Cleanup/Reset).
- **Concurrency-Safe**: Tests running in parallel must not share global or static data that could cause race conditions.

---

## STOP Gate

You MUST STOP if:

- Tests rely on hard-coded IDs or volatile UI selectors.
- Data cleanup is not automated, leading to "dirty" environments.
