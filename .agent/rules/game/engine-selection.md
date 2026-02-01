---
trigger: always_on
---

# Game Engine Selection Rules

## Purpose

This document defines a **decision framework** for selecting a game engine.
The goal is to minimize long-term risk, not maximize short-term comfort.

Engine choice is a **structural decision** and MUST be justified.

---

## Mandatory Pre-Conditions

Before selecting an engine, you MUST know:

- Target platform(s)
- Game scope and genre
- Team size and experience
- Performance constraints
- Tooling and pipeline needs

If any are unknown → STOP.

---

## Decision Tree

### Choose **Unity** if:

- Target: Mobile / Cross-platform
- Team: Small to mid-size
- Gameplay: 2D / 3D casual, simulation
- Need fast iteration and asset ecosystem

Avoid Unity if:

- Heavy multiplayer backend required
- Engine-level customization is critical

---

### Choose **Godot** if:

- Target: Indie, experimental, open-source friendly
- Team: Small, engineering-heavy
- Need full engine control
- 2D-first or lightweight 3D

Avoid Godot if:

- Large-scale AAA visuals required
- Heavy console certification constraints

---

### Choose **Unreal Engine** if:

- Target: High-end PC / Console
- Gameplay: Real-time 3D, FPS, open world
- Team: Medium to large
- Rendering fidelity is a priority

Avoid Unreal if:

- Team lacks C++ expertise
- Scope is small or casual

---

## Forbidden Selection Criteria

You MUST NOT select an engine based on:

- Popularity
- Personal preference
- Trend or hype
- Portfolio aesthetics

---

## Documentation Requirement

The selected engine MUST be recorded in:

- `{task-slug}.md`
- With rationale and rejected alternatives
