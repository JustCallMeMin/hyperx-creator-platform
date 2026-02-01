# Neon Platform Rules
> Serverless PostgreSQL Rules

These rules apply ONLY when using Neon.

---

## Strengths

Neon SHOULD be used for:

- Serverless architectures
- Workloads with spiky traffic
- Cost-sensitive environments

---

## Operational Rules

You MUST:

- Be aware of cold-start latency
- Avoid chatty transaction patterns
- Prefer fewer, larger queries over many small ones

---

## Migration Rules

You SHOULD:

- Design migrations for branching environments
- Validate schema consistency across branches

---

## Anti-Patterns

You MUST NOT:

- Assume persistent warm connections
- Ignore connection pooling behavior
