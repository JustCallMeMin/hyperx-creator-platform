# Software Supply Chain Security Rules

## Purpose

Mitigate risks associated with third-party dependencies, build processes, and CI/CD pipelines.

---

## Dependency Governance

- **Lockfiles**: Every project MUST use a lockfile (`package-lock.json`, `go.sum`) to ensure reproducible and signed builds.
- **Vulnerability Scanning**: Automated scanning (e.g., `npm audit`, `trivy`) MUST be performed in the CI pipeline.
- **Provenience**: Prefer official or verified repositories for all dependencies.

---

## CI/CD Security

- **Isolated Builds**: Use clean, ephemeral build environments for every pipeline run.
- **Signed Artifacts**: Critical build artifacts MUST be signed to prevent tampering during delivery.
- **Secrets in CI**: Use CI-native secret management; never expose secrets in build logs.

---

## SBOM (Software Bill of Materials)

- **Inventory**: Maintain a clear inventory of all third-party components and their versions.
- **Visibility**: Periodically audit the SBOM against known CVE databases.

---

## STOP Gate

You MUST STOP if:

- A high-severity vulnerability exists in a core dependency without a plan for mitigation.
- The CI pipeline environment has broad, unmonitored access to production secrets.
- Dependency lockfiles are missing or inconsistent.
