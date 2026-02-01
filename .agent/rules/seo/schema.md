# Structured Data & Schema Markup

## Purpose

Use JSON-LD to provide explicit clues about the meaning of a page to search engines and AI systems.

---

## Mandatory Schemas (where applicable)

- **Organization**: Define the entity behind the content.
- **WebSite / WebPage**: Basic structure for every page.
- **Article / BlogPosting**: For all content-driven pages, including `datePublished` and `author`.
- **FAQPage**: Extremely high value for GEO and Google Rich Results.
- **BreadcrumbList**: Helps search engines understand site hierarchy.

---

## Schema Quality Rules

- [ ] **Completeness**: Fill in all recommended fields, not just the required ones.
- [ ] **Consistency**: Ensure data in JSON-LD matches what is visible on the page.
- [ ] **Validity**: MUST pass the Schema Markup Validator or Google Rich Results Test.
- [ ] **Author Entity**: Link authors to their profile pages or social media to support E-E-A-T.

---

## Technical Implementation

- Prefer **JSON-LD** over Microdata or RDFa.
- Inject schema as early as possible in the `<head>` or at the end of the `<body>`.
- Use unique IDs (`@id`) to connect different entities across the site.
