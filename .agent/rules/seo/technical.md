# Technical SEO & Core Web Vitals

## Core Web Vitals Targets

Maintain a high performance bar to ensure indexability and user satisfaction.

| Metric | Good | Poor |
|:---|:---|:---|
| **LCP** (Largest Contentful Paint) | < 2.5s | > 4.0s |
| **INP** (Interaction to Next Paint) | < 200ms | > 500ms |
| **CLS** (Cumulative Layout Shift) | < 0.1 | > 0.25 |

---

## Technical SEO Checklist

All production pages MUST satisfy:

- [ ] **XML Sitemap**: Correctly structured and submitted to Search Console.
- [ ] **robots.txt**: Configured to allow crawling of essential assets while protecting sensitive paths.
- [ ] **Canonicalization**: Correct canonical tags to prevent duplicate content issues.
- [ ] **HTTPS**: Valid SSL/TLS certificate is mandatory for trust and ranking.
- [ ] **Mobile-Friendly**: Responsive design that passes the Mobile-Friendly test.
- [ ] **Indexability**: No accidental `noindex` tags on priority pages.
- [ ] **Performance**: Passing the Core Web Vitals thresholds defined above.

---

## Technical Audit Best Practices

- Use `lighthouse` or `pagespeed-insights` for performance measurement.
- Verify structured data validity using the Schema Markup Validator.
- Ensure all resources (CSS, JS, Images) are crawlable.
