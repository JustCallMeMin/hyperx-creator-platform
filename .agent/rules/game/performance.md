---
trigger: always_on
---

# Game Performance Rules

## Purpose

Performance is a **first-class feature**.
A game that does not hit its performance target is incomplete.

---

## Mandatory Performance Targets

You MUST define FPS targets per platform:

- Mobile: 30–60 FPS
- PC: 60 FPS minimum
- Competitive / Action: 120 FPS if applicable
- VR: Platform-specific (usually ≥90 FPS)

No code before targets are defined.

---

## Frame Budget Discipline

You MUST:

- Define frame budget (ms per frame)
- Allocate budget per system (render, logic, physics, IO)

Blind optimization is forbidden.

---

## Profiling Rules

You MUST:

- Profile before optimizing
- Use engine-native profilers
- Identify top 3 bottlenecks only

Optimizing beyond measured bottlenecks is forbidden.

---

## Memory Rules

- Object allocation per frame must be minimized
- Pooling is mandatory for frequently spawned objects
- GC spikes must be avoided in gameplay loops

---

## Performance STOP Gate

You MUST STOP and re-evaluate if:

- FPS drops below target consistently
- Frame spikes exceed budget
- Performance regressions appear after changes
