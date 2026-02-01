# Mobile Navigation Rules

## Purpose

Navigation defines user mental models.

---

## Structure Rules

- Keep navigation depth shallow
- Avoid hidden critical screens
- Navigation state must be predictable

---

## Pattern Selection

- Tabs → peer-level destinations
- Stack → drill-down flows
- Modals → temporary tasks only

---

## Deep Linking

If supported, you MUST:

- Define entry points clearly
- Handle invalid states safely
- Restore navigation state correctly

---

## STOP Gate

You MUST STOP if:

- Navigation hierarchy is unclear
- Deep link behavior is undefined
