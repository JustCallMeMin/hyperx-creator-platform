# Cloud & Infrastructure Security Rules

## Purpose

Define security standards for cloud resources, network configuration, and server hardening.

---

## Infrastructure as Code (IaC) Security

- **Hardening**: Use pre-hardened images and follow CSP best practices (e.g., AWS Well-Architected).
- **Static Analysis**: Run security benchmarks (e.g., Checkov, tfsec) on IaC templates before deployment.
- **Drift Detection**: Monitor for manual changes to infrastructure that bypass the IaC pipeline.

---

## Network Security

- **Micro-segmentation**: Use Security Groups and VPC Private Subnets to isolate critical assets.
- **Egress Filtering**: Restrict outbound traffic from servers to only necessary destinations.
- **Inbound Protection**: Use WAFs and Load Balancers to terminate SSL and filter common web attacks.

---

## Observability & Logging

- **Audit Logs**: Enable and retain audit logs for all cloud provider API calls (e.g., CloudTrail).
- **Centralized Logs**: Ship security logs to a centralized, tamper-resistant location.
- **Alerting**: Configure real-time alerts for critical events like unauthorized access or privilege escalation.

---

## STOP Gate

You MUST STOP if:

- Critical infrastructure is exposed to the public internet without passing through a security layer (LB/WAF).
- Default administrative credentials or open firewall ports (e.g., SSH/RDP to 0.0.0.0/0) are found.
