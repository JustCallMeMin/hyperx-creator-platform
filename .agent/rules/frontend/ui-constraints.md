---
trigger: always_on
---

# Frontend UI Constraints

## Purpose

This file defines **hard visual constraints**
to prevent cliché, low-effort, or trend-chasing UI.

These constraints are MANDATORY when loaded.

---

## Color Constraints

- Purple and violet hues are forbidden unless explicitly approved
- Overused gradients without semantic meaning are forbidden
- Color must encode meaning, not decoration

---

## Layout Constraints

Forbidden patterns:
- Bento-box layouts
- Overused card grids
- “Safe” center-aligned SaaS layouts
- Copy-paste landing page structures

---

## Component Usage Rules

- Components must serve content, not the other way around
- Do not introduce components without necessity
- Reusable components must justify abstraction

---

## Visual Noise Control

- No decorative icons without meaning
- No animations without purpose
- No effects added “because it looks cool”

If it does not improve clarity, remove it.

---

## Constraint Override Protocol

Overrides are allowed ONLY if:
- Business explicitly requires it
- Trade-offs are acknowledged
- Decision is documented in `{task-slug}.md`
