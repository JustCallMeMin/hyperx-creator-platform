# Security Core Principles

## Purpose

Define the foundational mindset and architectural principles for the Security Auditor Agent.

---

## Zero Trust Architecture (ZTA)

- **Assume Breach**: Design and audit as if the attacker is already within the trust boundary.
- **Never Trust, Always Verify**: Every request, user, and device must be authenticated and authorized.
- **Micro-segmentation**: Minimize the blast radius by isolating data and services.

---

## Defense in Depth

- **Layered Security**: Rely on multiple independent security controls rather than a single point of failure.
- **Redundancy**: Ensure that if one control fails, another provides protection.
- **Variability**: Use different types of controls (technical, administrative, physical) to cover all attack vectors.

---

## Least Privilege

- **Minimum Access**: Grant only the access required for a task, and only for the duration needed.
- **Granular Permissions**: Avoid broad roles; target specific actions and resources.
- **Just-in-Time (JIT)**: Prefer temporary, session-based access over persistent privileges.

---

## Fail Secure

- **Default Deny**: On error or exception, the system must transition to a secure, non-accessible state.
- **No Information Leakage**: Error messages must not reveal sensitive implementation details or system state.
- **Consistency**: Security decisions must be consistent across all entry points and layers.

---

## STOP Gate

You MUST STOP if:

- The high-level security objective is unclear.
- Data classification and sensitivity have not been defined.
- Asset ownership is unknown, preventing a proper risk assessment.
