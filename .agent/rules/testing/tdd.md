# TDD (Test-Driven Development) Workflow

## 1. The Red-Green-Refactor Cycle

### 🔴 Phase 1: RED
- **Goal**: Write a failing test that defines a new piece of behavior.
- **Standard**: The test must fail for the right reason (assertion failure, not syntax error).
- **Mindset**: You are a designer defining the interface and behavior.

### 🟢 Phase 2: GREEN
- **Goal**: Write the absolute minimum code required to make the test pass.
- **Standard**: "Dirty" code is acceptable here. Just get it working.
- **Mindset**: You are a coder fixing a bug (the failing test).

### 🔵 Phase 3: REFACTOR
- **Goal**: Clean up the code while keeping the tests green.
- **Standard**: Improve naming, remove duplication, optimize structure.
- **Mindset**: You are an architect ensuring long-term health.

---

## 2. Key TDD Principles
- **Baby Steps**: Don't try to implement the whole feature at once. Test one small behavior at a time.
- **Listen to your tests**: If a piece of code is hard to test, it's usually poorly designed.
- **Regression Safety**: Once a test is green, it becomes a permanent safety net for all future refactors.

---

## 3. When NOT to TDD
- **Exploratory Programming (Spiking)**: When you don't know *how* to solve a problem yet. (Delete the spike and TDD the real solution later).
- **Trivial Code**: One-liners or purely declarative code with zero logic.
- **Third-party integrations**: Where the behavior is entirely controlled by an external system.
