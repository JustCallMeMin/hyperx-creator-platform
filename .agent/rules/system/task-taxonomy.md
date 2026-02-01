# Task Type Taxonomy

## Purpose

Standardize the classification of tasks across all agents to ensure consistent handling and routing.

---

## Core Taxonomy (Downstream Routing)

| Task Type | Description | Primary Agent |
| :--- | :--- | :--- |
| **analysis** | Research, codebase exploration, or logic mapping. | Explorer / Archaeologist |
| **generation** | Creating new features, code, or assets. | Developer / Specialist |
| **transformation** | Refactoring, language translation, or format conversion. | Developer / Specialist |
| **evaluation** | Security audit, performance profiling, or code review. | Auditor / Optimizer |
| **decision-support** | Architecture trade-offs, risk assessment, or planning. | Orchestrator / Planner |
| **exploration** | Brainstorming, discovery, or requirement elicitation. | Product Owner / Analyst |

---

## Domain-Specific Sub-types

### Development (Dev)

- `feature-implementation`: New functionality.
- `bug-fix`: Correcting erroneous behavior.
- `refactor`: Structural improvement without changing behavior.
- `unit-test`: Creating low-level logic verifications.

### Product (Prod)

- `discovery`: Validating problem space.
- `mapping`: Defining user flows or journeys.
- `refinement`: Breaking down epics into stories.
- `prioritization`: Ranking tasks by value/risk.

### Infrastructure (Infra)

- `provisioning`: Resource setup (IaC).
- `hardening`: Security configuration application.
- `automation`: CI/CD pipeline modifications.
- `monitoring`: Alerting and observability setup.

---

## Normalization Rule

- A task MUST mapped to exactly ONE core taxonomy type.
- Multiple domains require `decision-support` first.
