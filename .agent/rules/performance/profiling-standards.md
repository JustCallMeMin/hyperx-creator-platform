# Performance Profiling Standards

## Purpose

Define how to baseline and measure system performance.

---

## Measurement Tools

- **Backend**: `pprof` (Go), `py-spy` (Python), `clinic.js` (Node).
- **Database**: `EXPLAIN ANALYZE`, Slow query logs.
- **Frontend**: Chrome DevTools Performance tab, Lighthouse.

---

## Baseline Requirements

Before optimization, you MUST record:

- Average response time (P50)
- Tail latency (P95, P99)
- CPU and Memory consumption under load
- Throughput (Requests per second)

---

## Optimization Gate

You MUST NOT proceed if:

- No stable baseline exists.
- The optimization is "speculative" without profiling evidence.
