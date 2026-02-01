# Mobile Performance Rules

## Purpose

Performance on mobile is part of UX.

---

## Mandatory Targets

- UI interactions: 60 FPS
- Avoid frame drops > 16ms
- Startup time must be measured

Targets must be defined before optimization.

---

## Rendering Discipline

You MUST:

- Avoid unnecessary re-renders
- Use memoization sparingly
- Prefer native components over heavy JS logic

---

## Memory Rules

- Minimize object allocation during interaction
- Avoid large in-memory lists
- Release resources on background

GC pauses during gestures are forbidden.

---

## Profiling Gate

You MUST:

- Profile before optimizing
- Identify top bottlenecks only
- Re-profile after changes

Optimization without measurement is forbidden.
