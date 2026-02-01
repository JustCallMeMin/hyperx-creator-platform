---
description: Deployment command for production releases. Pre-flight checks and deployment execution.
---

---
description: Controlled deployment workflow for preview and production releases with mandatory pre-flight checks and verification.
---

## Invocation Rules

This workflow MAY be invoked only when:
- Code is ready for preview or production release.
- All required checks have passed locally or in CI.
- A deployment decision has been approved.

This workflow MUST NOT be invoked when:
- Critical tests are failing.
- Required safety or security checks have not been performed.
- An active Incident Mode is in progress.

---

# /deploy — Controlled Deployment

$ARGUMENTS

---

## Purpose

Execute a controlled deployment with explicit pre-flight validation,
clear execution steps, and post-deployment verification.
This workflow prioritizes safety, traceability, and rollback readiness.

---

## Supported Commands

```
/deploy
/deploy check
/deploy preview
/deploy production
/deploy rollback
```

---

## Behavior

When `/deploy` is triggered:

1. **Pre-Flight Validation**
   - Verify code quality, tests, and build readiness.
   - Confirm security and configuration requirements.
   - Ensure documentation and change records are up to date.

2. **Deployment Preparation**
   - Build the application artifacts.
   - Select target environment (preview or production).
   - Confirm deployment parameters.

3. **Deployment Execution**
   - Deploy artifacts to the selected platform.
   - Apply environment-specific configuration.

4. **Post-Deployment Verification**
   - Run health checks.
   - Verify critical service availability.
   - Confirm successful rollout.

5. **Completion or Rollback**
   - Mark deployment complete if verification passes.
   - Offer rollback path if verification fails.

---

## Pre-Flight Checklist

Before any deployment, the following MUST be satisfied:

```markdown
### Code Quality
- [ ] Type checks pass
- [ ] Linting passes
- [ ] Automated tests pass

### Security
- [ ] No hardcoded secrets
- [ ] Environment variables documented
- [ ] Dependency audit completed

### Performance
- [ ] Bundle size reviewed
- [ ] Debug logging removed
- [ ] Assets optimized

### Documentation
- [ ] README updated
- [ ] CHANGELOG updated
- [ ] API documentation current
```

---

## Output Expectations

A deployment result MUST include:
- Target environment
- Version or release identifier
- Verification status
- Rollback availability

---

## Output Format

### Successful Deployment

```markdown
## Deployment Complete

### Summary
- Version: vX.Y.Z
- Environment: production
- Duration: [time]
- Platform: [platform name]

### Access
- Application URL: [url]
- Management Dashboard: [url]

### Changes
- [Brief summary of changes]

### Verification
- API responding as expected
- Dependent services healthy
- No critical errors detected
```

---

### Failed Deployment

```markdown
## Deployment Failed

### Failure Point
[Build / Deploy / Verification]

### Details
[Error messages or logs]

### Resolution Steps
1. Identify and fix the issue.
2. Re-run local and CI checks.
3. Retry deployment.

### Rollback
Previous stable version remains available.
Rollback may be executed if required.
```

---

## Platform Notes

| Platform | Deployment Method | Notes |
|---------|-------------------|-------|
| Vercel | Managed CLI | Common for web frontends |
| Railway | Managed CLI | Requires platform authentication |
| Fly.io | Managed CLI | Requires infrastructure configuration |
| Docker | Container-based | For self-hosted environments |

---

## Usage Examples

```
/deploy
/deploy check
/deploy preview
/deploy production
/deploy rollback
```

---

## Key Principles

- Never deploy without passing pre-flight checks.
- Verification is mandatory after deployment.
- Rollback must always be possible.
- Deployment decisions must be traceable.
