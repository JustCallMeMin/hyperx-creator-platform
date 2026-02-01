# OWASP Top 10:2025 Standard

## Purpose

Define the primary web and API risk categories and audit focal points based on OWASP 2025.

---

## Risk Categories & Audit Focus

| Category | Description | Primary Audit Goal |
| :--- | :--- | :--- |
| **A01: Broken Access Control** | Authorization gaps, IDOR, SSRF, bypasses. | Verify that permissions are enforced at the resource level. |
| **A02: Security Misconfiguration** | Cloud configs, headers, defaults, unnecessary services. | Ensure hardened configurations and secure communication. |
| **A03: Software Supply Chain** | Dependencies, CI/CD, integrity failures. | Audit lock files, dependency sources, and CI pipeline security. |
| **A04: Cryptographic Failures** | Weak crypto, exposed secrets, certificate issues. | Ensure strong algorithms and secure secret management. |
| **A05: Injection** | SQL, command, XSS, template injection. | Verify parameterized queries and proper output encoding. |
| **A06: Insecure Design** | Architecture flaws, logic vulnerabilities. | Evaluate the system design against threat models. |
| **A07: Authentication Failures** | Sessions, MFA, password handling. | Ensure secure session management and strong authentication. |
| **A08: Integrity Failures** | Tampered data, unsigned updates. | Verify data integrity with hashes and digital signatures. |
| **A09: Logging & Alerting** | Visibility gaps, silent failures. | Ensure critical security events are logged and monitored. |
| **A10: Exceptional Conditions** | Error handling leaks, fail-open states. | Verify that errors do not leak data or bypass controls. |

---

## STOP Gate

You MUST STOP if:

- The audit identifies a Critical OWASP vulnerability with a viable exploit path.
- The system lacks basic logging for security-sensitive events.
