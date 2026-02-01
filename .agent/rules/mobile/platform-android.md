# Android Platform Constraints

## Purpose

Handle Android fragmentation safely.

---

## Device Diversity

- Expect varied screen sizes
- Expect OEM differences
- Test across API levels

---

## Lifecycle Discipline

You MUST handle:

- Activity recreation
- Process death
- Permission revocation

---

## Performance & Stability

- Avoid main-thread blocking
- Handle low-memory callbacks
- Respect background execution limits

---

## STOP Gate

You MUST STOP if:

- Minimum SDK is unknown
- Lifecycle edge cases are ignored
