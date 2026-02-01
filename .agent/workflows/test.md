---
description: Test generation and test running command. Creates and executes tests for code.
---

---
description: Test generation and execution workflow for unit, integration, and E2E validation.
---

## Invocation Rules

This workflow MAY be invoked only when:
- Generating tests for a file, module, or feature.
- Executing existing test suites.
- Reviewing test coverage.

This workflow MUST NOT be invoked when:
- Writing production logic without test context.
- Performing UI or design prototyping.

---

# /test — Test Generation and Execution

$ARGUMENTS

---

## Purpose

Generate and execute tests to validate system behavior.
This workflow supports quality assurance but does not replace planning or debugging workflows.

---

## Supported Commands

```
/test
/test [file-or-feature]
/test coverage
/test watch
```

---

## Behavior

### Test Generation

When generating tests:

1. **Code Analysis**
   - Identify public behaviors.
   - Detect dependencies and boundaries.
   - Determine required mocks or fakes.

2. **Test Design**
   - Happy path cases
   - Error and edge cases
   - Integration paths if applicable

3. **Test Creation**
   - Follow project testing conventions.
   - Use the existing test framework.
   - Avoid testing implementation details.

---

### Test Execution

When running tests:
- Execute the relevant test suite.
- Report pass/fail summary.
- Surface failing assertions clearly.

---

## Output Expectations

Depending on command, output MUST include:
- Generated test file path (if applicable)
- Test execution summary
- Coverage report (if requested)

---

## Key Principles

- Test observable behavior, not implementation.
- Prefer clarity over coverage inflation.
- Use deterministic tests.
- Mock only external dependencies.
