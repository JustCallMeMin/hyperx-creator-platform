# Mocking & Test Doubles Strategy

## 1. Terminology
- **Dummy**: Objects passed around but never used.
- **Stub**: Provides canned answers to calls made during the test.
- **Spy**: Records information about how it was called.
- **Mock**: Objects pre-programmed with expectations of how they should be called.
- **Fake**: Working implementation with a shortcut (e.g., In-memory DB).

---

## 2. The Golden Rule
**"Mock only what you don't own."**
- Mock: External APIs, Third-party SDKs, File systems, Network.
- DO NOT Mock: Your own domain logic, small internal utilities.

---

## 3. Mocking Boundaries
- **Mock at the Boundary**: Mock the interface that talks to the external world, not the logic that processes the data.
- **Contract Verification**: When you mock an external service, ensure the mock remains in sync with the actual service contract.

---

## 4. Risks of Over-Mocking
- **Brittle Tests**: Tests that break every time the internal implementation changes.
- **False Confidence**: Tests pass but the system fails because the mock behavior doesn't match reality.

---

## 5. Mock Hygiene
- **Reset after each test**: Use `vi.clearAllMocks()` or equivalent.
- **Avoid logic in mocks**: Mocks should be as simple as possible.
- **Prefer Fakes for Databases**: Using an in-memory database is often more reliable than mocking every SQL call.
