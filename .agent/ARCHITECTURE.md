# Antigravity Kit Architecture

> Core Doctrine — Single Source of Truth

---

## 1. System Overview

Antigravity Kit is a modular AI capability expansion system designed around a strict separation of concerns between **Agents**, **Skills**, and **Workflows**.

This separation is foundational and non-negotiable.  
No component may assume responsibilities outside its designated layer.

### Core Layers

- **Agents** — Persistent decision-making authorities with domain responsibility
- **Skills** — Procedural and domain-specific execution modules
- **Workflows** — Deterministic slash-command orchestration flows

---

## 2. Directory Structure

```plaintext
.agent/
├── ARCHITECTURE.md          # Core doctrine (this file)
├── agents/                  # Agent definitions (.agent.md)
├── skills/                  # Skills (SKILL.md + resources)
├── workflows/               # Slash commands
├── rules/                   # Global behavioral rules
└── scripts/                 # System-level validation scripts
```

---

## 3. Authority Model (MANDATORY)

### 3.1 Authority Hierarchy

```text
User Intent
   ↓
Agent (decides)
   ↓
Skill (executes)
   ↓
Script / Reference (implements)
```

### 3.2 Authority Rules

- **Agents decide** what should be done and whether clarification is required
- **Skills execute** procedural logic and provide domain knowledge
- **Workflows orchestrate only**; they never decide intent or strategy
- **Scripts and references** are implementation details with no authority

No layer may override the authority of a higher layer.

---

## 4. Agent Layer (Primary Authority)

### 4.1 What an Agent Is

An **Agent** is a persistent decision authority, not a prompt bundle and not a workflow.

An agent:

- Interprets user intent
- Asks clarifying questions when required
- Selects appropriate skills
- Enforces quality gates and constraints

An agent does **not**:

- Contain procedural workflows
- Duplicate skill knowledge
- Hardcode technology choices
- Execute implementation steps directly

---

### 4.2 Agent Contract Standard

Every agent under `.agent/agents/` **MUST** follow this contract.

Required sections, in order:

```markdown
# Role

# Mission

# Core Principles

# Decision Rules

# Constraints

# Capability Scope

# Execution Model

# Review Checklist

# Quality Gates
```

This contract is mandatory to prevent agent drift and ensure composability.

---

### 4.3 Agent Lifecycle

Every agent operates according to the following lifecycle:

1. Interpret user intent
2. Determine whether clarification is required
3. Select eligible skills
4. Delegate execution to skills
5. Enforce review and quality gates
6. Produce final output

Agents may repeat or exit this cycle as needed.

---

### 4.4 Agent–Skill Relationship

Agents do **not** own skills.

Agents define **skill eligibility**, not skill binding.

Correct interpretation:

> “This agent MAY load the following skills when relevant.”

Agents decide:

- Whether a skill is needed
- When a skill should be loaded
- Whether multiple skills should be combined

---

## 5. Skill Layer (Procedural Execution)

### 5.1 What a Skill Is

A **Skill** is a modular, self-contained unit of procedural knowledge.

A skill:

- Encodes workflows, patterns, or domain rules
- May include scripts, references, and assets
- Is loaded on-demand via description matching

A skill does **not**:

- Decide intent
- Ask clarifying questions
- Enforce global policy
- Act as a persona

---

### 5.2 Skill Structure

```plaintext
skill-name/
├── SKILL.md           # Required metadata and instructions
├── scripts/           # Optional executable scripts
├── references/        # Optional documentation and schemas
└── assets/            # Optional templates or static files
```

### 5.3 Skill Loading Protocol

```text
User Request
   ↓
Agent Interpretation
   ↓
Skill Description Match
   ↓
Load SKILL.md
   ↓
Read references/ (if needed)
   ↓
Execute scripts/ (if needed)
```

Skills are loaded lazily to preserve context efficiency.

---

## 6. Workflow Layer (Deterministic Orchestration)

### 6.1 What a Workflow Is

A **Workflow** is a deterministic slash-command procedure.

A workflow:

- Executes a predefined sequence of steps
- Delegates decisions to agents
- Never interprets intent

A workflow does **not**:

- Choose skills
- Ask clarifying questions
- Override agent decisions

---

### 6.2 Workflow Authority Constraint

Workflows may call:

- Agents
- Skills (via agent direction)

Workflows may not:

- Replace agents
- Inject decision logic
- Modify agent constraints

---

## 7. Global Rules Layer

Global rules define system-wide constraints applied across all agents, skills, and workflows.

Examples:

- Security rules
- Safety policies
- Output formatting constraints
- Execution limits

Global rules override all other layers.

---

## 8. Validation & Quality Control

### 8.1 System Scripts

System-level scripts enforce correctness and readiness.

- `checklist.py` — Core checks during development
- `verify_all.py` — Full verification before release

### 8.2 Quality Philosophy

Antigravity prioritizes:

- Determinism over cleverness
- Explainability over opacity
- Safety over speed
- Long-term maintainability over short-term convenience

---

## 9. Scaling the System

### Adding an Agent

- Must follow the Agent Contract Standard
- Must not duplicate existing agent authority

### Adding a Skill

- Must remain procedural
- Must not encode persona or decision logic

### Adding a Workflow

- Must remain deterministic
- Must delegate decisions to agents

---

## 10. Final Doctrine

- **Agents decide**
- **Skills execute**
- **Workflows orchestrate**
- **Rules govern**

Violating these principles is considered a system design error.
