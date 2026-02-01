# Test Coverage Interpretation & Quality

## 1. Coverage is a Signal, Not a Goal
- **80% is high, 100% is usually a lie**: Don't chase 100% at the cost of testing trivial implementation details.
- **Target Critical Paths**: 100% coverage on core business logic (e.g., pricing, auth, data transformation) is mandatory.
- **Ignore the Boilerplate**: Don't waste time testing getters, setters, or library boilerplate.

---

## 2. Types of Coverage
- **Line Coverage**: Did the line run? (Weakest)
- **Branch Coverage**: Did every `if/else` and `switch` case run? (Better)
- **State Coverage**: Did the system enter every possible valid state? (Strongest)

---

## 3. The "Ghost Coverage" Trap
- Tests can pass through code without actually *verifying* anything. 
- High coverage with low-quality assertions is dangerous.

---

## 4. Improving Quality over Quantity
- **Mutation Testing**: (Advanced) Use tools like Stryker to see if your tests catch intentional bugs introduced into the code.
- **Failure-Focused Coverage**: Ensure every `catch` block and error condition has a corresponding test case.

---

## 5. Review Standards
- Never approve a PR that decreases coverage of critical modules without a clear technical justification.
- Coverage reports should be part of the PR feedback loop.
