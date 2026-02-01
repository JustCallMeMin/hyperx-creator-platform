---
trigger: always_on
---

# Event-Driven Architecture (EDA) Rules

> Contextual Architecture Rules for Event-Driven Systems

These rules apply when the system follows Event-Driven Architecture principles.

---

## 1. Event Fundamentals (MANDATORY)

You MUST treat events as:

- Immutable facts
- Historical records
- Statements of something that already happened

Events MUST NOT:

- Contain business commands
- Be used to request actions
- Be mutated after publication

---

## 2. Command vs Event Separation

You MUST:

- Separate **command handling** from **event handling**
- Validate business rules at command time
- Emit events only after successful state transitions

Commands express intent.  
Events express facts.

---

## 3. Immutability & Schema Discipline

You MUST:

- Treat event payloads as immutable
- Use versioned schemas when evolution is required
- Avoid breaking changes to existing event contracts

Once an event is published, it is part of history forever.

---

## 4. Schema Evolution & Versioning

You MUST handle schema changes gracefully:

- **Backward Compatibility**: New fields must be optional; existing fields must not be removed or renamed.
- **Versioning**: Use versioning in the event type name (e.g., `user.created.v2`) or a `version` field in the metadata.
- **Dual Writing/N-1 Support**: Support the previous version of a schema during deployments to avoid downtime.
- **Registries**: Use a Schema Registry when using binary formats like Protobuf or Avro.

---

## 4. Idempotency (MANDATORY)

You MUST assume:

- At-least-once delivery
- Duplicate events
- Out-of-order delivery

Therefore:

- Event handlers MUST be idempotent
- Side effects MUST be safe to repeat
- Deduplication strategies MUST be explicit

Non-idempotent consumers are production hazards.

---

## 5. Decoupling Rules

You MUST:

- Avoid coupling domain logic directly to brokers
- Keep messaging infrastructure at the boundary
- Prevent event consumers from knowing producers

Systems communicate via contracts, not shared code.

---

## 6. Error Handling & Retries

You MUST:

- Distinguish between transient and permanent failures
- Use retries for transient failures
- Avoid infinite retry loops without backoff
- Provide dead-letter handling strategies

Silent event loss is unacceptable.

---

## 7. Consistency & Transactions

You MUST:

- Avoid distributed transactions
- Accept eventual consistency
- Model compensating actions explicitly
- Never assume synchronous guarantees from events

Consistency is a design choice, not a default.

---

## 8. Observability

You SHOULD:

- Log event processing outcomes
- Correlate events via identifiers
- Expose metrics for failures and retries

Unobservable systems are unmaintainable systems.

---

## 9. Consumer Reliability Patterns

- **Dead Letter Box**: Move permanently failed events to a separate queue for manual inspection.
- **Compensating Actions**: Model "undo" logic for failures in eventual consistency flows.
- **Strict Ordering**: If business logic requires ordering, use partition keys (sharding) correctly at the broker level.
- **Poison Pill Handling**: Ensure one bad event doesn't stall the entire consumer pipeline.

---

## 9. Anti-Patterns (Forbidden)

You MUST NOT:

- Use events as RPC
- Encode workflow logic implicitly across events
- Hide side effects inside event constructors
- Depend on event ordering for correctness
- Mutate state during event publication

---

## Final Rule

Event-Driven systems trade immediacy for resilience and scalability.

Design for failure, duplication, and delay by default.
