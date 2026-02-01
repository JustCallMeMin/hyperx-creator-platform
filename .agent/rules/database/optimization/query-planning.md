# Query Planning Rules
> Execution Plan Optimization

---

## Mandatory Practice

You MUST:

- Use `EXPLAIN ANALYZE`
- Inspect nested loops and join strategies
- Optimize for total execution time

---

## Planner Awareness

You SHOULD:

- Understand planner cost estimates
- Avoid planner fences unless intentional
- Rewrite queries when planner choices are suboptimal

---

## Anti-Patterns

You MUST NOT:

- Optimize queries without observing plans
- Rely on ORM-generated queries blindly
