# Cryptography & Secret Management Rules

## Purpose

Standardize the application of cryptographic controls and the management of sensitive secrets.

---

## Approved Algorithms

- **Hashing**: Argon2id (for passwords), SHA-256 or SHA-3 (for integrity).
- **Encryption (At Rest)**: AES-256-GCM (Authenticated Encryption).
- **Encryption (In Transit)**: TLS 1.3 (minimum TLS 1.2 with secure cipher suites).
- **Signing**: Ed25519 or RSA-4096 (minimum 3072 bits).

---

## Secret Management

- **Centralized Vault**: Use dedicated secret managers (e.g., AWS Secrets Manager, Vault, Doppler).
- **No Hardcoding**: Secrets MUST NEVER be stored in source code, configuration files, or build artifacts.
- **Rotation**: Critical secrets (API keys, DB passwords) MUST have a defined rotation policy.

---

## Key Management

- **Separation of Duties**: Ensure different keys are used for different environments (Dev, Staging, Prod).
- **Access Control**: Limit key access to the minimum set of users and services required.
- **Auditing**: Log all key usage, rotation, and access modifications.

---

## STOP Gate

You MUST STOP if:

- Hardcoded secrets are discovered in the repository or plan files.
- Deprecated algorithms (e.g., DES, MD5, SHA1 for security) are in use.
- Encryption keys are stored alongside encrypted data.
