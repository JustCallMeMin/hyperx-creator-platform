---
trigger: always_on
---

# Platform Constraints Rules

## Purpose

Each platform imposes **hard constraints**
that shape game design and architecture.

Ignoring them leads to failure.

---

## Mobile

Constraints:

- Limited CPU/GPU
- Battery consumption
- Touch input

Rules:

- Simplify shaders
- Reduce draw calls
- Avoid heavy physics

---

## Web

Constraints:

- Browser performance
- Memory limits
- Input variability

Rules:

- Minimize bundle size
- Avoid heavy native dependencies
- Expect inconsistent performance

---

## PC / Console

Constraints:

- Certification requirements
- Hardware diversity (PC)
- Input devices

Rules:

- Abstract input systems
- Support scalable settings
- Handle platform-specific APIs carefully

---

## VR

Constraints:

- Very high FPS requirement
- Motion sickness risk
- Input precision

Rules:

- Maintain stable frame time
- Avoid sudden camera movement
- Optimize aggressively

---

## Platform STOP Gate

You MUST STOP if:

- Platform target is undefined
- Constraints conflict with gameplay goals
