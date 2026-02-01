# Workflow Design Guidelines

> Workflows are the **interface** for the user. Skills are the **brain** of the agent.

## 1. Structural Requirements

Every workflow file MUST contain:

- **YAML Frontmatter**: `description` field.
- **Invocation Rules**: `MAY` and `MUST NOT` conditions.
- **$ARGUMENTS**: Required placeholder for user input.
- **Output Format**: Clear Markdown templates for consistent UI.

## 2. Logic Boundary (The "Mask" Rule)

- **Workflows = Interaction**: Define *how* to talk and *what* to show.
- **Skills = Execution**: Define *how* to think and *how* to solve.
- **Prohibited**: Avoid embedding complex reasoning, agent selection logic, or multi-step algorithms inside workflows. Delegate these to specialist Skills.

## 3. Communication Style

- **Efficiency**: Minimize repetitive questioning. Rely on the `brainstorming` skill for initial discovery.
- **Presentation**: Use structured Markdown (tables, lists, code blocks).
- **Humanization**: Follow `GEMINI.md` Tier 0 (Professional, neutral tone).

## 4. Maintenance

- If a workflow starts "thinking too much", refactor the logic into a Skill.
