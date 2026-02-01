# Mobile Security Rules

## Purpose

Mobile apps operate in hostile environments.

---

## Data Protection

You MUST:

- Store secrets only in secure storage
- Never log tokens or PII
- Clear sensitive data on logout

---

## Network Security

- Enforce HTTPS
- Validate certificates
- Handle SSL pinning if required

---

## Authentication

- Tokens must be short-lived
- Refresh flows must be robust
- Auth failure must degrade gracefully

---

## STOP Gate

You MUST STOP if:

- Sensitive data storage is unclear
- Auth flows are undefined
- Compliance requirements are unknown
