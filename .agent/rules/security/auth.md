# Authentication & Authorization Rules

## Purpose

Define standards for identity management, session handling, and access control.

---

## Authentication Standards

- **Secure Storage**: Passwords MUST be hashed with strong, work-factor algorithms (e.g., Argon2, bcrypt).
- **MFA Required**: Multi-factor authentication is mandatory for all administrative and high-risk actions.
- **Fail-Safe Response**: Use generic error messages (e.g., "Invalid username or password") to prevent enumeration.

---

## Session Management

- **Secure Attributes**: Cookies MUST use `HttpOnly`, `Secure`, and `SameSite=Lax/Strict`.
- **Invalidation**: Sessions MUST be invalidated on logout and after a reasonable inactivity period.
- **Complexity**: Session tokens MUST be cryptographically strong, non-predictable, and sufficiently long.

---

## Authorization Models

- **RBAC / ABAC**: Use Role-Based or Attribute-Based Access Control to manage permissions.
- **Failure Mode**: Any failure in the authorization check MUST result in access being denied.
- **IDOR Prevention**: Always verify that the current user has ownership or permission to access specific record IDs.

---

## STOP Gate

You MUST STOP if:

- Secrets or passwords are stored in plain text or using weak/deprecated hashes (MD5, SHA1).
- Authentication bypass vulnerabilities are discovered in the core flow.
