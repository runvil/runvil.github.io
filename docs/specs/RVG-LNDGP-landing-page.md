# Specification — Runvil Landing Page

| Field       | Value                                       |
| ----------- | ------------------------------------------- |
| SpecID      | RVG-LNDGP                                   |
| Title       | Runvil Landing Page                         |
| Status      | Draft                                       |
| Date        | 2026-08-18                                  |
| Author      | Runvil Contributors                         |
| Domain      | Web — Landing Page                          |

## 1. Context

The landing page of the Runvil ecosystem (served at `https://runvil.github.io/`) needs to be a professional, visually stunning, and highly engaging showcase. It needs to reflect the framework's values: blazing fast speed, simplicity, and modern developer experience. 

## 2. Problem Statement

The current landing page is generated as a 1-page book using `mdbind` chapter layouts, which lacks marketing elements, visual appeal, premium aesthetics, and responsive layout spacing. It looks like a minimal text document rather than a high-selling-value gateway for a meta-framework.

## 3. Goals

- G1 — Create a modern, aesthetic landing page with dark/light mode support using custom properties.
- G2 — Present clear product values: Go native speed, Declarative SSG, Scoped CSS, and Book Builder.
- G3 — Provide clear getting started instructions and CTAs for repositories and documentation.
- G4 — Built declaratively using the framework's native `ssg.yaml` configuration.

## 4. Non-Goals

- NG1 — Multi-page marketing sub-routes (e.g. pricing, blog); the landing page is a single landing portal.
- NG2 — Custom theme systems independent of `framework/ui`.

## 5. Requirements

| ID         | Requirement                                                       | Priority |
| ---------- | ----------------------------------------------------------------- | -------- |
| LND-SSG-010 | The landing page must be generated using `ssg.yaml` (declarative SSG mode) instead of `manuscript/` (book mode). | Must |
| LND-HERO-020 | Display a premium Hero section with a bold headline ("Runvil"), a descriptive subtitle, a call-to-action to "/docs/" (Documentation) and a secondary action to the GitHub repository, along with a code install block. | Must |
| LND-FEAT-030 | Display a responsive grid of 4 core features: Go Native, Declarative SSG, Scoped CSS, and Book Builder, with clean layouts and icons. | Must |
| LND-CODE-040 | Include an interactive-looking code preview component showing component composition (e.g. a simple `<Header>` or custom layout) to demonstrate how Runvil works. | Must |
| LND-THEME-050 | Integrate the standard theme-toggle switcher (`Theme.Button`) in a clean top header. | Must |
| LND-DESIGN-060 | Use CSS variables (`--primary`, `--base-1`, etc.) with responsive padding, modern sans-serif typography, elegant border card glows, and subtle scale-up micro-animations on hover. | Must |
| LND-RESP-070 | Ensure layout adjusts perfectly for viewports from mobile (320px) up to ultra-wide (1440px+). | Must |
| LND-HERO-021 | Hero must sell the value proposition with a version/status badge, a strong headline, and a trust stats row (speed, zero-JS, open source) below the CTAs. | Must |
| LND-SOCIAL-080 | Display a social proof band between the hero and features with 4 key metrics (build time, JS runtime size, Go purity, license). | Must |
| LND-FEAT-031 | Feature cards must use inline SVG icons in gradient chips instead of emoji, add hover glow, and include "Learn more" links. | Must |
| LND-CODE-041 | Showcase both SSG mode and Book mode side by side with separate tab-styled code cards demonstrating the single `runvil.yaml` source of truth. | Must |
| LND-COMPARE-090 | Include a comparison section ("Runvil vs. the alternatives") positioning Runvil against JS frameworks and hand-rolled scripts. | Must |
| LND-CTA-100 | CTA section must be a gradient box with a headline, primary + secondary actions, and a meta line (Go version, license, OS support). | Must |

## 6. Non-Functional Requirements

- NFR1 — **Aesthetics.** The landing page must look modern and premium, using glassmorphism, subtle gradients, and custom spacing.
- NFR2 — **Standard-compliant.** Use semantic HTML5 elements (`<header>`, `<main>`, `<section>`, `<footer>`).
- NFR3 — **Performance.** Ensure clean, lightweight styling with fast load times.

## 7. Success Criteria

- S1 — Running `runvil build` in `runvil.github.io/` produces `index.html`, `404.html`, and `assets/style.css` matching the design.
- S2 — Theme switcher toggles light/dark modes correctly and colors adjust automatically.
- S3 — Elements scale responsively on mobile viewports.
- S4 — The page contains a social proof band, SVG-icon feature cards, a dual-mode code showcase, and a comparison section.

## 8. Related Specifications

| SpecID    | Title                                           |
| --------- | ----------------------------------------------- |
| [RVF-5K7PZ](https://github.com/runvil/framework/blob/main/docs/specs/RVF-5K7PZ-web-theming-system.md) | UI Theming System |
| [RVF-PN41Q](https://github.com/runvil/framework/blob/main/docs/specs/RVF-PN41Q-static-site-generator.md) | Static Site Generator |
