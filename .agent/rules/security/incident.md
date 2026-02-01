# Security Incident & Breach Response

## Purpose

Define protocols for high-pressure security situations and breach recovery.

---

## Breach Containment

- **Isolation**: Immediately isolate compromised systems or assets to prevent lateral movement.
- **Credential Rotation**: Assume all secrets in a compromised environment are leaked; initiate rotation.
- **Session Revocation**: Kill all active sessions for affected accounts or entire platforms if necessary.

---

## Evidence Preservation

- **Non-destructive**: Avoid actions that overwrite logs or modify file timestamps during initial response.
- **Imaging**: Capture disk and memory images of compromised systems before remediation.
- **Audit Trails**: Preserve all access logs, network traces, and application events related to the incident time window.

---

## Communication & Handoff

- **Truth over Speed**: Provide accurate, verified information rather than speculative updates.
- **Escalation**: Ensure legal, compliance, and senior leadership are notified per the organization's policy.
- **Remediation Plan**: Transition to a recovery phase only after the root cause is identified and patched.

---

## STOP Gate

You MUST STOP if:

- Incident response actions threaten to destroy critical evidence prematurely.
- The containment strategy is undefined or likely to cause more damage than the breach.
