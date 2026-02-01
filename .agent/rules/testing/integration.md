# Integration Testing Standards

## 1. Purpose
Integration tests verify that different modules or services work correctly together. They catch "boundary bugs" where the contract between two systems is misunderstood.

---

## 2. Boundaries & Scope
- **Database Integration**: Verify that SQL queries, migrations, and repositories behave correctly with a real (but isolated) DB.
- **API Integration**: Verify that endpoints return the correct status codes and payloads.
- **Service Interaction**: Verify that Service A can correctly call Service B and handle its responses.

---

## 3. Environment Management
- **Isolation**: Use a fresh database schema or container for each test run.
- **Cleanup**: Ensure tests leave the environment in a clean state (using transactions or `beforeEach`/`afterEach` cleanup).
- **External APIs**: Use tools like **MSW (Mock Service Worker)** or **Nock** to simulate real network responses without actually hitting the internet.

---

## 4. Data Setup
- **Factories over Fixtures**: Use factory functions to create test data dynamically rather than relying on static, fragile JSON files.
- **Minimal State**: Only set up the data required for the specific test case.

---

## 5. Performance
- Integration tests are slower than unit tests.
- **Parallelize**: Run integration tests in parallel whenever the underlying infrastructure supports it.
- **Selective Execution**: Only run relevant integration tests during development; run the full suite in CI.
