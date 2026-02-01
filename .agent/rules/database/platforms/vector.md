# Vector Database Rules
> Rules for Vector / Embedding Workloads

These rules apply when vector similarity search is used.

---

## Use Cases

Vector databases SHOULD be used for:

- Semantic search
- Recommendation systems
- AI embedding similarity

They MUST NOT replace relational truth.

---

## Integration Rules

You MUST:

- Keep vector data derived, not authoritative
- Maintain source-of-truth in relational tables
- Regenerate embeddings when source data changes

---

## Performance Rules

You SHOULD:

- Choose index type based on recall vs latency trade-offs
- Measure retrieval quality and latency explicitly

---

## Anti-Patterns

You MUST NOT:

- Store business-critical data only in vector form
- Use vector DBs as primary storage
