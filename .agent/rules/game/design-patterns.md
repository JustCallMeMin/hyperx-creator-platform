---
trigger: always_on
---

# Game Design Patterns Rules

## Purpose

This file defines **approved structural patterns**
for scalable and maintainable game systems.

Patterns exist to serve gameplay, not architecture purity.

---

## State Management

Use **State Machines** when:

- Game flow has discrete phases
- Player or AI behavior depends on state

Avoid state machines if:

- State transitions are highly dynamic
- Data-driven systems are more suitable

---

## ECS (Entity Component System)

Use ECS when:

- Many entities share behavior
- Performance and cache locality matter
- Systems operate over large entity sets

Avoid ECS if:

- Game logic is simple
- Team lacks ECS experience

---

## Object Pooling

Pooling is MANDATORY when:

- Objects are created/destroyed frequently
- Bullets, particles, enemies, effects

Avoid pooling for:

- Rarely created objects
- Long-lived unique entities

---

## Pattern Abuse Warning

You MUST NOT:

- Introduce patterns without necessity
- Stack multiple paradigms unnecessarily
- Over-abstract early gameplay code

Simplicity beats pattern purity.
