---
trigger: always_on
---

# Multiplayer Architecture Rules

## Purpose

Multiplayer decisions are **irreversible late-game**.
Architecture MUST be chosen carefully and early.

---

## Multiplayer Mode Classification

### Real-Time Multiplayer

Use when:

- Competitive or action gameplay
- Tight synchronization required

Costs:

- High latency sensitivity
- Complex state reconciliation

---

### Turn-Based Multiplayer

Use when:

- Strategy, async gameplay
- Latency tolerance is acceptable

Benefits:

- Simpler architecture
- Easier scaling

---

## Authority Model

You MUST choose:

- Server-authoritative (preferred)
- Client-authoritative (rare, risky)

Client-authoritative systems MUST justify cheating risks.

---

## Synchronization Rules

- Minimize state sync
- Prefer events over full state
- Handle desync explicitly

---

## Multiplayer STOP Gate

You MUST STOP if:

- Authority model is unclear
- Cheating risk is not addressed
- Latency budget is unknown
