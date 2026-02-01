---
description: Debugging command. Activates DEBUG mode for systematic problem investigation.
---

---
description: Systematic debugging workflow for investigating errors, failures, and unexpected behavior.
---

## Invocation Rules

This workflow MAY be invoked only when:
- An issue, error, or unexpected behavior has been reported.
- A structured root cause analysis is required.
- The system is functioning but producing incorrect results.

This workflow MUST NOT be invoked when:
- The task is a new feature request or enhancement.
- A task-plan (`{task-slug}.md`) is being created or executed.
- The system is in a critical outage or breach (use Incident Response).

---

# /debug — Systematic Problem Investigation

$ARGUMENTS

---

## Purpose

Enable a structured and hypothesis-driven investigation of system issues.
This workflow focuses on understanding the root cause, applying a fix,
and preventing recurrence.

---

## Behavior

When `/debug` is triggered:

1. **Information Gathering**
   - Capture error messages or symptoms.
   - Record reproduction steps.
   - Clarify expected versus actual behavior.
   - Identify recent changes or deployments.

2. **Hypothesis Formation**
   - List possible causes.
   - Order hypotheses by likelihood and impact.

3. **Systematic Investigation**
   - Test hypotheses one by one.
   - Inspect logs, data flow, and state transitions.
   - Eliminate causes based on evidence.

4. **Fix and Prevention**
   - Apply the minimal correct fix.
   - Explain the root cause clearly.
   - Define concrete prevention measures.

---

## Output Expectations

The debug result MUST include:
- A clearly identified root cause.
- Evidence supporting the conclusion.
- A fix that addresses the cause, not just the symptom.
- At least one prevention action.

---

## Output Format

```markdown
## Debug: [Issue Title]

### Symptom
[Description of what is happening]

### Information Gathered
- Error message: `[error message]`
- Location: `[file path and line number if applicable]`
- Environment: `[runtime / platform]`

### Hypotheses
1. [Hypothesis] Most likely cause
2. [Hypothesis] Alternative cause
3. [Hypothesis] Less likely cause

### Investigation

**Testing hypothesis 1**
[What was checked] → [Result]

**Testing hypothesis 2**
[What was checked] → [Result]

### Root Cause
[Clear explanation of why the issue occurred]

### Fix
```[language]
// Before
[problematic code]

// After
[corrected code]
```

### Prevention
[Actions to prevent recurrence, such as tests, validation, monitoring]
```

---

## Usage Examples

```
/debug login not working
/debug API returns 500
/debug form submission fails
/debug data not persisted
```

---

## Key Principles

- Gather evidence before forming conclusions.
- Test hypotheses systematically.
- Fix the root cause, not the symptom.
- Always define a prevention strategy.
