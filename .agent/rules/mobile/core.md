# Mobile Core Principles

## Purpose

Define foundational principles that apply to ALL mobile work.

---

## Mobile-First Reality

- Mobile devices are resource-constrained
- Network is unreliable by default
- Interruptions (calls, backgrounding) are normal

Design MUST assume:

- App can be paused or killed at any time
- Connectivity can drop mid-action

---

## UX Fundamentals

- Touch-first interaction
- No hover-dependent logic
- Feedback must be immediate
- Gestures must be discoverable or optional

---

## Lifecycle Awareness

You MUST handle:

- App background / foreground
- Screen rotation
- OS reclaiming memory

Ignoring lifecycle is forbidden.

---

## STOP Gate

You MUST STOP if:

- Target devices are unknown
- Offline behavior is undefined
- App lifecycle is not considered
