# Task: Build Top-Down Shooter Prototype
Slug: build-topdown-shooter
Primary Agent: @game-developer
Status: APPROVED

---

## 1. Business Context

### Objective
Build a playable prototype to validate the core combat loop
and performance feasibility on mid-range mobile devices.

This prototype is not intended for monetization or launch.

### Business Rules
- Gameplay must be understandable within 30 seconds
- Session length: 3–5 minutes
- No multiplayer in this phase

### Terminology
- Player: controllable character
- Enemy: AI-controlled hostile unit
- Wave: group of enemies spawned together

---

## 2. Technical Flow

### Current State
- No existing game code
- Empty project scaffold

### Proposed Flow
1. Player spawns in arena
2. Enemies spawn in timed waves
3. Player shoots enemies
4. Score increments on enemy defeat
5. Game ends when player HP reaches zero

### Impact
- Single-module game loop
- No backend dependency
- No persistence required

---

## 3. Detailed Task Breakdown

- [ ] Select engine based on constraints → @game-developer
- [ ] Define core gameplay loop → @game-developer
- [ ] Implement player movement & shooting → @game-developer
- [ ] Implement enemy spawn & AI → @game-developer
- [ ] Add scoring and simple UI → @frontend-specialist
- [ ] Performance profiling on target device → @game-developer

---

## 4. Performance & Constraints

### Platform
- Mobile (Android mid-range)

### Performance Targets
- Minimum 60 FPS
- No frame spikes > 20 ms

### Constraints
- Touch input only
- Limited GPU budget

---

## 5. Risk & Mitigation

### Risks
- Performance degradation with enemy count
- Input latency on low-end devices

### Mitigation
- Object pooling for enemies and bullets
- Cap enemy count per wave

---

## 6. Acceptance Criteria

- Game loop playable end-to-end
- Stable 60 FPS on target device
- No crashes during 10-minute session

---

## 7. Rollout / Rollback

- Rollout: Internal prototype only
- Rollback: Delete branch, no production impact
