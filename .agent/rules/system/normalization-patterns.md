# Normalization Patterns & Ambiguity Handling

## Purpose

Define patterns for the Prompt Normalizer to handle unclear human input and ensure high confidence before execution.

---

## 1. Ambiguity Detection Patterns

### Short-Circuit Triggers (Confidence = Low)

- **Single-word prompts**: e.g., "commit", "test", "deploy".
- **Pronoun-heavy prompts**: e.g., "fix it", "change that file".
- **Ambiguous verbs**: e.g., "make it better", "handle the thing".

### Conflict Triggers (Confidence = Medium/Low)

- **Tool/Goal mismatch**: e.g., "use git to fix the UI color" (Git is for versioning, not styling).
- **Overloaded intent**: e.g., "refactor and add a new feature" (Violates atomicity).

---

## 2. Clarification Loop Policy (STOP Rule)

If `Confidence = Low`, the Normalizer MUST:

1. **Invoke the STOP command**.
2. **Present the Ambiguity List**:
   - What is unclear?
   - What are the potential guesses? (limited to 2-3).
3. **Request specific missing pieces**:
   - "Please specify the [Target File/Domain/System]."
   - "Confirm if the goal is [A] or [B]."

### Forbidden Actions

- **Silent Correction**: Guessing the target file or intent without confirmation.
- **Defaulting to "Latest"**: Assuming the user means the most recently modified file.

---

## 3. Ambiguity Test Cases (Verification)

| Input | Expected Normalization Result | Required Clarification |
| :--- | :--- | :--- |
| "fix" | Task Type: `transformation` | Missing: Intent details, target file, bug description. |
| "optimize this" | Task Type: `evaluation` | Missing: Context (what is 'this'?), Metric (speed? memory?). |
| "add auth" | Task Type: `generation` | Missing: Scope (which system?), Constraints (JWT? OAuth?). |
| "check security" | Task Type: `evaluation` | Missing: Domain (Cloud? Code? Network?). |

---

## 4. Confidence Score Metric

- **High (1.0)**: Intent, Task Type, Context, and Constraints are all explicit.
- **Medium (0.5 - 0.9)**: Intent and Context are clear, but some constraints or requirements are implicit.
- **Low ( < 0.5)**: Missing target, unclear intent, or conflicting domain.

Normalization result MUST include a confidence score.
If score < 0.7, a **Caution** note must be appended.
If score < 0.4, a **STOP** is mandatory.
